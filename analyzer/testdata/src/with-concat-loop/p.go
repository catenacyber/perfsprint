package p

import ( // want "fiximports: Fix imports"
	"fmt"
)

func positive() {
	var s string
	words := []string{"one", "two", "three"}
	for w := range words {
		s += words[w] // want "concat-loop: string concatenation in a loop"
	}
	for w := range words {
		s = s + words[w] // want "concat-loop: string concatenation in a loop"
	}
	for w := 0; w < 10; w++ {
		s = s + "y" // want "concat-loop: string concatenation in a loop"
	}
	for w := 0; w < 10; w++ {
		if w%2 == 1 {
			s = s + "y" // want "concat-loop: string concatenation in a loop"
		}
	}
	nb := 0
	for w := 0; w < 10; w++ {
		if w%2 == 1 {
			nb += 1
		} else {
			s = s + "y" // want "concat-loop: string concatenation in a loop"
		}
	}
	for w := 0; w < 10; w++ {
		if w%2 == 1 {
			s = s + "x" // want "concat-loop: string concatenation in a loop"
		} else {
			s = s + "y"
		}
	}
	s2 := "prefix"
	for w := 0; w < 10; w++ {
		if w%2 == 1 {
			s2 = s2 + "x"
		} else {
			s = s + "y" // want "concat-loop: string concatenation in a loop"
		}
	}
	for w := 0; w < 10; w++ {
		s = s + "y" // want "concat-loop: string concatenation in a loop"
		if len(s)%3 == 1 {
			s = s + ","
		}
	}

	for w := 0; w < 10; w++ {
		s = fmt.Sprintf("%s+%s", s, "y")
		if w%2 == 1 {
			s = s + "," // want "concat-loop: string concatenation in a loop"
		}
	}
}

func negative() {
	for w := 0; w < 10; w++ {
		s := "local"
		s = s + "y"
	}
	for w := 0; w < 10; w++ {
		var s string
		s = s + "y"
	}
	for w := 0; w < 10; w++ {
		var s2, s string
		s = s + "y"
		_ = s
		_ = s2
	}
	for w := 0; w < 10; w++ {
		s2, s := "local", "same"
		s = s + "y"
		_ = s
		_ = s2
	}
	nb := 0
	for w := 0; w < 10; w++ {
		nb += w
	}
	for w := 0; w < 10; w++ {
		nb = nb + w
	}
	words := []string{"one", "two", "three"}
	var s string
	for w := range words {
		s = "toto" + words[w]
	}
	var s2 string
	for w := range words {
		s = s2 + words[w]
	}
	_ = s
	_ = s2
}
