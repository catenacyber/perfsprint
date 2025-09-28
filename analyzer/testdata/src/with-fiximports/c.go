package a

import (
	"fmt" // want "fiximports: Fix imports"
)

func loopconcat3() {
	// only fmt removed, but strings added
	s := ""
	for i := 0; i < 10; i++ {
		s += "toto" // want "concat-loop: string concatenation in a loop"
	}
	println("Hello, World!", fmt.Sprintf("%s", s)) // want "string-format: fmt.Sprintf can be replaced with just using the string"
}
