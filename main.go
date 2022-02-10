package main

import (
	"fmt"
	"protobuf-to-go/src/gen/simple"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My simple message",
		SampleList: []int32{1, 4, 7, 8},
	}

	fmt.Println(sm)

	sm.Name = "I renamed you"

	fmt.Println(sm)
	fmt.Println("The ID is:", sm.GetId())
}
