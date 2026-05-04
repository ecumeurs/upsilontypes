package skillgenerator

import (
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property/def"
	"github.com/ecumeurs/upsilontools/tools"
)

// produceCounter builds a counter skill: Behavior=Counter, Damage or ShieldPower.
func produceCounter(targetPSW int) skill.Skill {
	bp := newBlueprint()
	bp.setBehavior(def.BehaviorTypeCounter) // +30 SW

	targetPSW -= 30
	if targetPSW < 0 {
		targetPSW = 0
	}

	if tools.RandomInt(0, 100) < 50 {
		// Counter-damage
		bp.setTargetType(def.TargetTypeEnemyOnly)
		bp.addDamage(targetPSW)
	} else {
		// Counter-shield (absorb and reflect)
		bp.setTargetType(def.TargetTypeSelf)
		sp := targetPSW / 10
		if sp < 1 {
			sp = 1
		}
		bp.addDamage(0)
		bp.addShieldPower(sp)
	}
	return bp.build()
}
