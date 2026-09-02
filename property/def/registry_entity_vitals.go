package def

import "github.com/ecumeurs/upsilontypes/property"

// entityVitalsEntries is the entity life-state pool: HP/SP/MP counters, the
// Shield/Poison/Stun accumulated-state trio, and the IsDying countdown flag.
// See architecture/property_key_vocabulary.md §3.
// @spec-link [[upsilontypes:module_property_key_registry]]
var entityVitalsEntries = map[property.Key]Entry{
	// Absence means 10/10.
	// DUAL-SCOPE Item|Entity (ISS-142, buffability ruling 2026-08-27, reconfirmed 2026-09-01): items
	// can grant entity-attribute buffs; equipment composes HP onto the entity via applyItemAsBuff.
	property.HP: {
		Key: property.HP, Scopes: ScopeItem | ScopeEntity, Kind: KindIntCounter, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return HP() },
	},
	// Absence means 10/10.
	// DUAL-SCOPE Item|Entity (ISS-142, buffability ruling 2026-08-27, reconfirmed 2026-09-01): see
	// HP above.
	property.SP: {
		Key: property.SP, Scopes: ScopeItem | ScopeEntity, Kind: KindIntCounter, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return SP() },
	},
	// Absence means 10/10.
	// DUAL-SCOPE Item|Entity (ISS-142, buffability ruling 2026-08-27, reconfirmed 2026-09-01): see
	// HP above.
	property.MP: {
		Key: property.MP, Scopes: ScopeItem | ScopeEntity, Kind: KindIntCounter, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return MP() },
	},
	property.Shield: { // Absence means 0/0; damage-absorbing counter
		Key: property.Shield, Scopes: ScopeEntity, Kind: KindIntCounter, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return Shield() },
	},
	property.Poison: { // Absence means 0; accumulated poison state
		Key: property.Poison, Scopes: ScopeEntity, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return Poison() },
	},
	property.Stun: { // Absence means 0; accumulated stun state
		Key: property.Stun, Scopes: ScopeEntity, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return Stun() },
	},
	property.IsDying: { // Absence means -1 (not dying)
		Key: property.IsDying, Scopes: ScopeEntity, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.Public, New: func() property.Property { return IsDying() },
	},
}
