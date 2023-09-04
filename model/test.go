package model

import (
	"fmt"
	"strconv"
)

const (
	Enter_NotAnywhere = iota
	Enter_FirstLayer_Attributes
	Enter_FirstLayer_Events
	Enter_FirstLayer_InstrumentationLibrary
	Enter_FirstLayer_Parent
	Enter_FirstLayer_Resource
	Enter_FirstLayer_SpanContext
	Enter_FirstLayer_Status
	Enter_SecondLayer_Attributes
	Enter_SecondLayer_ValueString
)

// UnmarshalByGen is an automatically generated function that parses JSON data
// and populates the provided TracingData structure.
// It cannot be modified. (Do Not Edit ⛔️)
func UnmarshalByGen(jsonTracingLog []byte, tData *TracingData) (err error) {
	var resourceCount, attributesCount, eventsCount int
	var firstLayerKey, secondLayerKey int
	// Using DetectJsonString to extract the key from the JSON trace log.
	var positionNext, keyTail, keyLength int

	for positionNext = 0; positionNext < len(jsonTracingLog); positionNext++ {

		positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)

		// If no quoted key was found, return an error.
		if keyLength == 0 {
			err = fmt.Errorf("no quoted key found")
			return
		}

		// Extract the key from the trace log.
		key := string(jsonTracingLog[(keyTail - keyLength):keyTail])
		// fmt.Println(key)

		// Determine the which block based on the key.
		switch key {
		case "Attributes":
			if firstLayerKey == Enter_FirstLayer_Events {
				tData.Events[len(tData.Events)-1].Attributes = append(tData.Events[len(tData.Events)-1].Attributes, Attribute{})
			} else {
				tData.Attributes = append(tData.Attributes, Attribute{})
			}
		case "ChildSpanCount":
			positionNext, keyTail, keyLength = DetectJsonNonString(positionNext, jsonTracingLog)
			tData.ChildSpanCount, _ = strconv.Atoi(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
			firstLayerKey = Enter_NotAnywhere
		case "Code":
			positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
			tData.Status.Code = string(jsonTracingLog[(keyTail - keyLength):keyTail])
		case "Description":
			positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
			tData.Status.Description = string(jsonTracingLog[(keyTail - keyLength):keyTail])
		case "DroppedAttributeCount":
			if firstLayerKey == Enter_FirstLayer_Events {
				positionNext, keyTail, keyLength = DetectJsonNonString(positionNext, jsonTracingLog)
				tData.Events[len(tData.Events)-1].DroppedAttributeCount, _ = strconv.Atoi(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
			}
		case "DroppedAttributes":
			positionNext, keyTail, keyLength = DetectJsonNonString(positionNext, jsonTracingLog)
			tData.DroppedAttributes, _ = strconv.Atoi(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
			firstLayerKey = Enter_NotAnywhere
		case "DroppedEvents":
			positionNext, keyTail, keyLength = DetectJsonNonString(positionNext, jsonTracingLog)
			tData.DroppedEvents, _ = strconv.Atoi(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
			firstLayerKey = Enter_NotAnywhere
		case "DroppedLinks":
			positionNext, keyTail, keyLength = DetectJsonNonString(positionNext, jsonTracingLog)
			tData.DroppedLinks, _ = strconv.Atoi(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
			firstLayerKey = Enter_NotAnywhere
		case "EndTime":
			positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
			// tData.EndTime, err = time.Parse(time.RFC3339Nano, string(jsonTracingLog[(keyTail-keyLength):keyTail]))
			firstLayerKey = Enter_NotAnywhere
		case "Events":
			tData.Events = append(tData.Events, Event{})
			firstLayerKey = Enter_FirstLayer_Events
		case "InstrumentationLibrary":
			firstLayerKey = Enter_FirstLayer_InstrumentationLibrary
		case "Key":
			if firstLayerKey == Enter_FirstLayer_Resource {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.Resource[len(tData.Resource)-1].Key = string(jsonTracingLog[(keyTail - keyLength):keyTail])
				resourceCount++
				if resourceCount%2 == 0 {
					tData.Resource = append(tData.Resource, Attribute{})
				}
			}
			if firstLayerKey == Enter_FirstLayer_Attributes {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.Attributes[len(tData.Attributes)-1].Key = string(jsonTracingLog[(keyTail - keyLength):keyTail])
				attributesCount++
				if attributesCount%2 == 0 {
					tData.Attributes = append(tData.Attributes, Attribute{})
				}
			}
			if firstLayerKey == Enter_FirstLayer_Events {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.Events[len(tData.Events)-1].Attributes[len(tData.Events[len(tData.Events)-1].Attributes)-1].Key = string(jsonTracingLog[(keyTail - keyLength):keyTail])
				eventsCount++
				if eventsCount%3 == 0 {
					tData.Events[len(tData.Events)-1].Attributes = append(tData.Events[len(tData.Events)-1].Attributes, Attribute{})
					secondLayerKey = Enter_NotAnywhere
				}
			}
		case "Links":
			positionNext, keyTail, keyLength = DetectJsonNonString(positionNext, jsonTracingLog)
			tData.Links = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			firstLayerKey = Enter_NotAnywhere
		case "Name":
			if firstLayerKey == Enter_FirstLayer_Events {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.Events[len(tData.Events)-1].Name = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_FirstLayer_InstrumentationLibrary {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.InstrumentationLibrary.Name = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_NotAnywhere {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.Name = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
		case "Parent":
			firstLayerKey = Enter_FirstLayer_Parent
		case "Remote":
			if firstLayerKey == Enter_FirstLayer_Parent {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.Parent.Remote, _ = strconv.ParseBool(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
				firstLayerKey = Enter_FirstLayer_Parent
			} else if firstLayerKey == Enter_FirstLayer_SpanContext {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.SpanContext.Remote, _ = strconv.ParseBool(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
				firstLayerKey = Enter_FirstLayer_SpanContext
			}
		case "Resource":
			tData.Resource = append(tData.Resource, Attribute{})
			firstLayerKey = Enter_FirstLayer_Resource
		case "SchemaURL":
			if firstLayerKey == Enter_FirstLayer_InstrumentationLibrary {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.InstrumentationLibrary.SchemaURL = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
		case "SpanContext":
			firstLayerKey = Enter_FirstLayer_SpanContext
		case "SpanID":
			if firstLayerKey == Enter_FirstLayer_SpanContext {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.SpanContext.SpanID = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_FirstLayer_Parent {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.Parent.SpanID = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
		case "SpanKind":
			positionNext, keyTail, keyLength = DetectJsonNonString(positionNext, jsonTracingLog)
			tData.SpanKind, _ = strconv.Atoi(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
			firstLayerKey = Enter_NotAnywhere
		case "StartTime":
			positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
			// tData.StartTime, err = time.Parse(time.RFC3339Nano, string(jsonTracingLog[(keyTail-keyLength):keyTail]))
			firstLayerKey = Enter_NotAnywhere
		case "Status":
			firstLayerKey = Enter_FirstLayer_Status
		case "Time":
			if firstLayerKey == Enter_FirstLayer_Events {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				// tData.Events[len(tData.Events)-1].Time, _ = time.Parse(time.RFC3339Nano, string(jsonTracingLog[(keyTail-keyLength):keyTail]))
			}
		case "TraceFlags":
			if firstLayerKey == Enter_FirstLayer_SpanContext {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.SpanContext.TraceFlags = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			if firstLayerKey == Enter_FirstLayer_Parent {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.Parent.TraceFlags = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
		case "TraceID":
			if firstLayerKey == Enter_FirstLayer_SpanContext {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.SpanContext.TraceID = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_FirstLayer_Parent {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.Parent.TraceID = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
		case "TraceState":
			if firstLayerKey == Enter_FirstLayer_Parent {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.Parent.TraceState = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_FirstLayer_SpanContext {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.SpanContext.TraceState = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
		case "Type":
			if firstLayerKey == Enter_FirstLayer_Attributes && secondLayerKey == Enter_SecondLayer_ValueString {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.Attributes[len(tData.Attributes)-1].Value.Type = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_FirstLayer_Resource && secondLayerKey == Enter_SecondLayer_ValueString {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.Resource[len(tData.Resource)-1].Value.Type = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_FirstLayer_Events && secondLayerKey == Enter_SecondLayer_ValueString {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.Events[len(tData.Events)-1].Attributes[len(tData.Events[len(tData.Events)-1].Attributes)-1].Value.Type = string(jsonTracingLog[(keyTail - keyLength):keyTail])
				eventsCount++
				if eventsCount%3 == 0 {
					tData.Events[len(tData.Events)-1].Attributes = append(tData.Events[len(tData.Events)-1].Attributes, Attribute{})
					secondLayerKey = Enter_NotAnywhere
				}
			}
		case "ValueString":
			if firstLayerKey == Enter_FirstLayer_Attributes {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.Attributes[len(tData.Attributes)-1].Value.Value = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_FirstLayer_Events && secondLayerKey == Enter_SecondLayer_ValueString {
				if tData.Events[len(tData.Events)-1].Attributes[len(tData.Events[len(tData.Events)-1].Attributes)-1].Value.Type == "STRING" {
					positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
					tData.Events[len(tData.Events)-1].Attributes[len(tData.Events[len(tData.Events)-1].Attributes)-1].Value.Value = string(jsonTracingLog[(keyTail - keyLength):keyTail])
				}
				if tData.Events[len(tData.Events)-1].Attributes[len(tData.Events[len(tData.Events)-1].Attributes)-1].Value.Type == "INT64" {
					positionNext, keyTail, keyLength = DetectJsonNonString(positionNext, jsonTracingLog)
					tData.Events[len(tData.Events)-1].Attributes[len(tData.Events[len(tData.Events)-1].Attributes)-1].Value.Value = string(jsonTracingLog[(keyTail - keyLength):keyTail])
				}
				eventsCount++
				if eventsCount%3 == 0 {
					tData.Events[len(tData.Events)-1].Attributes = append(tData.Events[len(tData.Events)-1].Attributes, Attribute{})
					secondLayerKey = Enter_NotAnywhere
				}
			} else if firstLayerKey == Enter_FirstLayer_Events && secondLayerKey != Enter_SecondLayer_ValueString {
				secondLayerKey = Enter_SecondLayer_ValueString
			} else if firstLayerKey == Enter_FirstLayer_Resource && secondLayerKey == Enter_SecondLayer_ValueString {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.Resource[len(tData.Resource)-1].Value.Value = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_FirstLayer_Resource && secondLayerKey != Enter_SecondLayer_ValueString {
				secondLayerKey = Enter_SecondLayer_ValueString
			} else if firstLayerKey == Enter_FirstLayer_Attributes && secondLayerKey == Enter_SecondLayer_ValueString {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.Attributes[len(tData.Attributes)-1].Value.Value = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
		case "Version":
			if firstLayerKey == Enter_FirstLayer_InstrumentationLibrary {
				positionNext, keyTail, keyLength = DetectJsonString(positionNext, jsonTracingLog)
				tData.InstrumentationLibrary.Version = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
		}
	}
	return
}
