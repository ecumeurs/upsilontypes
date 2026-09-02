package def

import "github.com/ecumeurs/upsilontypes/property"

// itemEntries declares the item keys from
// architecture/property_key_vocabulary.md §5 (rows 52-64).
// @spec-link [[upsilontypes:module_property_key_registry]]
var itemEntries = map[property.Key]Entry{
	property.Durability: { // row 52: absence means 0 (invulnerable)
		Key: property.Durability, Scopes: ScopeItem, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.Public, New: func() property.Property { return Durability() },
	},
	property.Weight: { // row 53: absence means 0
		Key: property.Weight, Scopes: ScopeItem, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.Public, New: func() property.Property { return Weight() },
	},
	property.ItemType: { // row 54: validated string, absence means Misc
		Key: property.ItemType, Scopes: ScopeItem, Kind: KindString, Composition: CompositionNone,
		MinInfoLevel: property.OwnController,
		Allowed: []string{
			string(ItemTypeWearable), string(ItemTypeConsumable), string(ItemTypeUsable),
			string(ItemTypeThrowable), string(ItemTypeAmmunitions), string(ItemTypeNone),
		},
		New: func() property.Property { return DefaultItemType() },
	},
	// row 55: absence means 0; VALUE RENAMED from "Armor" (vocabulary §7).
	// DUAL-SCOPE Item|Entity: authored on items, but legitimately read off an Entity because
	// applyItemAsBuff flattens equipped items into Forever:true entity buffs (equipment is not a
	// live layer). Corrected 2026-08-31 in step 13 after the scope guard caught the real read at
	// effectapplicator.go:134 and ruler/rules/attack.go:59. See vocabulary §2 correction.
	property.ArmorRating: {
		Key: property.ArmorRating, Scopes: ScopeItem | ScopeEntity, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.Public, New: func() property.Property { return ArmorRating() },
	},
	property.WeaponType: { // row 56: validated string, absence means None
		Key: property.WeaponType, Scopes: ScopeItem, Kind: KindString, Composition: CompositionNone,
		MinInfoLevel: property.Public,
		Allowed: []string{
			string(NoWeapon), string(OneHandedMelee), string(TwoHandedMelee),
			string(OneHandedRanged), string(TwoHandedRanged),
		},
		New: func() property.Property { return DefaultWeaponTypeProperty() },
	},
	property.ArmorType: { // row 57: validated string, absence means None
		Key: property.ArmorType, Scopes: ScopeItem, Kind: KindString, Composition: CompositionNone,
		MinInfoLevel: property.Public,
		Allowed: []string{
			string(NoArmor), string(HeadSlot), string(BodySlot), string(HandsSlot),
			string(LegsSlot), string(FeetSlot), string(BeltSlot), string(NeckSlot), string(RingSlot),
		},
		New: func() property.Property { return DefaultArmorTypeProperty() },
	},
	property.ToolType: { // row 58: validated string, absence means None
		Key: property.ToolType, Scopes: ScopeItem, Kind: KindString, Composition: CompositionNone,
		MinInfoLevel: property.Public,
		Allowed:      []string{string(NoTool), string(SomeTool)},
		New:          func() property.Property { return DefaultToolTypeProperty() },
	},
	// row 59: absence means 0.
	// DUAL-SCOPE Item|Entity (corrected 2026-08-31, step 13): equipment-authored combat stat
	// composed onto the entity by applyItemAsBuff; read off the Entity at attack_checks.go:76.
	property.WeaponRange: {
		Key: property.WeaponRange, Scopes: ScopeItem | ScopeEntity, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.Public, New: func() property.Property { return WeaponRange() },
	},
	// row 60: absence means 0.
	// DUAL-SCOPE Item|Entity (corrected 2026-08-31, step 13): equipment-authored combat stat
	// composed onto the entity by applyItemAsBuff; read off the Entity at attack.go:49.
	property.WeaponBaseDamage: {
		Key: property.WeaponBaseDamage, Scopes: ScopeItem | ScopeEntity, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.Public, New: func() property.Property { return WeaponBaseDamage() },
	},
	property.Stackable: { // row 61: absence means false
		Key: property.Stackable, Scopes: ScopeItem, Kind: KindBool, Composition: CompositionAnd,
		MinInfoLevel: property.Public, New: func() property.Property { return Stackable() },
	},
	property.StackSize: { // row 62: absence means 0
		Key: property.StackSize, Scopes: ScopeItem, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.Public, New: func() property.Property { return StackSize() },
	},
	property.Effect: { // row 63: absence means nil; DefaultEffectProperty() already wraps a nil Effect
		Key: property.Effect, Scopes: ScopeItem, Kind: KindEffect, Composition: CompositionNone,
		MinInfoLevel: property.Analyser, New: func() property.Property { return DefaultEffectProperty() },
	},
	property.ItemValue: { // row 64: absence means 0; RENAMED identifier from Value (vocabulary §7); constructor renamed to match (slice 14B)
		Key: property.ItemValue, Scopes: ScopeItem, Kind: KindInt, Composition: CompositionAdd,
		MinInfoLevel: property.Public, New: func() property.Property { return ItemValue() },
	},
}
