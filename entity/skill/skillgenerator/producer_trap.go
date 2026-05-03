package skillgenerator

import (
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property/def"
)

// produceTrap builds a trap skill: Behavior=Trap, Damage, TargetType=Tile.
func produceTrap(targetPSW int) skill.Skill {
	bp := newBlueprint()
	bp.setBehavior(def.BehaviorTypeTrap)
	bp.setTargetType(def.TargetTypeTile)
	bp.addDamage(targetPSW)
	return bp.build()
}
