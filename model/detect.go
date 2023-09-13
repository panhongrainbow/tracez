package model

// DetectJsonElement combines the DetectJsonNonString and DetectJsonString functions in DetectJsonElement.
//
//go:inline
func DetectJsonElement(positionCurrent int, jsonTracingLog []byte) (positionNext, elementTail, elementLength int) {

	// It is divided into two halves:
	// One where if the next element is considered a string, it is processed using the DetectJsonString function,
	// and if it is not a string, it is processed using the DetectJsonNonString function.

	// The following code integration will be performed.
	for ; positionCurrent < len(jsonTracingLog); positionCurrent++ {
		b := jsonTracingLog[positionCurrent]

		if b == '[' || b == ']' || b == '{' || b == '}' {
			return positionCurrent + 1, positionCurrent + 1, 1
		}

		if b == '"' {
			positionCurrent--
			if positionCurrent < 0 {
				positionCurrent = 0
			}
			return DetectJsonString(positionCurrent, jsonTracingLog)
		} else if b != ' ' && b != ',' && b != '{' && b != '}' && b != ':' && b != '\t' && b != '\n' {
			positionCurrent--
			if positionCurrent < 0 {
				positionCurrent = 0
			}
			return DetectJsonNonString(positionCurrent, jsonTracingLog)
		}
	}

	// If the reading is completed, return the length of the entire JSON data.
	return len(jsonTracingLog), 0, 0
}

// DetectJsonNonString scans the bool, int value in JSON trace log.
//
//go:inline
func DetectJsonNonString(positionCurrent int, jsonTracingLog []byte) (positionNext, nonStringTail, nonStringLength int) {
	// Initialize some variables.
	var inBool bool
	var outBool int

	// Iterate through the JSON trace log bytes.
	for ; positionCurrent < len(jsonTracingLog); positionCurrent++ {

		b := jsonTracingLog[positionCurrent]
		if b == ' ' || b == ',' || b == '}' || b == ':' {
			outBool++
			if outBool == 1 {
				inBool = !inBool
			}
			if nonStringLength > 0 && !inBool {
				positionNext = positionCurrent // Exclude adding 1 here, as I need to start reading from the comma.
				break
			}
		} else if inBool {
			nonStringLength++
			outBool = 0
		}
	}

	// Determine the key value by using keyTail and keyLength.
	positionNext = positionCurrent
	nonStringTail = positionCurrent

	return
}

// DetectJsonString scans the string value in JSON trace log.
//
//go:inline
func DetectJsonString(positionCurrent int, jsonTracingLog []byte) (positionNext, keyValueTail, keyValueLength int) {
	// Initialize some variables.
	var inQuotes bool

	// Iterate through the JSON trace log bytes.
	for ; positionCurrent < len(jsonTracingLog); positionCurrent++ {
		b := jsonTracingLog[positionCurrent]
		if b == '"' {
			inQuotes = !inQuotes
			// if keyValueLength > 0 && !inQuotes {
			if !inQuotes {
				positionNext = positionCurrent + 1 // Next time, start counting from the next byte.
				break
			}
		} else if inQuotes {
			keyValueLength++
		}
	}

	// Determine the key value by using keyTail and keyLength.
	keyValueTail = positionCurrent

	return
}
