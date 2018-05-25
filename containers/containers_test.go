package containers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexp(t *testing.T) {
	s := PatternDockerID.FindStringSubmatch("docker-containerd-shim -namespace moby -workdir /var/lib/docker/containerd/daemon/io.containerd.runtime.v1.linux/moby/28e4253a8a17f2bf9cf7ff44546f90134ade2106a23a4361f0993872f3fbf398 -address /var/run/docker/containerd/docker-containerd.sock -containerd-binary /usr/bin/docker-containerd -runtime-root /var/run/docker/runtime-runc")
	assert.Len(t, s, 2)
	assert.Equal(t, "28e4253a8a17f2bf9cf7ff44546f90134ade2106a23a4361f0993872f3fbf398", s[1])
}
