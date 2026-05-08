package skillgenerator

import (
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property/def"
)

// produceStun builds a stun skill: StunChance + StunPower + Damage, EnemyOnly.
// StunChance is capped at 100% (200 PSW); remaining budget fills Damage.
// Both StunChance AND StunPower must be set for the effect applicator to apply stun.
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
	// StunPower must be paired with StunChance for the effect to apply (ISS-095).
	bp.addStunPower(stunChance / 5)
	bp.addDamage(dmgPSW)
	return bp.build()
}

// layerStun adds StunChance + StunPower to an existing skill within the given budget.
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
	// StunPower must be paired with StunChance for the effect to apply (ISS-095).
	bp.addStunPower(stunChance / 5)
	*sk = bp.sk
}
