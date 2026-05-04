package skillgenerator

import (
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property/def"
)

// produceTrap builds a trap skill: Behavior=Trap, Damage, TargetType=Tile.
func produceTrap(targetPSW int) skill.Skill {
	bp := newBlueprint()
	bp.setBehavior(def.BehaviorTypeTrap) // +40 SW
	bp.setTargetType(def.TargetTypeTile)

	targetPSW -= 40
	if targetPSW < 0 {
		targetPSW = 0
	}
	bp.addDamage(targetPSW)
	return bp.build()
}
