package a

import (
	"fmt" // want "fiximports: Fix imports"
)

func fmte() {
	// fmt removed as only import
	println("Hello, World!", fmt.Sprintf("%s", "toto")) // want "string-format: fmt.Sprintf can be replaced with just using the string"
}
