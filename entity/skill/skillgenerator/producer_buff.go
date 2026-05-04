package skillgenerator

import (
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property/def"
)

// produceBuff builds a buff skill: CritChance + optional ShieldPower, Duration=3, Self.
// CritChance is capped at 100% (200 PSW); ShieldPower fills any remaining budget.
func produceBuff(targetPSW int) skill.Skill {
	bp := newBlueprint()
	bp.setTargetType(def.TargetTypeSelf)
	bp.addDuration(3) // +40 PSW

	targetPSW -= 40
	if targetPSW < 0 {
		targetPSW = 0
	}

	critPSW := targetPSW
	if critPSW > 200 {
		critPSW = 200
	}
	critChance := critPSW / 2
	if critChance < 5 {
		critChance = 5
	}
	bp.addCritChance(critChance)

	// Fill remaining budget with ShieldPower (10 PSW/point)
	remaining := targetPSW - critChance*2
	if remaining >= 10 {
		bp.addShieldPower(remaining / 10)
	}
	return bp.build()
}

// layerBuff adds Duration + CritChance to an existing skill (self-targeting secondary).
func layerBuff(sk *skill.Skill, budget int) {
	durationWeight := 20 // for duration 2
	if budget < durationWeight+10 {
		return
	}
	budget -= durationWeight
	critChance := budget / 2
	if critChance > 50 {
		critChance = 50
	}
	bp := &blueprint{sk: *sk}
	bp.addCritChance(critChance)
	bp.addDuration(2)
	*sk = bp.sk
}
