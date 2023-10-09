# gostrconv
Gostrconv: Golang linter for performance, aiming at usages of `fmt.Sprintf` which have faster alternatives.

# Usage

./gostrconv file.go

Rewrites `fmt.Sprintf("%d",` into faster `strconv.Itoa` and such
