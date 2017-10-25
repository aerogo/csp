package csp

import "bytes"

// ContentSecurityPolicy ...
type ContentSecurityPolicy struct {
	policies map[string]string
	cache    string
}

// New creates a new content security policy.
func New() *ContentSecurityPolicy {
	return &ContentSecurityPolicy{
		policies: make(map[string]string),
	}
}

// Get retrieves a policy by key.
func (csp *ContentSecurityPolicy) Get(key string, value string) string {
	return csp.policies[key]
}

// Set sets a policy by key.
func (csp *ContentSecurityPolicy) Set(key string, value string) {
	csp.policies[key] = value
	csp.updateCache()
}

// SetMap overwrites all policies with the given map.
func (csp *ContentSecurityPolicy) SetMap(m map[string]string) {
	csp.policies = m
	csp.updateCache()
}

// String returns the actual string representation of the CSP header value.
func (csp *ContentSecurityPolicy) String() string {
	return csp.cache
}

// updateCache updates the internal cache.
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
