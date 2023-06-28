package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonData = []byte(`{"Name":"functionBSpan","SpanContext":{"TraceID":"77ea467445562b1afb250147b0ddc178","SpanID":"e7d262c286f5b660","TraceFlags":"01","TraceState":"","Remote":false},"Parent":{"TraceID":"77ea467445562b1afb250147b0ddc178","SpanID":"b178c20dce680429","TraceFlags":"01","TraceState":"","Remote":false},"SpanKind":1,"StartTime":"2023-05-29T01:49:32.011159939+08:00","EndTime":"2023-05-29T01:49:33.011344186+08:00","Attributes":[{"Key":"ParameterB","ValueString":{"Type":"STRING","ValueString":"ValueB"}}],"Events":[{"Name":"exception","Attributes":[{"Key":"ID","ValueString":{"Type":"INT64","ValueString":1}},{"Key":"postscript","ValueString":{"Type":"STRING","ValueString":"more details"}},{"Key":"exception.type","ValueString":{"Type":"STRING","ValueString":"*errors.errorString"}},{"Key":"exception.message","ValueString":{"Type":"STRING","ValueString":"error"}}],"DroppedAttributeCount":0,"Time":"2023-05-29T01:49:32.011163907+08:00"}],"Links":null,"Status":{"Code":"Error","Description":"functionB failed"},"DroppedAttributes":0,"DroppedEvents":0,"DroppedLinks":0,"ChildSpanCount":0,"Resource":[{"Key":"service.name","ValueString":{"Type":"STRING","ValueString":"unknown_service:___go_build_github_com_panhongrainbow_tracez_example_openTelemetry2file"}},{"Key":"telemetry.sdk.language","ValueString":{"Type":"STRING","ValueString":"go"}},{"Key":"telemetry.sdk.name","ValueString":{"Type":"STRING","ValueString":"opentelemetry"}},{"Key":"telemetry.sdk.version","ValueString":{"Type":"STRING","ValueString":"1.14.0"}}],"InstrumentationLibrary":{"Name":"functionBTracer","Version":"","SchemaURL":""}}`)

func Test_Check_TracingData(t *testing.T) {
	result := Unmarshal(jsonData)
	fmt.Println(result)
}

// 25268 ns/op
func Benchmark_Estimate_TracingData(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var tracingData TracingData
		_ = json.Unmarshal(jsonData, &tracingData)
	}
}

// 3199 ns/op
func Benchmark_Estimate_TracingData2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Unmarshal(jsonData)
	}
}
