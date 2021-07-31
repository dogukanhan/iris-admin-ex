package utility

import (
	"os/exec"
)

func RunCommand(command string) (string, error) {

	out, err := exec.Command(command).Output()

	if err != nil {
		return "", err
	} else {
		output := string(out[:])
		return output, nil
	}
}
