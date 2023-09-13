package model

import (
	"sync"
	"time"
)

// TracingData is the main struct for Tracing Log.
type TracingData struct {
	Name                   string                 `json:"Name"` // The name of the tracing data.
	SpanContext            SpanContext            // The context of the span, including trace ID, span ID, etc.
	Parent                 SpanContext            // The context of the parent span.
	SpanKind               int                    `json:"SpanKind"`               // The kind of the span, such as server, client, producer, etc.
	StartTime              time.Time              `json:"StartTime"`              // The start time of the span.
	EndTime                time.Time              `json:"EndTime"`                // The end time of the span.
	Attributes             []Attribute            `json:"Attributes"`             // The attributes of the span, such as labels, tags, etc.
	Events                 []Event                `json:"Events"`                 // The events that occurred during the span, such as logs, errors, etc.
	Links                  any                    `json:"Links"`                  // The links to other spans or traces.
	Status                 Status                 `json:"Status"`                 // The status of the span, such as OK, Error, etc.
	DroppedAttributes      int                    `json:"DroppedAttributes"`      // The number of attributes that were dropped due to limit.
	DroppedEvents          int                    `json:"DroppedEvents"`          // The number of events that were dropped due to limit.
	DroppedLinks           int                    `json:"DroppedLinks"`           // The number of links that were dropped due to limit.
	ChildSpanCount         int                    `json:"ChildSpanCount"`         // The number of child spans that were created by this span.
	Resource               []Attribute            `json:"Resource"`               // The resource information of the span, such as host name, service name, etc.
	InstrumentationLibrary InstrumentationLibrary `json:"InstrumentationLibrary"` // The instrumentation library information of the span, such as name, version, schema URL, etc.
}

// SpanContext is the struct for Span Context.
type SpanContext struct {
	TraceID    string `json:"TraceID"`    // The unique identifier of the trace.
	SpanID     string `json:"SpanID"`     // The unique identifier of the span within the trace.
	TraceFlags string `json:"TraceFlags"` // The flags that indicate properties of the trace, such as sampled or not.
	TraceState string `json:"TraceState"` // The additional vendor-specific trace information.
	Remote     bool   `json:"Remote"`     // Whether the span context was propagated from a remote parent.
}

// This is a sync pool for attributes to reduce memory allocation and garbage collection.
var attributePool = sync.Pool{
	New: func() interface{} {
		return make([]Attribute, 0, 10) // Create a new slice of attributes with capacity 10.
	},
}

// Attribute is the struct for Attribute.
type Attribute struct {
	Key   string `json:"Key"` // The key or name of the attribute.
	Value struct {
		Type  string      `json:"Type"`        // The type of the attribute value, such as string, int, bool, etc.
		Value interface{} `json:"ValueString"` // The value of the attribute in JSON format.
	} `json:"ValueString"`
}

// Event is the struct for Event.
type Event struct {
	Name                  string      `json:"Name"`                  // The name or description of the event.
	Attributes            []Attribute `json:"Attributes"`            // The attributes of the event, such as parameters, results, etc.
	DroppedAttributeCount int         `json:"DroppedAttributeCount"` // The number of attributes that were dropped due to limit.
	Time                  time.Time   `json:"Time"`                  // The time when the event occurred.
}

// Status is the struct for Status.
type Status struct {
	Code        string `json:"Code"`        // The code or indicator of the status, such as OK, Error, Cancelled, etc.
	Description string `json:"Description"` // The description or message of the status.
}

// InstrumentationLibrary is the struct for Instrumentation Library
type InstrumentationLibrary struct {
	Name      string `json:"Name"`      // The name of the instrumentation library that generated the span data.
	Version   string `json:"Version"`   // The version of the instrumentation library.
	SchemaURL string `json:"SchemaURL"` // The schema URL of the instrumentation library.
}
