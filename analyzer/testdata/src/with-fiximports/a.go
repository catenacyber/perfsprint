package a

func loopconcat() { // want "fiximports: Fix imports"
	// no imports, adding strings
	s := ""
	for i := 0; i < 10; i++ {
		s += "toto" // want "concat-loop: string concatenation in a loop"
	}
	println("Hello, World!", s)
}
