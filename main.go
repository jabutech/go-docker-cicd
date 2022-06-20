package main

import (
	"fmt"
	"go-docker-cicd/helper"
)

func main() {
	name := "Rizky"
	hello := HelloWorld(name)

	fmt.Println(hello)
}
