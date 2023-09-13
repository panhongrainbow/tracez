package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

// TestDetect_Check_JsonElement primarily focuses on verifying the extraction of elements in JSON.
func TestDetect_Check_JsonElement(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		// Test case 1: JSON object with a name field containing "John".
		{
			input:  `"name": "John"`,
			output: "name,John",
		},
		// Test case 2: JSON object with an age field containing the number 30.
		{
			input:  `"age": 30`,
			output: "age,30",
		},
		// Test case 3: JSON object with a city field containing "New York".
		{
			input:  `   {"city": "New York"}`,
			output: "city,New York",
		},
		// Test case 4: JSON object with a state field containing "California".
		{
			input:  `{"state": "California"}   `,
			output: "state,California,",
		},
		// Test case 5: JSON object with a status field containing true boolean value.
		{
			input:  `{"status": true}   `,
			output: "status,true,",
		},
		// Test case 6: JSON object with a status field containing false boolean value.
		{
			input:  `{"status": false}   `,
			output: "status,false,",
		},
		// Test case 6: JSON object with a status field containing the empty value.
		{
			input:  `{"value1":"","value2":"2"}   `,
			output: "value1,,value2,2,",
		},
		// Test case 7: Complex JSON object with nested fields and an array.
		{
			input: `{
  "name": "John Doe",
  "age": 30,
  "address": {
    "street": "123 Main St",
    "city": "New York",
    "zip": "10001"
  },
  "contacts": [
    {
      "type": "email",
      "value": "john@example.com"
    },
    {
      "type": "phone",
      "value": "+1 123-456-7890"
    }
  ]
}`,
			output: "name,John Doe,age,30,address,street,123 Main St,city,New York,zip,10001,contacts,[\n,type,email,value,john@example.com,type,phone,value,+1 123-456-7890,]\n",
		},
	}

	// Perform individual tests within the loop.
	for _, tt := range testCases {
		// Set initial variables.
		var positionNext, nonStringTail, nonStringLength int
		var keyValues []string
		// Loop through the input string to extract JSON elements.
		for ; positionNext < len(tt.input); positionNext++ {
			positionNext, nonStringTail, nonStringLength = DetectJsonElement(positionNext, []byte(tt.input))
			eachKeyValue := string([]byte(tt.input)[(nonStringTail - nonStringLength):nonStringTail])
			keyValues = append(keyValues, eachKeyValue)
		}

		// Check if the extracted keys match the expected output.
		require.Equal(t, tt.output, strings.Join(keyValues, ","))

		// Reset variables for the next test case.
		positionNext = 0
		nonStringTail = 0
		nonStringLength = 0
	}
}

// Test_Check_DetectJsonString performs testing for detecting JSON string values, comparing expected and actual results.
func Test_Check_DetectJsonString(t *testing.T) {
	// Initialize the starting position for testing.
	initPosition, _, _ := DetectJsonString(0, []byte(`{"key":`))

	tests := []struct {
		name            string
		jsonStr         []byte
		positionCurrent int
		expectedKey     string
		expectedNext    int
		expectedRest    string
	}{
		{
			name:            "detect string value in compact json string",
			jsonStr:         []byte(`{"key":"value","otherKey":"value"}`),
			positionCurrent: initPosition,
			expectedKey:     "value",
			expectedNext:    14,
			expectedRest:    ",\"otherKey\":\"value\"}",
		},
		{
			name:            "detect string value in json string with spaces",
			jsonStr:         []byte(`{"key"     :     "value"     ,"otherKey":"value"}`),
			positionCurrent: initPosition,
			expectedKey:     "value",
			expectedNext:    24,
			expectedRest:    "     ,\"otherKey\":\"value\"}",
		},
		{
			name:            "detect string value in json string with even spaces",
			jsonStr:         []byte(`{"key":          "value"          ,"otherKey": "value"}`),
			positionCurrent: initPosition,
			expectedKey:     "value",
			expectedNext:    24,
			expectedRest:    "          ,\"otherKey\": \"value\"}",
		},
		{
			name:            "detect string value with an emtpy string value",
			jsonStr:         []byte(`{"key":"","otherKey":"value"}`),
			positionCurrent: initPosition,
			expectedKey:     "",
			expectedNext:    9,
			expectedRest:    ",\"otherKey\":\"value\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function to detect JSON string and get the results.
			positionNext, keyTail, keyLength := DetectJsonString(tt.positionCurrent, tt.jsonStr)

			// Extract the actual key using calculated tail and length.
			actualKey := string(tt.jsonStr[(keyTail - keyLength):keyTail])

			// Retrieve the remaining content.
			remaining := string(tt.jsonStr[positionNext:])

			// Compare the actual and expected results for key and next position.
			assert.Equal(t, tt.expectedKey, actualKey, fmt.Sprintf("unexpected key: got %q, want %q", actualKey, tt.expectedKey))
			assert.Equal(t, tt.expectedNext, positionNext, fmt.Sprintf("unexpected next position: got %d, want %d", positionNext, tt.expectedNext))
			assert.Equal(t, tt.expectedRest, remaining, fmt.Sprintf("unexpected emaining content: got %s, want %s", remaining, tt.expectedRest))
		})
	}
}

