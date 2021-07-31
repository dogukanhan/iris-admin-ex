package utility

import (
	"os/exec"
)

func RunCommand(name string, arg ...string) (string, error) {

	out, err := exec.Command(name, arg...).Output()

	if err != nil {
		return "", err
	} else {
		output := string(out[:])
		return output, nil
	}
}
