package test

import (
	"testing"

	"github.com/jabutech/go-docker-cicd/helper"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	name := "Rizky"
	helloWorld := helper.HelloWorld(name)

	assert.Equal(t, "Hello Rizky", helloWorld)
}
