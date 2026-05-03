package skillgenerator

import (
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/entity/skill/skillweight"
	"github.com/ecumeurs/upsilontypes/property"
	"github.com/ecumeurs/upsilontypes/property/def"
	"github.com/ecumeurs/upsilontypes/property/defaultproperty"
	"github.com/ecumeurs/upsilonmapdata/grid/position/pattern"
)

// blueprint builds a skill while tracking positive SW spend.
type blueprint struct {
	sk skill.Skill
}

func newBlueprint() *blueprint {
	bp := &blueprint{sk: skill.New()}
	// Default TargetingMechanics is "Anywhere" which adds +40 PSW.
	// Override to LoS so budget calculations are deterministic.
	p := def.DefaultTargetingMechanics()
	p.SetS(string(def.TargetingMechanicsLOS))
	bp.sk.Targeting[property.TargetingMechanics.String()] = p
	return bp
}

func (b *blueprint) addDamage(dmg int) *blueprint {
	// Always set explicitly — absence defaults to 100 in skillweight, which inflates PSW.
	b.sk.Effect.Properties = append(b.sk.Effect.Properties,
		defaultproperty.MakeIntProperty(property.Damage, dmg, property.Public, property.Skill))
	return b
}

func (b *blueprint) addHeal(heal int) *blueprint {
	if heal <= 0 {
		return b
	}
	b.sk.Effect.Properties = append(b.sk.Effect.Properties,
		defaultproperty.MakeIntProperty(property.Heal, heal, property.Public, property.Skill))
	return b
}

func (b *blueprint) addShieldPower(sp int) *blueprint {
	if sp <= 0 {
		return b
	}
	b.sk.Effect.Properties = append(b.sk.Effect.Properties,
		defaultproperty.MakeIntProperty(property.ShieldPower, sp, property.Public, property.Skill))
	return b
}

func (b *blueprint) addStunChance(sc int) *blueprint {
	if sc <= 0 {
		return b
	}
	b.sk.Effect.Properties = append(b.sk.Effect.Properties,
		defaultproperty.MakeIntProperty(property.StunChance, sc, property.Public, property.Skill))
	return b
}

func (b *blueprint) addCritChance(cc int) *blueprint {
	if cc <= 0 {
		return b
	}
	if cc > 100 {
		cc = 100
	}
	b.sk.Effect.Properties = append(b.sk.Effect.Properties,
		defaultproperty.MakeIntProperty(property.CriticalChance, cc, property.Public, property.Skill))
	return b
}

func (b *blueprint) addPoisonPower(pp int) *blueprint {
	if pp <= 0 {
		return b
	}
	b.sk.Effect.Properties = append(b.sk.Effect.Properties,
		defaultproperty.MakeIntProperty(property.PoisonPower, pp, property.Public, property.Skill))
	return b
}

func (b *blueprint) addDuration(turns int) *blueprint {
	if turns <= 0 {
		return b
	}
	b.sk.Effect.Properties = append(b.sk.Effect.Properties,
		defaultproperty.MakeIntCounterProperty(property.Duration, 0, turns, property.Public, property.Skill))
	return b
}

// setRange sets the Range targeting property (>1 contributes +10 PSW per extra cell).
func (b *blueprint) setRange(r int) *blueprint {
	if r < 1 {
		r = 1
	}
	b.sk.Targeting[property.Range.String()] =
		defaultproperty.MakeIntCounterProperty(property.Range, r, r, property.Public, property.Skill)
	return b
}

// setTargetType sets the TargetType targeting property.
func (b *blueprint) setTargetType(tt def.TargetTypes) *blueprint {
	p := def.DefaultTargetType()
	p.SetS(string(tt))
	b.sk.Targeting[property.TargetType.String()] = p
	return b
}

// setZoneLine sets a Line AoE zone with the given number of cells (must be >= 2).
func (b *blueprint) setZoneLine(cells int) *blueprint {
	if cells < 2 {
		return b
	}
	b.sk.Targeting[property.Zone.String()] =
		def.MakeZoneProperty(pattern.Line(cells), "Line")
	return b
}

// setBehavior sets the skill behavior.
func (b *blueprint) setBehavior(bt def.BehaviorType) *blueprint {
	p := def.DefaultBehavior()
	p.SetS(string(bt))
	b.sk.Behavior = p
	return b
}

// setCooldown adds a cooldown cost (reduces net SW but not PSW).
func (b *blueprint) setCooldown(turns int) *blueprint {
	if turns <= 0 {
		return b
	}
	b.sk.Costs[property.Cooldown.String()] =
		defaultproperty.MakeIntCounterProperty(property.Cooldown, 0, turns, property.Public, property.Skill)
	return b
}

// psw returns the current positive SW of the skill as built so far.
func (b *blueprint) psw() int {
	pSW, _, _ := skillweight.Calculate(b.sk)
	return pSW
}

func (b *blueprint) build() skill.Skill {
	return b.sk
}
