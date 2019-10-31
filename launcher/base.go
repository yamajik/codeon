package launcher

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/yamajik/codeon/utils"
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
		codeProgram = "code"
	}
	return
}

// ValidateCodeProgram bulabula
func ValidateCodeProgram(codeProgram string) (err error) {
	if codeProgram == "code" {
		return utils.Exec(codeProgram, "--version")
	}
	stat, err := os.Stat(codeProgram)
	if err != nil {
		return
	}
	if stat.IsDir() {
		err = fmt.Errorf("Program is a directory: %s", codeProgram)
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
	err = l.Validate()
	if err != nil {
		return
	}
	return l.Exec()
}

// Exec bulabula
func (l *VscodeLauncher) Exec(args ...string) (err error) {
	return utils.Exec(l.codeProgram, args...)
}

// Validate bulabula
func (l *VscodeLauncher) Validate() (err error) {
	return ValidateCodeProgram(l.codeProgram)
}
