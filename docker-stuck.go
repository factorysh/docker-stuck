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
	good, bad, err := inspector.InspectAll()
	if err != nil {
		panic(err)
	}
	fmt.Println("Good\n")
	for _, container := range good {
		fmt.Println(container)
	}

	fmt.Println("Bad\n")
	for _, container := range bad {
		fmt.Println(container)
	}
}
