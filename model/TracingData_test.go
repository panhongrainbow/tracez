package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonData = []byte(`{"Name":"functionBSpan","SpanContext":{"TraceID":"77ea467445562b1afb250147b0ddc178","SpanID":"e7d262c286f5b660","TraceFlags":"01","TraceState":"","Remote":false},"Parent":{"TraceID":"77ea467445562b1afb250147b0ddc178","SpanID":"b178c20dce680429","TraceFlags":"01","TraceState":"","Remote":false},"SpanKind":1,"StartTime":"2023-05-29T01:49:32.011159939+08:00","EndTime":"2023-05-29T01:49:33.011344186+08:00","Attributes":[{"Key":"ParameterB","Value":{"Type":"STRING","Value":"ValueB"}}],"Events":[{"Name":"exception","Attributes":[{"Key":"ID","Value":{"Type":"INT64","Value":1}},{"Key":"postscript","Value":{"Type":"STRING","Value":"more details"}},{"Key":"exception.type","Value":{"Type":"STRING","Value":"*errors.errorString"}},{"Key":"exception.message","Value":{"Type":"STRING","Value":"error"}}],"DroppedAttributeCount":0,"Time":"2023-05-29T01:49:32.011163907+08:00"}],"Links":null,"Status":{"Code":"Error","Description":"functionB failed"},"DroppedAttributes":0,"DroppedEvents":0,"DroppedLinks":0,"ChildSpanCount":0,"Resource":[{"Key":"service.name","Value":{"Type":"STRING","Value":"unknown_service:___go_build_github_com_panhongrainbow_tracez_example_openTelemetry2file"}},{"Key":"telemetry.sdk.language","Value":{"Type":"STRING","Value":"go"}},{"Key":"telemetry.sdk.name","Value":{"Type":"STRING","Value":"opentelemetry"}},{"Key":"telemetry.sdk.version","Value":{"Type":"STRING","Value":"1.14.0"}}],"InstrumentationLibrary":{"Name":"functionBTracer","Version":"","SchemaURL":""}}`)

func Test_Check_TracingData(t *testing.T) {
	result := Unmarshal(jsonData)
	fmt.Println(result)
}

// 25576 ns/op
func Benchmark_Estimate_TracingData(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var tracingData TracingData
		_ = json.Unmarshal(jsonData, &tracingData)
	}
}

// 2861 ns/op
func Benchmark_Estimate_TracingData2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Unmarshal(jsonData)
	}
}
