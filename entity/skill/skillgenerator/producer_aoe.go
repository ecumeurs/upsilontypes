package skillgenerator

import (
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property/def"
)

// produceAOE builds an AoE skill: Line zone pattern + Damage, EnemyOnly.
// Each extra cell (beyond 1) costs 50 PSW.
func produceAOE(targetPSW int) skill.Skill {
	bp := newBlueprint()
	bp.setTargetType(def.TargetTypeEnemyOnly)

	// How many extra cells can we afford?
	extraCells := targetPSW / 50
	if extraCells < 1 {
		extraCells = 1
	}
	if extraCells > 4 {
		extraCells = 4
	}
	zonePSW := extraCells * 50
	dmg := targetPSW - zonePSW
	if dmg < 0 {
		dmg = 0
	}

	bp.setZoneLine(extraCells + 1) // +1 because Line(n) starts at cell 0 (caster)
	bp.addDamage(dmg)               // explicit 0 overrides default-100 PSW

	return bp.build()
}

// layerAOE adds a Line AoE zone to an existing skill within the given PSW budget.
func layerAOE(sk *skill.Skill, budget int) {
	if budget < 50 {
		return
	}
	extraCells := budget / 50
	if extraCells > 3 {
		extraCells = 3
	}
	bp := &blueprint{sk: *sk}
	bp.setZoneLine(extraCells + 1)
	*sk = bp.sk
}
