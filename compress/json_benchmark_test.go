package compress

import (
	"sync"
	"sync/atomic"
	"testing"
	"unsafe"
)

/*
Benchmark_OneLine benchmarks OneLine function;
raw string with tabs/newlines inputted, function called, result assigned;
loop iterates; measure performance concatenating into single string.
*/
func Benchmark_OneLine(b *testing.B) {
	// Raw string input with tabs and newlines
	raw := []byte{
		'\t', '\n', 'A', '\t', '\n',
		'B', '\t', '\n', 'C', '\t', '\n',
		'D', '\t', '\n', 'E', '\t', '\n',
		'F', '\t', '\n', 'G', '\t', '\n',
		'H', '\t', '\n', 'I', '\t', '\n',
		'J', '\t', '\n', 'K', '\t', '\n',
		'L', '\t', '\n', 'M', '\t', '\n',
		'N', '\t', '\n', 'O', '\t', '\n',
		'P', '\t', '\n', 'Q', '\t', '\n',
		'R', '\t', '\n', 'S', '\t', '\n',
		'T', '\t', '\n', 'U', '\t', '\n',
		'V', '\t', '\n', 'W', '\t', '\n',
		'X', '\t', '\n', 'Y', '\t', '\n',
		'Z', '\t', '\n',
	}

	// Loop for b.N iterations
	for i := 0; i < b.N; i++ {
		// Call OneLine() function and abandon the result
		_ = OneLine(string(raw))
	}

	// Take a glance of the result
	// fmt.Println(string(raw))
}

/*
Benchmark_UnsafeOneLine_WithMux benchmarks OneLine function by using unsafe pointer and sync.Mutex;
raw string with tabs/newlines inputted, function called, result assigned;
loop iterates; measure performance concatenating into single string.
*/
func Benchmark_UnsafeOneLine_WithMux(b *testing.B) {
	// Create a sync.Mutex
	var mu sync.Mutex

	// Raw string input with tabs and newlines
	raw := []byte{
		'\t', '\n', 'A', '\t', '\n',
		'B', '\t', '\n', 'C', '\t', '\n',
		'D', '\t', '\n', 'E', '\t', '\n',
		'F', '\t', '\n', 'G', '\t', '\n',
		'H', '\t', '\n', 'I', '\t', '\n',
		'J', '\t', '\n', 'K', '\t', '\n',
		'L', '\t', '\n', 'M', '\t', '\n',
		'N', '\t', '\n', 'O', '\t', '\n',
		'P', '\t', '\n', 'Q', '\t', '\n',
		'R', '\t', '\n', 'S', '\t', '\n',
		'T', '\t', '\n', 'U', '\t', '\n',
		'V', '\t', '\n', 'W', '\t', '\n',
		'X', '\t', '\n', 'Y', '\t', '\n',
		'Z', '\t', '\n',
	}

	// Loop for b.N iterations
	for i := 0; i < b.N; i++ {
		// Call UnsafeOneLine() function and abandon the result
		mu.Lock()
		UnsafeOneLine(raw, unsafe.Pointer(&raw))
		mu.Unlock()
	}

	// Take a glance of the result
	// fmt.Println(string(raw))
}

