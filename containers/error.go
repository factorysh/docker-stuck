package containers

type ContainerError interface {
	Container() *Container
}

type InspectorError struct {
	err       error
	container *Container
}

func (i *InspectorError) Error() string {
	return i.err.Error()
}

func (t *InspectorError) Container() *Container {
	return t.container
}
