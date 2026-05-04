package skillgenerator

import (
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property/def"
)

// produceDot builds a damage-over-time skill: PoisonPower, EnemyOnly.
// PoisonPower costs 15 PSW per point.
func produceDot(targetPSW int) skill.Skill {
	bp := newBlueprint()
	bp.setTargetType(def.TargetTypeEnemyOnly)

	pp := targetPSW / 15
	if pp < 1 {
		pp = 1
	}
	bp.addPoisonPower(pp)
	return bp.build()
}

// layerDot adds PoisonPower to an existing skill within the given budget.
func layerDot(sk *skill.Skill, budget int) {
	if budget < 15 {
		return
	}
	pp := budget / 15
	if pp > 4 {
		pp = 4
	}
	bp := &blueprint{sk: *sk}
	bp.addPoisonPower(pp)
	*sk = bp.sk
}
