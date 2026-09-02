package skill

import (
	"github.com/ecumeurs/upsilontypes/property"
	"github.com/ecumeurs/upsilontypes/property/def"
	"github.com/ecumeurs/upsilontypes/property/effect"
	"github.com/google/uuid"
)

type Skill struct {
	ID   uuid.UUID
	Name string

	Behavior property.Property

	Targeting map[string]property.Property
	Costs     map[string]property.Property
	Effect    effect.Effect
	Cooldown  int
}

// New
func New() Skill {
	return Skill{
		ID:        uuid.New(),
		Name:      "New Skill",
		Behavior:  def.DefaultBehavior(),
		Targeting: make(map[string]property.Property),
		Costs:     make(map[string]property.Property),
		Effect:    *effect.New(),
		Cooldown:  0, // current cooldown value
	}
}

// NewSkill
func NewSkill(name string, targeting, cost map[string]property.Property, effect effect.Effect) Skill {
	return Skill{
		ID:        uuid.New(),
		Name:      name,
		Behavior:  def.DefaultBehavior(),
		Targeting: targeting,
		Costs:     cost,
		Effect:    effect,
	}
}

// IsDirect
func (s *Skill) IsDirect() bool {
	return s.Behavior.Get().(string) == string(def.BehaviorTypeDirect)
}

// IsReaction
func (s *Skill) IsReaction() bool {
	return s.Behavior.Get().(string) == string(def.BehaviorTypeReaction)
}

// IsPassive
func (s *Skill) IsPassive() bool {
	return s.Behavior.Get().(string) == string(def.BehaviorTypePassive)
}

// IsCounter
func (s *Skill) IsCounter() bool {
	return s.Behavior.Get().(string) == string(def.BehaviorTypeCounter)
}

// IsTrap
func (s *Skill) IsTrap() bool {
	return s.Behavior.Get().(string) == string(def.BehaviorTypeTrap)
}

// HasProperty
func (s Skill) HasProperty(p property.Key) bool {
	pstr := property.PropertyToString(p)
	for _, v := range s.Targeting {
		if v.Name(property.GameMaster) == pstr {
			return true
		}
	}
	for _, v := range s.Costs {
		if v.Name(property.GameMaster) == pstr {
			return true
		}
	}
	return s.Effect.HasProperty(p)
}

// HasProperties
func (s Skill) HasAnyProperties(p ...property.Key) bool {
	one := false
	for _, v := range p {
		if s.HasProperty(v) {
			one = true
		}
	}
	return one
}

func (s Skill) HasAllProperties(p ...property.Key) bool {
	for _, v := range p {
		if !s.HasProperty(v) {
			return false
		}
	}
	return true
}

// GetProperty
func (s Skill) GetProperty(p property.Key) property.Property {
	pstr := property.PropertyToString(p)
	for _, v := range s.Targeting {
		if v.Name(property.GameMaster) == pstr {
			return v
		}
	}
	for _, v := range s.Costs {
		if v.Name(property.GameMaster) == pstr {
			return v
		}
	}
	for _, v := range s.Effect.Properties {
		if v.Name(property.GameMaster) == pstr {
			return v
		}
	}

	return def.DefaultProperty(p)
}

// GetProperty
func (s Skill) GetPropertyI(p property.Key) property.IntProperty {
	return s.GetProperty(p).(property.IntProperty)
}

// GetProperty
func (s Skill) GetPropertyF(p property.Key) property.FloatProperty {
	return s.GetProperty(p).(property.FloatProperty)
}

// GetProperty
func (s Skill) GetPropertyC(p property.Key) property.IntCounterProperty {
	return s.GetProperty(p).(property.IntCounterProperty)
}
