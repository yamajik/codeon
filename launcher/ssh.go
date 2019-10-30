package launcher

import (
	"github.com/yamajik/codeon/ssh"
)

// VscodeSSHLauncher bulabula
type VscodeSSHLauncher struct {
	VscodeLauncher
	sshProgram    string
	sshConfigFile string
}

// NewVscodeSSHLauncher bulabula
func NewVscodeSSHLauncher() (l *VscodeSSHLauncher, err error) {
	sshConfigFile, err := ssh.UserConfigFile()
	if err != nil {
		return
	}
	l = &VscodeSSHLauncher{
		VscodeLauncher: VscodeLauncher{codeProgram: "code"},
		sshProgram:     "ssh",
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

// Launch bulabula
func (l *VscodeSSHLauncher) Launch(host string, folder string, hostsJSONString string) (err error) {
	config, err := ssh.LoadUserConfig()
	if err != nil {
		return
	}
	if hostsJSONString != "" {
		err = config.AddHostsFromJSON(hostsJSONString)
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
