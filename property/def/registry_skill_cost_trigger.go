package def

import "github.com/ecumeurs/upsilontypes/property"

// skillCostTriggerEntries declares the Skill-vocabulary Cost (§4c, rows 40-46),
// Reposition (§4d, rows 47-48) and Trigger (§4e, rows 49-51) keys.
// @spec-link [[upsilontypes:module_property_key_registry]]
var skillCostTriggerEntries = map[property.Key]Entry{
	property.Delay: { // row 40: cast delay counter, absence means 0/500
		Key: property.Delay, Scopes: ScopeSkill, Kind: KindIntCounter, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return Delay() },
	},
	property.Channeling: { // row 41: channeling counter, absence means 0/0
		Key: property.Channeling, Scopes: ScopeSkill, Kind: KindIntCounter, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return Channeling() },
	},
	property.HPLeech: { // row 42: absence means 0
		Key: property.HPLeech, Scopes: ScopeSkill, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return HPLeech() },
	},
	property.MPLeech: { // row 43: absence means 0
		Key: property.MPLeech, Scopes: ScopeSkill, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return MPLeech() },
	},
	property.SPLeech: { // row 44: absence means 0
		Key: property.SPLeech, Scopes: ScopeSkill, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return SPLeech() },
	},
	property.MovementCost: { // row 45: current identifier is MovementCost; constructor renamed to match (slice 14B)
		Key: property.MovementCost, Scopes: ScopeSkill, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return MovementCost() },
	},
	property.Cooldown: { // row 46: absence means 0/3; Value = remaining at battle start, MaxValue = applied on use
		Key: property.Cooldown, Scopes: ScopeSkill, Kind: KindIntCounter, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return Cooldown() },
	},
	property.RepositionSubject: { // row 47: who a movement skill displaces, validated set
		Key: property.RepositionSubject, Scopes: ScopeSkill, Kind: KindString, Composition: CompositionNone,
		MinInfoLevel: property.FriendlyController, Allowed: []string{"Self", "Target"},
		New: func() property.Property { return RepositionSubject(RepositionSubjectSelf) },
	},
	property.RepositionDistance: { // row 48: signed tiles along the caster->target ray, absence means 0
		Key: property.RepositionDistance, Scopes: ScopeSkill, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return RepositionDistance(0) },
	},
	property.TriggerType: { // row 49: when a positional effect fires, validated set
		Key: property.TriggerType, Scopes: ScopeSkill, Kind: KindString, Composition: CompositionNone,
		MinInfoLevel: property.FriendlyController, Allowed: []string{"OnEnter", "OnExit", "OnStep", "OnTurn", "OnDeath"},
		New: func() property.Property { return TriggerType(property.TriggerOnEnter) },
	},
	property.RemoveOnTrigger: { // row 50: consumed after firing once, absence means true
		Key: property.RemoveOnTrigger, Scopes: ScopeSkill, Kind: KindBool, Composition: CompositionAnd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return RemoveOnTrigger(true) },
	},
	property.TriggerCount: { // row 51: fire count, absence means 1 (0 = unlimited)
		Key: property.TriggerCount, Scopes: ScopeSkill, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return TriggerCount(1) },
	},
}
