package compress

import (
	"sync"
	"sync/atomic"
	"testing"
	"unsafe"
)

// Test_Race_OneLine tests the race condition in the OneLine function by running 1000 goroutines.
func Test_Race_OneLine(t *testing.T) {
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

	// Loop for 1000 iterations
	for i := 0; i < 1000; i++ {
		go func() {
			// Convert to string
			str := string(raw)
			// Call OneLine() function
			str = OneLine(str)
		}()
	}
}

/*
Test_Race_UnsafeOneLine_WithMux tests the race condition in the OneLine function
by running 1000 goroutines and using sync.Mutex.
*/
func Test_Race_UnsafeOneLine_WithMux(t *testing.T) {
	// Create a sync.Mutex
	var mu sync.Mutex

	// Raw JSON string input
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

	// Loop for 1000 iterations
	for i := 0; i < 1000; i++ {
		go func() {
			// Call UnsafeOneLine() function
			mu.Lock()
			UnsafeOneLine(raw, unsafe.Pointer(&raw))
			mu.Unlock()
		}()
	}
}

/*
Test_Race_UnsafeOneLine_Enhanced tests the race condition in the OneLine function
by running 1000 goroutines and using atomic.
*/
func Test_Race_UnsafeOneLine_Enhanced(t *testing.T) {
	// Create a CAS Lock
	var locked int32 = 0

	// Raw JSON string input
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

	// Loop for 1000 iterations
	for i := 0; i < 1000; i++ {
		go func() {
			// Call UnsafeOneLine() function
			if atomic.CompareAndSwapInt32(&locked, 0, 1) {
				UnsafeOneLine(raw, unsafe.Pointer(&raw))
				atomic.StoreInt32(&locked, 0)
			}
		}()
	}
}

// Test_Race_Separate tests the race condition in the Separate() function by running 1000 goroutines.
func Test_Race_Separate(t *testing.T) {
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

	// Loop for 1000 iterations
	for i := 0; i < 1000; i++ {
		go func() {
			// Call Separate() function and abandon the result
			_ = Separate(string(raw))
		}()
	}
}

/*
Test_Race_Separate_WithMux tests the race condition in the Separate() function
by running 1000 goroutines and using sync.Mutex.
*/
func Test_Race_Separate_WithMux(t *testing.T) {
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

	for i := 0; i < 1000; i++ {
		go func() {
			// Call UnsafeSeparate() function and abandon the result
			mu.Lock()
			_ = UnsafeSeparate(raw)
			mu.Unlock()
		}()
	}
}

/*
Test_Race_Separate_Enhanced tests the race condition in the Separate() function
by running 1000 goroutines and using atomic.
*/
func Test_Race_Separate_Enhanced(t *testing.T) {
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

	// Loop for 1000 iterations
	for i := 0; i < 1000; i++ {
		go func() {
			// Call UnsafeSeparate() function and abandon the result
			if atomic.CompareAndSwapInt32(&locked, 0, 1) {
				_ = UnsafeSeparate(raw)
				atomic.StoreInt32(&locked, 0)
			}
		}()
	}
}
