package skillgenerator

import (
	"fmt"

	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/entity/skill/skillweight"
	"github.com/ecumeurs/upsilontypes/property"
	"github.com/ecumeurs/upsilontypes/property/defaultproperty"
	"github.com/ecumeurs/upsilontools/tools"
)

// gradeBand maps grade name to [pswLo, pswHi, secondaryChancePct].
var gradeBand = map[string][3]int{
	"I":   {60, 150, 25},
	"II":  {151, 300, 50},
	"III": {301, 500, 65},
	"IV":  {501, 750, 75},
	"V":   {751, 1000, 85},
}

// GenerateRequest parameterises skill generation.
// @spec-link [[api_skill_generation]]
type GenerateRequest struct {
	TargetGrade string   `json:"grade"`        // "I"…"V"; empty defaults to "I"
	AllowedTags []string `json:"allowed_tags"` // empty = any category
	ForbidTags  []string `json:"forbid_tags"`  // exclude categories
}

var allProducerTags = []string{
	"melee", "ranged", "aoe", "heal", "shield",
	"buff", "debuff", "dot", "stun", "trap",
	"passive",
	// DISABLED: "counter", "reaction", "mobility" — engine does not implement
	// trigger/activation systems for these behavior types yet (ISS-095 #6).
}

type producerFn func(targetPSW int) skill.Skill

var producers = map[string]producerFn{
	"melee":   produceMelee,
	"ranged":  produceRanged,
	"aoe":     produceAOE,
	"heal":    produceHeal,
	"shield":  produceShield,
	"buff":    produceBuff,
	"debuff":  produceDebuff,
	"dot":     produceDot,
	"stun":    produceStun,
	"trap":    produceTrap,
	"passive": producePassive,
	// DISABLED: "counter", "reaction", "mobility" — pending engine support (ISS-095 #6).
}

type layerFn func(sk *skill.Skill, budget int)

var secondaryLayers = map[string]layerFn{
	"dot":    layerDot,
	"aoe":    layerAOE,
	"stun":   layerStun,
	"crit":   layerCrit,
	"debuff": layerDebuff,
	"buff":   layerBuff,
}

// Generate returns (skill, orderedTags, error).
// Replaces GenerateRandomSkill(); kept as alias: Generate(GenerateRequest{TargetGrade:"I"}).
// @spec-link [[mech_skill_generator_core]]
func Generate(req GenerateRequest) (skill.Skill, []string, error) {
	grade := req.TargetGrade
	if grade == "" {
		grade = "I"
	}
	band, ok := gradeBand[grade]
	if !ok {
		return skill.Skill{}, nil, fmt.Errorf("unknown grade: %s", grade)
	}

	pswLo, pswHi, secondaryPct := band[0], band[1], band[2]
	// Use a conservative inner range so that producer rounding (floor/ceil, ±14)
	// never produces PSW outside [pswLo, pswHi].
	targetLo := pswLo + 15
	targetHi := pswHi - 14
	if targetLo > targetHi {
		targetLo = pswLo
		targetHi = pswHi
	}
	targetPSW := tools.RandomInt(targetLo, targetHi+1)

	allowed := buildAllowedList(req.AllowedTags, req.ForbidTags)
	if len(allowed) == 0 {
		return skill.Skill{}, nil, fmt.Errorf("no allowed producers after applying ForbidTags")
	}

	primaryTag := allowed[tools.RandomInt(0, len(allowed))]
	produceFn, ok := producers[primaryTag]
	if !ok {
		return skill.Skill{}, nil, fmt.Errorf("no producer for tag: %s", primaryTag)
	}

	sk := produceFn(targetPSW)

	// Secondary layer
	if tools.RandomInt(0, 100) < secondaryPct {
		pSW, _, _ := skillweight.Calculate(&sk)
		remaining := pswHi - pSW
		if remaining >= 30 {
			candidates := secondaryLayerCandidates(primaryTag)
			if len(candidates) > 0 {
				secTag := candidates[tools.RandomInt(0, len(candidates))]
				if lFn, ok := secondaryLayers[secTag]; ok {
					lFn(&sk, remaining)
				}
			}
		}
	}

	applyDelayCloser(&sk)

	tags := Classify(sk)
	if len(tags) > 0 {
		sk.Name = Name(tags[0], tags[1:], grade)
	}

	return sk, tags, nil
}

// GenerateRandomSkill is kept for backwards compatibility.
func GenerateRandomSkill() skill.Skill {
	sk, _, _ := Generate(GenerateRequest{TargetGrade: "I"})
	return sk
}

func buildAllowedList(allowed, forbid []string) []string {
	forbidSet := make(map[string]bool, len(forbid))
	for _, t := range forbid {
		forbidSet[t] = true
	}
	source := allProducerTags
	if len(allowed) > 0 {
		source = allowed
	}
	result := make([]string, 0, len(source))
	for _, t := range source {
		if !forbidSet[t] {
			result = append(result, t)
		}
	}
	return result
}

func applyDelayCloser(sk *skill.Skill) {
	// Zero Delay first so the calculation below reflects only other costs.
	sk.Costs[property.Delay.String()] = defaultproperty.MakeIntCounterProperty(
		property.Delay, 0, 0, property.Public, property.Skill)
	pSW, nSW, _ := skillweight.Calculate(sk)
	// netSW is now pSW + all non-Delay costs; set Delay to absorb the remainder.
	delay := pSW + nSW
	if delay < 0 {
		delay = 0
	}
	sk.Costs[property.Delay.String()] = defaultproperty.MakeIntCounterProperty(
		property.Delay, 0, delay, property.Public, property.Skill)
}

func secondaryLayerCandidates(primaryTag string) []string {
	result := make([]string, 0, len(secondaryLayers))
	for tag := range secondaryLayers {
		if tag != primaryTag {
			result = append(result, tag)
		}
	}
	return result
}
