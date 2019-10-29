package launcher

import (
	"bytes"
	"fmt"
	"os/exec"
)

// VscodeLauncher bulabula
type VscodeLauncher struct {
	codeProgram string
}

// NewVscodeLauncher bulabula
func NewVscodeLauncher() *VscodeLauncher {
	return &VscodeLauncher{codeProgram: "code"}
}

// CodeProgram bulabula
func (l *VscodeLauncher) CodeProgram(codeProgram string) *VscodeLauncher {
	if codeProgram != "" {
		l.codeProgram = codeProgram
	}
	return l
}

// Launch bulabula
func (l *VscodeLauncher) Launch() (err error) {
	fmt.Println("Vscode launch.")
	return
}

// Exec bulabula
func (l *VscodeLauncher) Exec(args ...string) (err error) {
	cmd := exec.Command(l.codeProgram, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	return cmd.Run()
}
