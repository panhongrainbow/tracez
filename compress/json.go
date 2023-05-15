package compress

import (
	"reflect"
	"strings"
	"unsafe"
)

// OneLine
func OneLine(jsonStr string) (oneLine string) {
	// Replace all tab characters with empty strings
	oneLine = strings.Replace(jsonStr, "\t", "", -1)
	// Replace all newline characters with empty strings
	oneLine = strings.Replace(oneLine, "\n", "", -1)
	return
}

// UnsafeOneLine
func UnsafeOneLine(jsonStr []byte, pointer unsafe.Pointer) {
	//
	jsonStrPtr := unsafe.Pointer(&jsonStr[0])

	//
	var preserve int
	var length = len(jsonStr)
	for i := 0; i < len(jsonStr); i++ {
		//
		jsonStr[preserve] = jsonStr[i]
		//
		character := *(*byte)(unsafe.Pointer(uintptr(jsonStrPtr) + uintptr(i)))
		//
		if character != byte('\t') && character != byte('\n') {
			//
			preserve++
		} else {
			//
			length--
		}
	}

	//
	header := (*reflect.SliceHeader)(pointer)
	header.Len = length

	//
	return
}

func Seperate(jsonStr string) (oneLine []string) {
	oneLine = strings.Split(jsonStr, "}{")
	if len(oneLine) >= 2 {
		oneLine[0] = oneLine[0] + "}"
		oneLine[len(oneLine)-1] = "{" + oneLine[len(oneLine)-1]
		for i := 1; i < len(oneLine)-1; i++ {
			oneLine[i] = "{" + oneLine[i] + "}"
		}
	}
	return
}

func UnsafeSeperate(jsonStr []byte) (breakPoints []int) {
	jsonStrPtr := unsafe.Pointer(&jsonStr[0])
	breakPoints = append(breakPoints, 0)

	for i := 0; i < len(jsonStr)-1; i++ {
		char1 := *(*byte)(unsafe.Pointer(uintptr(jsonStrPtr) + uintptr(i)))
		char2 := *(*byte)(unsafe.Pointer(uintptr(jsonStrPtr) + uintptr(i+1)))
		if char1 == byte('}') && char2 == byte('{') {
			breakPoints = append(breakPoints, i+1)
		}
	}

	breakPoints = append(breakPoints, len(jsonStr))

	return
}
