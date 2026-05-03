package skillgenerator

import (
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property/def"
	"github.com/ecumeurs/upsilontools/tools"
)

// produceRanged builds a ranged skill: Range 2–4, EnemyOnly, Damage.
func produceRanged(targetPSW int) skill.Skill {
	bp := newBlueprint()

	rangeVal := tools.RandomInt(2, 5) // 2, 3, or 4
	rangePSW := (rangeVal - 1) * 10
	dmg := targetPSW - rangePSW
	if dmg < 10 {
		dmg = 10
		rangeVal = 2 // fall back to shortest range
		rangePSW = 10
	}

	bp.setRange(rangeVal)
	bp.setTargetType(def.TargetTypeEnemyOnly)
	bp.addDamage(dmg)

	return bp.build()
}