// Test_Check_DetectJsonNonString performs testing for detecting JSON boolean values, comparing expected and actual results.
func Test_Check_DetectJsonNonString(t *testing.T) {
	// Initialize the starting position for testing.
	initPosition, _, _ := DetectJsonString(0, []byte(`{"key":`))

	tests := []struct {
		name            string
		jsonStr         []byte
		positionCurrent int
		expectedKey     string
		expectedNext    int
		expectedRest    string
	}{
		{
			name:            "detect true boolean value in compact json string",
			jsonStr:         []byte(`{"key":true,"otherKey":"value"}`),
			positionCurrent: initPosition,
			expectedKey:     "true",
			expectedNext:    11,
			expectedRest:    ",\"otherKey\":\"value\"}",
		},
		{
			name:            "detect true boolean value in abnormal json string",
			jsonStr:         []byte(`{"key"     :     true     ,"otherKey":"value"}`),
			positionCurrent: initPosition,
			expectedKey:     "true",
			expectedNext:    21,
			expectedRest:    "     ,\"otherKey\":\"value\"}",
		},
		{
			name:            "detect true boolean value in json string with even spaces",
			jsonStr:         []byte(`{"key":          true          ,"otherKey": "value"}`),
			positionCurrent: initPosition,
			expectedKey:     "true",
			expectedNext:    21,
			expectedRest:    "          ,\"otherKey\": \"value\"}",
		},
		{
			name:            "detect false boolean value in compact json string",
			jsonStr:         []byte(`{"key":false,"otherKey":"value"}`),
			positionCurrent: initPosition,
			expectedKey:     "false",
			expectedNext:    12,
			expectedRest:    ",\"otherKey\":\"value\"}",
		},
		{
			name:            "detect int value in compact json string",
			jsonStr:         []byte(`{"key":123,"otherKey":"value"}`),
			positionCurrent: initPosition,
			expectedKey:     "123",
			expectedNext:    10,
			expectedRest:    ",\"otherKey\":\"value\"}",
		},
		{
			name:            "detect float value in compact json string",
			jsonStr:         []byte(`{"key":12.3,"otherKey":"value"}`),
			positionCurrent: initPosition,
			expectedKey:     "12.3",
			expectedNext:    11,
			expectedRest:    ",\"otherKey\":\"value\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function to detect JSON boolean and get the results.
			positionNext, keyTail, keyLength := DetectJsonNonString(tt.positionCurrent, tt.jsonStr)

			// Extract the actual key using calculated tail and length.
			actualKey := string(tt.jsonStr[(keyTail - keyLength):keyTail])

			// Retrieve the remaining content.
			remaining := string(tt.jsonStr[positionNext:])

			// Compare the actual and expected results for key and next position.
			assert.Equal(t, tt.expectedKey, actualKey, fmt.Sprintf("unexpected key: got %q, want %q", actualKey, tt.expectedKey))
			assert.Equal(t, tt.expectedNext, positionNext, fmt.Sprintf("unexpected next position: got %d, want %d", positionNext, tt.expectedNext))
			assert.Equal(t, tt.expectedRest, remaining, fmt.Sprintf("unexpected emaining content: got %s, want %s", remaining, tt.expectedRest))
		})
	}
}
