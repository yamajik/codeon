package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yamajik/codeon/launcher"
)

func main() {
	sshCmd := flag.NewFlagSet("ssh", flag.ExitOnError)
	codePath := sshCmd.String("codepath", "", "codepath")
	sshPath := sshCmd.String("sshpath", "", "sshpath")
	sshConfigFile := sshCmd.String("sshconfig", "", "sshconfig")
	sshHosts := sshCmd.String("config", "", "config")

	if len(os.Args) < 4 {
		fmt.Println("expected 'ssh host path'")
		os.Exit(1)
	}

	host, path := os.Args[2], os.Args[3]
	sshCmd.Parse(os.Args[4:])

	err := launcher.NewVscodeSSHLauncher().CodeProgram(*codePath).SSHProgram(*sshPath).SSHConfigFile(*sshConfigFile).Launch(host, path, *sshHosts)
	if err != nil {
		panic(err)
	}
}
