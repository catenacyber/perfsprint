# perfsprint

[![CI](https://github.com/catenacyber/perfsprint/actions/workflows/ci.yml/badge.svg)](https://github.com/catenacyber/perfsprint/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/catenacyber/perfsprint)](https://goreportcard.com/report/github.com/catenacyber/perfsprint?dummy=unused)

Golang linter for performance, aiming at usages of `fmt.Sprintf` which have faster alternatives.

## Installation

```sh
go get github.com/catenacyber/perfsprint@latest
```

## Usage

```sh
perfsprint --fix ./...
```

### Replacements

```
fmt.Sprintf("%s", strVal)  ->  strVal
fmt.Sprintf("%t", boolVal) ->  strconv.FormatBool(boolBal)
fmt.Sprintf("%x", hash)    ->  hex.EncodeToString(hash)
fmt.Sprintf("%d", id)      ->  strconv.Itoa(id)
fmt.Sprintf("%v", version) ->  strconv.FormatUint(uint64(version), 10)
```

More in [tests](./analyzer/testdata/src/p/p.go).
