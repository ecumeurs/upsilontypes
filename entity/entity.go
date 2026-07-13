package entity

import (
	"fmt"
	"time"

	"github.com/ecumeurs/upsilonmapdata/grid/position"
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property"
	"github.com/ecumeurs/upsilontypes/property/def"
	"github.com/google/uuid"
)

type EntityType int

// @spec-link [[mechanic_temporary_entity_system]]
const (
	Character EntityType = 0
	Monster   EntityType = 1
	// TimeBased is a channeling entity or delayed-effect entity that acts on its own turn.
	TimeBased EntityType = 2
	// Trap is an entity that triggers OnStep. It typically has no controller but has positional effects.
	Trap EntityType = 3
	// AreaEffect is an entity that affects a zone each turn.
	AreaEffect EntityType = 4
	// Obstacle is a player-created wall/barrier: blocks movement, has HP, may decay.
	// NOTE: This is distinct from cell.CellType.Obstacle (immutable map geometry).
	Obstacle EntityType = 5
	Others   EntityType = 6
)

type EntityOrientation int

const (
	Up    EntityOrientation = 0 // X:0 Y:1
	Right EntityOrientation = 1 // X:1 Y:0
	Down  EntityOrientation = 2 // X:0 Y:-1
	Left  EntityOrientation = 3 // X:-1 Y:0
)

// String
func (e EntityOrientation) String() string {
	return [...]string{"Up", "Right", "Down", "Left"}[e]
}

type Entity struct {
	ID           uuid.UUID
	ControllerID uuid.UUID
	Type         EntityType
	LastActivity time.Time
	Position     position.Position
	Name         string
	CurrentDelay int
	Orientation  EntityOrientation
	Properties   map[string]property.Property
	Buffs        []property.TemporaryProperties
	Skills       map[uuid.UUID]skill.Skill
	// IsCasting holds the in-flight channeled skill while the entity is locked in a
	// channel. nil means the entity is not channeling. @spec-link [[upsilonbattle:mechanic_channeling_mechanic]]
	IsCasting *CastingState
}

// CastingState tracks an in-flight channeled skill on its caster.
//
// The target is captured per the skill's targeting mode: entity-target skills must
// FOLLOW their target (it can move before the channel resolves) so they store
// TargetEntity; tile/self skills resolve at a fixed coordinate so they store
// TargetPos. Exactly one of the two is set (the other is uuid.Nil / nil).
//
// @spec-link [[upsilonbattle:mechanic_channeling_mechanic]]
type CastingState struct {
	SkillID      uuid.UUID          // the channeled skill being cast
	TargetEntity uuid.UUID          // entity-target channels; uuid.Nil for tile/self
	TargetPos    *position.Position // tile/self channels; nil for entity-target
	Interruption int                // 0-100; accumulates 10 per 1 damage; >=100 fails the channel
}

// NewEntity
func New() Entity {
	return Entity{
		ID:           uuid.New(),
		ControllerID: uuid.Nil,
		Type:         Others,
		LastActivity: time.Now(),
		CurrentDelay: 0,
		Orientation:  Up,
		Properties:   make(map[string]property.Property),
		Buffs:        make([]property.TemporaryProperties, 0),
		Skills:       make(map[uuid.UUID]skill.Skill),
	}
}

// String
func (e Entity) String() string {
	return fmt.Sprintf("[E %s-%s]", e.ID.String()[0:8], e.Name)
}

func (e Entity) PrettyString() string {
	res := fmt.Sprintf("%s %s %s\n", e.String(), e.Position.String(), e.Orientation.String())
	for _, v := range e.Properties {
		res += fmt.Sprintf(" - %s\n", property.PrettyPrint(v, property.GameMaster))
	}

	return res
}

// FaceToward will change the orientation of the entity to face given position.
// Decide facing based on angle toward position, with 0 being facing straight up.
// Allow UP to be set from -45 to 45 degrees, RIGHT from 45 to 135, DOWN from 135 to 225, LEFT from 225 to 315.
func (e *Entity) FaceToward(p position.Position) {
	angle := e.Position.AngleTo(p)
	if angle < 45 || angle > 315 {
		e.Orientation = Up
	} else if angle < 135 {
		e.Orientation = Right
	} else if angle < 225 {
		e.Orientation = Down
	} else {
		e.Orientation = Left
	}
}

// IsBackstabbing returns true if the current entity is attacking the target from behind.
// Uses the same orientation mapping as FaceToward for consistency.
// @spec-link [[mechanic_backstab_detection_algorithm]]
func (e Entity) IsBackstabbing(target Entity) bool {
	// Calculate the angle from the target to the attacker to see if it's in the target's "back" sector.
	backAngle := target.Position.AngleTo(e.Position)

	var backAngleMin, backAngleMax int

	// Mapping based on FaceToward's thresholds (where 0 is East/Right):
	// Up (0): Faces 0, Back 180 -> Range [135, 225]
	// Right (1): Faces 90, Back 270 -> Range [225, 315]
	// Down (2): Faces 180, Back 0 -> Range [315, 45]
	// Left (3): Faces 270, Back 90 -> Range [45, 135]
	switch target.Orientation {
	case Up:
		backAngleMin, backAngleMax = 135, 225
	case Right:
		backAngleMin, backAngleMax = 225, 315
	case Down:
		backAngleMin, backAngleMax = 315, 45
	case Left:
		backAngleMin, backAngleMax = 45, 135
	}

	if backAngleMin > backAngleMax { // Wrap around case (Down)
		return backAngle >= backAngleMin || backAngle <= backAngleMax
	}

	return backAngle >= backAngleMin && backAngle <= backAngleMax
}

