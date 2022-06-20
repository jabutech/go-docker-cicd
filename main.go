package main

import (
	"fmt"

	"github.com/jabutech/go-docker-cicd/helper"
)

func main() {
	name := "Rizky"
	hello := helper.HelloWorld(name)

	fmt.Println(hello)
}
