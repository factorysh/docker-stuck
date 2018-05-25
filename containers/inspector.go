package containers

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Inspector struct {
	client *client.Client
	ctx    context.Context
}

type ContainerInspection struct {
	Container *Container
	Inspect   *types.ContainerJSON
	Err       error
}

func NewInspector() (*Inspector, error) {
	client, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return &Inspector{
		client: client,
		ctx:    context.Background(),
	}, nil
}

func (i *Inspector) inspect(container *Container, ci chan *ContainerInspection) {
	c := make(chan *ContainerInspection, 1)
	go func() {
		insp, err := i.client.ContainerInspect(i.ctx, container.DockerID)
		c <- &ContainerInspection{
			Container: container,
			Inspect:   &insp,
			Err:       err,
		}
	}()
	select {
	case insp := <-c:
		ci <- insp
	case <-time.After(10 * time.Second):
		ci <- &ContainerInspection{
			Container: container,
			Err:       fmt.Errorf("Timeout"),
		}
	}
}

func (i *Inspector) InspectAll() (int, chan *ContainerInspection, error) {
	containers, err := ContainersFromPs()
	if err != nil {
		return 0, nil, err
	}
	ci := make(chan *ContainerInspection, 1)
	for _, container := range containers {
		go i.inspect(container, ci)
	}
	return len(containers), ci, nil
}
