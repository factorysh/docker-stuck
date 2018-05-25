package containers

import (
	"regexp"
	"strings"

	ps "github.com/aelsabbahy/go-ps"
)

var PatternDockerID *regexp.Regexp

func init() {
	//docker-containerd-shim -namespace moby -workdir /var/lib/docker/containerd/daemon/io.containerd.runtime.v1.linux/moby/28e4253a8a17f2bf9cf7ff44546f90134ade2106a23a4361f0993872f3fbf398 -address /var/run/docker/containerd/docker-containerd.sock -containerd-binary /usr/bin/docker-containerd -runtime-root /var/run/docker/runtime-runc
	PatternDockerID = regexp.MustCompile("-workdir +/var/lib/docker/containerd/daemon/io.containerd.runtime.v1.linux/moby/([0-9a-f]+) ")
}

type Container struct {
	DockerID string
	Pid      int
}

func processes2containers(procs []ps.Process) ([]*Container, error) {
	containers := make([]*Container, 0)
	for _, proc := range procs {
		e := proc.Executable()
		if strings.HasPrefix(e, "docker-containerd-shim") {
			s := PatternDockerID.FindStringSubmatch(e)
			if len(s) > 1 {
				containers = append(containers, &Container{
					Pid:      proc.Pid(),
					DockerID: s[1],
				})
			}
		}
	}
	return containers, nil
}

func ContainersFromPs() ([]*Container, error) {
	procs, err := ps.Processes()
	if err != nil {
		return nil, err
	}
	return processes2containers(procs)
}
