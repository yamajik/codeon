package launcher

import (
	"bytes"
	"fmt"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
)

// VscodeLauncher bulabula
type VscodeLauncher struct {
	codeProgram string
}

// DefaultCodeProgram bulabula
func DefaultCodeProgram() (codeProgram string, err error) {
	switch runtime.GOOS {
	case "darwin", "linux":
		codeProgram = filepath.Join("/", "usr", "local", "bin", "code")
	case "windows":
		user, getCurrentUserErr := user.Current()
		if getCurrentUserErr != nil {
			err = getCurrentUserErr
			return
		}
		codeProgram = filepath.Join(user.HomeDir, "AppData", "Local", "Programs", "Microsoft VS Code", "bin", "code")
	default:
		codeProgram = "ssh"
	}
	return
}

// NewVscodeLauncher bulabula
func NewVscodeLauncher() (l *VscodeLauncher, err error) {
	codeProgram, err := DefaultCodeProgram()
	if err != nil {
		return
	}
	l = &VscodeLauncher{codeProgram: codeProgram}
	return
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
	fmt.Println(cmd)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	return cmd.Run()
}
