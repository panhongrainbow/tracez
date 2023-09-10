package model

import (
	"fmt"
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
		fmt.Println(key)

		if key == "SchemaURL" {
			fmt.Println()
		}

		switch key {
		case "Attributes":
			if firstLayerKey != Enter_FirstLayer_Attributes && firstLayerKey != Enter_FirstLayer_Events {
				firstLayerKey = Enter_FirstLayer_Attributes
			}
			if firstLayerKey == Enter_FirstLayer_Events {
				secondLayerKey = Enter_SecondLayer_Attributes
			}
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "ChildSpanCount":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Code":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Description":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "DroppedAttributeCount":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "DroppedAttributes":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "DroppedEvents":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "DroppedLinks":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "EndTime":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Events":
			firstLayerKey = Enter_FirstLayer_Events
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "InstrumentationLibrary":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Key":
			fmt.Println("    位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
			if firstLayerKey == Enter_FirstLayer_Attributes {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Attributes) - 1
				tData.Attributes[length].Key = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			if firstLayerKey == Enter_FirstLayer_Events {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Events) - 1
				length2 := len(tData.Events[length].Attributes) - 1
				tData.Events[length].Attributes[length2].Key = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
		case "Links":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Name":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Parent":
			firstLayerKey = Enter_FirstLayer_Parent
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Remote":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Resource":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "SchemaURL":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "SpanContext":
			firstLayerKey = Enter_FirstLayer_SpanContext
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "SpanID":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "SpanKind":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "StartTime":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Status":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Time":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "TraceFlags":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "TraceID":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "TraceState":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "Type":
			fmt.Println("     位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
			if firstLayerKey == Enter_FirstLayer_Attributes {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Attributes) - 1
				tData.Attributes[length].Value.Type = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			if firstLayerKey == Enter_FirstLayer_Events {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Events) - 1
				length2 := len(tData.Events[length].Attributes) - 1
				tData.Events[length].Attributes[length2].Value.Type = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
		case "ValueString":
			fmt.Println("    位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
			if firstLayerKey == Enter_FirstLayer_Attributes &&
				secondLayerKey == Enter_NotAnywhere &&
				countcurlyBrace == 3 &&
				countBracket == 1 {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Attributes) - 1
				tData.Attributes[length].Value.Value = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			if firstLayerKey == Enter_FirstLayer_Events &&
				secondLayerKey == Enter_SecondLayer_Attributes &&
				countcurlyBrace == 4 &&
				countBracket == 2 {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Events) - 1
				length2 := len(tData.Events[length].Attributes) - 1
				tData.Events[length].Attributes[length2].Value.Value = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
		case "Version":
			fmt.Println("    位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
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

			countcurlyBrace++
			fmt.Println("    位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "}":
			countcurlyBrace--
			if (countBracket + countcurlyBrace) == 1 {
				firstLayerKey = NotInAnyBlock
			}
			fmt.Println("    位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		case "[":
			countBracket++
			fmt.Println("    位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
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
			fmt.Println("    位置:", firstLayerKey, secondLayerKey, countcurlyBrace, countBracket)
		}
	}
	return
}
