package p

import "fmt"

func notSelectorFuncAtAll() {
	_ = int32(23)
}

func notSprintfFuncAtAll() {
	fmt.Printf("test")
}

func sprintfDifferentNumberOfArgs() {
	fmt.Sprintf("test")
}

func sprintfNoBasciLit() {
	val := "val%d"
	fmt.Sprintf(val, 32)
}

func prinfLikeFunc() {
	fmt.Sprintf("%d", 32) // want "Sprintf can be replaced with faster function from strconv"
}
