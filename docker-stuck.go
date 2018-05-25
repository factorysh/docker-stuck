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
	fmt.Println("Good")
	for _, container := range good {
		fmt.Printf("\tDocker: %s PID: %d\n", container.DockerID, container.Pid)
	}

	fmt.Println("Bad")
	for _, container := range bad {
		fmt.Printf("\tDocker: %s PID: %d\n", container.DockerID, container.Pid)
	}
}
