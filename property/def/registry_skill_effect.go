package def

import "github.com/ecumeurs/upsilontypes/property"

// skillEffectEntries declares the skill Effect keys from
// architecture/property_key_vocabulary.md §4b (rows 30-39).
// @spec-link [[upsilontypes:module_property_key_registry]]
var skillEffectEntries = map[property.Key]Entry{
	property.DamageScale: { // row 30: percentage scaling of attack, absence means 100%
		Key: property.DamageScale, Scopes: ScopeSkill, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return DamageScale() },
	},
	property.Heal: { // row 31: absence means 0
		Key: property.Heal, Scopes: ScopeSkill, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return Heal() },
	},
	property.ShieldPower: { // row 32: skill-scope shield power, distinct from entity Shield
		Key: property.ShieldPower, Scopes: ScopeSkill, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return ShieldPower() },
	},
	property.StunPower: { // row 33: skill-scope stun power, distinct from entity Stun
		Key: property.StunPower, Scopes: ScopeSkill, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return StunPower() },
	},
	property.StunChance: { // row 34: percent chance, absence means 0
		Key: property.StunChance, Scopes: ScopeSkill, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return StunChance() },
	},
	property.CriticalChance: { // row 35: dual-scope (Skill|Entity), ISS-145 Defect 1
		Key: property.CriticalChance, Scopes: ScopeSkill | ScopeEntity, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return CriticalChance() },
	},
	property.CriticalMultiplier: { // row 36: dual-scope (Skill|Entity), absence means 100%
		Key: property.CriticalMultiplier, Scopes: ScopeSkill | ScopeEntity, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return CriticalMultiplier() },
	},
	property.Duration: { // row 37: buff duration, distinct from entity EntityDuration
		Key: property.Duration, Scopes: ScopeSkill, Kind: KindIntCounter, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return Duration() },
	},
	property.PoisonPower: { // row 38: skill-scope poison power, distinct from entity Poison
		Key: property.PoisonPower, Scopes: ScopeSkill, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return PoisonPower() },
	},
	property.PoisonChance: { // row 39: percent chance, absence means 0
		Key: property.PoisonChance, Scopes: ScopeSkill, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return PoisonChance() },
	},
}
