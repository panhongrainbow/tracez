package model

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"sync"
	"text/template"
	"time"
	"unsafe"
)

type SpanContext struct {
	TraceID    string `json:"TraceID"`
	SpanID     string `json:"SpanID"`
	TraceFlags string `json:"TraceFlags"`
	TraceState string `json:"TraceState"`
	Remote     bool   `json:"Remote"`
}

type Attribute struct {
	Key   string `json:"Key"`
	Value struct {
		Type  string      `json:"Type"`
		Value interface{} `json:"ValueString"`
	} `json:"ValueString"`
}

type Event struct {
	Name                  string      `json:"Name"`
	Attributes            []Attribute `json:"Attributes"`
	DroppedAttributeCount int         `json:"DroppedAttributeCount"`
	Time                  time.Time   `json:"Time"`
	digitTime             [8]int
}

type Status struct {
	Code        string `json:"Code"`
	Description string `json:"Description"`
}

type InstrumentationLibrary struct {
	Name      string `json:"Name"`
	Version   string `json:"Version"`
	SchemaURL string `json:"SchemaURL"`
}

type TracingData struct {
	Name                   string `json:"Name"`
	SpanContext            SpanContext
	Parent                 SpanContext
	SpanKind               int       `json:"SpanKind"`
	StartTime              time.Time `json:"StartTime"`
	digitStartTime         [8]int
	EndTime                time.Time `json:"EndTime"`
	digitEndTime           [8]int
	Attributes             []Attribute            `json:"Attributes"`
	Events                 []Event                `json:"Events"`
	Links                  any                    `json:"Links"`
	Status                 Status                 `json:"Status"`
	DroppedAttributes      int                    `json:"DroppedAttributes"`
	DroppedEvents          int                    `json:"DroppedEvents"`
	DroppedLinks           int                    `json:"DroppedLinks"`
	ChildSpanCount         int                    `json:"ChildSpanCount"`
	Resource               []Attribute            `json:"Resource"`
	InstrumentationLibrary InstrumentationLibrary `json:"InstrumentationLibrary"`
}

type Analyze struct {
	Collect []SubAnalyze
	AreaKey []string
}

type SubAnalyze struct {
	Path          []string
	VarPath       string
	ConditionPath string
	Key           string
	Type          string
	ShowKey       bool
}

// NewAnalyze analyzeJSON
func (an *Analyze) NewAnalyze(data map[string]interface{}, prefix string) {
	//
	collect := new([]SubAnalyze)
	gatherJsonKeyAndPath(data, prefix, collect)
	sort.Sort(JSONInfoSlice(*collect))
	//
	an.Collect = *collect
	an.marksDuplicatedJsonKeys()

	fmt.Println()
}

// gatherJsonKeyAndPath iterates through JSON data, gathering keys with type info recursively.
func gatherJsonKeyAndPath(data map[string]interface{}, prefix string, collect *[]SubAnalyze) {
	// Iterate through each key-value pair in the data map
	for key, value := range data {
		// Create a new prefix by combining the current prefix and key
		newPrefix := fmt.Sprintf("%s.%s", prefix, key)

		// Separate the path and the key value, and it will be easier to handle later on.
		arr := strings.Split(newPrefix, ".")

		// Determine the type of the current value using reflection
		valueType := reflect.TypeOf(value)

		// Create a JSONInfo struct to store information about the current key
		subAnalyze := SubAnalyze{
			Path:    arr[0 : len(arr)-1],
			VarPath: newPrefix,
			Key:     arr[len(arr)-1],
			ShowKey: true,
		}

		subArr := arr[1 : len(arr)-1]
		if len(subArr) == 1 {
			subAnalyze.ConditionPath = "if " + subArr[0] + " == true {"
		} else if len(subArr) > 1 {
			subAnalyze.ConditionPath = "if " + strings.Join(subArr, " == true && ") + " == true {"
		}

		// If the value type is not nil, assign its string representation to the info struct
		if valueType != nil {
			subAnalyze.Type = valueType.String()
		}

		// Append the info struct to the result slice
		*collect = append(*collect, subAnalyze)

		// [recursive function ↩️]
		// Check the type of the value and recursively analyze nested maps or arrays
		switch value.(type) {
		case map[string]interface{}:
			gatherJsonKeyAndPath(value.(map[string]interface{}), newPrefix, collect)
		case []interface{}:
			gatherJsonKeyAndPath(value.([]interface{})[0].(map[string]interface{}), newPrefix, collect)
		}
	}
}

