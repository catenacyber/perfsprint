package analyzer_test

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"strconv"
	"testing"

	"github.com/catenacyber/perfsprint/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	t.Parallel()

	t.Run("default", func(t *testing.T) {
		a := analyzer.New()
		analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), a, "default")
	})

	defaultAnalyzer := analyzer.New()
	defaultAnalyzer.Flags.VisitAll(func(f *flag.Flag) {
		if f.Name == "loop-other-ops" {
			return
		}
		name := f.Name
		var changedVal string
		switch f.DefValue {
		case "false":
			changedVal = "true"
			name = "with-" + f.Name
		case "true":
			changedVal = "false"
			name = "without-" + f.Name
		default:
			t.Fatalf("default value neither false or true")
		}

		t.Run(name, func(t *testing.T) {
			a := analyzer.New()
			if f.Name == "fiximports" {
				name = "with-fiximports"
			} else if f.Name == "concat-loop" {
				name = "with-concat-loop"
				err := a.Flags.Set("loop-other-ops", "true")
				if err != nil {
					t.Fatalf("failed to set %q flag", f.Name)
				}
			} else {
				err := a.Flags.Set(f.Name, changedVal)
				if err != nil {
					t.Fatalf("failed to set %q flag", f.Name)
				}
			}
			analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), a, name)
		})
	})
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
