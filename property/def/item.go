package def

import (
	"github.com/ecumeurs/upsilontypes/property"
	"github.com/ecumeurs/upsilontypes/property/defaultproperty"
	"github.com/ecumeurs/upsilontypes/property/effect"
)

func Durability() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.Durability, 0, property.Public, property.Item)
}

func Weight() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.Weight, 0, property.Public, property.Item)
}

func ArmorRating() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.ArmorRating, 0, property.Public, property.Item)
}

func WeaponRange() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.WeaponRange, 0, property.Public, property.Item)
}

func WeaponBaseDamage() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.WeaponBaseDamage, 0, property.Public, property.Item)
}

func Stackable() *defaultproperty.DefaultBoolProperty {
	return defaultproperty.MakeBoolProperty(property.Stackable, false, property.Public, property.Item)
}

func StackSize() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.StackSize, 0, property.Public, property.Item)
}

// ItemValue builds the item's monetary value property; absence means 0.
func ItemValue() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.ItemValue, 0, property.Public, property.Item)
}

// ItemType         Key = "ItemType"         // Absence means None (out of Wearable, Consumable, Usable, Throwable, Ammunitions and None)

type ItemTypes string

const (
	ItemTypeWearable    ItemTypes = "Wearable"
	ItemTypeConsumable  ItemTypes = "Consumable"
	ItemTypeUsable      ItemTypes = "Usable"
	ItemTypeThrowable   ItemTypes = "Throwable"
	ItemTypeAmmunitions ItemTypes = "Ammunitions"
	ItemTypeNone        ItemTypes = "Misc"
)

// DefaultItemTypeProperty
func DefaultItemType() *defaultproperty.DefaultStringProperty {
	return defaultproperty.MakeValidatedStringProperty(property.ItemType, string(ItemTypeNone), property.OwnController, property.Item, []string{
		string(ItemTypeWearable),
		string(ItemTypeConsumable),
		string(ItemTypeUsable),
		string(ItemTypeThrowable),
		string(ItemTypeAmmunitions),
		string(ItemTypeNone),
	})
}

// 	Effect           Key = "Effect"           // Absence means nil: No effect. Effects are Skills. (except None)

type EffectProperty struct {
	property.Property
	Effect              *effect.Effect
	minInformationLevel property.InformationLevel
}

// MakeEffectProperty
func MakeEffectProperty(e *effect.Effect, minInfoLevel property.InformationLevel) *EffectProperty {
	return &EffectProperty{
		Effect:              e,
		minInformationLevel: minInfoLevel,
	}
}

// DefaultEffectProperty
func DefaultEffectProperty() *EffectProperty {
	return MakeEffectProperty(nil, property.Analyser)
}

func (bh *EffectProperty) Name(i property.InformationLevel) string {
	if i >= bh.minInformationLevel {
		return "Effect"
	}
	return ""
}

func (bh *EffectProperty) UserFriendlyGet(i property.InformationLevel) interface{} {
	if i >= bh.minInformationLevel {
		return bh.Effect
	}
	return nil
}

func (bh *EffectProperty) Get() interface{} {
	return bh.Effect
}

func (bh *EffectProperty) Set(p interface{}) {
}

func (bh *EffectProperty) Increase() {
}

func (bh *EffectProperty) GetType() property.PropertyType {
	return property.Item
}

func (bh *EffectProperty) Duplicate() property.Property {
	return &EffectProperty{
		Effect:              bh.Effect,
		minInformationLevel: bh.minInformationLevel,
	}
}

func (bh *EffectProperty) ApplyBuff(p property.Property) property.Property {
	return bh.Duplicate() // property.Item can't be buffed.
}

func (bh *EffectProperty) UnapplyBuff(p property.Property) property.Property {
	return bh.Duplicate() // property.Item can't be buffed.
}

// 	WeaponType       Key = "WeaponType"       // Absence means 0: no weapon type (only for Wearable)

type WeaponTypes string

const (
	NoWeapon WeaponTypes = "None"
	// Melee
	OneHandedMelee WeaponTypes = "One-Handed Melee"
	TwoHandedMelee WeaponTypes = "Two-Handed Melee"
	// Ranged
	OneHandedRanged WeaponTypes = "One-Handed Ranged"
	TwoHandedRanged WeaponTypes = "Two-Handed Ranged"
)

// DefaultWeaponTypeProperty
func DefaultWeaponTypeProperty() *defaultproperty.DefaultStringProperty {
	return defaultproperty.MakeValidatedStringProperty(property.WeaponType, string(NoWeapon), property.Public, property.Item, []string{
		string(NoWeapon),
		string(OneHandedMelee),
		string(TwoHandedMelee),
		string(OneHandedRanged),
		string(TwoHandedRanged),
	})
}

//ArmorType        Key = "ArmorType"        // Absence means 0: no armor type (only for Wearable)

type ArmorTypes string

const (
	NoArmor   ArmorTypes = "None"
	HeadSlot  ArmorTypes = "Head"
	BodySlot  ArmorTypes = "Body"
	HandsSlot ArmorTypes = "Hands"
	LegsSlot  ArmorTypes = "Legs"
	FeetSlot  ArmorTypes = "Feet"
	BeltSlot  ArmorTypes = "Belt"
	NeckSlot  ArmorTypes = "Neck"
	RingSlot  ArmorTypes = "Ring"
)

// DefaultArmorTypeProperty
func DefaultArmorTypeProperty() *defaultproperty.DefaultStringProperty {
	return defaultproperty.MakeValidatedStringProperty(property.ArmorType, string(NoArmor), property.Public, property.Item, []string{
		string(NoArmor),
		string(HeadSlot),
		string(BodySlot),
		string(HandsSlot),
		string(LegsSlot),
		string(FeetSlot),
		string(BeltSlot),
		string(NeckSlot),
		string(RingSlot),
	})
}

//ToolType         Key = "ToolType"         // Absence means 0: no tool type (only for Wearable)

type ToolTypes string

const (
	NoTool   ToolTypes = "None"
	SomeTool ToolTypes = "SomeTool"
)

// DefaultToolTypeProperty
func DefaultToolTypeProperty() *defaultproperty.DefaultStringProperty {
	return defaultproperty.MakeValidatedStringProperty(property.ToolType, string(NoTool), property.Public, property.Item, []string{
		string(NoTool),
		string(SomeTool),
	})
}

// ItemProperty resolves the default Property for k, registry-backed and
// scope-filtered to Item. Returns nil for an unknown key or a key not scoped
// to Item (crash-early: no invented fallback for a wrong-scope key).
func ItemProperty(k property.Key) property.Property {
	entry, ok := Lookup(k)
	if !ok || entry.Scopes&ScopeItem == 0 {
		return nil
	}
	return entry.New()
}
