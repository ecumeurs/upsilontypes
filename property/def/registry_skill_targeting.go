package def

import "github.com/ecumeurs/upsilontypes/property"

// skillTargetingEntries declares the Skill-vocabulary targeting keys from
// architecture/property_key_vocabulary.md §4a (rows 21-29).
// @spec-link [[upsilontypes:module_property_key_registry]]
var skillTargetingEntries = map[property.Key]Entry{
	property.Behavior: { // row 21: broad skill category, validated set
		Key: property.Behavior, Scopes: ScopeSkill, Kind: KindString, Composition: CompositionNone,
		MinInfoLevel: property.FriendlyController, Allowed: []string{"Direct", "Reaction", "Passive", "Counter", "Trap"},
		New: func() property.Property { return DefaultBehavior() },
	},
	property.Range: { // row 22: skill range, absence means 1/1
		Key: property.Range, Scopes: ScopeSkill, Kind: KindIntCounter, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return DefaultRange() },
	},
	property.Zone: { // row 23: area of effect; min info level blank in table, substituted Public (flag for ruling)
		Key: property.Zone, Scopes: ScopeSkill, Kind: KindZone, Composition: CompositionReplace,
		MinInfoLevel: property.Public, New: func() property.Property { return DefaultZone() },
	},
	property.TargetNumber: { // row 24: live gap closed by construction, absence means all in zone
		Key: property.TargetNumber, Scopes: ScopeSkill, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return TargetNumber() },
	},
	property.Accuracy: { // row 25: dual-scope (ISS-145 Defect 1), only declaring slice for this key
		Key: property.Accuracy, Scopes: ScopeSkill | ScopeEntity, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return Accuracy() },
	},
	property.Dodge: { // row 26: dual-scope (ISS-145 Defect 1), only declaring slice for this key
		Key: property.Dodge, Scopes: ScopeSkill | ScopeEntity, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return Dodge() },
	},
	property.Parry: { // row 27: stays skill-only, ISS-148/149 out of round
		Key: property.Parry, Scopes: ScopeSkill, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.FriendlyController, New: func() property.Property { return Parry() },
	},
	property.TargetType: { // row 28: validated set of targetable shapes
		Key: property.TargetType, Scopes: ScopeSkill, Kind: KindString, Composition: CompositionNone,
		MinInfoLevel: property.FriendlyController, Allowed: []string{"Entity", "FriendOnly", "EnemyOnly", "Tile", "EntityOrTile", "Self"},
		New: func() property.Property { return DefaultTargetType() },
	},
	property.TargetingMechanics: { // row 29: validated set, note the space in "Line of Sight"
		Key: property.TargetingMechanics, Scopes: ScopeSkill, Kind: KindString, Composition: CompositionNone,
		MinInfoLevel: property.FriendlyController, Allowed: []string{"Anywhere", "Line of Sight"},
		New: func() property.Property { return DefaultTargetingMechanics() },
	},
}
