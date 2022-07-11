package main

import (
	"fmt"

	"pt.example/grcp-test/test/test"
)

func main() {
	v := test.Vertex{3, 4}
	fmt.Println(v.Abs())
}
