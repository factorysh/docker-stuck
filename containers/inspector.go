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

func (i *Inspector) inspect(container *Container, containers chan types.ContainerJSON, errors chan ContainerError) {
	ci := make(chan types.ContainerJSON, 1)
	ce := make(chan ContainerError, 1)
	go func() {
		insp, err := i.client.ContainerInspect(i.ctx, container.DockerID)
		if err == nil {
			ci <- insp
		} else {
			e := &InspectorError{err, container}
			ce <- e
		}
	}()
	select {
	case insp := <-ci:
		containers <- insp
	case err := <-ce:
		errors <- err
	case <-time.After(10 * time.Second):
		errors <- &InspectorError{fmt.Errorf("Timeout"), container}
	}
}

func (i *Inspector) InspectAll() ([]*Container, []*Container, error) {
	containers, err := ContainersFromPs()
	if err != nil {
		return nil, nil, err
	}
	inspects := make(chan types.ContainerJSON, 1)
	errors := make(chan ContainerError, 1)
	dico := make(map[string]*Container)
	for _, container := range containers {
		dico[container.DockerID] = container
		go i.inspect(container, inspects, errors)
	}
	cpt := len(containers)
	bad := make([]*Container, 0)
	good := make([]*Container, 0)
	for cpt > 0 {
		select {
		case insp := <-inspects:
			good = append(good, dico[insp.ID])
		case err := <-errors:
			bad = append(bad, err.Container())
		}
		cpt--
	}
	return good, bad, nil
}
