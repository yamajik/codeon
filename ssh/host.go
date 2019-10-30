package ssh

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/kevinburke/ssh_config"
)

// Host bulabula
type Host struct {
	Patterns []*Pattern
	Mappings []*Mapping
}

// LoadHost bulabula
func LoadHost(host *ssh_config.Host) (h *Host, err error) {
	if host == nil {
		err = errors.New("Host is nil")
	}

	var mappings []*Mapping
	for _, node := range host.Nodes {
		switch t := node.(type) {
		case *ssh_config.Empty:
			continue
		case *ssh_config.Include:
			fmt.Println("Not support include: ignored.")
			continue
		case *ssh_config.KV:
			lkey := strings.ToLower(t.Key)
			if lkey == "match" {
				err = errors.New("Can't handle Match directives")
				return
			}
			mappings = append(mappings, t)
		}
	}

	var patterns []*Pattern
	for _, p := range host.Patterns {
		pattern, newPatternErr := NewPattern(p.String())
		if newPatternErr != nil {
			err = newPatternErr
			return
		}
		patterns = append(patterns, pattern)
	}

	h = &Host{
		Patterns: patterns,
		Mappings: mappings,
	}
	return
}

// NewHost bulabula
func NewHost(patterns []string, mappings map[string]string) (h *Host, err error) {
	var hostPatterns []*Pattern
	for _, p := range patterns {
		pattern, newPatternErr := NewPattern(p)
		if newPatternErr != nil {
			err = newPatternErr
			return
		}
		hostPatterns = append(hostPatterns, pattern)
	}

	var hostMappings []*Mapping
	for k, v := range mappings {
		hostMappings = append(hostMappings, NewMapping(k, v))
	}

	h = &Host{
		Patterns: hostPatterns,
		Mappings: hostMappings,
	}
	return
}

// NewDefaultHost bulabula
func NewDefaultHost() (h *Host, err error) {
	return NewHost([]string{"*"}, map[string]string{})
}

// HostJSONStruct bulabula
type HostJSONStruct struct {
	Patterns []string
	Mappings map[string]string
}

// LoadHostFromStruct bulabula
func LoadHostFromStruct(hostStruct HostJSONStruct) (host *Host, err error) {
	return NewHost(hostStruct.Patterns, hostStruct.Mappings)
}

// LoadHostsFromStruct bulabula
func LoadHostsFromStruct(hostStructs []HostJSONStruct) (hosts []*Host, err error) {
	for _, hostStruct := range hostStructs {
		host, newHostErr := NewHost(hostStruct.Patterns, hostStruct.Mappings)
		if newHostErr != nil {
			err = newHostErr
			return
		}
		hosts = append(hosts, host)
	}
	return
}

// LoadHostFromJSON bulabula
func LoadHostFromJSON(jsonString string) (host *Host, err error) {
	hostStruct := HostJSONStruct{}
	err = json.Unmarshal([]byte(jsonString), &hostStruct)
	if err != nil {
		return
	}
	return LoadHostFromStruct(hostStruct)
}

// LoadHostsFromJSON bulabula
func LoadHostsFromJSON(jsonString string) (hosts []*Host, err error) {
	hostStructs := []HostJSONStruct{}
	err = json.Unmarshal([]byte(jsonString), &hostStructs)
	if err != nil {
		return
	}
	return LoadHostsFromStruct(hostStructs)
}

// String bulabula
func (h *Host) String() string {
	patternStrings := []string{"Host"}
	for _, p := range h.Patterns {
		patternStrings = append(patternStrings, p.String())
	}
	patternString := strings.Join(patternStrings, " ")

	mappingStrings := []string{}
	for _, m := range h.Mappings {
		mappingString := "    " + strings.Join([]string{m.Key, m.Value}, " ")
		mappingStrings = append(mappingStrings, mappingString)
	}

	hostStrings := append([]string{patternString}, mappingStrings...)
	return strings.Join(hostStrings, "\n")
}

// Match bulabula
func (h *Host) Match(name string) bool {
	return h.MatchPattern(name, false) != nil
}

// FindPattern bulabula
func (h *Host) FindPattern(name string) *Pattern {
	return h.MatchPattern(name, true)
}

// MatchPattern bulabula
func (h *Host) MatchPattern(name string, strict bool) (pattern *Pattern) {
	for _, p := range h.Patterns {
		if p.Match(name, strict) {
			pattern = p
			return
		}
	}
	return
}

// AddPattern bulabula
func (h *Host) AddPattern(pattern string) (err error) {
	p, err := NewPattern(pattern)
	h.Patterns = append(h.Patterns, p)
	return
}

// RemovePattern bulabula
func (h *Host) RemovePattern(pattern string) (err error) {
	p, err := NewPattern(pattern)
	h.Patterns = append(h.Patterns, p)
	return
}

// FindMapping bulabula
func (h *Host) FindMapping(key string) (m *Mapping) {
	for _, m := range h.Mappings {
		if m.Key == key {
			return m
		}
	}
	return
}

// Update bulabula
func (h *Host) Update(host *Host) {
	for _, pattern := range host.Patterns {
		if h.FindPattern(pattern.String()) == nil {
			h.Patterns = append(h.Patterns, pattern)
		}
	}
	for _, mapping := range host.Mappings {
		hostMapping := h.FindMapping(mapping.Key)
		if hostMapping != nil {
			hostMapping.Key = mapping.Key
			hostMapping.Value = mapping.Value
			hostMapping.Comment = mapping.Comment
		} else {
			h.Mappings = append(h.Mappings, mapping)
		}
	}
}
