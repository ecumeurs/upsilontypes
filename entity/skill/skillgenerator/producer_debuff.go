package skillgenerator

import (
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property/def"
)

// produceDebuff builds a debuff skill: PoisonPower + PoisonChance + Duration, EnemyOnly.
// PoisonPower costs 15 PSW per 1 dmg/turn.
// Both PoisonPower AND PoisonChance must be set for the effect applicator to apply poison.
func produceDebuff(targetPSW int) skill.Skill {
	bp := newBlueprint()
	bp.setTargetType(def.TargetTypeEnemyOnly)
	bp.addDuration(3) // +40 PSW

	targetPSW -= 40
	if targetPSW < 0 {
		targetPSW = 0
	}

	pp := targetPSW / 15
	if pp < 1 {
		pp = 1
	}
	bp.addDamage(0)
	bp.addPoisonPower(pp)
	// PoisonChance must be paired with PoisonPower for the effect to apply (ISS-095).
	bp.addPoisonChance(70)
	return bp.build()
}

// layerDebuff adds PoisonPower + PoisonChance + Duration to an existing enemy-targeted skill.
func layerDebuff(sk *skill.Skill, budget int) {
	durationWeight := 20 // for duration 2
	if budget < durationWeight+15 {
		return
	}
	budget -= durationWeight
	pp := budget / 15
	if pp > 5 {
		pp = 5
	}
	bp := &blueprint{sk: *sk}
	bp.addPoisonPower(pp)
	// PoisonChance must be paired with PoisonPower for the effect to apply (ISS-095).
	bp.addPoisonChance(70)
	bp.addDuration(2)
	*sk = bp.sk
}
