package effect

import (
	"github.com/ecumeurs/upsilontypes/property"
	"github.com/google/uuid"
)

type Effect struct {
	Properties []property.Property
	Name       string
	CasterID   uuid.UUID
}

// New
func New() *Effect {
	return &Effect{
		Properties: []property.Property{},
		Name:       "New Effect",
		CasterID:   uuid.Nil,
	}
}

// HasProperty
func (e Effect) HasProperty(p property.Key) bool {
	pstr := property.PropertyToString(p)
	for _, v := range e.Properties {
		if v.Name(property.GameMaster) == pstr {
			return true
		}
	}
	return false
}

// HasPositiveProperty returns true if the property exists and its value is strictly positive.
func (s Effect) HasPositiveProperty(p property.Key) bool {
	pstr := property.PropertyToString(p)
	for _, v := range s.Properties {
		if v.Name(property.GameMaster) == pstr {
			if ip, ok := v.(property.IntProperty); ok {
				return ip.I() > 0
			}
		}
	}
	return false
}

// HasNegativeProperty
func (s Effect) HasNegativeProperty(p property.Key) bool {
	pstr := property.PropertyToString(p)
	for _, v := range s.Properties {
		if v.Name(property.GameMaster) == pstr {
			if v.(property.IntProperty).I() < 0 {
				return true
			}
		}
	}
	return false
}

// GetProperty
func (e Effect) GetProperty(p property.Key) property.Property {
	pstr := property.PropertyToString(p)
	for _, v := range e.Properties {
		if v.Name(property.GameMaster) == pstr {
			return v
		}
	}

	return nil
}

// GetProperty
func (e Effect) GetPropertyI(p property.Key) property.IntProperty {
	return e.GetProperty(p).(property.IntProperty)
}

// GetProperty
func (e Effect) GetPropertyF(p property.Key) property.FloatProperty {
	return e.GetProperty(p).(property.FloatProperty)
}

// GetProperty
func (e Effect) GetPropertyC(p property.Key) property.IntCounterProperty {
	return e.GetProperty(p).(property.IntCounterProperty)
}

// IsDamaging
func (s Effect) IsDamaging() bool {
	return (s.HasPositiveProperty(property.DamageScale) ||
		s.HasPositiveProperty(property.StunPower) ||
		s.HasPositiveProperty(property.PoisonPower) ||
		s.HasNegativeProperty(property.ShieldPower))
}

// IsHealing returns true if the effect provides healing or defensive buffs without offensive damage.
func (s Effect) IsHealing() bool {
	return (!s.HasPositiveProperty(property.DamageScale) ||
		s.HasNegativeProperty(property.StunPower) ||
		s.HasNegativeProperty(property.PoisonPower) ||
		s.HasPositiveProperty(property.ShieldPower) ||
		s.HasPositiveProperty(property.Heal))
}

// IsOverTime (poison, stun, etc) (Buff/Curse)
func (s Effect) IsOverTime() bool {
	return s.HasPositiveProperty(property.Duration)
}
