package def

import (
	"fmt"

	"github.com/ecumeurs/upsilontypes/property"
	"github.com/ecumeurs/upsilontypes/property/defaultproperty"
	"github.com/ecumeurs/upsilonmapdata/grid/position/pattern"
)

// Prepare default Properties.

func TargetNumber() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.TargetNumber, 0, property.FriendlyController, property.Skill)
}

func Accuracy() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.Accuracy, 100, property.FriendlyController, property.Skill)
}

func Dodge() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.Dodge, 0, property.FriendlyController, property.Skill)
}

func Parry() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.Parry, 0, property.FriendlyController, property.Skill)
}

func Damage() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.Damage, 100, property.FriendlyController, property.Skill)
}

func Heal() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.Heal, 0, property.FriendlyController, property.Skill)
}

func ShieldPower() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.ShieldPower, 0, property.FriendlyController, property.Skill)
}

func StunPower() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.StunPower, 0, property.FriendlyController, property.Skill)
}

func StunChance() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.StunChance, 0, property.FriendlyController, property.Skill)
}

func CriticalChance() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.CriticalChance, 0, property.FriendlyController, property.Skill)
}

func CriticalMultiplier() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.CriticalMultiplier, 100, property.FriendlyController, property.Skill)
}

func Duration() *defaultproperty.DefaultIntCounterProperty {
	return defaultproperty.MakeIntCounterProperty(property.Duration, 0, 0, property.FriendlyController, property.Skill)
}

func PoisonPower() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.PoisonPower, 0, property.FriendlyController, property.Skill)
}

func PoisonChance() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.PoisonChance, 0, property.FriendlyController, property.Skill)
}

func Delay() *defaultproperty.DefaultIntCounterProperty {
	return defaultproperty.MakeIntCounterProperty(property.Delay, 0, 500, property.FriendlyController, property.Skill)
}

func Channeling() *defaultproperty.DefaultIntCounterProperty {
	return defaultproperty.MakeIntCounterProperty(property.Channeling, 0, 0, property.FriendlyController, property.Skill)
}

func HPLeech() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.HPLeech, 0, property.FriendlyController, property.Skill)
}

func MPLeech() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.MPLeech, 0, property.FriendlyController, property.Skill)
}

func SPLeech() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.SPLeech, 0, property.FriendlyController, property.Skill)
}

func MvtCost() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.MvtCost, 0, property.FriendlyController, property.Skill)
}

// Cooldown() Default to 0, 3 ;Special note: Cool down is stored as a counter, minValue represent initial cooldown at battle start. MaxValue represent the cooldown value when used.
func Cooldown() *defaultproperty.DefaultIntCounterProperty {

	return defaultproperty.MakeIntCounterProperty(property.Cooldown, 0, 3, property.FriendlyController, property.Skill)
}

// Behavior property.Property: 	Behavior SkillProperties = "Behavior" property.Skill broad category: Direct, Reaction, Passive, Counter

// Note: Behavior is a property because ...
// Expect buffs to have an impact on behavior (e.g. a buff that makes a skill a counter skill)

type BehaviorType int

const (
	BehaviorTypeDirect BehaviorType = iota
	BehaviorTypeReaction
	BehaviorTypePassive
	BehaviorTypeCounter
	BehaviorTypeTrap
)

type BehaviorProperty struct {
	property.Property
	BehaviorType BehaviorType
}

// IsDirect
func (bh *BehaviorProperty) IsDirect() bool {
	return bh.BehaviorType == BehaviorTypeDirect
}

// IsReaction
func (bh *BehaviorProperty) IsReaction() bool {
	return bh.BehaviorType == BehaviorTypeReaction
}

// IsPassive
func (bh *BehaviorProperty) IsPassive() bool {
	return bh.BehaviorType == BehaviorTypePassive
}

// IsCounter
func (bh *BehaviorProperty) IsCounter() bool {
	return bh.BehaviorType == BehaviorTypeCounter
}

// IsTrap
func (bh *BehaviorProperty) IsTrap() bool {
	return bh.BehaviorType == BehaviorTypeTrap
}

// MakeBehaviorProperty creates a BehaviorProperty
func MakeBehaviorProperty(bh BehaviorType) *BehaviorProperty {
	return &BehaviorProperty{
		BehaviorType: bh,
	}
}

// DefaultBehaviorProperty
func DefaultBehavior() *BehaviorProperty {
	return MakeBehaviorProperty(BehaviorTypeDirect)
}

func (bh *BehaviorProperty) Name(i property.InformationLevel) string {
	return "Behavior"
}

func (bh *BehaviorProperty) UserFriendlyGet(i property.InformationLevel) interface{} {
	return bh.BehaviorType
}

func (bh *BehaviorProperty) Get() interface{} {
	return bh.BehaviorType
}

func (bh *BehaviorProperty) Set(p interface{}) {
	// shouldn't be altered.
	bh.BehaviorType = p.(BehaviorType)
}

func (bh *BehaviorProperty) Increase() {
	// shouldn't be upgradable.... well maybe convert a Direct skill to a Reaction or Counter ? fun...
}

func (bh *BehaviorProperty) GetType() property.PropertyType {
	return property.Skill
}

func (db *BehaviorProperty) Duplicate() property.Property {
	return &BehaviorProperty{
		BehaviorType: db.BehaviorType,
	}
}

func (bh *BehaviorProperty) ApplyBuff(p property.Property) property.Property {
	res := bh.Duplicate().(*BehaviorProperty)
	// replace
	res.BehaviorType = p.(*BehaviorProperty).BehaviorType
	return res
}

