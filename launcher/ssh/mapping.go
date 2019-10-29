package ssh

import "github.com/kevinburke/ssh_config"

// Mapping bulabula
type Mapping = ssh_config.KV

// NewMapping bulabula
func NewMapping(key string, value string) *Mapping {
	return &Mapping{
		Key:   key,
		Value: value,
	}
}
