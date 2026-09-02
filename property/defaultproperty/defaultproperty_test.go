// Package defaultproperty provides unit tests for the crash-early guards on
// the default Property constructors.
package defaultproperty

import (
	"testing"

	"github.com/ecumeurs/upsilontypes/property"
	"github.com/stretchr/testify/require"
)

// TestMakeIntCounterProperty_EmptyKeyPanics verifies that MakeIntCounterProperty
// panics on an empty property key instead of silently returning a nil Property
// (CODING_RULE §3: crash early, no undefined behavior far from its cause).
// @test-link [[upsilontypes:module_property_key_registry]]
func TestMakeIntCounterProperty_EmptyKeyPanics(t *testing.T) {
	require.Panics(t, func() {
		MakeIntCounterProperty(property.Key(""), 0, 0, property.Public, property.Skill)
	})
}

// TestMakeValidatedStringProperty_EmptyKeyPanics verifies that
// MakeValidatedStringProperty panics on an empty property key instead of
// silently returning a nil Property (CODING_RULE §3).
// @test-link [[upsilontypes:module_property_key_registry]]
func TestMakeValidatedStringProperty_EmptyKeyPanics(t *testing.T) {
	require.Panics(t, func() {
		MakeValidatedStringProperty(property.Key(""), "", property.Public, property.Skill, nil)
	})
}

// TestMakeStringProperty_EmptyKeyPanics verifies that MakeStringProperty, a
// thin pass-through to MakeValidatedStringProperty, inherits the same
// crash-early guard on an empty property key.
// @test-link [[upsilontypes:module_property_key_registry]]
func TestMakeStringProperty_EmptyKeyPanics(t *testing.T) {
	require.Panics(t, func() {
		MakeStringProperty(property.Key(""), "", property.Public, property.Skill)
	})
}