// marksDuplicatedJsonKeys marks duplicated JSON keys.
func (an *Analyze) marksDuplicatedJsonKeys() {
	var perviousKey string
	for i := 0; i < len(an.Collect); i++ {
		if perviousKey == an.Collect[i].Key {
			an.Collect[i].ShowKey = false
		}
		perviousKey = an.Collect[i].Key
	}
}

func generateUnmarshal(slice []SubAnalyze) string {
	tmpl := `
// UnmarshalByGen is an automatically generated function that parses JSON data
// and populates the provided TracingData structure.
// It cannot be modified. (Do Not Edit ⛔️)
func UnmarshalByGen(jsonTracingLog []byte, result *TracingData) (err error) {
	// Using DetectJsonString to extract the key from the JSON trace log.
	var keyTail, keyLength int
	positionNext, keyTail, keyLength = DetectJsonString(positionCurrent, jsonTracingLog)

	// If no quoted key was found, return an error.
	if keyLength == 0 {
		err = fmt.Errorf("no quoted key found")
		block = Block_Error_Happens
		return
	}

	// Extract the key from the trace log.
	key := string(jsonTracingLog[(keyTail - keyLength):keyTail])

	// Determine the which block based on the key.
	switch key {
	{{- range . }}
		{{- if eq .ShowKey true }}
		case "{{ .Key }}":
		{{- end }}
		{{- if ne .ConditionPath "" }}
			{{ .ConditionPath }}
				// handle1
			}
		{{- end }}
		{{- if eq .ConditionPath "" }}
			if firstLayer == true {
				// handle2
			}
		{{- end }}
	{{- end }}
	}
}
	`
	t, err := template.New("switchTemplate").Parse(tmpl)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return ""
	}

	var code strings.Builder
	err = t.Execute(&code, slice)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return ""
	}

	return code.String()
}

type JSONInfoSlice []SubAnalyze

func (s JSONInfoSlice) Len() int           { return len(s) }
func (s JSONInfoSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s JSONInfoSlice) Less(i, j int) bool { return s[i].Key < s[j].Key }

// DetectJsonNonString scans the bool, int value in JSON trace log.
//
//go:inline
func DetectJsonNonString(positionCurrent int, jsonTracingLog []byte) (positionNext, boolTail, boolLength int) {
	// Initialize some variables.
	var inBool bool
	var outBool int

	// Iterate through the JSON trace log bytes.
	for ; positionCurrent < len(jsonTracingLog); positionCurrent++ {

		b := jsonTracingLog[positionCurrent]
		if b == ' ' || b == ',' || b == '}' || b == ':' {
			outBool++
			if outBool == 1 {
				inBool = !inBool
			}
			if boolLength > 0 && !inBool {
				positionNext = positionCurrent // Exclude adding 1 here, as I need to start reading from the comma.
				break
			}
		} else if inBool {
			boolLength++
			outBool = 0
		}
	}

	// Determine the key value by using keyTail and keyLength.
	boolTail = positionCurrent

	return
}

// DetectJsonString scans the string value in JSON trace log.
//
//go:inline
func DetectJsonString(positionCurrent int, jsonTracingLog []byte) (positionNext, keyValueTail, keyValueLength int) {
	// Initialize some variables.
	var inQuotes bool

	// Iterate through the JSON trace log bytes.
	for ; positionCurrent < len(jsonTracingLog); positionCurrent++ {
		b := jsonTracingLog[positionCurrent]
		if b == '"' {
			inQuotes = !inQuotes
			if keyValueLength > 0 && !inQuotes {
				positionNext = positionCurrent + 1 // Next time, start counting from the next byte.
				break
			}
		} else if inQuotes {
			keyValueLength++
		}
	}

	// Determine the key value by using keyTail and keyLength.
	keyValueTail = positionCurrent

	return
}

// >>>>> >>>>> >>>>> >>>>> >>>>> old

