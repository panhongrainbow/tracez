package file

import (
	"strings"
	"testing"
)

var jsonData = `{
	"Name": "functionBSpan",
	"SpanContext": {
		"TraceID": "673a3e3c000be1d70868e4addee05720",
		"SpanID": "3868c5686e347e6f",
		"TraceFlags": "01",
		"TraceState": "",
		"Remote": false
	},
	"Parent": {
		"TraceID": "673a3e3c000be1d70868e4addee05720",
		"SpanID": "ab06b5370da36ef5",
		"TraceFlags": "01",
		"TraceState": "",
		"Remote": false
	},
	"SpanKind": 1,
	"StartTime": "2023-04-18T20:49:10.595897864+08:00",
	"EndTime": "2023-04-18T20:49:11.596005054+08:00",
	"Attributes": null,
	"Events": null,
	"Links": null,
	"Status": {
		"Code": "Unset",
		"Description": ""
	},
	"DroppedAttributes": 0,
	"DroppedEvents": 0,
	"DroppedLinks": 0,
	"ChildSpanCount": 0,
	"Resource": [
		{
			"Key": "service.name",
			"Value": {
				"Type": "STRING",
				"Value": "unknown_service:openTelemetry2file"
			}
		},
		{
			"Key": "telemetry.sdk.language",
			"Value": {
				"Type": "STRING",
				"Value": "go"
			}
		},
		{
			"Key": "telemetry.sdk.name",
			"Value": {
				"Type": "STRING",
				"Value": "opentelemetry"
			}
		},
		{
			"Key": "telemetry.sdk.version",
			"Value": {
				"Type": "STRING",
				"Value": "1.14.0"
			}
		}
	],
	"InstrumentationLibrary": {
		"Name": "functionBTracer",
		"Version": "",
		"SchemaURL": ""
	}
}
{
	"Name": "functionASpan",
	"SpanContext": {
		"TraceID": "673a3e3c000be1d70868e4addee05720",
		"SpanID": "ab06b5370da36ef5",
		"TraceFlags": "01",
		"TraceState": "",
		"Remote": false
	},
	"Parent": {
		"TraceID": "673a3e3c000be1d70868e4addee05720",
		"SpanID": "a88ecbd7a8632a6f",
		"TraceFlags": "01",
		"TraceState": "",
		"Remote": false
	},
	"SpanKind": 1,
	"StartTime": "2023-04-18T20:49:10.595895339+08:00",
	"EndTime": "2023-04-18T20:49:11.596116357+08:00",
	"Attributes": null,
	"Events": null,
	"Links": null,
	"Status": {
		"Code": "Unset",
		"Description": ""
	},
	"DroppedAttributes": 0,
	"DroppedEvents": 0,
	"DroppedLinks": 0,
	"ChildSpanCount": 1,
	"Resource": [
		{
			"Key": "service.name",
			"Value": {
				"Type": "STRING",
				"Value": "unknown_service:openTelemetry2file"
			}
		},
		{
			"Key": "telemetry.sdk.language",
			"Value": {
				"Type": "STRING",
				"Value": "go"
			}
		},
		{
			"Key": "telemetry.sdk.name",
			"Value": {
				"Type": "STRING",
				"Value": "opentelemetry"
			}
		},
		{
			"Key": "telemetry.sdk.version",
			"Value": {
				"Type": "STRING",
				"Value": "1.14.0"
			}
		}
	],
	"InstrumentationLibrary": {
		"Name": "functionATracer",
		"Version": "",
		"SchemaURL": ""
	}
}
{
	"Name": "mainSpan",
	"SpanContext": {
		"TraceID": "673a3e3c000be1d70868e4addee05720",
		"SpanID": "a88ecbd7a8632a6f",
		"TraceFlags": "01",
		"TraceState": "",
		"Remote": false
	},
	"Parent": {
		"TraceID": "00000000000000000000000000000000",
		"SpanID": "0000000000000000",
		"TraceFlags": "00",
		"TraceState": "",
		"Remote": false
	},
	"SpanKind": 1,
	"StartTime": "2023-04-18T20:49:10.595883748+08:00",
	"EndTime": "2023-04-18T20:49:11.596161752+08:00",
	"Attributes": null,
	"Events": [
		{
			"Name": "mainEvent",
			"Attributes": null,
			"DroppedAttributeCount": 0,
			"Time": "2023-04-18T20:49:11.596156649+08:00"
		}
	],
	"Links": null,
	"Status": {
		"Code": "Unset",
		"Description": ""
	},
	"DroppedAttributes": 0,
	"DroppedEvents": 0,
	"DroppedLinks": 0,
	"ChildSpanCount": 1,
	"Resource": [
		{
			"Key": "service.name",
			"Value": {
				"Type": "STRING",
				"Value": "unknown_service:openTelemetry2file"
			}
		},
		{
			"Key": "telemetry.sdk.language",
			"Value": {
				"Type": "STRING",
				"Value": "go"
			}
		},
		{
			"Key": "telemetry.sdk.name",
			"Value": {
				"Type": "STRING",
				"Value": "opentelemetry"
			}
		},
		{
			"Key": "telemetry.sdk.version",
			"Value": {
				"Type": "STRING",
				"Value": "1.14.0"
			}
		}
	],
	"InstrumentationLibrary": {
		"Name": "mainTracer",
		"Version": "",
		"SchemaURL": ""
	}
}`

func compressToThreeLines(jsonStr string) (jsonStrs []string) {
	jsonStr = strings.Replace(jsonStr, "\t", "", -1)
	jsonStr = strings.Replace(jsonStr, "\n", "", -1)
	// jsonStrs = append(jsonStrs, jsonStr)

	jsonStrs = strings.Split(jsonStr, "}{")

	jsonStrs[0] = jsonStrs[0] + "}"

	for i := 1; i < len(jsonStrs)-1; i++ {
		//
		jsonStrs[i] = "{" + jsonStrs[i] + "}"
	}

	jsonStrs[len(jsonStrs)-1] = "{" + jsonStrs[len(jsonStrs)-1]

	return
}

func compressToThreeLines1(jsonStr string) (jsonStrs []string) {
	jsonStr = strings.Replace(jsonStr, "\t", "", -1)
	jsonStr = strings.Replace(jsonStr, "\n", "", -1)
	// jsonStrs = append(jsonStrs, jsonStr)

	jsonStrs = strings.Split(jsonStr, "}{")

	jsonStrs[0] = jsonStrs[0] + "}"

	for i := 1; i < len(jsonStrs)-1; i++ {
		//
		jsonStrs[i] = "{" + jsonStrs[i] + "}"
	}

	jsonStrs[len(jsonStrs)-1] = "{" + jsonStrs[len(jsonStrs)-1]

	return
}

func Test_Race_fileClient(t *testing.T) {
	// test := compressToThreeLines1(jsonData)
	// fmt.Println(test)
	Search()
}

func Benchmark_Test(b *testing.B) {
	//
}
