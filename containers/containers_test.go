package containers

import (
	"testing"

	ps "github.com/aelsabbahy/go-ps"
	"github.com/stretchr/testify/assert"
)

type TestProcess struct {
	pid  int
	ppid int
	exe  string
}

func (t *TestProcess) Pid() int {
	return t.pid

}

func (t *TestProcess) PPid() int {
	return t.ppid
}

func (t *TestProcess) Executable() string {
	return t.exe
}

func TestContainers(t *testing.T) {
	procs := []ps.Process{
		&TestProcess{
			pid: 23940,
			exe: "docker-containerd-shim -namespace moby -workdir /var/lib/docker/containerd/daemon/io.containerd.runtime.v1.linux/moby/28e4253a8a17f2bf9cf7ff44546f90134ade2106a23a4361f0993872f3fbf398 -address /var/run/docker/containerd/docker-containerd.sock -containerd-binary /usr/bin/docker-containerd -runtime-root /var/run/docker/runtime-runc",
		},
		&TestProcess{
			pid: 23423,
			exe: "grep docker",
		},
	}
	c, err := processes2containers(procs)
	assert.NoError(t, err)
	t.Log(c)
	assert.Len(t, c, 1)
	assert.Equal(t, 23940, c[0].Pid)
	assert.Equal(t, "28e4253a8a17f2bf9cf7ff44546f90134ade2106a23a4361f0993872f3fbf398", c[0].DockerID)
}
