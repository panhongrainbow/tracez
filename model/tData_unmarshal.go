package model

import (
	"strconv"
	"time"
)

const (
	Enter_NotAnywhere = iota
	Enter_Attributes
	Enter_Events
	Enter_InstrumentationLibrary
	Enter_Parent
	Enter_Resource
	Enter_SpanContext
	Enter_Status
	// Enter_ValueString // Temporarily not in use
)

// UnmarshalByTuning is an automatically generated function that parses JSON data
// and populates the provided TracingData structure.
func UnmarshalByTuning(jsonTracingLog []byte, tData *TracingData) (err error) {
	// Set the position variable first.
	var firstLayerKey, secondLayerKey int
	var countCurlyBrace int
	var countBracket int

	// Retrieve an empty slice from the sync pool.
	tData.Attributes = attributePool.Get().([]Attribute)

	// Using DetectJsonString to extract the key from the JSON trace log.
	var positionNext, keyTail, keyLength int

	// Here, you don't need positionNext++ because 'detect' will perform the movement.
	for positionNext = 0; positionNext < len(jsonTracingLog); {

		positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)

		// Extract the key from the trace log.
		key := string(jsonTracingLog[(keyTail - keyLength):keyTail])
		// fmt.Println(key)

		switch key {
		case "Attributes":
			if firstLayerKey != Enter_Attributes &&
				firstLayerKey != Enter_Events &&
				firstLayerKey != Enter_Resource {
				// If not in the above three areas
				firstLayerKey = Enter_Attributes
			}
			if firstLayerKey == Enter_Events {
				secondLayerKey = Enter_Attributes
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "ChildSpanCount":
			if firstLayerKey == Enter_NotAnywhere {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.ChildSpanCount, err = strconv.Atoi(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
				if err != nil {
					panic("Parsing ChildSpanCount error: " + err.Error())
				}
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "Code":
			if firstLayerKey == Enter_Status {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.Status.Code = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "Description":
			if firstLayerKey == Enter_Status {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.Status.Description = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "DroppedAttributeCount":
			if firstLayerKey == Enter_Events {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Events) - 1
				tData.Events[length].DroppedAttributeCount, err = strconv.Atoi(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
				if err != nil {
					panic("Parsing DroppedAttributeCount error: " + err.Error())
				}
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "DroppedAttributes":
			if firstLayerKey == Enter_NotAnywhere {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.DroppedAttributes, err = strconv.Atoi(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
				if err != nil {
					panic("Parsing DroppedAttributes error: " + err.Error())
				}
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "DroppedEvents":
			if firstLayerKey == Enter_NotAnywhere {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.DroppedEvents, err = strconv.Atoi(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
				if err != nil {
					panic("Parsing DroppedEvents error: " + err.Error())
				}
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "DroppedLinks":
			if firstLayerKey == Enter_NotAnywhere {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.DroppedLinks, err = strconv.Atoi(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
				if err != nil {
					panic("Parsing DroppedLinks error: " + err.Error())
				}
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "EndTime":
			if firstLayerKey == Enter_NotAnywhere {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.EndTime, err = time.Parse(time.RFC3339Nano, string(jsonTracingLog[(keyTail-keyLength):keyTail]))
				if err != nil {
					panic("Parsing EndTime error: " + err.Error())
				}
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "Events":
			firstLayerKey = Enter_Events
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "InstrumentationLibrary":
			firstLayerKey = Enter_InstrumentationLibrary
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "Key":
			if firstLayerKey == Enter_Attributes {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Attributes) - 1
				tData.Attributes[length].Key = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_Events {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Events) - 1
				length2 := len(tData.Events[length].Attributes) - 1
				tData.Events[length].Attributes[length2].Key = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_Resource {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Resource) - 1
				tData.Resource[length].Key = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "Links":
			if firstLayerKey == Enter_NotAnywhere {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.Links = string(jsonTracingLog[(keyTail - keyLength):keyTail])
				if err != nil {
					panic("Parsing Links error: " + err.Error())
				}
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "Name":
			if firstLayerKey == Enter_Events {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Events) - 1
				tData.Events[length].Name = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_InstrumentationLibrary {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.InstrumentationLibrary.Name = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_NotAnywhere {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.Name = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "Parent":
			firstLayerKey = Enter_Parent
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "Remote":
			if firstLayerKey == Enter_Parent {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.Parent.Remote, err = strconv.ParseBool(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
				if err != nil {
					panic("Parsing Parent's Remote error: " + err.Error())
				}
			} else if firstLayerKey == Enter_SpanContext {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.SpanContext.Remote, err = strconv.ParseBool(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
				if err != nil {
					panic("Parsing SpanContext's Remote error: " + err.Error())
				}
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "Resource":
			firstLayerKey = Enter_Resource
			// retrieve an empty slice from the sync pool
			tData.Resource = attributePool.Get().([]Attribute)
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "SchemaURL":
			if firstLayerKey == Enter_InstrumentationLibrary {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.InstrumentationLibrary.SchemaURL = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "SpanContext":
			firstLayerKey = Enter_SpanContext
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "SpanID":
			if firstLayerKey == Enter_SpanContext {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.SpanContext.SpanID = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_Parent {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.Parent.SpanID = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "SpanKind":
			positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
			tData.SpanKind, err = strconv.Atoi(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
			if err != nil {
				panic("Parsing SpanKind error: " + err.Error())
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "StartTime":
			if firstLayerKey == Enter_NotAnywhere {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.StartTime, err = time.Parse(time.RFC3339Nano, string(jsonTracingLog[(keyTail-keyLength):keyTail]))
				if err != nil {
					panic("Parsing StartTime error: " + err.Error())
				}
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "Status":
			if firstLayerKey == Enter_NotAnywhere {
				firstLayerKey = Enter_Status
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "Time":
			if firstLayerKey == Enter_Events {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Events) - 1
				tData.Events[length].Time, err = time.Parse(time.RFC3339Nano, string(jsonTracingLog[(keyTail-keyLength):keyTail]))
				if err != nil {
					panic("Parsing StartTime error: " + err.Error())
				}
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "TraceFlags":
			if firstLayerKey == Enter_SpanContext {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.SpanContext.TraceFlags = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_Parent {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.Parent.TraceFlags = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "TraceID":
			if firstLayerKey == Enter_SpanContext {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.SpanContext.TraceID = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_Parent {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.Parent.TraceID = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "TraceState":
			if firstLayerKey == Enter_Parent {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.Parent.TraceState = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_SpanContext {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.SpanContext.TraceState = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "Type":
			if firstLayerKey == Enter_Attributes {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Attributes) - 1
				tData.Attributes[length].Value.Type = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_Events {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Events) - 1
				length2 := len(tData.Events[length].Attributes) - 1
				tData.Events[length].Attributes[length2].Value.Type = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_Resource {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Resource) - 1
				tData.Resource[length].Value.Type = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "ValueString":
			if firstLayerKey == Enter_Attributes &&
				secondLayerKey == Enter_NotAnywhere &&
				countCurlyBrace == 3 &&
				countBracket == 1 {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Attributes) - 1
				tData.Attributes[length].Value.Value = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_Events &&
				secondLayerKey == Enter_Attributes &&
				countCurlyBrace == 4 &&
				countBracket == 2 {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Events) - 1
				length2 := len(tData.Events[length].Attributes) - 1
				tData.Events[length].Attributes[length2].Value.Value = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			} else if firstLayerKey == Enter_Resource &&
				secondLayerKey == Enter_NotAnywhere &&
				countCurlyBrace == 3 &&
				countBracket == 1 {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				length := len(tData.Resource) - 1
				tData.Resource[length].Value.Value = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "Version":
			if firstLayerKey == Enter_InstrumentationLibrary {
				positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
				tData.InstrumentationLibrary.Version = string(jsonTracingLog[(keyTail - keyLength):keyTail])
			}
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "{":
			if firstLayerKey == Enter_Attributes &&
				secondLayerKey == Enter_NotAnywhere &&
				countCurlyBrace == 1 &&
				countBracket == 1 {
				tData.Attributes = append(tData.Attributes, Attribute{})
			}

			if firstLayerKey == Enter_Events &&
				secondLayerKey == Enter_Attributes &&
				countCurlyBrace == 2 &&
				countBracket == 2 {
				length := len(tData.Events) - 1
				tData.Events[length].Attributes = append(tData.Events[length].Attributes, Attribute{})
			}

			if firstLayerKey == Enter_Events &&
				secondLayerKey == Enter_NotAnywhere &&
				countCurlyBrace == 1 &&
				countBracket == 1 {
				tData.Events = append(tData.Events, Event{})
				length := len(tData.Events) - 1
				// retrieve an empty slice from the sync pool
				tData.Events[length].Attributes = attributePool.Get().([]Attribute)
			}

			if firstLayerKey == Enter_Resource &&
				secondLayerKey == Enter_NotAnywhere &&
				countCurlyBrace == 1 &&
				countBracket == 1 {
				tData.Resource = append(tData.Resource, Attribute{})
			}

			countCurlyBrace++
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "}":
			if (countBracket + countCurlyBrace) == 2 {
				firstLayerKey = Enter_NotAnywhere
			}

			countCurlyBrace--
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "[":
			countBracket++
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		case "]":
			if firstLayerKey == Enter_Events &&
				secondLayerKey == Enter_Attributes &&
				countCurlyBrace == 2 &&
				countBracket == 2 {
				secondLayerKey = Enter_NotAnywhere
			}
			if (countBracket + countCurlyBrace) == 2 {
				firstLayerKey = Enter_NotAnywhere
			}

			countBracket--
			/*fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)*/
		}
	}
	return
}

// CleanTracingData resets and returns all slices in tData to reduce memory usage.
func CleanTracingData(tData *TracingData) {
	// Set the Attributes slice of tData to an empty slice.
	tData.Attributes = tData.Attributes[:0]
	// Return the Attributes slice to the attributePool for reuse.
	attributePool.Put(tData.Attributes)

	// Loop through all the Events in tData.
	for i := 0; i < len(tData.Events); i++ {
		// Set the Attributes slice of each Event to an empty slice.
		tData.Events[i].Attributes = tData.Events[i].Attributes[:0]
		// Return the Attributes slice of each Event to the attributePool for reuse.
		attributePool.Put(tData.Events[i].Attributes)
	}

	// Set the Resource slice of tData to an empty slice.
	tData.Resource = tData.Resource[:0]
	// Return the Resource slice to the attributePool for reuse.
	attributePool.Put(tData.Resource)
}
