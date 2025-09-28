package a

import (
	"strings"
)

func loopconcat2() {
	// already strings
	s := ""
	for i := 0; i < 10; i++ {
		s += "toto" // want "concat-loop: string concatenation in a loop"
	}
	println("Hello, World!", strings.Fields(s))
}
