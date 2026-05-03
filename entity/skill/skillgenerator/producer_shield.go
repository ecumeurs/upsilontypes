package skillgenerator

import (
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property/def"
	"github.com/ecumeurs/upsilontools/tools"
)

// produceShield builds a shield skill: ShieldPower, Self or FriendOnly.
// ShieldPower costs 10 PSW per point.
func produceShield(targetPSW int) skill.Skill {
	bp := newBlueprint()

	sp := targetPSW / 10
	if sp < 1 {
		sp = 1
	}

	// Randomly target self or friendly
	if tools.RandomInt(0, 100) < 50 {
		bp.setTargetType(def.TargetTypeSelf)
	} else {
		bp.setTargetType(def.TargetTypeFriendOnly)
	}
	bp.addDamage(0)
	bp.addShieldPower(sp)
	return bp.build()
}
