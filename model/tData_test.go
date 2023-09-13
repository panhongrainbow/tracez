package model

import (
	"encoding/json"
	thirdJson "github.com/goccy/go-json"
	"github.com/stretchr/testify/require"
	"testing"
)

// var jsonTracingLog = []byte(`{"Name":"functionBSpan","SpanContext":{"TraceID":"77ea467445562b1afb250147b0ddc178","SpanID":"e7d262c286f5b660","TraceFlags":"01","TraceState":"CCC","Remote":true},"Parent":{"TraceID":"77ea467445562b1afb250147b0ddc178","SpanID":"b178c20dce680429","TraceFlags":"01","TraceState":"DDD","Remote":true},"SpanKind":1,"StartTime":"2023-05-29T01:49:32.011159939+08:00","EndTime":"2023-05-29T01:49:33.011344186+08:00","Attributes":[{"Key":"ParameterB","ValueString":{"Type":"STRING","ValueString":"ValueB"}}],"Events":[{"Name":"exception","Attributes":[{"Key":"ID","ValueString":{"Type":"INT64","ValueString":1}},{"Key":"postscript","ValueString":{"Type":"STRING","ValueString":"more details"}},{"Key":"exception.type","ValueString":{"Type":"STRING","ValueString":"*errors.errorString"}},{"Key":"exception.message","ValueString":{"Type":"STRING","ValueString":"error"}}],"DroppedAttributeCount":12,"Time":"2023-05-29T01:49:32.011163907+08:00"}],"Links":null,"Status":{"Code":"Error","Description":"functionB failed"},"DroppedAttributes":13,"DroppedEvents":14,"DroppedLinks":15,"ChildSpanCount":11,"Resource":[{"Key":"service.name","ValueString":{"Type":"STRING","ValueString":"unknown_service:___go_build_github_com_panhongrainbow_tracez_example_openTelemetry2file"}},{"Key":"telemetry.sdk.language","ValueString":{"Type":"STRING","ValueString":"go"}},{"Key":"telemetry.sdk.name","ValueString":{"Type":"STRING","ValueString":"opentelemetry"}},{"Key":"telemetry.sdk.version","ValueString":{"Type":"STRING","ValueString":"1.14.0"}}],"InstrumentationLibrary":{"Name":"functionBTracer","Version":"YYY","SchemaURL":"XXX"}}`)

var jsonTracingLog = []byte(`{"Name":"functionBSpan","SpanContext":{"TraceID":"77ea467445562b1afb250147b0ddc178","SpanID":"e7d262c286f5b660","TraceFlags":"01","TraceState":"","Remote":false},"Parent":{"TraceID":"77ea467445562b1afb250147b0ddc178","SpanID":"b178c20dce680429","TraceFlags":"01","TraceState":"","Remote":false},"SpanKind":1,"StartTime":"2023-05-29T01:49:32.011159939+08:00","EndTime":"2023-05-29T01:49:33.011344186+08:00","Attributes":[{"Key":"ParameterB","ValueString":{"Type":"STRING","ValueString":"ValueB"}},{"Key":"ParameterC","ValueString":{"Type":"STRING","ValueString":"ValueC"}},{}],"Events":[{"Name":"exception","Attributes":[{"Key":"ID","ValueString":{"Type":"INT64","ValueString":1}},{"Key":"postscript","ValueString":{"Type":"STRING","ValueString":"more details"}},{"Key":"exception.type","ValueString":{"Type":"STRING","ValueString":"*errors.errorString"}},{"Key":"exception.message","ValueString":{"Type":"STRING","ValueString":"error"}}],"DroppedAttributeCount":12,"Time":"2023-05-29T01:49:32.011163907+08:00"},{}],"Links":null,"Status":{"Code":"Error","Description":"functionB failed"},"DroppedAttributes":0,"DroppedEvents":0,"DroppedLinks":0,"ChildSpanCount":11,"Resource":[{"Key":"service.name","ValueString":{"Type":"STRING","ValueString":"unknown_service:___go_build_github_com_panhongrainbow_tracez_example_openTelemetry2file"}},{"Key":"telemetry.sdk.language","ValueString":{"Type":"STRING","ValueString":"go"}},{"Key":"telemetry.sdk.name","ValueString":{"Type":"STRING","ValueString":"opentelemetry"}},{"Key":"telemetry.sdk.version","ValueString":{"Type":"STRING","ValueString":"1.14.0"}}],"InstrumentationLibrary":{"Name":"functionBTracer","Version":"","SchemaURL":""}}`)

func Test_Check_Unmarshal(t *testing.T) {
	tData := TracingData{}
	err := UnmarshalByTuning(jsonTracingLog, &tData)
	require.NoError(t, err)
	require.Equal(t, err, tData.Name)
}

// Benchmark_Performance_TracingData measures the performance of three different versions for processing tracing data.
func Benchmark_Performance_TracingData(b *testing.B) {
	b.ResetTimer()
	b.Run("Official Version", func(b *testing.B) {
		// The measured speed is 32478 ns/op.
		for i := 0; i < b.N; i++ {
			var tracingData TracingData
			_ = json.Unmarshal(jsonTracingLog, &tracingData)
		}
	})
	b.ResetTimer()
	b.Run("Customized Version", func(b *testing.B) {
		// The measured speed is 7860 ns/op.
		for i := 0; i < b.N; i++ {
			var tracingData TracingData
			_ = UnmarshalByTuning(jsonTracingLog, &tracingData)
			CleanTracingData(&tracingData)
		}
	})
	b.ResetTimer()
	b.Run("Third-Party Version", func(b *testing.B) {
		// The measured speed is 8249 ns/op.
		for i := 0; i < b.N; i++ {
			var tracingData TracingData
			_ = thirdJson.Unmarshal(jsonTracingLog, &tracingData)
		}
	})

}
