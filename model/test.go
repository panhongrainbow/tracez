package model

import (
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
	// Enter_SecondLayer_ValueString // Temporarily not in use
)

// UnmarshalByGen is an automatically generated function that parses JSON data
// and populates the provided TracingData structure.
// It cannot be modified. (Do Not Edit ⛔️)
func UnmarshalByGen(jsonTracingLog []byte, tData *TracingData) (err error) {
	// var resourceCount, attributesCount, eventsCount int
	var firstLayerKey, secondLayerKey int
	var countcurlyBrace int
	var countBracket int
	/*var countAttributes int
	inFirstLayer := true*/

	// Using DetectJsonString to extract the key from the JSON trace log.
	var positionNext, keyTail, keyLength int

	for positionNext = 0; positionNext < len(jsonTracingLog); { // 這裡不用 positionNext++

		positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)

		// Extract the key from the trace log.
		key := string(jsonTracingLog[(keyTail - keyLength):keyTail])
		// fmt.Println(key)

		switch key {
		case "Attributes":
			if firstLayerKey != Enter_FirstLayer_Attributes && firstLayerKey != Enter_FirstLayer_Events {
				firstLayerKey = Enter_FirstLayer_Attributes
			}
			if firstLayerKey == Enter_FirstLayer_Events {
				secondLayerKey = Enter_SecondLayer_Attributes
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "ChildSpanCount":
			if firstLayerKey == Enter_NotAnywhere {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.ChildSpanCount, err = strconv.Atoi(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
				if err != nil {
					panic("Parsing ChildSpanCount error: " + err.Error())
				}
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Code":
			if firstLayerKey == Enter_FirstLayer_Status {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.Status.Code = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Description":
			if firstLayerKey == Enter_FirstLayer_Status {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.Status.Description = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "DroppedAttributeCount":
			if firstLayerKey == Enter_FirstLayer_Events {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Events) - 1
				tData.Events[length].DroppedAttributeCount, err = strconv.Atoi(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
				if err != nil {
					panic("Parsing DroppedAttributeCount error: " + err.Error())
				}
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "DroppedAttributes":
			if firstLayerKey == Enter_NotAnywhere {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.DroppedAttributes, err = strconv.Atoi(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
				if err != nil {
					panic("Parsing DroppedAttributes error: " + err.Error())
				}
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "DroppedEvents":
			if firstLayerKey == Enter_NotAnywhere {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.DroppedEvents, err = strconv.Atoi(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
				if err != nil {
					panic("Parsing DroppedEvents error: " + err.Error())
				}
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "DroppedLinks":
			if firstLayerKey == Enter_NotAnywhere {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.DroppedLinks, err = strconv.Atoi(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
				if err != nil {
					panic("Parsing DroppedLinks error: " + err.Error())
				}
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "EndTime":
			if firstLayerKey == Enter_NotAnywhere {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				// tData.EndTime, err = time.Parse(time.RFC3339Nano, string(jsonTracingLog[(keyTail-keyLength):keyTail]))
				if err != nil {
					panic("Parsing EndTime error: " + err.Error())
				}
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Events":
			firstLayerKey = Enter_FirstLayer_Events
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "InstrumentationLibrary":
			firstLayerKey = Enter_FirstLayer_InstrumentationLibrary
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Key":
			if firstLayerKey == Enter_FirstLayer_Attributes {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Attributes) - 1
				tData.Attributes[length].Key = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_FirstLayer_Events {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Events) - 1
				length2 := len(tData.Events[length].Attributes) - 1
				tData.Events[length].Attributes[length2].Key = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_FirstLayer_Resource {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Resource) - 1
				tData.Resource[length].Key = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			// fmt.Println("    位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Links":
			if firstLayerKey == Enter_NotAnywhere {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.Links = string(jsonTracingLog[(keyTail - keyLength):keyTail])
				if err != nil {
					panic("Parsing Links error: " + err.Error())
				}
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Name":
			if firstLayerKey == Enter_FirstLayer_Events {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Events) - 1
				tData.Events[length].Name = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_FirstLayer_InstrumentationLibrary {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.InstrumentationLibrary.Name = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_NotAnywhere {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.Name = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Parent":
			firstLayerKey = Enter_FirstLayer_Parent
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Remote":
			if firstLayerKey == Enter_FirstLayer_Parent {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.Parent.Remote, err = strconv.ParseBool(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
				if err != nil {
					panic("Parsing Parent's Remote error: " + err.Error())
				}
			} else if firstLayerKey == Enter_FirstLayer_SpanContext {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.SpanContext.Remote, err = strconv.ParseBool(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
				if err != nil {
					panic("Parsing SpanContext's Remote error: " + err.Error())
				}
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Resource":
			firstLayerKey = Enter_FirstLayer_Resource
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "SchemaURL":
			if firstLayerKey == Enter_FirstLayer_InstrumentationLibrary {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.InstrumentationLibrary.SchemaURL = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "SpanContext":
			firstLayerKey = Enter_FirstLayer_SpanContext
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "SpanID":
			if firstLayerKey == Enter_FirstLayer_SpanContext {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.SpanContext.SpanID = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_FirstLayer_Parent {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.Parent.SpanID = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "SpanKind":
			positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
			tData.SpanKind, err = strconv.Atoi(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
			if err != nil {
				panic("Parsing SpanKind error: " + err.Error())
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "StartTime":
			if firstLayerKey == Enter_NotAnywhere {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				// tData.StartTime, err = time.Parse(time.RFC3339Nano, string(jsonTracingLog[(keyTail-keyLength):keyTail]))
				if err != nil {
					panic("Parsing StartTime error: " + err.Error())
				}
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Status":
			if firstLayerKey == Enter_NotAnywhere {
				firstLayerKey = Enter_FirstLayer_Status
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Time":
			if firstLayerKey == Enter_FirstLayer_Events {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				// length := len(tData.Events) - 1
				// tData.Events[length].Time, _ = time.Parse(time.RFC3339Nano, string(jsonTracingLog[(keyTail-keyLength):keyTail]))
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "TraceFlags":
			if firstLayerKey == Enter_FirstLayer_SpanContext {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.SpanContext.TraceFlags = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_FirstLayer_Parent {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.Parent.TraceFlags = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "TraceID":
			if firstLayerKey == Enter_FirstLayer_SpanContext {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.SpanContext.TraceID = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_FirstLayer_Parent {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.Parent.TraceID = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "TraceState":
			if firstLayerKey == Enter_FirstLayer_Parent {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.Parent.TraceState = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_FirstLayer_SpanContext {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.SpanContext.TraceState = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Type":
			if firstLayerKey == Enter_FirstLayer_Attributes {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Attributes) - 1
				tData.Attributes[length].Value.Type = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_FirstLayer_Events {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Events) - 1
				length2 := len(tData.Events[length].Attributes) - 1
				tData.Events[length].Attributes[length2].Value.Type = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_FirstLayer_Resource {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Resource) - 1
				tData.Resource[length].Value.Type = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			// fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "ValueString":
			if firstLayerKey == Enter_FirstLayer_Attributes &&
				secondLayerKey == Enter_NotAnywhere &&
				countcurlyBrace == 3 &&
				countBracket == 1 {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Attributes) - 1
				tData.Attributes[length].Value.Value = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_FirstLayer_Events &&
				secondLayerKey == Enter_SecondLayer_Attributes &&
				countcurlyBrace == 4 &&
				countBracket == 2 {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Events) - 1
				length2 := len(tData.Events[length].Attributes) - 1
				tData.Events[length].Attributes[length2].Value.Value = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_FirstLayer_Resource &&
				secondLayerKey == Enter_NotAnywhere &&
				countcurlyBrace == 3 &&
				countBracket == 1 {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Resource) - 1
				tData.Resource[length].Value.Value = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			// fmt.Println("    位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Version":
			if firstLayerKey == Enter_FirstLayer_InstrumentationLibrary {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.InstrumentationLibrary.Version = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			// fmt.Println("    位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "{":
			if firstLayerKey == Enter_FirstLayer_Attributes &&
				secondLayerKey == Enter_NotAnywhere &&
				countcurlyBrace == 1 &&
				countBracket == 1 {
				tData.Attributes = append(tData.Attributes, Attribute{})
			}

			if firstLayerKey == Enter_FirstLayer_Events &&
				secondLayerKey == Enter_SecondLayer_Attributes &&
				countcurlyBrace == 2 &&
				countBracket == 2 {
				length := len(tData.Events) - 1
				tData.Events[length].Attributes = append(tData.Events[length].Attributes, Attribute{})
			}

			if firstLayerKey == Enter_FirstLayer_Events &&
				secondLayerKey == Enter_NotAnywhere &&
				countcurlyBrace == 1 &&
				countBracket == 1 {
				tData.Events = append(tData.Events, Event{})
			}

			if firstLayerKey == Enter_FirstLayer_Resource &&
				secondLayerKey == Enter_NotAnywhere &&
				countcurlyBrace == 1 &&
				countBracket == 1 {
				tData.Resource = append(tData.Resource, Attribute{})
			}

			countcurlyBrace++
			// fmt.Println("    位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "}":
			countcurlyBrace--
			if (countBracket + countcurlyBrace) == 1 {
				firstLayerKey = NotInAnyBlock
			}
			// fmt.Println("    位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "[":
			countBracket++
			// fmt.Println("    位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "]":
			if firstLayerKey == Enter_FirstLayer_Events &&
				secondLayerKey == Enter_SecondLayer_Attributes &&
				countcurlyBrace == 2 &&
				countBracket == 2 {
				secondLayerKey = Enter_NotAnywhere
			}
			countBracket--
			if (countBracket + countcurlyBrace) == 1 {
				firstLayerKey = NotInAnyBlock
			}
			// fmt.Println("    位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		}
	}
	return
}
