package compress

import (
	"github.com/stretchr/testify/require"
	"testing"
	"unsafe"
)

/*
Test_Check_OneLine tests OneLine, safe/unsafe;
string/JSON input with tabs/newlines;
call function, validate concatenated result;
convert and combine into single string.
*/
func Test_Check_OneLine(t *testing.T) {
	// Run test case for safe one line separation
	t.Run("Safe one line", func(t *testing.T) {
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

		// Convert to string
		str := string(raw)
		// Call OneLine() function
		str = OneLine(str)
		// Check result
		require.Equal(t, str, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	})
	// Run test case for unsafe one line separation
	t.Run("Unsafe one line", func(t *testing.T) {
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
		// Call UnsafeOneLine() function
		UnsafeOneLine(raw, unsafe.Pointer(&raw))
		// Check result
		require.Equal(t, string(raw), "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	})
}

/*
Test_Check_Separate tests Separate, safe and unsafe;
validate JSON separated, checking output.
*/
func Test_Check_Separate(t *testing.T) {
	// Run test case for safe separation
	t.Run("Safe separate", func(t *testing.T) {
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
		// Call Separate() function
		arr := Separate(string(raw))
		// Validate each line is not empty
		for i := 0; i < len(arr); i++ {
			require.NotEmpty(t, arr[i])
		}
		// Check each separated line
		require.Equal(t, arr[0], "{ABC}")
		require.Equal(t, arr[1], "{DEF}")
		require.Equal(t, arr[2], "{GHI}")
		require.Equal(t, arr[3], "{JKL}")
		require.Equal(t, arr[4], "{MNO}")
		require.Equal(t, arr[5], "{PQR}")
		require.Equal(t, arr[6], "{STU}")
		require.Equal(t, arr[7], "{VWX}")
		require.Equal(t, arr[8], "{YZ}")
	})
	// Run test case for unsafe separation
	t.Run("Unsafe one line", func(t *testing.T) {
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
		// Call UnsafeSeparate() function
		b := UnsafeSeparate(raw)
		//Validate each separated line is not empty
		for i := 0; i < len(b)-1; i++ {
			require.NotEmpty(t, string(raw[b[i]:b[i+1]]))
		}
		// Check each separated line
		require.Equal(t, string(raw[b[0]:b[0+1]]), "{ABC}")
		require.Equal(t, string(raw[b[1]:b[1+1]]), "{DEF}")
		require.Equal(t, string(raw[b[2]:b[2+1]]), "{GHI}")
		require.Equal(t, string(raw[b[3]:b[3+1]]), "{JKL}")
		require.Equal(t, string(raw[b[4]:b[4+1]]), "{MNO}")
		require.Equal(t, string(raw[b[5]:b[5+1]]), "{PQR}")
		require.Equal(t, string(raw[b[6]:b[6+1]]), "{STU}")
		require.Equal(t, string(raw[b[7]:b[7+1]]), "{VWX}")
		require.Equal(t, string(raw[b[8]:b[8+1]]), "{YZ}")
	})
}
