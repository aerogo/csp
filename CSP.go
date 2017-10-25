package csp

import "bytes"

// ContentSecurityPolicy ...
type ContentSecurityPolicy struct {
	policies map[string]string
	cache    string
}

// New ...
func New() *ContentSecurityPolicy {
	return &ContentSecurityPolicy{
		policies: make(map[string]string),
	}
}

// Get ...
func (csp *ContentSecurityPolicy) Get(key string, value string) string {
	return csp.policies[key]
}

// Set ...
func (csp *ContentSecurityPolicy) Set(key string, value string) {
	csp.policies[key] = value
	csp.updateCache()
}

// SetMap ...
func (csp *ContentSecurityPolicy) SetMap(m map[string]string) {
	csp.policies = m
	csp.updateCache()
}

// String ...
func (csp *ContentSecurityPolicy) String() string {
	return csp.cache
}

// updateCache ...
func (csp *ContentSecurityPolicy) updateCache() {
	buffer := bytes.Buffer{}

	for key, value := range csp.policies {
		buffer.WriteString(key)
		buffer.WriteByte(' ')
		buffer.WriteString(value)
		buffer.WriteByte(';')
	}

	csp.cache = buffer.String()
}
