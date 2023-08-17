package model

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var jsonTracingLog = []byte(`{"Name":"functionBSpan","SpanContext":{"TraceID":"77ea467445562b1afb250147b0ddc178","SpanID":"e7d262c286f5b660","TraceFlags":"01","TraceState":"","Remote":false},"Parent":{"TraceID":"77ea467445562b1afb250147b0ddc178","SpanID":"b178c20dce680429","TraceFlags":"01","TraceState":"","Remote":false},"SpanKind":1,"StartTime":"2023-05-29T01:49:32.011159939+08:00","EndTime":"2023-05-29T01:49:33.011344186+08:00","Attributes":[{"Key":"ParameterB","ValueString":{"Type":"STRING","ValueString":"ValueB"}}],"Events":[{"Name":"exception","Attributes":[{"Key":"ID","ValueString":{"Type":"INT64","ValueString":1}},{"Key":"postscript","ValueString":{"Type":"STRING","ValueString":"more details"}},{"Key":"exception.type","ValueString":{"Type":"STRING","ValueString":"*errors.errorString"}},{"Key":"exception.message","ValueString":{"Type":"STRING","ValueString":"error"}}],"DroppedAttributeCount":0,"Time":"2023-05-29T01:49:32.011163907+08:00"}],"Links":null,"Status":{"Code":"Error","Description":"functionB failed"},"DroppedAttributes":0,"DroppedEvents":0,"DroppedLinks":0,"ChildSpanCount":0,"Resource":[{"Key":"service.name","ValueString":{"Type":"STRING","ValueString":"unknown_service:___go_build_github_com_panhongrainbow_tracez_example_openTelemetry2file"}},{"Key":"telemetry.sdk.language","ValueString":{"Type":"STRING","ValueString":"go"}},{"Key":"telemetry.sdk.name","ValueString":{"Type":"STRING","ValueString":"opentelemetry"}},{"Key":"telemetry.sdk.version","ValueString":{"Type":"STRING","ValueString":"1.14.0"}}],"InstrumentationLibrary":{"Name":"functionBTracer","Version":"","SchemaURL":""}}`)

func Test_Check_Unmarshal(t *testing.T) {
	var result = new(TracingData)
	var err error
	err = Unmarshal(jsonTracingLog, result)
	require.NoError(t, err)
}

// Test_Check_DetectJSONProcessingBlock checks the functionality of the DetectJSONProcessingBlock function.
func Test_Check_DetectJSONProcessingBlock(t *testing.T) {
	// Initialize test parameters.
	var positionCurrent int
	var block uint8
	var key string
	var err error

	// Call the function under test.
	positionCurrent, block, key, err = DetectJSONProcessingBlock(positionCurrent, jsonTracingLog)

	// Assert no error occurred and match the expected value.
	assert.NoError(t, err)
	assert.Equal(t, 8, positionCurrent)
	assert.Equal(t, Block_Json_Others, block)
	assert.Equal(t, "Name", key)
}

// >>>>> >>>>> >>>>> >>>>> >>>>> old

func Test_Check_TracingData(t *testing.T) {
	var result TracingData
	_ = Unmarshal(jsonTracingLog, &result)
	fmt.Println(result)
}

