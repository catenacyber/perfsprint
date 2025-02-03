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

The 5 following options cover all optimizations proposed by the linter.

Some have suboptions for specific cases.

- integer-format (formatting integer with the package `strconv`)
    - int-conversion : disable when the optimization adds a int/uint cast (readability)
- error-format (formatting errors)
    - errorf : known behavior change avoiding panic
    - err-error : known behavior change panicking for nil errors
- string-format (formatting strings)
   - sprintf1 : known behavior change avoiding panic
   - strconcat : disable turning some `fmt.Sprintf` to a string concatenation (readability)
- bool-format (formatting bool with `strconv.FormatBool`)
- hex-format (formatting bytes with `hex.EncodeToString`)


To disable `fmt.Errorf` optimization, you can use the flag `-errorf=false`
This optimization is not always equivalent.
The code
```
msg := "format string attack %s"
fmt.Errorf(msg)
```
will panic when its optimized version will not (so it should be safer).

To disable `fmt.Sprintf("toto")` optimization, you can use the flag `-sprintf1=false`
This optimization is not always equivalent.
The code
```
msg := "format string attack %s"
fmt.Sprintf(msg)
```
will panic when its optimized version will not (so it should be safer).

To enable `err.Error()` optimization, you can use the flag `-err-error=true`
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
