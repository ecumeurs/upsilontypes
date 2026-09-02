// Package def provides a conformance test pinning ISS-142's buffability
// ruling into the property registry's scope bits.
package def

import (
	"testing"

	"github.com/ecumeurs/upsilontypes/property"
	"github.com/stretchr/testify/require"
)

// buffableEntityAttributes is the 8-key set ISS-142's buffability ruling
// (user, 2026-08-27, reconfirmed 2026-09-01) declares item-grantable: the
// true combat attributes (Attack, Defense, AttackRange, Movement,
// JumpHeight) plus the three resources (HP, SP, MP).
var buffableEntityAttributes = []property.Key{
	property.Attack, property.Defense, property.AttackRange,
	property.Movement, property.JumpHeight,
	property.HP, property.SP, property.MP,
}

// nonBuffableEntityStatuses is the Shield/Poison/Stun trio ISS-142 ruled OUT
// of the buffable set: effectapplicator already applies them directly via
// ShieldPower/PoisonPower/StunPower, so a parallel item-buff path would
// reproduce the ISS-147 escalation.
var nonBuffableEntityStatuses = []property.Key{property.Shield, property.Poison, property.Stun}

// nonBuffableEntityFlags is the 9-key flag/plumbing set ISS-142 ruled OUT of
// the buffable set: these are not attributes, and letting an item grant them
// would be an exploit, not a feature.
var nonBuffableEntityFlags = []property.Key{
	property.TeamID, property.IsDying, property.HasMoved, property.HasActed,
	property.EntityDuration, property.ExpiresWithCaster, property.WalkThrough,
	property.Invisible, property.AIArchetype,
}

// TestBuffability_ExactlyEightEntityAttributesAreItemGrantable pins ISS-142's
// buffability ruling (user, 2026-08-27, reconfirmed 2026-09-01): among the 20
// keys declared in the ENTITY theme files (registry_entity_core.go,
// registry_entity_movement.go, registry_entity_vitals.go — the "ENTITY scope
// (20)" table in architecture/property_key_vocabulary.md §3), exactly the 8
// combat-attribute and resource keys carry the dual ScopeItem|ScopeEntity
// bitmask, so an equipped item can grant them as a buff (see upsilonapi/bridge
// TestArenaInit_EquippedItemsBecomeBuffs). No other entity-declared key may
// drift into item-buffability without this test going red.
//
// Scoped to the entity theme files rather than the full merged registry
// deliberately: three ITEM-declared keys (ArmorRating, WeaponRange,
// WeaponBaseDamage in registry_item.go) are ALSO ScopeItem|ScopeEntity, but
// that is the pre-existing, unrelated step-13 "equipment composes onto the
// entity" correction — not part of ISS-142's buffability ruling — so they
// are intentionally excluded from this partition.
// @test-link [[upsilontypes:module_property_key_registry]]
func TestBuffability_ExactlyEightEntityAttributesAreItemGrantable(t *testing.T) {
	// Setup: union the three entity theme maps and collect every key among
	// them that also carries ScopeItem.
	entityDeclaredEntries := mergeThemes(entityCoreEntries, entityMovementEntries, entityVitalsEntries)
	var dualScopeKeys []property.Key
	for k, entry := range entityDeclaredEntries {
		if entry.Scopes&ScopeItem != 0 {
			dualScopeKeys = append(dualScopeKeys, k)
		}
	}

	// Execution: build want/got sets as key->bool maps so map iteration order
	// never affects the comparison.
	want := map[property.Key]bool{}
	for _, k := range buffableEntityAttributes {
		want[k] = true
	}
	got := map[property.Key]bool{}
	for _, k := range dualScopeKeys {
		got[k] = true
	}

	// Validation: the two sets must match exactly. A miss means ISS-142's
	// ruling silently stopped granting a buffable attribute; an extra means
	// an unruled key started allowing item-authored buffs.
	require.Len(t, got, len(want), "expected exactly %d entity-scoped keys to be item-buffable "+
		"per ISS-142's 2026-08-27 ruling (reconfirmed 2026-09-01), found %d: %v",
		len(want), len(got), dualScopeKeys)
	for _, k := range buffableEntityAttributes {
		require.True(t, got[k], "%q must be ScopeItem|ScopeEntity: ISS-142 rules it a buffable "+
			"combat attribute or resource, item-grantable via applyItemAsBuff", k)
	}
	for k := range got {
		require.True(t, want[k], "%q is item-buffable but is not one of the 8 keys ISS-142's "+
			"ruling names as buffable — an unruled key must not silently gain item scope", k)
	}
}

// TestBuffability_StatusesStayEntityOnly pins ISS-142/ISS-147: Shield,
// Poison and Stun must remain ScopeEntity only. ISS-147 was exactly this
// defect — an item carrying the key "Poison" fell through the bridge's
// ITEM-scope lookup into an ENTITY-scope lookup and injected a real Poison
// status effect. Dual-scoping these keys would silently reopen that hole.
// @test-link [[upsilontypes:module_property_key_registry]]
func TestBuffability_StatusesStayEntityOnly(t *testing.T) {
	for _, k := range nonBuffableEntityStatuses {
		// Setup/Execution: resolve the registry entry for the status key.
		entry, ok := Lookup(k)
		require.True(t, ok, "%q must be a registered key", k)

		// Validation: entity scope present, item scope absent.
		require.NotZero(t, entry.Scopes&ScopeEntity, "%q must be ScopeEntity", k)
		require.Zero(t, entry.Scopes&ScopeItem, "%q must NOT be ScopeItem (ISS-147 regression "+
			"guard): an item carrying this key must fail its ITEM-scope lookup rather than fall "+
			"through and inject a real entity status effect", k)
	}
}

// TestBuffability_FlagsAndPlumbingStayEntityOnly pins ISS-142: the 9
// flag/plumbing entity keys must remain ScopeEntity only. These are not
// combat attributes; letting an item grant them would be an exploit — e.g.
// an item silently flipping TeamID or Invisible — not a feature.
// @test-link [[upsilontypes:module_property_key_registry]]
func TestBuffability_FlagsAndPlumbingStayEntityOnly(t *testing.T) {
	for _, k := range nonBuffableEntityFlags {
		// Setup/Execution: resolve the registry entry for the flag/plumbing key.
		entry, ok := Lookup(k)
		require.True(t, ok, "%q must be a registered key", k)

		// Validation: entity scope present, item scope absent.
		require.NotZero(t, entry.Scopes&ScopeEntity, "%q must be ScopeEntity", k)
		require.Zero(t, entry.Scopes&ScopeItem, "%q must NOT be ScopeItem: buffing %q from an "+
			"item would be an exploit (e.g. TeamID or Invisible), not a feature — ISS-142 ruled "+
			"flags/plumbing out of the buffable set", k, k)
	}
}
