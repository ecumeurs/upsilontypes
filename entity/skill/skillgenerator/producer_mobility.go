package skillgenerator

import (
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property"
	"github.com/ecumeurs/upsilontypes/property/def"
	"github.com/ecumeurs/upsilontypes/property/defaultproperty"
)

// produceMobility builds a mobility skill: Self, CritChance + ShieldPower overflow, Duration.
// CritChance is capped at 100% (200 PSW); ShieldPower fills remaining budget.
// Movement property is added as a classifier marker.
func produceMobility(targetPSW int) skill.Skill {
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

	remaining := targetPSW - critChance*2
	if remaining >= 10 {
		bp.addShieldPower(remaining / 10)
	}

	// Movement marker for the classifier (no PSW contribution).
	// Use addMovement if it existed, but for now just use a safe set.
	bp.sk.Effect.Properties = append(bp.sk.Effect.Properties,
		defaultproperty.MakeIntProperty(property.Movement, 1, property.Public, property.Skill))

	return bp.build()
}

// layerCrit adds CriticalChance to an existing skill within the given budget.
func layerCrit(sk *skill.Skill, budget int) {
	if budget < 4 {
		return
	}
	critChance := budget / 2
	if critChance > 50 {
		critChance = 50
	}
	bp := &blueprint{sk: *sk}
	bp.addCritChance(critChance)
	*sk = bp.sk
}
