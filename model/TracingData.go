package model

import (
	"time"
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
		Value interface{} `json:"Value"`
	} `json:"Value"`
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

func Value(jsonData []byte, start int) (result string, end int) {
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

func Unmarshal(jsonData []byte) (result TracingData) {
	var entryNameTimes uint8
	var entrySpanContext bool
	var start, attribute int

	for ; start < len(jsonData); start++ {
		if jsonData[start] == ',' || jsonData[start] == '{' || jsonData[start] == '}' {
			attribute = start
		}
		if jsonData[start] == ':' {
			if start-1 > attribute+2 {
				key := string(jsonData[attribute+2 : start-1])
				switch key {
				case "Name":
					if entryNameTimes == 0 {
						result.Name, start = Value(jsonData, start)
						entryNameTimes++
					}
				case "SpanContext":
					entrySpanContext = true
				case "TraceID":
					if entrySpanContext {
						result.SpanContext.TraceID, start = Value(jsonData, start)
					}
				case "SpanID":
					if entrySpanContext {
						result.SpanContext.SpanID, start = Value(jsonData, start)
					}
				case "TraceFlags":
					if entrySpanContext {
						result.SpanContext.TraceFlags, start = Value(jsonData, start)
					}
				case "TraceState":
					if entrySpanContext {
						result.SpanContext.TraceState, start = Value(jsonData, start)
					}
				case "Remote":
					if entrySpanContext {
						var remote string
						remote, start = Value(jsonData, start)
						if remote == "true" {
							result.SpanContext.Remote = true
						}
						entrySpanContext = false
					}
				}
			}
		}
	}
	return
}
