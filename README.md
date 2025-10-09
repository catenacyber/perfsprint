# perfsprint

[![CI](https://github.com/catenacyber/perfsprint/actions/workflows/ci.yml/badge.svg)](https://github.com/catenacyber/perfsprint/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/catenacyber/perfsprint)](https://goreportcard.com/report/github.com/catenacyber/perfsprint?dummy=unused)

Golang linter for performance that replaces uses of `fmt.Sprintf` and `fmt.Errorf` with better (both in CPU and memory) alternatives.

## Installation

If you use `golangci-lint`, you can add it to your `.golangci.yml`:

```yaml
version: "2"
linters:
    enable:
        - perfsprint
```

## Options

The 6 options below cover all optimizations proposed by the linter. 

Some have suboptions for specific cases, including cases where the linter proposes a behavior change.

- integer-format (formatting integer with the package `strconv`)
    - int-conversion : disable when the optimization adds a int/uint cast (readability)
- error-format (formatting errors)
    - errorf : turns `fmt.Errorf` into `errors.New`, known behavior change, avoiding panic
    - err-error : turns `fmt.Sprintf(err)` and like into `err.Error()`, known behavior change, panicking for nil errors
- string-format (formatting strings)
    - sprintf1 : turns `fmt.Sprintf(msg)` and like into `msg`, known behavior change, avoiding panic
    - strconcat : disable turning some `fmt.Sprintf` to a string concatenation (readability)
- bool-format (formatting bool with `strconv.FormatBool`)
- hex-format (formatting bytes with `hex.EncodeToString`)
- concat-loop (replacing string concatenation in a loop by `strings.Builder`)
    - loop-other-ops : matches also if the loop has other operations than concatenation on the string

There is also a `fix-imports` option that should auto-fix the imports section.
It will add a comment `//TODO FIXME` if a package with the same name is already used.

The `errorf` optimization is not always equivalent:
```
msg := "format string attack %s"
// fmt.Errorf(msg) // original, panics
errors.New(msg) // optimized, does not panic
```

The `sprintf1` optimization is not always equivalent:
```
msg := "format string attack %s"
// a := fmt.Sprintf(msg) // original, panics
a := msg // optimized, does not panic
```

The `err-error` optimization is not always equivalent:
```
var err error
// fmt.Sprintf(err) // original, does not panic, prints <nil>
err.Error() // optimized, panics !
```
This optimization only works when the error is not nil, otherwise the resulting code will panic.

The `loop-other-ops` optimization is not always equivalent.
The proposed fix will likely fail to compile.
Here is an example where the linter will rightly trigger but fail to propose a good fix.
```
s := ""
for i:=0; i<10; i++ {
    s += "ab"
    if len(s) > 10 { // not a concatenation, no autofix
        return s // not a concatenation, no autofix
    }
}
```

## Replacements

In general, using `fmt.Sprintf` is slow because it has to parse the arguments and format them according to various supported verbs (`%x`, `%d`, `%v`, etc.). 

This linter proposes the following replacements that are faster and allocate less memory.

```
fmt.Sprintf("%s", strVal)  ->  strVal
fmt.Sprintf("%t", boolVal) ->  strconv.FormatBool(boolBal)
fmt.Sprintf("%x", hash)    ->  hex.EncodeToString(hash)
fmt.Sprintf("%d", id)      ->  strconv.Itoa(id)
fmt.Sprintf("%v", version) ->  strconv.FormatUint(uint64(version), 10)
```

To know how fast each replacement is, run `make bench`. You will see something like this for each replacement:

```
cpu: Apple M4 Max
BenchmarkStringFormatting/fmt.Sprint            227844582               25.39 ns/op            5 B/op            1 allocs/op
BenchmarkStringFormatting/fmt.Sprintf           222438842               27.40 ns/op            5 B/op            1 allocs/op
BenchmarkStringFormatting/REPLACEMENT:just_string               1000000000               0.2421 ns/op            0 B/op          0 allocs/op
```

The replacement is 100x faster (25 ns per operation vs 0.23 nanoseconds per operation) and allocates no memory (5 Bytes per operation vs 0 Bytes per operation).

More in [tests](./analyzer/testdata/src) and in this blog: https://philpearl.github.io/post/bad_go_sprintf/
