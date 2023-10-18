# gostrconv

[![CI](https://github.com/catenacyber/gostrconv/actions/workflows/ci.yml/badge.svg)](https://github.com/catenacyber/gostrconv/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/catenacyber/gostrconv)](https://goreportcard.com/report/github.com/catenacyber/gostrconv?dummy=unused)

Golang linter for performance, aiming at usages of `fmt.Sprintf` which have faster alternatives.

## Installation

```sh
go get github.com/catenacyber/gostrconv@latest
```

## Usage

```sh
gostrconv --fix ./...
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
