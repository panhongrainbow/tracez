package model

import (
	"encoding/json"
	"log"
	"reflect"
	"sync"
	"testing"
	"unsafe"
)

// nonOmitemptySample is the test sample struct with omitempty.
type omitemptySample struct {
	Pair []inner `json:"pair,omitempty"`
}

type inner struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

var jsonStr = `{"pair":[{"key":"key1","value":"value1"},{"key":"key2","value":"value2"}]}`

// Test_Estimate_omitemptySample checks the usage of omitemptySample data.
func Test_Estimate_omitemptySample(t *testing.T) {
	var sample omitemptySample
	var err error
	err = json.Unmarshal(stringToBytes(jsonStr), &sample)
	if err != nil {
		log.Fatal(err)
	}

	var jsonData []byte
	jsonData, err = json.Marshal(sample)
	if err != nil {
		log.Fatal(err)
	}

	if *(*string)(unsafe.Pointer(&jsonData)) != jsonStr {
		log.Fatal("not equal")
	}
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(inner)
	},
}

// Test_Estimate_omitemptySample_SyncPool checks the usage of omitemptySample data and sync Pool.
func Test_Estimate_omitemptySample_SyncPool(t *testing.T) {
	var sample = pool.Get()
	var err error
	err = json.Unmarshal(stringToBytes(jsonStr), &sample)
	if err != nil {
		log.Fatal(err)
	}

	var jsonData []byte
	jsonData, err = json.Marshal(sample)
	if err != nil {
		log.Fatal(err)
	}

	if *(*string)(unsafe.Pointer(&jsonData)) != jsonStr {
		log.Fatal("not equal")
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

		_, err = json.Marshal(sample)
		/*var jsonData []byte
		jsonData, err = json.Marshal(sample)
		if err != nil {
			log.Fatal(err)
		}

		if *(*string)(unsafe.Pointer(&jsonData)) != jsonStr {
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

func Benchmark_Estimate_omitemptySample_syncPool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sample = pool.Get()
		var err error
		err = json.Unmarshal(stringToBytes(jsonStr), &sample)
		if err != nil {
			log.Fatal(err)
		}

		_, err = json.Marshal(sample)
		/*var jsonData []byte
		jsonData, err = json.Marshal(sample)
		if err != nil {
			log.Fatal(err)
		}

		if *(*string)(unsafe.Pointer(&jsonData)) != jsonStr {
			log.Fatal("not equal")
		}*/
	}
}
