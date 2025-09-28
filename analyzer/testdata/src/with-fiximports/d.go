package a

import (
	"fmt" // want "fiximports: Fix imports"
	"strings"
)

func loopconcat4() {
	// only fmt removed
	s := ""
	for i := 0; i < 10; i++ {
		s += "toto" // want "concat-loop: string concatenation in a loop"
	}
	println("Hello, World!", strings.Fields(fmt.Sprintf("%s", s))) // want "string-format: fmt.Sprintf can be replaced with just using the string"
}
