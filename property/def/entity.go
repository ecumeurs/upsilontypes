package def

import (
	"github.com/ecumeurs/upsilontypes/property"
	"github.com/ecumeurs/upsilontypes/property/defaultproperty"
)

// Default Properties Declaration

func HP() *defaultproperty.DefaultIntCounterProperty {
	return defaultproperty.MakeIntCounterProperty(property.HP, 10, 10, property.FriendlyController, property.Character)
}

func Movement() *defaultproperty.DefaultIntCounterProperty {
	return defaultproperty.MakeIntCounterProperty(property.Movement, 3, 3, property.FriendlyController, property.Character)
}

func SP() *defaultproperty.DefaultIntCounterProperty {
	return defaultproperty.MakeIntCounterProperty(property.SP, 10, 10, property.FriendlyController, property.Character)
}

func MP() *defaultproperty.DefaultIntCounterProperty {
	return defaultproperty.MakeIntCounterProperty(property.MP, 10, 10, property.FriendlyController, property.Character)
}

func Attack() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.Attack, 1, property.FriendlyController, property.Character)
}

func Defense() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.Defense, 0, property.FriendlyController, property.Character)
}

func JumpHeight() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.JumpHeight, 2, property.FriendlyController, property.Character)
}

func IsDying() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.IsDying, -1, property.Public, property.Character)
}

func HasMoved() *defaultproperty.DefaultBoolProperty {
	return defaultproperty.MakeBoolProperty(property.HasMoved, false, property.GameMaster, property.Character)
}

func HasActed() *defaultproperty.DefaultBoolProperty {
	return defaultproperty.MakeBoolProperty(property.HasActed, false, property.GameMaster, property.Character)
}

func AttackRange() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.AttackRange, 1, property.FriendlyController, property.Character)
}

func Shield() *defaultproperty.DefaultIntCounterProperty {
	return defaultproperty.MakeIntCounterProperty(property.Shield, 0, 0, property.FriendlyController, property.Character)
}

func Poison() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.Poison, 0, property.FriendlyController, property.Character)
}

func Stun() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.Stun, 0, property.FriendlyController, property.Character)
}

func TeamID() *defaultproperty.DefaultIntProperty {
	return defaultproperty.MakeIntProperty(property.TeamID, 0, property.Public, property.Character)
}

func EntityDuration() *defaultproperty.DefaultIntCounterProperty {
	return defaultproperty.MakeIntCounterProperty(property.EntityDuration, 0, 0, property.Public, property.Character)
}

func ExpiresWithCaster() *defaultproperty.DefaultBoolProperty {
	return defaultproperty.MakeBoolProperty(property.ExpiresWithCaster, false, property.Public, property.Character)
}

func WalkThrough() *defaultproperty.DefaultBoolProperty {
	return defaultproperty.MakeBoolProperty(property.WalkThrough, false, property.Public, property.Character)
}

func Invisible() *defaultproperty.DefaultBoolProperty {
	return defaultproperty.MakeBoolProperty(property.Invisible, false, property.Public, property.Character)
}

func AIArchetype() *defaultproperty.DefaultStringProperty {
	return defaultproperty.MakeStringProperty(property.AIArchetype, "", property.Public, property.Character)
}

// note: futher properties may be added per entity basis.
// @spec-link [[upsilontypes:module_property_key_registry]]
func PropertiesForCharacter() []property.Property {
	return []property.Property{
		defaultproperty.MakeIntCounterProperty(property.HP, 10, 10, property.Public, property.Character),
		defaultproperty.MakeIntCounterProperty(property.Movement, 3, 3, property.Public, property.Character),
		defaultproperty.MakeIntCounterProperty(property.SP, 10, 10, property.Public, property.Character),
		defaultproperty.MakeIntCounterProperty(property.MP, 10, 10, property.Public, property.Character),
		defaultproperty.MakeIntCounterProperty(property.Shield, 0, 0, property.Public, property.Character),
		defaultproperty.MakeIntProperty(property.Attack, 3, property.Public, property.Character),
		defaultproperty.MakeIntProperty(property.AttackRange, 1, property.Public, property.Character),
		defaultproperty.MakeIntProperty(property.Defense, 0, property.Public, property.Character),
		defaultproperty.MakeIntProperty(property.JumpHeight, 2, property.Public, property.Character),
		defaultproperty.MakeIntProperty(property.IsDying, -1, property.Public, property.Character),
		defaultproperty.MakeIntProperty(property.TeamID, 0, property.Public, property.Character),
		defaultproperty.MakeBoolProperty(property.HasMoved, false, property.GameMaster, property.Character),
		defaultproperty.MakeBoolProperty(property.HasActed, false, property.GameMaster, property.Character),
	}
}

// EntityProperty resolves the default Property for k, registry-backed and
// scope-filtered to Entity. Returns nil for an unknown key or a key not
// scoped to Entity (crash-early: no invented fallback for a wrong-scope key).
// @spec-link [[upsilontypes:module_property_key_registry]]
func EntityProperty(k property.Key) property.Property {
	entry, ok := Lookup(k)
	if !ok || entry.Scopes&ScopeEntity == 0 {
		return nil
	}
	return entry.New()
}
