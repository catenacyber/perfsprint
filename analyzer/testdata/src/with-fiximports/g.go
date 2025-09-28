package a

import ( // want "fiximports: Fix imports"
	"fmt"
)

func fmtg() {
	// fmt still used after removal, but new strconv
	fmt.Printf("Hello, World!", fmt.Sprintf("%d", 123)) // want "integer-format: fmt.Sprintf can be replaced with faster strconv.Itoa"
}
