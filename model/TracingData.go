package model

import (
	"fmt"
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

func Unmarshal(jsonData []byte) (result TracingData) {
	start := 0
	attribute := 0

	for ; start < len(jsonData); start++ {
		if jsonData[start] == ',' {
			attribute = start
		}
		if jsonData[start] == ':' {
			key := string(jsonData[attribute+2 : start])
			fmt.Println("key->", key)
			valueStart := start + 2

			for valueEnd := valueStart; valueEnd < len(jsonData); valueEnd++ {
				if jsonData[valueEnd] == '"' {
					value := string(jsonData[valueStart:valueEnd])
					fmt.Println("value->", value)
					break
				}
			}

			attribute = start + 1
			start = start + 1
		}
	}
	return
}
