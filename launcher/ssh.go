package launcher

import (
	"path/filepath"
	"runtime"

	"github.com/yamajik/codeon/ssh"
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
	case "darwin":
	case "linux":
		sshProgram = filepath.Join("/", "usr", "bin", "ssh")
	case "windows":
	default:
		sshProgram = "ssh"
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
