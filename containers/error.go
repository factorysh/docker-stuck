package containers

import "fmt"

type ContainerError interface {
	ContainerID() string
}

type InspectorTimeout struct {
	containerID string
}

func (t *InspectorTimeout) Error() string {
	return fmt.Sprintf("Timeout for container %s", t.containerID)
}

func (t *InspectorTimeout) ContainerID() string {
	return t.containerID
}

type InspectorError struct {
	err         error
	containerID string
}

func (i *InspectorError) Error() string {
	return i.err.Error()
}

func (t *InspectorError) ContainerID() string {
	return t.containerID
}
