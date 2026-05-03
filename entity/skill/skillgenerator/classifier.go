package skillgenerator

import (
	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property"
	"github.com/ecumeurs/upsilontypes/property/def"
)

// Classify derives the ordered tag list for a skill.
// Order: behavior tag → effect family → delivery → modifiers.
// @spec-link [[shared:req_skill_generation_overhaul]]
func Classify(sk skill.Skill) []string {
	tags := make([]string, 0, 4)

	// ── Layer 1: Behavior tag (non-Direct only) ─────────────────
	behavior := sk.Behavior.Get().(string)
	switch def.BehaviorType(behavior) {
	case def.BehaviorTypeTrap:
		tags = append(tags, "trap")
	case def.BehaviorTypeCounter:
		tags = append(tags, "counter")
	case def.BehaviorTypeReaction:
		tags = append(tags, "reaction")
	case def.BehaviorTypePassive:
		tags = append(tags, "passive")
	}

	// ── Layer 2: Effect family ───────────────────────────────────
	heal := sk.GetPropertyI(property.Heal).I()
	if heal > 0 {
		tags = append(tags, "heal")
	}

	shieldPower := sk.GetPropertyI(property.ShieldPower).I()
	if shieldPower > 0 {
		tags = append(tags, "shield")
	}

	poisonPower := sk.GetPropertyI(property.PoisonPower).I()
	if poisonPower > 0 {
		tags = append(tags, "dot")
	}

	stunPower := sk.GetPropertyI(property.StunPower).I()
	stunChance := sk.GetPropertyI(property.StunChance).I()
	if stunPower > 0 || stunChance > 0 {
		tags = append(tags, "stun")
	}

	// buff/debuff: Duration > 0 with targeting
	duration := sk.GetPropertyC(property.Duration).GetMaxValue()
	targetType := targetTypeStr(sk)

	if duration > 0 {
		critChance := sk.GetPropertyI(property.CriticalChance).I()
		switch targetType {
		case string(def.TargetTypeSelf), string(def.TargetTypeFriendOnly):
			// buff: any positive non-damage effect present
			if heal > 0 || shieldPower > 0 || critChance > 0 {
				tags = append(tags, "buff")
			}
		case string(def.TargetTypeEnemyOnly):
			// debuff: negative status present
			if poisonPower > 0 || stunChance > 0 || stunPower > 0 {
				tags = append(tags, "debuff")
			}
		}
	}

	// ── Layer 3: Delivery ────────────────────────────────────────
	zoneProp, hasZone := sk.Targeting[property.Zone.String()]
	if hasZone {
		if zp, ok := zoneProp.(*def.ZoneProperty); ok && len(zp.ZonePattern) > 1 {
			tags = append(tags, "aoe")
		}
	}

	rangeProp := sk.GetPropertyC(property.Range).GetMaxValue()
	damage := sk.GetPropertyI(property.Damage).I()

	if rangeProp >= 2 && damage > 0 {
		tags = append(tags, "ranged")
	} else if rangeProp == 1 && (damage > 0 || stunPower > 0) {
		tags = append(tags, "melee")
	}

	// ── Layer 4: Modifiers ───────────────────────────────────────
	critChance := sk.GetPropertyI(property.CriticalChance).I()
	if critChance >= 25 {
		tags = append(tags, "crit")
	}

	channeling := sk.GetPropertyC(property.Channeling).GetMaxValue()
	if channeling > 0 {
		tags = append(tags, "channeled")
	}

	delay := sk.GetPropertyC(property.Delay).GetMaxValue()
	if delay <= 100 && channeling == 0 {
		tags = append(tags, "instant")
	}

	// mobility: Self target + Movement in effects + Duration > 0
	if targetType == string(def.TargetTypeSelf) && duration > 0 {
		for _, p := range sk.Effect.Properties {
			if p.Name(property.GameMaster) == string(property.Movement) {
				tags = append(tags, "mobility")
				break
			}
		}
	}

	if len(tags) == 0 {
		tags = append(tags, "melee") // fallback
	}

	return tags
}

func targetTypeStr(sk skill.Skill) string {
	if tp, ok := sk.Targeting[property.TargetType.String()]; ok {
		return tp.Get().(string)
	}
	return string(def.TargetTypeEntity)
}
