package def

import "github.com/ecumeurs/upsilontypes/property"

// DefaultProperty resolves the default Property for k directly from the
// property key registry, with no scope filter. Returns nil for an unknown
// key (crash-early: no invented fallback).
// @spec-link [[upsilontypes:module_property_key_registry]]
func DefaultProperty(k property.Key) property.Property {
	entry, ok := Lookup(k)
	if !ok {
		return nil
	}
	return entry.New()
}
