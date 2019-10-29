package ssh

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/kevinburke/ssh_config"
)

// Config bulabula
type Config struct {
	config *ssh_config.Config
	Hosts  []*Host
}

// UserConfigFile bulabula
func UserConfigFile() string {
	return filepath.Join(os.Getenv("HOME"), ".ssh", "config")
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
		config: config,
		Hosts:  hosts,
	}
	return
}

// LoadUserConfig bulabula
func LoadUserConfig() (cfg *Config, err error) {
	return LoadConfig(UserConfigFile())
}

// SaveConfig bulabula
func SaveConfig(file string, config *Config) (err error) {
	return config.Save(file)
}

// SaveUserConfig bulabula
func SaveUserConfig(config *Config) (err error) {
	return SaveConfig(UserConfigFile(), config)
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

// AddHostsFromJSON bulabula
func (c *Config) AddHostsFromJSON(jsonString string) (err error) {
	hosts, err := LoadHostsFromJSON(jsonString)
	if err != nil {
		return
	}
	for _, host := range hosts {
		err = c.AddHost(host)
		if err != nil {
			return
		}
	}
	return
}
