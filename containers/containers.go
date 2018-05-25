package containers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	ps "github.com/aelsabbahy/go-ps"
)

var PatternDockerID *regexp.Regexp

func init() {
	//docker-containerd-shim -namespace moby -workdir /var/lib/docker/containerd/daemon/io.containerd.runtime.v1.linux/moby/28e4253a8a17f2bf9cf7ff44546f90134ade2106a23a4361f0993872f3fbf398 -address /var/run/docker/containerd/docker-containerd.sock -containerd-binary /usr/bin/docker-containerd -runtime-root /var/run/docker/runtime-runc
	PatternDockerID = regexp.MustCompile("docker-containerd-shim .*-workdir +/var/lib/docker/containerd/daemon/io.containerd.runtime.v1.linux/moby/([0-9a-f]+) ")
}

type Container struct {
	DockerID string
	Pid      int
}

// Read the cmdline of a pid, and use ' ' as a separator
func cmdLine(pid int) ([]byte, error) {
	cmd, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
	if err != nil {
		return nil, err
	}
	// \000 => ' '
	return bytes.Replace(cmd, []byte{0}, []byte{32}, -1), nil
}

func ContainersFromPs() ([]*Container, error) {
	procs, err := ps.Processes()
	if err != nil {
		return nil, err
	}
	containers := make([]*Container, 0)
	for _, proc := range procs {
		e := proc.Executable()
		if strings.HasPrefix(e, "docker-contain") { // Yes, it's truncated, and ugly
			cmd, err := cmdLine(proc.Pid())
			if err != nil {
				return nil, err
			}
			s := PatternDockerID.FindSubmatch(cmd)
			if len(s) > 1 {
				containers = append(containers, &Container{
					Pid:      proc.Pid(),
					DockerID: string(s[1]),
				})
			}
		}
	}
	return containers, nil
}
