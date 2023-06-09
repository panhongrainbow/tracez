package model

import (
	"sync"
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

// Unmarshal function processes JSON strings in the tracing log faster than standard packages.
// Performing benchmark testing in a desktop environment is actually troublesome.
// It is because it makes the performance unstable.
// The best method is to do comparisons.
// For JSON processing, third-party packages will be about twice as fast as standard packages,
// and self-written parsers will be about 6 to 7 times faster than benchmark packages.
// In reality, it was about 8.3 times faster, which is an acceptable result.
func Unmarshal(jsonData []byte) (TracingData, error) {
	result := TracingData{
		Attributes: AttributePool.Get().([]Attribute),
		Resource:   AttributePool.Get().([]Attribute),
	}
	var err error

	var doubleQuotes, previousDoubleQuotes, colon int // comma, leftBrace, rightBrace int
	var position int
	var passAttributesBlock, passEventsAttributesBlock bool

	for i := 0; i < len(jsonData); i++ {
		switch jsonData[i] {
		case '"':
			previousDoubleQuotes = doubleQuotes
			doubleQuotes = i
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

		if colon > doubleQuotes && doubleQuotes > previousDoubleQuotes {
			attr := ByteArrayToString(jsonData[previousDoubleQuotes+1 : doubleQuotes])
			switch attr {
			case "SpanContext":
				position = SpanContextBlock
			case "Parent":
				position = ParentBlock
			case "Attributes":
				position = AttributesBlock
			case "Events":
				position = EventsAttributesBlock
				passAttributesBlock = true
			case "Status":
				position = StatusBlock
			case "Resource":
				position = ResourceBlock
				passEventsAttributesBlock = true
			case "InstrumentationLibrary":
				position = InstrumentationLibraryBlock
			case "SpanKind", "DroppedAttributes", "DroppedEvents", "DroppedLinks", "ChildSpanCount":
				result.SpanKind, i, err = ByteArrayToFindInt(i, &jsonData)
				if err != nil {
					return TracingData{}, err
				}
			case "DroppedAttributeCount":
				result.Events[len(result.Events)-1].DroppedAttributeCount, i, err = ByteArrayToFindInt(i, &jsonData)
				if err != nil {
					return TracingData{}, err
				}
			default:
			LOOP2:
				for ; i < len(jsonData); i++ {
					if jsonData[i] == '"' {
						previousDoubleQuotes = doubleQuotes
						doubleQuotes = i
						if previousDoubleQuotes > colon {
							valueByte2 := jsonData[previousDoubleQuotes+1 : doubleQuotes]
							value := ByteArrayToString(valueByte2)
							switch attr {
							case "Name":
								switch position {
								case NotInAnyBlock:
									result.Name = value
								case EventsAttributesBlock:
									result.Events = append(result.Events, Event{Name: value, Attributes: AttributePool.Get().([]Attribute)})
								case InstrumentationLibraryBlock:
									result.InstrumentationLibrary.Name = value
								}
							case "TraceID":
								switch position {
								case SpanContextBlock:
									result.SpanContext.TraceID = value
								case ParentBlock:
									result.Parent.TraceID = value
								}
							case "SpanID":
								switch position {
								case SpanContextBlock:
									result.SpanContext.SpanID = value
								case ParentBlock:
									result.Parent.SpanID = value
								}
							case "TraceFlags":
								switch position {
								case SpanContextBlock:
									result.SpanContext.TraceFlags = value
								case ParentBlock:
									result.Parent.TraceFlags = value
								}
							case "TraceState":
								switch position {
								case SpanContextBlock:
									result.SpanContext.TraceState = value
								case ParentBlock:
									result.Parent.TraceState = value
								}
							case "Remote":
								var valueBool bool
								if value == "true" || value == "True" {
									valueBool = true

									switch position {
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

								if passEventsAttributesBlock {
									result.Resource = append(result.Resource, tmp)
								} else if passAttributesBlock {
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
							}
							break LOOP2
						}
					}
				}
			}
		}

	}

	AttributePool.Put(result.Attributes[:0])
	AttributePool.Put(result.Resource[:0])
	for j := 0; j < len(result.Events); j++ {
		AttributePool.Put(result.Events[j].Attributes[:0])
	}

	return result, err
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