// getBasePropertyOrDefault
func (e Entity) getBasePropertyOrDefault(name interface{}) property.Property {
	nname := property.PropertyToString(name)

	var prop property.Property

	if _, found := e.Properties[nname]; found {
		prop = e.Properties[nname].Duplicate()
	} else {
		prop = def.DefaultProperty(name)
	}

	return prop
}

// GetProperty will return the property with the given name, or default
func (e Entity) GetProperty(name interface{}) property.Property {
	prop := e.getBasePropertyOrDefault(name)

	buffs := e.GetBuffsFor(name)

	for _, buff := range buffs {
		prop = prop.ApplyBuff(buff)
	}

	return prop
}

// GetPropertyI will return the property with the given name, or default
func (e Entity) GetPropertyI(name interface{}) property.IntProperty {
	return e.GetProperty(name).(property.IntProperty)
}

// GetPropertyF will return the property with the given name, or default
func (e Entity) GetPropertyF(name interface{}) property.FloatProperty {
	return e.GetProperty(name).(property.FloatProperty)
}

// GetPropertyC will return the property with the given name, or default
func (e Entity) GetPropertyC(name interface{}) property.IntCounterProperty {
	return e.GetProperty(name).(property.IntCounterProperty)
}

func (e *Entity) UpdateProperty(p property.Property) {
	e.Properties[p.Name(property.GameMaster)] = p
}

func (e *Entity) RegisterSkill(s skill.Skill) {
	e.Skills[s.ID] = s
}

func (e *Entity) RegisterBuff(b property.TemporaryProperties) {
	e.Buffs = append(e.Buffs, b)
}

// RemoveBuffsByOrigin removes all buffs from the entity that originated from a specific source ID.
// @spec-link [[mechanic_item_buff_application]]
func (e *Entity) RemoveBuffsByOrigin(originID uuid.UUID) {
	kept := make([]property.TemporaryProperties, 0, len(e.Buffs))
	for _, b := range e.Buffs {
		if b.OriginEntityID != originID {
			kept = append(kept, b)
		}
	}
	e.Buffs = kept
}

func (e Entity) GetBuffsFor(name interface{}) []property.Property {
	nname := property.PropertyToString(name)

	res := make([]property.Property, 0)
	for _, v := range e.Buffs {
		if _, found := v.Properties[nname]; found {
			res = append(res, v.Properties[nname])
		}
	}
	return res
}

func (e *Entity) BuffTickDown() {
	nbbuf := make([]property.TemporaryProperties, 0)
	for _, buff := range e.Buffs {
		if !buff.TickDown() {
			nbbuf = append(nbbuf, buff)
		}
	}
	e.Buffs = nbbuf
}

// SkillCooldownTickDown decrements every equipped skill's cooldown counter by
// one, floored at 0. It mirrors BuffTickDown but operates on skill cooldowns
// instead of buffs: a skill's Cooldown is set to its max value when cast and
// must count back down to 0 (one step per elapsed turn of its owner) before it
// can be cast again. Called once per turn for the entity whose turn is ending.
// @spec-link [[upsilonbattle:mech_skill_validation]]
func (e *Entity) SkillCooldownTickDown() {
	for id, sk := range e.Skills {
		if sk.Cooldown > 0 {
			sk.Cooldown--
			e.Skills[id] = sk
		}
	}
}

func (e Entity) HasActed() bool {
	return e.GetProperty(property.HasActed).Get().(bool)
}

// IsChanneling reports whether the entity is locked in a channeled skill cast.
// @spec-link [[upsilonbattle:mechanic_channeling_mechanic]]
func (e Entity) IsChanneling() bool {
	return e.IsCasting != nil
}

func (e Entity) HasMoved() bool {
	return e.GetProperty(property.HasMoved).Get().(bool)
}

func (e Entity) HasProperty(name interface{}) bool {
	_, found := e.Properties[property.PropertyToString(name)]
	return found
}

// RepsertPropertyValue will insert property if unknown!
func (e *Entity) RepsertPropertyValue(p interface{}, value interface{}) {
	prop := e.GetProperty(p)
	prop.Set(value)
	e.Properties[prop.Name(property.GameMaster)] = prop
}

// RepsertPropertyCMaxValue inserts or updates a counter property maximum value.
func (e *Entity) RepsertPropertyCMaxValue(p interface{}, maxvalue int) {
	prop := e.GetProperty(p).(property.IntCounterProperty)
	prop.SetMaxValue(maxvalue)
	e.Properties[prop.Name(property.GameMaster)] = prop
}

// RepsertPropertyCValue inserts or updates a counter property current value.
func (e *Entity) RepsertPropertyCValue(p interface{}, value int) {
	prop := e.GetProperty(p).(property.IntCounterProperty)
	prop.SetValue(value)
	e.Properties[prop.Name(property.GameMaster)] = prop
}

// UpdatePropertyValue Will only update value if known to the entity (wont affect buffs)
func (e *Entity) UpdatePropertyValue(p interface{}, value interface{}) {
	prop := e.GetProperty(p)
	prop.Set(value)
	e.UpdateProperty(prop)
}

// UpdatePropertyCMaxValue Will only update max value if known to the entity (wont affect buffs)
func (e *Entity) UpdatePropertyCMaxValue(p interface{}, maxvalue int) {
	prop := e.GetProperty(p).(property.IntCounterProperty)
	prop.SetMaxValue(maxvalue)
	e.UpdateProperty(prop)
}

// UpdatePropertyCValue Will only update current value if known to the entity (wont affect buffs)
func (e *Entity) UpdatePropertyCValue(p interface{}, value int) {
	prop := e.GetProperty(p).(property.IntCounterProperty)
	prop.SetValue(value)
	e.UpdateProperty(prop)
}
