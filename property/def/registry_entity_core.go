package def

import "github.com/ecumeurs/upsilontypes/property"

// entityCoreEntries declares the combat-basics and entity identity/lifetime keys
// from architecture/property_key_vocabulary.md §3 (rows 5, 6, 8, 12-14, 16, 17).
// @spec-link [[upsilontypes:module_property_key_registry]]
var entityCoreEntries = map[property.Key]Entry{
	// row 5: basic attack, absence means 1.
	// DUAL-SCOPE Item|Entity (ISS-142, buffability ruling 2026-08-27, reconfirmed 2026-09-01): items
	// can grant entity-attribute buffs; equipment composes Attack onto the entity via applyItemAsBuff.
	property.Attack: {
		Key: property.Attack, Scopes: ScopeItem | ScopeEntity, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return Attack() },
	},
	// row 6: basic defense, absence means 0.
	// DUAL-SCOPE Item|Entity (ISS-142, buffability ruling 2026-08-27, reconfirmed 2026-09-01): see
	// Attack above.
	property.Defense: {
		Key: property.Defense, Scopes: ScopeItem | ScopeEntity, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return Defense() },
	},
	// row 12: basic attack range, absence means 1.
	// DUAL-SCOPE Item|Entity (ISS-142, buffability ruling 2026-08-27, reconfirmed 2026-09-01): see
	// Attack above.
	property.AttackRange: {
		Key: property.AttackRange, Scopes: ScopeItem | ScopeEntity, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return AttackRange() },
	},
	property.TeamID: { // row 8: team affiliation; Add composition declared as observed (vocabulary §8.5)
		Key: property.TeamID, Scopes: ScopeEntity, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.Public, New: func() property.Property { return TeamID() },
	},
	property.AIArchetype: { // row 17: free-form archetype slug, not a validated set
		Key: property.AIArchetype, Scopes: ScopeEntity, Kind: KindString, Composition: CompositionNone,
		MinInfoLevel: property.Public, New: func() property.Property { return AIArchetype() },
	},
	property.EntityDuration: { // row 13: temporary-entity lifetime counter, distinct from skill Duration
		Key: property.EntityDuration, Scopes: ScopeEntity, Kind: KindIntCounter, Composition: CompositionAdd,
		MinInfoLevel: property.Public, New: func() property.Property { return EntityDuration() },
	},
	property.ExpiresWithCaster: { // row 14: removed with owned positional effects when caster dies
		Key: property.ExpiresWithCaster, Scopes: ScopeEntity, Kind: KindBool, Composition: CompositionAnd,
		MinInfoLevel: property.Public, New: func() property.Property { return ExpiresWithCaster() },
	},
	property.Invisible: { // row 16: not sent in client-facing state snapshots
		Key: property.Invisible, Scopes: ScopeEntity, Kind: KindBool, Composition: CompositionAnd,
		MinInfoLevel: property.Public, New: func() property.Property { return Invisible() },
	},
}
