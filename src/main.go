package main

import (
	"fmt"
	"example.simple"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := example_simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My simple message",
		SampleList: []int32{1, 4, 7, 8},
	}

	//fmt.Println(&sm)
	fmt.Println(sm)
}