func ValueString(jsonData []byte, start int) (result string, end int) {
	var valueStart, valueEnd, count int
	// end = start
	for start = start + 1; start < len(jsonData); start++ {
		if jsonData[start] == '"' {
			count++
			if count == 1 {
				valueStart = start
			}
			if count == 2 {
				valueEnd = start
				result = string(jsonData[valueStart+1 : valueEnd])
				end = valueEnd
				return
			}
		}
	}
	end = valueEnd
	return
}

func ValueTime(jsonData []byte, start int) (result time.Time, end int) {
	var valueStart, valueEnd, count int
	layout := "2006-01-02T15:04:05.000000000Z07:00"
	// end = start
	for start = start + 1; start < len(jsonData); start++ {
		if jsonData[start] == '"' {
			count++
			if count == 1 {
				valueStart = start
			}
			if count == 2 {
				valueEnd = start
				result, _ = time.Parse(layout, string(jsonData[valueStart+1:valueEnd]))
				end = valueEnd
				return
			}
		}
	}
	end = valueEnd
	return
}

func ValueBool(jsonData []byte, start int) (result bool, end int) {
	end = start + 1
	emtpy := start + 5
	for ; end < emtpy; end++ {
		if jsonData[end] == ' ' {
			continue
		}
	}
	if string(jsonData[end:end+5]) == "true" {
		result = true
		end = end + 5
	}
	return
}

func ValueInt(jsonData []byte, start int) (result int, end int) {
	end = start + 1
	emtpy := start + 5
	for ; end < emtpy; end++ {
		if jsonData[end] == ' ' {
			continue
		}
	}
	result = 1
	return
}

const (
	NotInAnyBlock = iota
	SpanContextBlock
	ParentBlock
	AttributesBlock
	EventsAttributesBlock
	StatusBlock
	ResourceBlock
	InstrumentationLibraryBlock
)

var AttributePool = sync.Pool{
	New: func() interface{} {
		return make([]Attribute, 0, 5)
	},
}

