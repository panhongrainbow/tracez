package model

import (
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
	SpanKind               int                    `json:"SpanKind"`
	StartTime              time.Time              `json:"StartTime"`
	EndTime                time.Time              `json:"EndTime"`
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

// This one is better.
func Unmarshal(jsonData []byte) (result TracingData) {
	var doubleQuotes, previousDoubleQuotes, colon int //, comma, leftBrace, rightBrace int

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
			case "SpanContext", "TraceFlags", "TraceState", "Remote", "Parent", "SpanKind", "Attributes", "ValueString", "DroppedAttributeCount", "Links", "Status", "DroppedAttributes", "DroppedEvents", "DroppedLinks", "ChildSpanCount", "Resource", "InstrumentationLibrary", "Version", "SchemaURL":
				//
			default:
			LOOP1:
				for ; i < len(jsonData); i++ {
					if jsonData[i] == '"' {
						previousDoubleQuotes = doubleQuotes
						doubleQuotes = i
						if previousDoubleQuotes > colon {
							// fmt.Println(attr, string(jsonData[previousDoubleQuotes+1:doubleQuotes]))
							break LOOP1
						}
					}
				}
			}
		}

	}

	// fmt.Println(comma, leftBrace, rightBrace)

	return
}

func ByteArrayToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// This approach is not good.
func Unmarshal_deprecated(jsonData []byte) (result TracingData) {
	var entryNameTimes uint8
	var entrySpanContext, entryParent bool
	var i, attribute int

	for ; i < len(jsonData); i++ {
		if jsonData[i] == ',' || jsonData[i] == '{' || jsonData[i] == '}' {
			// if jsonData[i] == '{' {
			attribute = i
		}
		if jsonData[i] == ':' {
			if i-1 > attribute+2 {
				key := string(jsonData[attribute+2 : i-1])
				switch key {
				case "Name":
					if entryNameTimes == 0 {
						result.Name, i = ValueString(jsonData, i)
						entryNameTimes++
					}
				case "SpanContext":
					entrySpanContext = true
				case "TraceID":
					if entrySpanContext {
						result.SpanContext.TraceID, i = ValueString(jsonData, i)
					}
					if entryParent {
						result.Parent.TraceID, i = ValueString(jsonData, i)
					}
				case "SpanID":
					if entrySpanContext {
						result.SpanContext.SpanID, i = ValueString(jsonData, i)
					}
					if entryParent {
						result.Parent.SpanID, i = ValueString(jsonData, i)
					}
				case "TraceFlags":
					if entrySpanContext {
						result.SpanContext.TraceFlags, i = ValueString(jsonData, i)
					}
					if entryParent {
						result.Parent.TraceFlags, i = ValueString(jsonData, i)
					}
				case "TraceState":
					if entrySpanContext {
						result.SpanContext.TraceState, i = ValueString(jsonData, i)
					}
					if entryParent {
						result.Parent.TraceState, i = ValueString(jsonData, i)
					}
				case "Remote":
					if entrySpanContext {
						result.SpanContext.Remote, i = ValueBool(jsonData, i)
						entrySpanContext = false
					}
					if entryParent {
						result.Parent.Remote, i = ValueBool(jsonData, i)
						entryParent = false
					}
				case "SpanKind":
					result.SpanKind, i = ValueInt(jsonData, i)
				case "Parent":
					entryParent = true
				case "StartTime":
					result.StartTime, i = ValueTime(jsonData, i)
				}
			}
		}
	}
	return
}