func (bh *BehaviorProperty) UnapplyBuff(p property.Property) property.Property {
	res := bh.Duplicate().(*BehaviorProperty)
	// TODO ... :)
	res.BehaviorType = p.(*BehaviorProperty).BehaviorType
	return res
}

func DefaultRange() *defaultproperty.DefaultIntCounterProperty {
	return defaultproperty.MakeIntCounterProperty(property.Range, 1, 1, property.FriendlyController, property.Skill)
}

// Zone         TargetingProperties = "Zone"  // Area of Effect
// ZoneProperty expects to be casted to be used.
// @spec-link [[mech_actor_pattern]]
type ZoneProperty struct {
	property.Property
	ZonePattern pattern.Pattern
}

// MakeZoneProperty
func MakeZoneProperty(zp pattern.Pattern) *ZoneProperty {
	return &ZoneProperty{
		ZonePattern: zp,
	}
}

// DefaultZoneProperty
func DefaultZone() *ZoneProperty {
	return MakeZoneProperty(pattern.Single())
}

func (bh *ZoneProperty) ApplyBuff(p property.Property) property.Property {
	res := bh.Duplicate().(*ZoneProperty)
	// replace
	res.ZonePattern = p.(*ZoneProperty).ZonePattern
	return res
}

func (bh *ZoneProperty) UnapplyBuff(p property.Property) property.Property {
	res := bh.Duplicate().(*ZoneProperty)
	// TODO
	res.ZonePattern = p.(*ZoneProperty).ZonePattern
	return res
}

func (bh *ZoneProperty) Name(i property.InformationLevel) string {
	return "Zone"
}

func (bh *ZoneProperty) UserFriendlyGet(i property.InformationLevel) interface{} {
	return "" // will be implemented later
}

func (bh *ZoneProperty) Get() interface{} {
	return bh
}

func (bh *ZoneProperty) Set(p interface{}) {
	// will be altered directly.
}

func (bh *ZoneProperty) Increase() {
	// shouldn't be upgradable.... well maybe convert a Direct skill to a Reaction or Counter ? fun...
}

func (bh *ZoneProperty) GetType() property.PropertyType {
	return property.Skill
}

func (bh *ZoneProperty) Duplicate() property.Property {
	return &ZoneProperty{
		ZonePattern: bh.ZonePattern,
	}
}

// TargetType         TargetingProperties = "TargetType"         // Entity, Tile, Both, Self

type TargetTypes string

const (
	TargetTypeEntity       TargetTypes = "Entity"
	TargetTypeFriendOnly   TargetTypes = "FriendOnly"
	TargetTypeEnemyOnly    TargetTypes = "EnemyOnly"
	TargetTypeTile         TargetTypes = "Tile"
	TargetTypeEntityOrTile TargetTypes = "EntityOrTile"
	TargetTypeSelf         TargetTypes = "Self"
)

// DefaultTargetTypeProperty
func DefaultTargetType() *defaultproperty.DefaultStringProperty {
	return defaultproperty.MakeValidatedStringProperty(property.TargetType, string(TargetTypeEntity), property.FriendlyController, property.Skill, []string{
		string(TargetTypeEntity),
		string(TargetTypeFriendOnly),
		string(TargetTypeEnemyOnly),
		string(TargetTypeTile),
		string(TargetTypeEntityOrTile),
		string(TargetTypeSelf),
	})
}

// TargetingMechanics TargetingProperties = "TargetingMechanics" // Anywhere, Line of Sight, and maybe other mechanics later.

type TargetingMechanicsType string

const (
	TargetingMechanicsAnywhere TargetingMechanicsType = "Anywhere"
	TargetingMechanicsLOS      TargetingMechanicsType = "Line of Sight"
)

// DefaultTargetingMechanicsProperty
func DefaultTargetingMechanics() *defaultproperty.DefaultStringProperty {
	return defaultproperty.MakeValidatedStringProperty(property.TargetingMechanics, string(TargetingMechanicsAnywhere), property.FriendlyController, property.Skill, []string{
		string(TargetingMechanicsAnywhere),
		string(TargetingMechanicsLOS),
	})
}

func SkillProperty(ps property.SkillProperties) property.Property {
	switch ps {
	case property.Accuracy:
		return Accuracy()
	case property.Behavior:
		return DefaultBehavior()
	case property.Range:
		return DefaultRange()
	case property.Zone:
		return DefaultZone()
	case property.TargetType:
		return DefaultTargetType()
	case property.TargetingMechanics:
		return DefaultTargetingMechanics()
	case property.Dodge:
		return Dodge()
	case property.Parry:
		return Parry()
	case property.Damage:
		return Damage()
	case property.Heal:
		return Heal()
	case property.ShieldPower:
		return ShieldPower()
	case property.StunPower:
		return StunPower()
	case property.StunChance:
		return StunChance()
	case property.CriticalChance:
		return CriticalChance()
	case property.CriticalMultiplier:
		return CriticalMultiplier()
	case property.Duration:
		return Duration()
	case property.PoisonPower:
		return PoisonPower()
	case property.PoisonChance:
		return PoisonChance()
	case property.Delay:
		return Delay()
	case property.Channeling:
		return Channeling()
	case property.Cooldown:
		return Cooldown()
	case property.HPLeech:
		return HPLeech()
	case property.MPLeech:
		return MPLeech()
	case property.SPLeech:
		return SPLeech()
	case property.MvtCost:
		return MvtCost()
	}
	return nil
}
