package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/yamajik/codeon/launcher"
	"github.com/yamajik/codeon/ssh"
)

// SSHConfigJSONStruct bulabula
type SSHConfigJSONStruct struct {
	Host          string
	Path          string
	AdditionHosts []ssh.HostJSONStruct
}

var (
	codePath      string
	sshPath       string
	sshConfigFile string
)

var rootCmd = &cobra.Command{
	Use:   "codeon",
	Short: "codeon [URL]",
	Long:  "codeon [URL]",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		re := regexp.MustCompile(`codeon:(\w+)/(\w+)`)
		match := re.FindStringSubmatch(url)
		if len(match) != 3 {
			panic(fmt.Errorf("invalid url specified: %s", url))
		}
		codeonType, codeonConfigBase64 := match[1], match[2]
		switch codeonType {
		case "ssh":
			{
				var sshConfigStructs SSHConfigJSONStruct
				codeonConfigString, err := base64.StdEncoding.DecodeString(codeonConfigBase64)
				if err != nil {
					panic(err)
				}
				err = json.Unmarshal([]byte(codeonConfigString), &sshConfigStructs)
				if err != nil {
					panic(err)
				}
				additionHosts, err := ssh.LoadHostsFromStruct(sshConfigStructs.AdditionHosts)
				if err != nil {
					panic(err)
				}
				l, err := launcher.NewVscodeSSHLauncher()
				if err != nil {
					panic(err)
				}
				err = l.CodeProgram(codePath).SSHProgram(sshPath).SSHConfigFile(sshConfigFile).SSHAdditionHosts(additionHosts).Launch(sshConfigStructs.Host, sshConfigStructs.Path)
				if err != nil {
					panic(err)
				}
			}
		default:
			{
				panic(fmt.Errorf("invalid type specified: %s", codeonType))
			}

		}
	},
}

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "ssh [host] [path]",
	Long:  "ssh [host] [path]",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		host, path := args[0], args[1]
		l, err := launcher.NewVscodeSSHLauncher()
		if err != nil {
			panic(err)
		}
		err = l.CodeProgram(codePath).SSHProgram(sshPath).SSHConfigFile(sshConfigFile).Launch(host, path)
		if err != nil {
			panic(err)
		}
	},
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVarP(&codePath, "code-path", "", "", "Code Path")

	sshCmd.Flags().StringVarP(&codePath, "code-path", "", "", "Code Path")
	sshCmd.Flags().StringVarP(&sshPath, "ssh-path", "", "", "SSH Path")
	sshCmd.Flags().StringVarP(&sshConfigFile, "ssh-config-path", "", "", "SSH Config Path")

	rootCmd.AddCommand(sshCmd)
}

func initConfig() {

}
