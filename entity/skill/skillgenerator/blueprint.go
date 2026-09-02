package skillgenerator

import (
	"github.com/ecumeurs/upsilonmapdata/grid/position/pattern"
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/entity/skill/skillweight"
	"github.com/ecumeurs/upsilontypes/property"
	"github.com/ecumeurs/upsilontypes/property/def"
	"github.com/ecumeurs/upsilontypes/property/defaultproperty"
)

// blueprint builds a skill while tracking positive SW spend.
// @spec-link [[mech_skill_generator_blueprint]]
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

	// Explicitly initialize critical properties to 0 to avoid inheriting non-zero engine defaults.
	bp.addDamage(0)
	bp.setCooldown(0)
	// Delay is handled by applyDelayCloser but initialize for safety.
	bp.sk.Costs[property.Delay.String()] = defaultproperty.MakeIntCounterProperty(
		property.Delay, 0, 0, property.Public, property.Skill)

	return bp
}

func (b *blueprint) setEffectProperty(p property.Key, val int, counter bool) {
	pstr := p.String()
	var newP property.Property
	if counter {
		newP = defaultproperty.MakeIntCounterProperty(p, 0, val, property.Public, property.Skill)
	} else {
		newP = defaultproperty.MakeIntProperty(p, val, property.Public, property.Skill)
	}

	for i, v := range b.sk.Effect.Properties {
		if v.Name(property.GameMaster) == pstr {
			if counter {
				if _, ok := v.(property.IntCounterProperty); ok {
					newP = defaultproperty.MakeIntCounterProperty(p, 0, val, property.Public, property.Skill)
				}
			} else {
				if oldP, ok := v.(property.IntProperty); ok {
					newP = defaultproperty.MakeIntProperty(p, oldP.I()+val, property.Public, property.Skill)
				}
			}
			b.sk.Effect.Properties[i] = newP
			return
		}
	}
	b.sk.Effect.Properties = append(b.sk.Effect.Properties, newP)
}

func (b *blueprint) addDamage(dmg int) *blueprint {
	if dmg < 0 {
		dmg = 0
	}
	b.setEffectProperty(property.DamageScale, dmg, false)
	return b
}

func (b *blueprint) addHeal(heal int) *blueprint {
	if heal <= 0 {
		return b
	}
	b.setEffectProperty(property.Heal, heal, false)
	return b
}

func (b *blueprint) addShieldPower(sp int) *blueprint {
	if sp <= 0 {
		return b
	}
	b.setEffectProperty(property.ShieldPower, sp, false)
	return b
}

func (b *blueprint) addStunChance(sc int) *blueprint {
	if sc <= 0 {
		return b
	}
	b.setEffectProperty(property.StunChance, sc, false)
	return b
}

func (b *blueprint) addStunPower(sp int) *blueprint {
	if sp <= 0 {
		return b
	}
	b.setEffectProperty(property.StunPower, sp, false)
	return b
}

func (b *blueprint) addCritChance(cc int) *blueprint {
	if cc <= 0 {
		return b
	}
	if cc > 100 {
		cc = 100
	}
	b.setEffectProperty(property.CriticalChance, cc, false)
	return b
}

func (b *blueprint) addPoisonPower(pp int) *blueprint {
	if pp <= 0 {
		return b
	}
	b.setEffectProperty(property.PoisonPower, pp, false)
	return b
}

func (b *blueprint) addPoisonChance(pc int) *blueprint {
	if pc <= 0 {
		return b
	}
	if pc > 100 {
		pc = 100
	}
	b.setEffectProperty(property.PoisonChance, pc, false)
	return b
}

func (b *blueprint) addDuration(turns int) *blueprint {
	if turns <= 0 {
		return b
	}
	b.setEffectProperty(property.Duration, turns, true)
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
	if turns < 0 {
		turns = 0
	}
	b.sk.Costs[property.Cooldown.String()] =
		defaultproperty.MakeIntCounterProperty(property.Cooldown, 0, turns, property.Public, property.Skill)
	return b
}

// psw returns the current positive SW of the skill as built so far.
func (b *blueprint) psw() int {
	pSW, _, _ := skillweight.Calculate(&b.sk)
	return pSW
}

func (b *blueprint) build() skill.Skill {
	return b.sk
}
