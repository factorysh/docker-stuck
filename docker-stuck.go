package main

import (
	"fmt"
	"os"

	"github.com/factorysh/docker-stuck/containers"
)

func main() {
	inspector, err := containers.NewInspector()
	if err != nil {
		panic(err)
	}
	n, ci, err := inspector.InspectAll()
	if err != nil {
		panic(err)
	}

	kill := len(os.Args) > 1 && os.Args[1] == "--kill"

	for i := 0; i < n; i++ {
		insp := <-ci
		status := "bad "
		if insp.Err == nil {
			status = "good"
		}
		fmt.Printf("[%s] Docker: %s PID: %d", status, insp.Container.DockerID, insp.Container.Pid)
		if insp.Err != nil {
			fmt.Printf(" %s", insp.Err.Error())
		}
		fmt.Println()
		if kill {
			p, err := os.FindProcess(insp.Container.Pid)
			if err != nil {
				panic(err)
			}
			p.Kill()
		}
	}
}