// Unmarshal_Old function processes JSON strings in the tracing log faster than standard packages.
// Performing benchmark testing in a desktop environment is actually troublesome.
// It is because it makes the performance unstable.
// The best method is to do comparisons.
// For JSON processing, third-party packages will be about twice as fast as standard packages,
// and self-written parsers will be about 6 to 7 times faster than benchmark packages.
// In reality, it was about 8.3 times faster, which is an acceptable result.
func Unmarshal_Old(jsonData []byte, result *TracingData) error {
	// Initialize the return value first.
	result.Attributes = AttributePool.Get().([]Attribute)
	result.Resource = AttributePool.Get().([]Attribute)
	var err error

	// Record the processing location of the JSON string.
	var positionCurrent int
	var positionDoubleQuotes, positionPreviousDoubleQuotes, colon int
	var positionPassAttributesBlock, positionPassEventsAttributesBlock bool

	for i := 0; i < len(jsonData); i++ {
		switch jsonData[i] {
		case '"':
			positionPreviousDoubleQuotes = positionDoubleQuotes
			positionDoubleQuotes = i
			continue
		case ':':
			colon = i
		/*case ',':
			comma = i
			continue
		case '{':
			leftBrace = i
			continue
		case '}':
			rightBrace = i
			continue*/
		default:
			continue
		}

		if colon > positionDoubleQuotes && positionDoubleQuotes > positionPreviousDoubleQuotes {
			attr := ByteArrayToString(jsonData[positionPreviousDoubleQuotes+1 : positionDoubleQuotes])
			switch attr {
			case "SpanContext":
				positionCurrent = SpanContextBlock
			case "Parent":
				positionCurrent = ParentBlock
			case "Attributes":
				positionCurrent = AttributesBlock
			case "Events":
				positionCurrent = EventsAttributesBlock
				positionPassAttributesBlock = true
			case "Status":
				positionCurrent = StatusBlock
			case "Resource":
				positionCurrent = ResourceBlock
				positionPassEventsAttributesBlock = true
			case "InstrumentationLibrary":
				positionCurrent = InstrumentationLibraryBlock
			case "SpanKind":
				result.SpanKind, i, err = ByteArrayToFindInt(i, &jsonData)
				if err != nil {
					return err
				}
			case "DroppedAttributes":
				result.DroppedAttributes, i, err = ByteArrayToFindInt(i, &jsonData)
				if err != nil {
					return err
				}
			case "DroppedEvents":
				result.DroppedEvents, i, err = ByteArrayToFindInt(i, &jsonData)
				if err != nil {
					return err
				}
			case "DroppedLinks":
				result.DroppedLinks, i, err = ByteArrayToFindInt(i, &jsonData)
				if err != nil {
					return err
				}
			case "ChildSpanCount":
				result.ChildSpanCount, i, err = ByteArrayToFindInt(i, &jsonData)
				if err != nil {
					return err
				}
			case "DroppedAttributeCount":
				result.Events[len(result.Events)-1].DroppedAttributeCount, i, err = ByteArrayToFindInt(i, &jsonData)
				if err != nil {
					return err
				}
			default:
			LOOP2:
				for ; i < len(jsonData); i++ {
					if jsonData[i] == '"' {
						positionPreviousDoubleQuotes = positionDoubleQuotes
						positionDoubleQuotes = i
						if positionPreviousDoubleQuotes > colon {
							valueByte2 := jsonData[positionPreviousDoubleQuotes+1 : positionDoubleQuotes]
							value := ByteArrayToString(valueByte2)
							switch attr {
							case "Name":
								switch positionCurrent {
								case NotInAnyBlock:
									result.Name = value
								case EventsAttributesBlock:
									result.Events = append(result.Events, Event{Name: value, Attributes: AttributePool.Get().([]Attribute)})
								case InstrumentationLibraryBlock:
									result.InstrumentationLibrary.Name = value
								}
							case "TraceID":
								switch positionCurrent {
								case SpanContextBlock:
									result.SpanContext.TraceID = value
								case ParentBlock:
									result.Parent.TraceID = value
								}
							case "SpanID":
								switch positionCurrent {
								case SpanContextBlock:
									result.SpanContext.SpanID = value
								case ParentBlock:
									result.Parent.SpanID = value
								}
							case "TraceFlags":
								switch positionCurrent {
								case SpanContextBlock:
									result.SpanContext.TraceFlags = value
								case ParentBlock:
									result.Parent.TraceFlags = value
								}
							case "TraceState":
								switch positionCurrent {
								case SpanContextBlock:
									result.SpanContext.TraceState = value
								case ParentBlock:
									result.Parent.TraceState = value
								}
							case "Remote":
								var valueBool bool
								if value == "true" || value == "True" {
									valueBool = true

									switch positionCurrent {
									case SpanContextBlock:
										result.SpanContext.Remote = valueBool
									case ParentBlock:
										result.Parent.Remote = valueBool
									}
								}
							case "StartTime":
								result.digitStartTime[0] = ByteArrayToInt(valueByte2[0:4])   // year
								result.digitStartTime[1] = ByteArrayToInt(valueByte2[5:7])   // month
								result.digitStartTime[2] = ByteArrayToInt(valueByte2[8:10])  // day
								result.digitStartTime[3] = ByteArrayToInt(valueByte2[11:13]) // hour
								result.digitStartTime[4] = ByteArrayToInt(valueByte2[14:16]) // minute
								result.digitStartTime[5] = ByteArrayToInt(valueByte2[17:19]) // second
								result.digitStartTime[6] = ByteArrayToInt(valueByte2[20:29]) // micro second
								// result.digitStartTime[7] = ByteArrayToInt(valueByte2[30:32]) // zone
							case "EndTime":
								result.digitEndTime[0] = ByteArrayToInt(valueByte2[0:4])   // year
								result.digitEndTime[1] = ByteArrayToInt(valueByte2[5:7])   // month
								result.digitEndTime[2] = ByteArrayToInt(valueByte2[8:10])  // day
								result.digitEndTime[3] = ByteArrayToInt(valueByte2[11:13]) // hour
								result.digitEndTime[4] = ByteArrayToInt(valueByte2[14:16]) // minute
								result.digitEndTime[5] = ByteArrayToInt(valueByte2[17:19]) // second
								result.digitEndTime[6] = ByteArrayToInt(valueByte2[20:29]) // micro second
							// result.digitEndTime[7] = ByteArrayToInt(valueByte2[30:32]) // zone
							case "Key":
								i++
								valueType, valueValue, valueI := ByteArrayToValueString(i, &jsonData)

								tmp := Attribute{
									Key: value,
									Value: struct {
										Type  string      `json:"Type"`
										Value interface{} `json:"ValueString"`
									}{
										Type:  valueType,
										Value: valueValue,
									},
								}

								if positionPassEventsAttributesBlock {
									result.Resource = append(result.Resource, tmp)
								} else if positionPassAttributesBlock {
									result.Events[len(result.Events)-1].Attributes = append(result.Events[len(result.Events)-1].Attributes, tmp)
								} else {
									result.Attributes = append(result.Attributes, tmp)
								}
								i = valueI
							case "Time":
								result.Events[len(result.Events)-1].digitTime[0] = ByteArrayToInt(valueByte2[0:4])   // year
								result.Events[len(result.Events)-1].digitTime[1] = ByteArrayToInt(valueByte2[5:7])   // month
								result.Events[len(result.Events)-1].digitTime[2] = ByteArrayToInt(valueByte2[8:10])  // day
								result.Events[len(result.Events)-1].digitTime[3] = ByteArrayToInt(valueByte2[11:13]) // hour
								result.Events[len(result.Events)-1].digitTime[4] = ByteArrayToInt(valueByte2[14:16]) // minute
								result.Events[len(result.Events)-1].digitTime[5] = ByteArrayToInt(valueByte2[17:19]) // second
								result.Events[len(result.Events)-1].digitTime[6] = ByteArrayToInt(valueByte2[20:29]) // micro second
							case "Code":
								result.Status.Code = ByteArrayToString(jsonData[positionPreviousDoubleQuotes+1 : positionDoubleQuotes])
							case "Description":
								result.Status.Description = ByteArrayToString(jsonData[positionPreviousDoubleQuotes+1 : positionDoubleQuotes])
							}
							break LOOP2
						}
					}
				}
			}
		}

	}

	return err
}

