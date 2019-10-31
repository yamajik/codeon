package utils

import (
	"bytes"
	"os/exec"
)

// Exec bulabula
func Exec(program string, args ...string) error {
	cmd := exec.Command(program, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	return cmd.Run()
}
