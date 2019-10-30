package ssh

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/kevinburke/ssh_config"
)

// Config bulabula
type Config struct {
	Hosts []*Host
}

// UserConfigFile bulabula
func UserConfigFile() (file string, err error) {
	user, err := user.Current()
	if err != nil {
		return
	}
	file = filepath.Join(user.HomeDir, ".ssh", "config")
	return
}

// LoadConfig bulabula
func LoadConfig(file string) (cfg *Config, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()

	config, err := ssh_config.Decode(f)
	if err != nil {
		return
	}

	var hosts []*Host
	for _, host := range config.Hosts[1:] {
		h, loadHostErr := LoadHost(host)
		if err != nil {
			err = loadHostErr
			return
		}
		hosts = append(hosts, h)
	}

	cfg = &Config{
		Hosts: hosts,
	}
	if !cfg.HasDefaultHost() {
		err = cfg.AddDefaultHost()
	}
	return
}

// NewConfig bulabula
func NewConfig() (cfg *Config, err error) {
	cfg = &Config{
		Hosts: []*Host{},
	}
	err = cfg.AddDefaultHost()
	return
}

// LoadUserConfig bulabula
func LoadUserConfig() (cfg *Config, err error) {
	file, err := UserConfigFile()
	if err != nil {
		return
	}
	return LoadConfig(file)
}

// SaveConfig bulabula
func SaveConfig(file string, config *Config) (err error) {
	return config.Save(file)
}

// SaveUserConfig bulabula
func SaveUserConfig(config *Config) (err error) {
	file, err := UserConfigFile()
	if err != nil {
		return
	}
	return SaveConfig(file, config)
}

// String bulabula
func (c *Config) String() string {
	var hostStrings []string
	for _, h := range c.Hosts {
		hostStrings = append(hostStrings, h.String())
	}
	return strings.Join(hostStrings, "\n\n")
}

// Save bulabula
func (c *Config) Save(file string) (err error) {
	return ioutil.WriteFile(file, []byte(c.String()), 0644)
}

// MatchHost bulabula
func (c *Config) MatchHost(name string) (host *Host) {
	for _, h := range c.Hosts {
		if h.Match(name) {
			host = h
			return
		}
	}
	return
}

// FindHost bulabula
func (c *Config) FindHost(name string) (host *Host) {
	for _, h := range c.Hosts {
		if h.FindPattern(name) != nil {
			host = h
			return
		}
	}
	return
}

// AddHost bulabula
func (c *Config) AddHost(host *Host) (err error) {
	var h *Host
	for _, pattern := range host.Patterns {
		h = c.FindHost(pattern.String())
		if h != nil {
			break
		}
	}
	if h != nil {
		h.Update(host)
	} else {
		c.Hosts = append(c.Hosts, host)
	}
	return
}

// HasDefaultHost bulabula
func (c *Config) HasDefaultHost() bool {
	return c.FindHost("*") != nil
}

// AddDefaultHost bulabula
func (c *Config) AddDefaultHost() (err error) {
	defaultHost, err := NewDefaultHost()
	if err != nil {
		return
	}
	c.Hosts = append([]*Host{defaultHost}, c.Hosts...)
	return
}
