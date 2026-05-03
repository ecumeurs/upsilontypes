package skillgenerator

import (
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property/def"
)

// produceStun builds a stun skill: StunChance + Damage, EnemyOnly.
// StunChance is capped at 100% (200 PSW); remaining budget fills Damage.
func produceStun(targetPSW int) skill.Skill {
	bp := newBlueprint()
	bp.setRange(1)
	bp.setTargetType(def.TargetTypeEnemyOnly)

	stunPSW := targetPSW * 6 / 10
	if stunPSW > 200 { // StunChance max 100% = 200 PSW
		stunPSW = 200
	}
	dmgPSW := targetPSW - stunPSW

	stunChance := stunPSW / 2
	if stunChance < 5 {
		stunChance = 5
	}
	bp.addStunChance(stunChance)
	bp.addDamage(dmgPSW)
	return bp.build()
}

// layerStun adds StunChance to an existing skill within the given budget.
func layerStun(sk *skill.Skill, budget int) {
	if budget < 4 {
		return
	}
	stunChance := budget / 2
	if stunChance > 50 {
		stunChance = 50
	}
	bp := &blueprint{sk: *sk}
	bp.addStunChance(stunChance)
	*sk = bp.sk
}
