package compress

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_OneLine(t *testing.T) {
	t.Run("safe one line", func(t *testing.T) {
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

		str := string(raw)
		str = OneLine(str)
		require.Equal(t, str, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	})

	t.Run("unsafe one line", func(t *testing.T) {
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
		Seperate(string(raw))
		require.Equal(t, string(raw), "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	})
}

func Test_Separate(t *testing.T) {
	t.Run("safe separate", func(t *testing.T) {
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

		arr := Seperate(string(raw))
		for i := 0; i < len(arr); i++ {
			fmt.Println(arr[i])
		}
	})

	t.Run("unsafe one line", func(t *testing.T) {
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
		b := UnsafeSeperate(raw)
		for i := 0; i < len(b)-1; i++ {
			fmt.Println(string(raw[b[i]:b[i+1]]))
		}
	})
}
