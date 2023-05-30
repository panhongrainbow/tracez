package model

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
	"unsafe"
)

// nonOmitemptySample is the test sample struct without omitempty.
type nonOmitemptySample struct {
	Key     string      `json:"key"`
	Value   interface{} `json:"value"`
	Value1  string      `json:"value1"`
	Value2  string      `json:"value2"`
	Value3  string      `json:"value3"`
	Value4  string      `json:"value4"`
	Value5  string      `json:"value5"`
	Value6  string      `json:"value6"`
	Value7  string      `json:"value7"`
	Value8  string      `json:"value8"`
	Value9  string      `json:"value9"`
	Value10 string      `json:"value10"`
}

// nonOmitemptySample is the test sample struct with omitempty.
type omitemptySample struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Value1  string `json:"value1,omitempty"`
	Value2  string `json:"value2,omitempty"`
	Value3  string `json:"value3,omitempty"`
	Value4  string `json:"value4,omitempty"`
	Value5  string `json:"value5,omitempty"`
	Value6  string `json:"value6,omitempty"`
	Value7  string `json:"value7,omitempty"`
	Value8  string `json:"value8,omitempty"`
	Value9  string `json:"value9,omitempty"`
	Value10 string `json:"value10,omitempty"`
}

// jsonStr is the test json sample string.
var jsonStr = `{"key":"key","value":"value"}`

// Benchmark_Estimate_nonOmitemptySample performs benchmark testing using nonOmitemptySample data.
func Benchmark_Estimate_nonOmitemptySample(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sample nonOmitemptySample
		var err error
		err = json.Unmarshal(stringToBytes(jsonStr), &sample)
		if err != nil {
			log.Fatal(err)
		}

		_, err = json.Marshal(sample)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Benchmark_Estimate_omitemptySample performs benchmark testing using omitemptySample data.
func Benchmark_Estimate_omitemptySample(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sample omitemptySample
		var err error
		err = json.Unmarshal(stringToBytes(jsonStr), &sample)
		if err != nil {
			log.Fatal(err)
		}

		/*var jsonData []byte
		jsonData, err = json.Marshal(sample)*/
		_, err = json.Marshal(sample)
		if err != nil {
			log.Fatal(err)
		}

		/*if *(*string)(unsafe.Pointer(&jsonData)) != jsonStr {
			log.Fatal("not equal")
		}*/
	}
}

// stringToBytes function converts string to []bytes with no copying.
func stringToBytes(s string) []byte {
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sliceHeader := reflect.SliceHeader{
		Data: strHeader.Data,
		Len:  strHeader.Len,
		Cap:  strHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}
