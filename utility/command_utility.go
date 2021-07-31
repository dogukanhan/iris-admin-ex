package utility

import (
	"os/exec"
)

func RunCommand(cmd string) (string, error) {

	out, err := exec.Command("bash","-c",cmd).Output()

	if err != nil {
		return "", err
	} else {
		output := string(out[:])
		return output, nil
	}
}
