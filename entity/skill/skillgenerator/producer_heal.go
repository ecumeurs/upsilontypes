package skillgenerator

import (
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property/def"
)

// produceHeal builds a healing skill: FriendOnly.
// Heal PSW = (heal/10) * 15, so heal = targetPSW * 10 / 15.
func produceHeal(targetPSW int) skill.Skill {
	bp := newBlueprint()
	bp.setTargetType(def.TargetTypeFriendOnly)

	healVal := targetPSW * 10 / 15 // approximate; rounds down
	if healVal < 10 {
		healVal = 10
	}
	// Round to nearest multiple of 10 to keep PSW exact.
	healVal = (healVal / 10) * 10

	bp.addDamage(0)
	bp.addHeal(healVal)
	return bp.build()
}
