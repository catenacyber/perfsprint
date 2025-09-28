package a

import (
	"fmt"
)

func fmtf() {
	// fmt still used after removal
	fmt.Printf("Hello, World!", fmt.Sprintf("%s", "toto")) // want "string-format: fmt.Sprintf can be replaced with just using the string"
}
