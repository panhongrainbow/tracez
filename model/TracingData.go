package model

import "time"

// TracingData struct represents a distributed tracing document
type TracingData struct {
	Name        string `json:"Name"`
	SpanContext struct {
		TraceID    string `json:"TraceID"`
		SpanID     string `json:"SpanID"`
		TraceFlags string `json:"TraceFlags"`
		TraceState string `json:"TraceState"`
		Remote     bool   `json:"Remote"`
	} `json:"SpanContext"`
	Parent struct {
		TraceID    string `json:"TraceID"`
		SpanID     string `json:"SpanID"`
		TraceFlags string `json:"TraceFlags"`
		TraceState string `json:"TraceState"`
		Remote     bool   `json:"Remote"`
	} `json:"Parent"`
	SpanKind   int       `json:"SpanKind"`
	StartTime  time.Time `json:"StartTime"`
	EndTime    time.Time `json:"EndTime"`
	Attributes []struct {
		Key   string `json:"Key"`
		Value struct {
			Type  string `json:"Type"`
			Value string `json:"Value"`
		} `json:"Value"`
	} `json:"Attributes"`
	Events []struct {
		Name       string `json:"Name"`
		Attributes []struct {
			Key   string `json:"Key"`
			Value struct {
				Type  string `json:"Type"`
				Value string `json:"Value"`
			} `json:"Value"`
		} `json:"Attributes"`
		DroppedAttributeCount int       `json:"DroppedAttributeCount"`
		Time                  time.Time `json:"Time"`
	} `json:"Events"`
	Links  any `json:"Links"`
	Status struct {
		Code        string `json:"Code"`
		Description string `json:"Description"`
	} `json:"Status"`
	DroppedAttributes int `json:"DroppedAttributes"`
	DroppedEvents     int `json:"DroppedEvents"`
	DroppedLinks      int `json:"DroppedLinks"`
	ChildSpanCount    int `json:"ChildSpanCount"`
	Resource          []struct {
		Key   string `json:"Key"`
		Value struct {
			Type  string `json:"Type"`
			Value string `json:"Value"`
		} `json:"Value"`
	} `json:"Resource"`
	InstrumentationLibrary struct {
		Name      string `json:"Name"`
		Version   string `json:"Version"`
		SchemaURL string `json:"SchemaURL"`
	} `json:"InstrumentationLibrary"`
}
