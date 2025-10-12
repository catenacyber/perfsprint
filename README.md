# perfsprint

[![CI](https://github.com/catenacyber/perfsprint/actions/workflows/ci.yml/badge.svg)](https://github.com/catenacyber/perfsprint/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/catenacyber/perfsprint)](https://goreportcard.com/report/github.com/catenacyber/perfsprint?dummy=unused)

Golang linter for performance, aiming at usages of `fmt.Sprintf` which have faster alternatives.

## Installation

```sh
go install github.com/catenacyber/perfsprint@latest
```

## Usage

```sh
perfsprint --fix ./...
```

## Options

The 6 following options cover all optimizations proposed by the linter.

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

### Replacements

```
fmt.Sprintf("%s", strVal)  ->  strVal
fmt.Sprintf("%t", boolVal) ->  strconv.FormatBool(boolBal)
fmt.Sprintf("%x", hash)    ->  hex.EncodeToString(hash)
fmt.Sprintf("%d", id)      ->  strconv.Itoa(id)
fmt.Sprintf("%v", version) ->  strconv.FormatUint(uint64(version), 10)
```

More in [tests](./analyzer/testdata/src/p/p.go).