//go:inline
func ByteArrayToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//go:inline
func ByteArrayToInt(b []byte) (number int) {
	for i := 0; i < len(b); i++ {
		number = number*10 + int(b[i]-48)
	}
	return
}

//go:inline
func ByteArrayToFindInt(i int, jsonData *[]byte) (number int, next int, err error) {
	var numberStart int
LOOP1:
	for ; i < len(*jsonData); i++ {
		if (*jsonData)[i] == ',' {
			number = ByteArrayToInt((*jsonData)[numberStart:i])
			if err != nil {
				return
			}
			break LOOP1
		}
		if (*jsonData)[i] != ' ' {
			numberStart = i
		}
	}
	next = i
	return
}

//go:inline
func ByteArrayToValueString(i int, jsonData *[]byte) (valueType string, value string, next int) {

	var countDoubleQuotes, previousDoubleQuotes, doubleQuotes int

	for ; i < len(*jsonData); i++ {
		if (*jsonData)[i] == '"' {
			countDoubleQuotes++
			previousDoubleQuotes = doubleQuotes
			doubleQuotes = i
			if countDoubleQuotes == 6 {
				valueByte := (*jsonData)[previousDoubleQuotes+1 : doubleQuotes]
				valueType = ByteArrayToString(valueByte)
			}
			if countDoubleQuotes == 8 {
				var previousDoubleQuotes2, doubleQuotes2, arabicStart2, arabicEnd2 int
				for ; i < len(*jsonData); i++ {
					if (*jsonData)[i] == '"' {
						previousDoubleQuotes2 = doubleQuotes2
						doubleQuotes2 = i
					}
					if (*jsonData)[i] >= 49 && (*jsonData)[i] <= 57 {
						if arabicStart2 == 0 {
							arabicStart2 = i
						}
						arabicEnd2 = i
					}
					if (*jsonData)[i] == '}' {
						switch valueType {
						case "STRING":
							value = ByteArrayToString((*jsonData)[previousDoubleQuotes2+1 : doubleQuotes2])
						case "INT64":
							value = ByteArrayToString((*jsonData)[arabicStart2 : arabicEnd2+1])
						}

						next = i
						return
					}
				}
			}
		}

	}

	return
}
