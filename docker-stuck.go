package main

import (
	"fmt"

	"github.com/factorysh/docker-stuck/containers"
)

func main() {
	inspector, err := containers.NewInspector()
	if err != nil {
		panic(err)
	}
	_, bad, err := inspector.InspectAll()
	if err != nil {
		panic(err)
	}
	for _, container := range bad {
		fmt.Println(container)
	}
}
