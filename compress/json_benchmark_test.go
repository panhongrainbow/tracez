package compress

import (
	"testing"
	"unsafe"
)

func Benchmark_OneLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
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
		_ = OneLine(string(raw))
	}
}

func Benchmark_UnsafeOneLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
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
		UnsafeOneLine(raw, unsafe.Pointer(&raw))
	}
}

func Benchmark_Separate(b *testing.B) {
	for i := 0; i < b.N; i++ {
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
		_ = Seperate(string(raw))
	}
}

func Benchmark_UnsafeSeparate(b *testing.B) {
	for i := 0; i < b.N; i++ {
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
		_ = UnsafeSeperate(raw)
	}
}
