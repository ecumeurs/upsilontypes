package skillgenerator

// @test-link [[mech_skill_classifier]]

import (
	"testing"

	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/property"
	"github.com/ecumeurs/upsilontypes/property/def"
	"github.com/ecumeurs/upsilontypes/property/defaultproperty"
	"github.com/ecumeurs/upsilonmapdata/grid/position/pattern"
)

func TestClassify_Melee(t *testing.T) {
	sk := skill.New()
	// Range == 1 (default) + Damage effect → melee
	sk.Targeting[property.Range.String()] =
		defaultproperty.MakeIntCounterProperty(property.Range, 1, 1, property.Public, property.Skill)
	sk.Effect.Properties = append(sk.Effect.Properties,
		defaultproperty.MakeIntProperty(property.Damage, 80, property.Public, property.Skill),
	)
	tags := Classify(sk)
	found := false
	for _, tag := range tags {
		if tag == "melee" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected 'melee' in tags, got %v", tags)
	}
}

func TestClassify_Passive(t *testing.T) {
	sk := skill.New()
	p := def.DefaultBehavior()
	p.SetS(string(def.BehaviorTypePassive))
	sk.Behavior = p
	tags := Classify(sk)
	if len(tags) == 0 || tags[0] != "passive" {
		t.Errorf("expected primary tag 'passive', got %v", tags)
	}
}

func TestClassify_Trap(t *testing.T) {
	sk := skill.New()
	p := def.DefaultBehavior()
	p.SetS(string(def.BehaviorTypeTrap))
	sk.Behavior = p
	sk.Effect.Properties = append(sk.Effect.Properties,
		defaultproperty.MakeIntProperty(property.Heal, 60, property.Public, property.Skill),
	)
	tags := Classify(sk)
	// Expected: [trap, heal, ...] — behavior first, effect second
	if len(tags) < 2 || tags[0] != "trap" || tags[1] != "heal" {
		t.Errorf("expected [trap, heal, ...], got %v", tags)
	}
}

func TestClassify_DotAoeOrdering(t *testing.T) {
	// AoE + PoisonPower → tags should have dot before aoe (effect before delivery)
	sk := skill.New()
	// Use Neighbours() which returns a 3x3x3 pattern (27 cells > 1 → AoE)
	sk.Targeting[property.Zone.String()] = def.MakeZoneProperty(pattern.Neighbours(), "Neighbours")
	sk.Effect.Properties = append(sk.Effect.Properties,
		defaultproperty.MakeIntProperty(property.PoisonPower, 5, property.Public, property.Skill),
	)
	tags := Classify(sk)
	// Expected: dot appears before aoe
	dotIdx, aoeIdx := -1, -1
	for i, tag := range tags {
		if tag == "dot" {
			dotIdx = i
		}
		if tag == "aoe" {
			aoeIdx = i
		}
	}
	if dotIdx < 0 {
		t.Errorf("expected 'dot' in tags, got %v", tags)
	}
	if aoeIdx < 0 {
		t.Errorf("expected 'aoe' in tags, got %v", tags)
	}
	if dotIdx >= 0 && aoeIdx >= 0 && dotIdx > aoeIdx {
		t.Errorf("expected 'dot' before 'aoe', got %v", tags)
	}
}

func TestClassify_StunTag(t *testing.T) {
	sk := skill.New()
	sk.Targeting[property.Range.String()] =
		defaultproperty.MakeIntCounterProperty(property.Range, 1, 1, property.Public, property.Skill)
	sk.Effect.Properties = append(sk.Effect.Properties,
		defaultproperty.MakeIntProperty(property.StunChance, 30, property.Public, property.Skill),
	)
	tags := Classify(sk)
	found := false
	for _, tag := range tags {
		if tag == "stun" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected 'stun' in tags when StunChance > 0, got %v", tags)
	}
}

func TestClassify_BuffTag(t *testing.T) {
	sk := skill.New()
	tp := def.DefaultTargetType()
	tp.SetS(string(def.TargetTypeSelf))
	sk.Targeting[property.TargetType.String()] = tp
	sk.Effect.Properties = append(sk.Effect.Properties,
		defaultproperty.MakeIntCounterProperty(property.Duration, 0, 3, property.Public, property.Skill),
		defaultproperty.MakeIntProperty(property.CriticalChance, 25, property.Public, property.Skill),
	)
	tags := Classify(sk)
	found := false
	for _, tag := range tags {
		if tag == "buff" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected 'buff' in tags for Self+Duration+CritChance, got %v", tags)
	}
}

func TestClassify_DebuffTag(t *testing.T) {
	sk := skill.New()
	tp := def.DefaultTargetType()
	tp.SetS(string(def.TargetTypeEnemyOnly))
	sk.Targeting[property.TargetType.String()] = tp
	sk.Effect.Properties = append(sk.Effect.Properties,
		defaultproperty.MakeIntCounterProperty(property.Duration, 0, 3, property.Public, property.Skill),
		defaultproperty.MakeIntProperty(property.PoisonPower, 5, property.Public, property.Skill),
	)
	tags := Classify(sk)
	found := false
	for _, tag := range tags {
		if tag == "debuff" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected 'debuff' in tags for EnemyOnly+Duration+PoisonPower, got %v", tags)
	}
}