/*
Benchmark_UnsafeOneLine_Enhanced benchmarks OneLine function by using unsafe pointer and atomic;
raw string with tabs/newlines inputted, function called, result assigned;
loop iterates; measure performance concatenating into single string.
*/
func Benchmark_UnsafeOneLine_Enhanced(b *testing.B) {
	// Create a CAS Lock
	var locked int32 = 0

	// Raw string input with tabs and newlines
	raw := []byte{
		'\t', '\n', 'A', '\t', '\n',
		'B', '\t', '\n', 'C', '\t', '\n',
		'D', '\t', '\n', 'E', '\t', '\n',
		'F', '\t', '\n', 'G', '\t', '\n',
		'H', '\t', '\n', 'I', '\t', '\n',
		'J', '\t', '\n', 'K', '\t', '\n',
		'L', '\t', '\n', 'M', '\t', '\n',
		'N', '\t', '\n', 'O', '\t', '\n',
		'P', '\t', '\n', 'Q', '\t', '\n',
		'R', '\t', '\n', 'S', '\t', '\n',
		'T', '\t', '\n', 'U', '\t', '\n',
		'V', '\t', '\n', 'W', '\t', '\n',
		'X', '\t', '\n', 'Y', '\t', '\n',
		'Z', '\t', '\n',
	}

	// Loop for b.N iterations
	for i := 0; i < b.N; i++ {
		if atomic.CompareAndSwapInt32(&locked, 0, 1) {
			// Call UnsafeOneLine() function and abandon the result
			UnsafeOneLine(raw, unsafe.Pointer(&raw))
			atomic.StoreInt32(&locked, 0)
		}
	}

	// Take a glance of the result
	// fmt.Println(string(raw))
}

/*
Benchmark_Separate benchmarks Separate function;
raw JSON input, function call, result discarded;
measure performance separating string.
*/
func Benchmark_Separate(b *testing.B) {
	// Loop for b.N iterations
	for i := 0; i < b.N; i++ {
		// Raw JSON string input
		raw := []byte{
			'{', 'A', 'B', 'C', '}',
			'{', 'D', 'E', 'F', '}',
			'{', 'G', 'H', 'I', '}',
			'{', 'J', 'K', 'L', '}',
			'{', 'M', 'N', 'O', '}',
			'{', 'P', 'Q', 'R', '}',
			'{', 'S', 'T', 'U', '}',
			'{', 'V', 'W', 'X', '}',
			'{', 'Y', 'Z', '}',
		}
		// Call Separate() function and abandon the result
		_ = Separate(string(raw))
	}
}

/*
Benchmark_UnsafeSeparate_WithMux benchmarks Separate function by using unsafe pointer and sync.Mutex;
raw JSON input, function call, result discarded;
measure performance separating string.
*/
func Benchmark_UnsafeSeparate_WithMux(b *testing.B) {
	// Create a sync.Mutex
	var mu sync.Mutex

	// Raw JSON string input
	raw := []byte{
		'{', 'A', 'B', 'C', '}',
		'{', 'D', 'E', 'F', '}',
		'{', 'G', 'H', 'I', '}',
		'{', 'J', 'K', 'L', '}',
		'{', 'M', 'N', 'O', '}',
		'{', 'P', 'Q', 'R', '}',
		'{', 'S', 'T', 'U', '}',
		'{', 'V', 'W', 'X', '}',
		'{', 'Y', 'Z', '}',
	}

	// Loop for b.N iterations
	for i := 0; i < b.N; i++ {

		// Call UnsafeSeparate() function and abandon the result
		mu.Lock()
		_ = UnsafeSeparate(raw)
		mu.Unlock()
	}
}

/*
Benchmark_UnsafeSeparate_Enhanced benchmarks Separate function by using unsafe pointer and enhanced;
raw JSON input, function call, result discarded;
measure performance separating string.
*/
func Benchmark_UnsafeSeparate_Enhanced(b *testing.B) {
	// Create a CAS Lock
	var locked int32 = 0

	// Raw JSON string input
	raw := []byte{
		'{', 'A', 'B', 'C', '}',
		'{', 'D', 'E', 'F', '}',
		'{', 'G', 'H', 'I', '}',
		'{', 'J', 'K', 'L', '}',
		'{', 'M', 'N', 'O', '}',
		'{', 'P', 'Q', 'R', '}',
		'{', 'S', 'T', 'U', '}',
		'{', 'V', 'W', 'X', '}',
		'{', 'Y', 'Z', '}',
	}

	// Loop for b.N iterations
	for i := 0; i < b.N; i++ {
		if atomic.CompareAndSwapInt32(&locked, 0, 1) {
			// Call UnsafeSeparate() function and abandon the result
			_ = UnsafeSeparate(raw)
			atomic.StoreInt32(&locked, 0)
		}
	}
}
