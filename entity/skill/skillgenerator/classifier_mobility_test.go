package skillgenerator

import (
	"testing"

	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property/def"
)

// TestClassifyRepositionIsMobility verifies that any skill which displaces a subject
// (RepositionDistance != 0) is tagged "mobility", regardless of subject or target type.
// @test-link [[mech_movement_reposition]]
func TestClassifyRepositionIsMobility(t *testing.T) {
	cases := []struct {
		name    string
		subject def.RepositionSubjectType
		dist    int
	}{
		{"self dash", def.RepositionSubjectSelf, 3},
		{"target push", def.RepositionSubjectTarget, 2},
		{"target pull", def.RepositionSubjectTarget, -2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			sk := skill.New()
			sk.Effect.Properties = append(sk.Effect.Properties,
				def.RepositionSubject(tc.subject),
				def.RepositionDistance(tc.dist),
			)

			tags := Classify(sk)
			if !containsTag(tags, "mobility") {
				t.Errorf("expected 'mobility' tag for %s, got %v", tc.name, tags)
			}
		})
	}
}

// TestClassifyNoRepositionNotMobility verifies a plain skill is not tagged mobility.
// @test-link [[mech_movement_reposition]]
func TestClassifyNoRepositionNotMobility(t *testing.T) {
	sk := skill.New()
	if containsTag(Classify(sk), "mobility") {
		t.Errorf("plain skill should not be tagged mobility")
	}
}

func containsTag(tags []string, want string) bool {
	for _, tag := range tags {
		if tag == want {
			return true
		}
	}
	return false
}
