package skillgenerator

import (
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property/def"
)

// produceDebuff builds a debuff skill: PoisonPower + Duration, EnemyOnly.
// PoisonPower costs 15 PSW per 1 dmg/turn.
func produceDebuff(targetPSW int) skill.Skill {
	bp := newBlueprint()
	bp.setTargetType(def.TargetTypeEnemyOnly)
	bp.addDuration(3)

	pp := targetPSW / 15
	if pp < 1 {
		pp = 1
	}
	bp.addDamage(0)
	bp.addPoisonPower(pp)
	return bp.build()
}

// layerDebuff adds PoisonPower + Duration to an existing enemy-targeted skill.
func layerDebuff(sk *skill.Skill, budget int) {
	if budget < 15 {
		return
	}
	pp := budget / 15
	if pp > 5 {
		pp = 5
	}
	bp := &blueprint{sk: *sk}
	bp.addPoisonPower(pp)
	bp.addDuration(2)
	*sk = bp.sk
}
