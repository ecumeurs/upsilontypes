package skillgenerator

import (
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property/def"
	"github.com/ecumeurs/upsilontools/tools"
)

// produceMelee builds a melee skill: Range=1, EnemyOnly, Damage (and optionally StunChance).
func produceMelee(targetPSW int) skill.Skill {
	bp := newBlueprint()
	bp.setRange(1)
	bp.setTargetType(def.TargetTypeEnemyOnly)

	// Split budget between Damage and StunChance if enough room.
	// StunChance costs 2 PSW/% — use up to 25% of budget for stun.
	if targetPSW >= 80 && tools.RandomInt(0, 100) < 50 {
		stunPSW := tools.RandomInt(20, targetPSW/3+1)
		stunChance := stunPSW / 2
		if stunChance < 1 {
			stunChance = 1
		}
		bp.addStunChance(stunChance)
		bp.addDamage(targetPSW - stunChance*2)
	} else {
		bp.addDamage(targetPSW)
	}

	return bp.build()
}