func Test_Check_ByteArrayToValueString(t *testing.T) {
	// Define a slice of test cases
	tests := []struct {
		name      string
		jsonData  []byte
		valueType string
		value     string
		next      int
	}{
		{
			name:      "test case 0",
			valueType: "STRING",
			value:     "ValueB",
			next:      79,
		},
		{
			name:      "test case 1",
			valueType: "INT64",
			value:     "1",
			next:      59,
		},
		{
			name:      "test case 2",
			valueType: "STRING",
			value:     "more details",
			next:      73,
		},
		{
			name:      "test case 3",
			valueType: "STRING",
			value:     "*errors.errorString",
			next:      80,
		},
		{
			name:      "test case 4",
			valueType: "STRING",
			value:     "error",
			next:      66,
		},
		{
			name:      "test case 5",
			valueType: "STRING",
			value:     "unknown_service:___go_build_github_com_panhongrainbow_tracez_example_openTelemetry2file",
			next:      160,
		},
		{
			name:      "test case 6",
			valueType: "STRING",
			value:     "go",
			next:      75,
		},
		{
			name:      "test case 7",
			valueType: "STRING",
			value:     "opentelemetry",
			next:      86,
		},
		{
			name:      "test case 8",
			valueType: "STRING",
			value:     "1.14.0",
			next:      76,
		},
		{
			name:      "test case 9",
			valueType: "INT64",
			value:     "123456",
			next:      75,
		},
	}

	str := `"ValueString": {
        "Type": "STRING",
        "ValueString": "ValueB"
    }`
	tests[0].jsonData = []byte(str)

	str = `"ValueString": {
		"Type": "INT64",
		"ValueString": 1
    }`
	tests[1].jsonData = []byte(str)

	str = `"ValueString": {
		"Type": "STRING",
		"ValueString": "more details"
    }`
	tests[2].jsonData = []byte(str)

	str = `"ValueString": {
		"Type": "STRING",
		"ValueString": "*errors.errorString"
    }`
	tests[3].jsonData = []byte(str)

	str = `"ValueString": {
		"Type": "STRING",
		"ValueString": "error"
    }`
	tests[4].jsonData = []byte(str)

	str = `"ValueString": {
        "Type": "STRING",
        "ValueString": "unknown_service:___go_build_github_com_panhongrainbow_tracez_example_openTelemetry2file"
    }`
	tests[5].jsonData = []byte(str)

	str = `"ValueString": {
        "Type": "STRING",
        "ValueString": "go"
    }`
	tests[6].jsonData = []byte(str)

	str = `"ValueString": {
        "Type": "STRING",
        "ValueString": "opentelemetry"
    }`
	tests[7].jsonData = []byte(str)

	str = `"ValueString": {
        "Type": "STRING",
        "ValueString": "1.14.0"
	}`
	tests[8].jsonData = []byte(str)

	str = `"ValueString": {
        "Type": "INT64",
        "ValueString": "123456"
	}`
	tests[9].jsonData = []byte(str)

	for _, test := range tests {
		valueType, value, next := ByteArrayToValueString(0, &test.jsonData)
		assert.Equal(t, test.valueType, valueType, test.name+"'s valueType error !")
		assert.Equal(t, test.value, value, test.name+"'s value error !")
		assert.Equal(t, test.next, next, test.name+"'s next error !")
	}
}

// 24585 ns/op
func Benchmark_Estimate_TracingData(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var tracingData TracingData
		_ = json.Unmarshal(jsonTracingLog, &tracingData)
	}
}

// 3632 ns/op
func Benchmark_Estimate_TracingData2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var tracingData TracingData
		_ = Unmarshal(jsonTracingLog, &tracingData)

		AttributePool.Put(tracingData.Attributes[:0])
		AttributePool.Put(tracingData.Resource[:0])
		for j := 0; j < len(tracingData.Events); j++ {
			AttributePool.Put(tracingData.Events[j].Attributes[:0])
		}
	}
}

func Benchmark_Estimate_ByteArrayToValueString(b *testing.B) {

	/*str := `"ValueString": {
	    "Type": "STRING",
	    "ValueString": "unknown_service:___go_build_github_com_panhongrainbow_tracez_example_openTelemetry2file"
	}`*/
	str := `"ValueString": {
        "Type": "STRING",
        "ValueString": "go"
    }`
	jsonTracingLog = []byte(str)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = ByteArrayToValueString(0, &jsonTracingLog)
	}
}

func Benchmark_Estimate_ByteArrayToValueString2(b *testing.B) {

	str := `"ValueString": {
	    "Type": "STRING",
	    "ValueString": "unknown_service:___go_build_github_com_panhongrainbow_tracez_example_openTelemetry2file"
	}`
	/*str := `"ValueString": {
	    "Type": "STRING",
	    "ValueString": "go"
	}`*/
	jsonTracingLog = []byte(str)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(jsonTracingLog); i++ {
			//
		}
	}
}
