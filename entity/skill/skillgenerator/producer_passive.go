package skillgenerator

import (
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property/def"
	"github.com/ecumeurs/upsilontools/tools"
)

// producePassive builds a passive skill: Behavior=Passive, Self target.
// CritChance is capped at 100% (200 PSW); ShieldPower fills any remaining budget.
func producePassive(targetPSW int) skill.Skill {
	bp := newBlueprint()
	bp.setBehavior(def.BehaviorTypePassive)
	bp.setTargetType(def.TargetTypeSelf)
	bp.addDamage(0)

	if tools.RandomInt(0, 100) < 50 {
		// CritChance path (capped at 200 PSW) + ShieldPower overflow
		critPSW := targetPSW
		if critPSW > 200 {
			critPSW = 200
		}
		critChance := critPSW / 2
		if critChance < 5 {
			critChance = 5
		}
		bp.addCritChance(critChance)
		remaining := targetPSW - critChance*2
		if remaining >= 10 {
			bp.addShieldPower(remaining / 10)
		}
	} else {
		// Pure ShieldPower (no cap issue)
		sp := targetPSW / 10
		if sp < 1 {
			sp = 1
		}
		bp.addShieldPower(sp)
	}
	return bp.build()
}
