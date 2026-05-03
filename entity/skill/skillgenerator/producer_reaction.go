package skillgenerator

import (
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property/def"
	"github.com/ecumeurs/upsilontools/tools"
)

// produceReaction builds a reaction skill: Behavior=Reaction, Damage or Heal.
func produceReaction(targetPSW int) skill.Skill {
	bp := newBlueprint()
	bp.setBehavior(def.BehaviorTypeReaction)

	if tools.RandomInt(0, 100) < 50 {
		bp.setTargetType(def.TargetTypeEnemyOnly)
		bp.addDamage(targetPSW)
	} else {
		bp.setTargetType(def.TargetTypeSelf)
		healVal := (targetPSW * 10 / 15 / 10) * 10
		if healVal < 10 {
			healVal = 10
		}
		bp.addDamage(0)
		bp.addHeal(healVal)
	}
	return bp.build()
}
