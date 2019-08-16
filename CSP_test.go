package csp_test

import (
	"testing"

	"github.com/aerogo/csp"
	"github.com/akyoto/assert"
)

func TestCSP(t *testing.T) {
	policy := csp.New()
	assert.Equal(t, policy.String(), "")

	// SetMap
	policy.SetMap(csp.Map{
		"image-src": "https:",
	})
	assert.Equal(t, "https:", policy.Get("image-src"))

	// Set
	policy.Set("image-src", "'self'")
	assert.Equal(t, "'self'", policy.Get("image-src"))

	// String
	assert.Equal(t, "image-src 'self';", policy.String())
}
