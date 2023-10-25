package noconv

import (
	"errors"
	"fmt"
	"io"
)

var errSentinel = errors.New("connection refused")

func negative() {
	var err error
	fmt.Sprintf("%s", errSentinel)
	fmt.Sprintf("%v", errSentinel)
	fmt.Sprint(errSentinel)
	fmt.Sprintf("%s", io.EOF)
	fmt.Sprintf("%v", io.EOF)
	fmt.Sprint(io.EOF)
	fmt.Sprintf("%s", err)
	fmt.Sprintf("%v", err)
	fmt.Sprint(err)
	var errNil error
	fmt.Sprint(errNil)
}
