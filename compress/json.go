package compress

import (
	"reflect"
	"strings"
	"unsafe"
)

// OneLine replaces tabs and newlines, compacting multiline JSON string into a single line.
func OneLine(jsonStr string) (oneLine string) {
	// Replace all tab characters with empty strings
	oneLine = strings.Replace(jsonStr, "\t", "", -1)
	// Replace all newline characters with empty strings
	oneLine = strings.Replace(oneLine, "\n", "", -1)
	return
}

/*
UnsafeOneLine Compacts JSON string by copying non-newlines forward,
decrementing length for each tab/newline encountered.
*/
func UnsafeOneLine(jsonStr []byte, pointer unsafe.Pointer) {
	// Get pointer to first byte of jsonStr
	jsonStrPtr := unsafe.Pointer(&jsonStr[0])

	// Keep track of non-tab/newline characters
	var preserve int
	//Keep track of total length
	var length = len(jsonStr)
	// Iterate over bytes in jsonStr
	for i := 0; i < len(jsonStr); i++ {
		// Copy non tab/newline bytes into the front of the slice (不是\t和\n就往前移动)
		jsonStr[preserve] = jsonStr[i]
		// If not tab or newline, increment preserve (只要不是 \n 和 \t，大家复制后，指针一起往后移)
		character := *(*byte)(unsafe.Pointer(uintptr(jsonStrPtr) + uintptr(i)))
		// If not tab or newline, increment preserve (所有指针一起往后移)
		if character != byte('\t') && character != byte('\n') {
			preserve++
			// If tab or newline, decrement total length
		} else {
			//
			length--
		}
	}

	// Update slice header length
	header := (*reflect.SliceHeader)(pointer)
	header.Len = length

	// Return
	return
}

// Separate splits JSON on }{, bracketing remnants, producing properly JSON lines.
func Separate(jsonStr string) (oneLine []string) {
	// Split the JSON string into lines based on }{
	oneLine = strings.Split(jsonStr, "}{")
	// If there are 2 or more lines
	if len(oneLine) >= 2 {
		// Add a closing } to the first line
		oneLine[0] = oneLine[0] + "}"
		// Add a closing { to the last line
		oneLine[len(oneLine)-1] = "{" + oneLine[len(oneLine)-1]
		// For all lines in between, add "{" and "}"
		for i := 1; i < len(oneLine)-1; i++ {
			oneLine[i] = "{" + oneLine[i] + "}"
		}
	}
	// Return the separated lines
	return
}

// UnsafeSeparate scans JSON string bytes, appending indices where }{ found, yielding split points.
func UnsafeSeparate(jsonStr []byte) (breakPoints []int) {
	// If input is empty, return immediately
	if len(jsonStr) == 0 {
		return
	}

	// Get pointer to first byte
	jsonStrPtr := unsafe.Pointer(&jsonStr[0])
	// Initial break point is index 0
	breakPoints = append(breakPoints, 0)

	// Iterate over bytes in the string
	for i := 0; i < len(jsonStr)-1; i++ {
		// Get current byte and next byte
		char1 := *(*byte)(unsafe.Pointer(uintptr(jsonStrPtr) + uintptr(i)))
		char2 := *(*byte)(unsafe.Pointer(uintptr(jsonStrPtr) + uintptr(i+1)))
		// If current byte is } and next byte is {, append next index as break point
		if char1 == byte('}') && char2 == byte('{') {
			breakPoints = append(breakPoints, i+1)
		}
	}

	// Add final break point at end of string
	breakPoints = append(breakPoints, len(jsonStr))

	// Return break points
	return
}
