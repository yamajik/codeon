package launcher

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/yamajik/codeon/ssh"
	"github.com/yamajik/codeon/utils"
)

// VscodeSSHLauncher bulabula
type VscodeSSHLauncher struct {
	VscodeLauncher
	sshProgram       string
	sshConfigFile    string
	sshConfig        *ssh.Config
	sshAdditionHosts []*ssh.Host
}

// DefaultSSHProgram bulabula
func DefaultSSHProgram() (sshProgram string) {
	switch runtime.GOOS {
	case "darwin", "linux":
		sshProgram = filepath.Join("/", "usr", "bin", "ssh")
	case "windows":
		sshProgram = filepath.Join("C:", "Windows", "System32", "OpenSSH")
	default:
		sshProgram = "ssh"
	}
	return
}

// ValidateSSHProgram bulabula
func ValidateSSHProgram(sshProgram string) (err error) {
	if sshProgram == "ssh" {
		return utils.Exec(sshProgram, "-V")
	}
	stat, err := os.Stat(sshProgram)
	if err != nil {
		return
	}
	if stat.IsDir() {
		err = fmt.Errorf("Program is a directory: %s", sshProgram)
	}
	return
}

// NewVscodeSSHLauncher bulabula
func NewVscodeSSHLauncher() (l *VscodeSSHLauncher, err error) {
	codeProgram, err := DefaultCodeProgram()
	if err != nil {
		return
	}
	sshConfigFile, err := ssh.UserConfigFile()
	if err != nil {
		return
	}
	l = &VscodeSSHLauncher{
		VscodeLauncher: VscodeLauncher{codeProgram: codeProgram},
		sshProgram:     DefaultSSHProgram(),
		sshConfigFile:  sshConfigFile,
	}
	return
}

// CodeProgram bulabula
func (l *VscodeSSHLauncher) CodeProgram(codeProgram string) *VscodeSSHLauncher {
	if codeProgram != "" {
		l.codeProgram = codeProgram
	}
	return l
}

// SSHProgram bulabula
func (l *VscodeSSHLauncher) SSHProgram(sshProgram string) *VscodeSSHLauncher {
	if sshProgram != "" {
		l.sshProgram = sshProgram
	}
	return l
}

// SSHConfigFile bulabula
func (l *VscodeSSHLauncher) SSHConfigFile(sshConfigFile string) *VscodeSSHLauncher {
	if sshConfigFile != "" {
		l.sshConfigFile = sshConfigFile
	}
	return l
}

// SSHAdditionHosts bulabula
func (l *VscodeSSHLauncher) SSHAdditionHosts(sshAdditionHosts []*ssh.Host) *VscodeSSHLauncher {
	l.sshAdditionHosts = append(l.sshAdditionHosts, sshAdditionHosts...)
	return l
}

// Launch bulabula
func (l *VscodeSSHLauncher) Launch(host string, folder string) (err error) {
	err = l.Validate()
	if err != nil {
		return
	}
	config, err := ssh.LoadUserConfig()
	if err != nil {
		return
	}
	for _, h := range l.sshAdditionHosts {
		err = config.AddHost(h)
		if err != nil {
			return
		}
	}
	err = ssh.SaveUserConfig(config)
	if err != nil {
		return
	}
	return l.Exec("--remote", "ssh-remote+"+host, folder)
}

// Validate bulabula
func (l *VscodeSSHLauncher) Validate() (err error) {
	err = ValidateCodeProgram(l.codeProgram)
	if err != nil {
		return
	}
	return ValidateSSHProgram(l.sshProgram)
}
