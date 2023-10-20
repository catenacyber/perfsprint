package analyzer_test

import (
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"testing"

	"github.com/catenacyber/perfsprint/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	t.Parallel()
	analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), analyzer.New(), "p")
}

func TestAnalyzerNoConv(t *testing.T) {
	t.Parallel()
	a := analyzer.New()
	err := a.Flags.Set("int-conversion", "false")
	if err != nil {
		t.Fatalf("failed to set int-conversion flag")
	}
	analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), a, "noconv")
}

func TestReplacements(t *testing.T) {
	t.Parallel()

	cases := []struct {
		before, after string
	}{
		{before: fmt.Sprintf("%s", "hello"), after: "hello"}, //nolint:gosimple //https://staticcheck.io/docs/checks#S1025
		{before: fmt.Sprintf("%v", "hello"), after: "hello"},
		{before: fmt.Sprint("hello"), after: "hello"}, //nolint:gosimple //https://staticcheck.io/docs/checks#S1039

		{before: fmt.Sprintf("%s", io.EOF), after: io.EOF.Error()},
		{before: fmt.Sprintf("%v", io.EOF), after: io.EOF.Error()},

		{before: fmt.Sprintf("%t", true), after: strconv.FormatBool(true)},
		{before: fmt.Sprintf("%v", true), after: strconv.FormatBool(true)},
		{before: fmt.Sprint(true), after: strconv.FormatBool(true)},
		{before: fmt.Sprintf("%t", false), after: strconv.FormatBool(false)},
		{before: fmt.Sprintf("%v", false), after: strconv.FormatBool(false)},
		{before: fmt.Sprint(false), after: strconv.FormatBool(false)},

		{before: fmt.Sprintf("%x", []byte{'a', 'b', 'c'}), after: hex.EncodeToString([]byte{'a', 'b', 'c'})},

		{before: fmt.Sprintf("%d", 42), after: strconv.Itoa(42)},
		{before: fmt.Sprintf("%v", 42), after: strconv.Itoa(42)},
		{before: fmt.Sprint(42), after: strconv.Itoa(42)},
		{before: fmt.Sprintf("%d", int8(42)), after: strconv.Itoa(int(int8(42)))},
		{before: fmt.Sprintf("%v", int8(42)), after: strconv.Itoa(int(int8(42)))},
		{before: fmt.Sprint(int8(42)), after: strconv.Itoa(int(int8(42)))},
		{before: fmt.Sprintf("%d", int16(42)), after: strconv.Itoa(int(int16(42)))},
		{before: fmt.Sprintf("%v", int16(42)), after: strconv.Itoa(int(int16(42)))},
		{before: fmt.Sprint(int16(42)), after: strconv.Itoa(int(int16(42)))},
		{before: fmt.Sprintf("%d", int32(42)), after: strconv.Itoa(int(int32(42)))},
		{before: fmt.Sprintf("%v", int32(42)), after: strconv.Itoa(int(int32(42)))},
		{before: fmt.Sprint(int32(42)), after: strconv.Itoa(int(int32(42)))},
		{before: fmt.Sprintf("%d", int64(42)), after: strconv.FormatInt(int64(42), 10)},
		{before: fmt.Sprintf("%v", int64(42)), after: strconv.FormatInt(int64(42), 10)},
		{before: fmt.Sprint(int64(42)), after: strconv.FormatInt(int64(42), 10)},

		{before: fmt.Sprintf("%d", uint(42)), after: strconv.FormatUint(uint64(uint(42)), 10)},
		{before: fmt.Sprintf("%v", uint(42)), after: strconv.FormatUint(uint64(uint(42)), 10)},
		{before: fmt.Sprint(uint(42)), after: strconv.FormatUint(uint64(uint(42)), 10)},
		{before: fmt.Sprintf("%d", uint8(42)), after: strconv.FormatUint(uint64(uint8(42)), 10)},
		{before: fmt.Sprintf("%v", uint8(42)), after: strconv.FormatUint(uint64(uint8(42)), 10)},
		{before: fmt.Sprint(uint8(42)), after: strconv.FormatUint(uint64(uint8(42)), 10)},
		{before: fmt.Sprintf("%d", uint16(42)), after: strconv.FormatUint(uint64(uint16(42)), 10)},
		{before: fmt.Sprintf("%v", uint16(42)), after: strconv.FormatUint(uint64(uint16(42)), 10)},
		{before: fmt.Sprint(uint16(42)), after: strconv.FormatUint(uint64(uint16(42)), 10)},
		{before: fmt.Sprintf("%d", uint32(42)), after: strconv.FormatUint(uint64(uint32(42)), 10)},
		{before: fmt.Sprintf("%v", uint32(42)), after: strconv.FormatUint(uint64(uint32(42)), 10)},
		{before: fmt.Sprint(uint32(42)), after: strconv.FormatUint(uint64(uint32(42)), 10)},
		{before: fmt.Sprintf("%d", uint64(42)), after: strconv.FormatUint(uint64(42), 10)},
		{before: fmt.Sprint(uint64(42)), after: strconv.FormatUint(uint64(42), 10)},
		{before: fmt.Sprintf("%v", uint64(42)), after: strconv.FormatUint(uint64(42), 10)},
	}
	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			if tt.before != tt.after {
				t.Fatalf("%s != %s", tt.before, tt.after)
			}
		})
	}
}
