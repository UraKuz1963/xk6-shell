package shell

import (
	"os/exec"

	"go.k6.io/k6/js/modules"
)

const version = "v0.0.1"

type Shell struct {
	Version string
}

func (ls *Shell) Apply(cmdcons string) (string, error) {
	out, err := exec.Command(cmdcons).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil

}
func init() {
	modules.Register("k6/x/shell", &Shell{
		Version: version,
	})
}
