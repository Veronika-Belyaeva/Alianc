package main

import (
	"Laba1/stack"
	"fmt"
)

func main() {
	s := Stack{}
	s.push("A")
	s.push("B")
	s.push("C")
	result, _ := s.pop()
	fmt.Print(result)

}