package def

import "github.com/ecumeurs/upsilontypes/property"

// entityMovementEntries is the entity movement and turn-state pool: Movement,
// JumpHeight, WalkThrough and the per-turn HasMoved/HasActed flags. See
// architecture/property_key_vocabulary.md §3.
// @spec-link [[upsilontypes:module_property_key_registry]]
var entityMovementEntries = map[property.Key]Entry{
	// Absence means 3/3 (constructor returns 5/5 today; ISS defect, fix deferred).
	// DUAL-SCOPE Item|Entity (ISS-142, buffability ruling 2026-08-27, reconfirmed 2026-09-01): items
	// can grant entity-attribute buffs; equipment composes Movement onto the entity via applyItemAsBuff.
	property.Movement: {
		Key: property.Movement, Scopes: ScopeItem | ScopeEntity, Kind: KindIntCounter, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return Movement() },
	},
	// Absence means 2.
	// DUAL-SCOPE Item|Entity (ISS-142, buffability ruling 2026-08-27, reconfirmed 2026-09-01): see
	// Movement above.
	property.JumpHeight: {
		Key: property.JumpHeight, Scopes: ScopeItem | ScopeEntity, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return JumpHeight() },
	},
	property.WalkThrough: { // Absence means false
		Key: property.WalkThrough, Scopes: ScopeEntity, Kind: KindBool, Composition: CompositionAnd,
		MinInfoLevel: property.Public, New: func() property.Property { return WalkThrough() },
	},
	property.HasMoved: { // Absence means false
		Key: property.HasMoved, Scopes: ScopeEntity, Kind: KindBool, Composition: CompositionAnd,
		MinInfoLevel: property.GameMaster, New: func() property.Property { return HasMoved() },
	},
	property.HasActed: { // Absence means false
		Key: property.HasActed, Scopes: ScopeEntity, Kind: KindBool, Composition: CompositionAnd,
		MinInfoLevel: property.GameMaster, New: func() property.Property { return HasActed() },
	},
}
