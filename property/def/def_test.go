// Package def provides unit tests for the property definition constructors.
// @test-link [[mech_actor_pattern]]
package def

import (
	"testing"
)

// TestHP verifies that the HP property constructor returns a non-nil property.
func TestHP(t *testing.T) {
	skp := HP()
	if skp == nil {
		t.Error("HP() should not return nil")
	}
}

// TestZoneProperty_Single verifies that "Single" produces a single-cell pattern.
func TestZoneProperty_Single(t *testing.T) {
	z := DefaultZone()
	z.Set("Single")
	if z.PatternType != "Single" {
		t.Errorf("expected PatternType Single, got %s", z.PatternType)
	}
	if len(z.ZonePattern) != 1 {
		t.Errorf("Single pattern must contain exactly 1 cell, got %d", len(z.ZonePattern))
	}
}

// TestZoneProperty_Neighbours verifies that "Neighbours" produces a 3x3x3 cube (27 cells).
func TestZoneProperty_Neighbours(t *testing.T) {
	z := DefaultZone()
	z.Set("Neighbours")
	if z.PatternType != "Neighbours" {
		t.Errorf("expected PatternType Neighbours, got %s", z.PatternType)
	}
	// Neighbours() = Square(1,1,1) => (2*1+1)^3 = 27 cells
	if len(z.ZonePattern) != 27 {
		t.Errorf("Neighbours pattern must contain 27 cells, got %d", len(z.ZonePattern))
	}
}

// TestZoneProperty_Circle verifies that "Circle:N" produces a spherical pattern.
// A Circle(radius) with radius=1 must include the origin and its 6 face-adjacent cells
// plus the 12 edge-adjacent cells that still satisfy x²+y²+z² <= 1 — actually exactly 7 cells.
func TestZoneProperty_Circle(t *testing.T) {
	z := DefaultZone()
	z.Set("Circle:1")
	if z.PatternType != "Circle:1" {
		t.Errorf("expected PatternType Circle:1, got %s", z.PatternType)
	}
	// Circle(1): only (0,0,0) and the 6 pure face-adjacent cells satisfy x²+y²+z²<=1.
	if len(z.ZonePattern) != 7 {
		t.Errorf("Circle:1 pattern must contain 7 cells, got %d", len(z.ZonePattern))
	}

	// Circle:2 test — must be strictly larger than Circle:1
	z2 := DefaultZone()
	z2.Set("Circle:2")
	if len(z2.ZonePattern) <= len(z.ZonePattern) {
		t.Errorf("Circle:2 must be larger than Circle:1; got %d vs %d", len(z2.ZonePattern), len(z.ZonePattern))
	}
}

// TestZoneProperty_Square verifies that "Square:N" produces a symmetric cubic pattern.
// Square:1 => Square(1,1,1) = (2*1+1)^3 = 27 cells.
// Square:2 => Square(2,2,2) = (2*2+1)^3 = 125 cells.
func TestZoneProperty_Square(t *testing.T) {
	z1 := DefaultZone()
	z1.Set("Square:1")
	if z1.PatternType != "Square:1" {
		t.Errorf("expected PatternType Square:1, got %s", z1.PatternType)
	}
	if len(z1.ZonePattern) != 27 {
		t.Errorf("Square:1 pattern must contain 27 cells, got %d", len(z1.ZonePattern))
	}

	z2 := DefaultZone()
	z2.Set("Square:2")
	if len(z2.ZonePattern) != 125 {
		t.Errorf("Square:2 pattern must contain 125 cells, got %d", len(z2.ZonePattern))
	}
}

// TestZoneProperty_Line verifies that "Line:N" produces a 1D line pattern of N cells.
func TestZoneProperty_Line(t *testing.T) {
	z := DefaultZone()
	z.Set("Line:3")
	if z.PatternType != "Line:3" {
		t.Errorf("expected PatternType Line:3, got %s", z.PatternType)
	}
	if len(z.ZonePattern) != 3 {
		t.Errorf("Line:3 pattern must contain 3 cells, got %d", len(z.ZonePattern))
	}
}

// TestZoneProperty_UnknownPatternPanics verifies Crash-Early behaviour:
// an unrecognised pattern string must panic, NOT silently fall back to Single.
func TestZoneProperty_UnknownPatternPanics(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic for unknown pattern, but none occurred — silent fallback is not allowed")
		}
	}()
	z := DefaultZone()
	// "Cone:3" is not a recognised pattern; the engine must reject it loudly.
	z.Set("Cone:3")
}

// TestZoneProperty_MalformedParamPanics verifies Crash-Early for a parameterised pattern
// that is missing its required integer argument.
func TestZoneProperty_MalformedParamPanics(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic for malformed Circle pattern, but none occurred")
		}
	}()
	z := DefaultZone()
	z.Set("Circle")
}

// TestZoneProperty_InvalidIntParamPanics verifies that a non-integer parameter panics.
func TestZoneProperty_InvalidIntParamPanics(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic for Circle:abc, but none occurred")
		}
	}()
	z := DefaultZone()
	z.Set("Circle:abc")
}

// TestZoneProperty_ZeroParamPanics verifies that a zero or negative parameter panics.
func TestZoneProperty_ZeroParamPanics(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic for Square:0, but none occurred")
		}
	}()
	z := DefaultZone()
	z.Set("Square:0")
}
