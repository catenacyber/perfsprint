package analyzer_test

import (
	"context"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"testing"
)

func BenchmarkStringFormatting(b *testing.B) {
	b.Run("fmt.Sprint", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprint("hello") //nolint:gosimple //https://staticcheck.io/docs/checks#S1039
		}
	})

	b.Run("fmt.Sprintf", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprintf("%s", "hello") //nolint:gosimple //https://staticcheck.io/docs/checks#S1025
		}
	})

	b.Run("just string", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = "hello"
		}
	})
}

func BenchmarkErrorFormatting(b *testing.B) {
	b.Run("fmt.Sprint", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprint(context.DeadlineExceeded)
		}
	})

	b.Run("fmt.Sprintf", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprintf("%s", context.DeadlineExceeded)
		}
	})

	b.Run("Error()", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = context.DeadlineExceeded.Error()
		}
	})
}

func BenchmarkBoolFormatting(b *testing.B) {
	b.Run("fmt.Sprint", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprint(true)
		}
	})

	b.Run("fmt.Sprintf", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprintf("%t", true)
		}
	})

	b.Run("strconv.FormatBool", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = strconv.FormatBool(true)
		}
	})
}

func BenchmarkHexEncoding(b *testing.B) {
	b.Run("fmt.Sprintf", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprintf("%x", []byte{'a', 'b', 'c'})
		}
	})

	b.Run("hex.EncodeToString", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = hex.EncodeToString([]byte{'a', 'b', 'c'})
		}
	})
}

func BenchmarkIntFormatting(b *testing.B) {
	b.Run("fmt.Sprint", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprint(math.MaxInt)
		}
	})

	b.Run("fmt.Sprintf", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprintf("%d", math.MaxInt)
		}
	})

	b.Run("strconv.Itoa", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = strconv.Itoa(math.MaxInt)
		}
	})
}

func BenchmarkUintFormatting(b *testing.B) {
	b.Run("fmt.Sprint", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprint(uint64(math.MaxUint))
		}
	})

	b.Run("fmt.Sprintf", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = fmt.Sprintf("%d", uint64(math.MaxUint))
		}
	})

	b.Run("strconv.FormatUint", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = strconv.FormatUint(math.MaxUint, 10)
		}
	})
}
