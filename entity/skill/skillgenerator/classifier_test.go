package skillgenerator

// @test-link [[shared:req_skill_generation_overhaul]]

// TODO: uncomment and implement once classifier.go is written.
//
// import (
// 	"testing"
//
// 	"github.com/ecumeurs/upsilontypes/entity/skill"
// 	"github.com/ecumeurs/upsilontypes/property"
// 	"github.com/ecumeurs/upsilontypes/property/defaultproperty"
// 	"github.com/ecumeurs/upsilontypes/property/def"
// )
//
// func TestClassify_Melee(t *testing.T) {
// 	sk := skill.New()
// 	// Range == 1 (default) + Damage effect → melee
// 	sk.Effect.Properties = append(sk.Effect.Properties,
// 		defaultproperty.MakeIntProperty(property.Damage, 80, property.Public, property.Skill),
// 	)
// 	tags := Classify(sk)
// 	if len(tags) == 0 || tags[0] != "melee" {
// 		t.Errorf("expected primary tag 'melee', got %v", tags)
// 	}
// }
//
// func TestClassify_Passive(t *testing.T) {
// 	sk := skill.New()
// 	sk.Behavior = def.DefaultBehavior()
// 	sk.Behavior.(*defaultproperty.DefaultStringProperty).Set(string(def.BehaviorTypePassive))
// 	tags := Classify(sk)
// 	if len(tags) == 0 || tags[0] != "passive" {
// 		t.Errorf("expected primary tag 'passive', got %v", tags)
// 	}
// }
//
// func TestClassify_Trap_Heal(t *testing.T) {
// 	sk := skill.New()
// 	sk.Behavior.(*defaultproperty.DefaultStringProperty).Set(string(def.BehaviorTypeTrap))
// 	sk.Effect.Properties = append(sk.Effect.Properties,
// 		defaultproperty.MakeIntProperty(property.Heal, 60, property.Public, property.Skill),
// 	)
// 	tags := Classify(sk)
// 	// Expected: [trap, heal] — behavior first, effect second
// 	if len(tags) < 2 || tags[0] != "trap" || tags[1] != "heal" {
// 		t.Errorf("expected [trap, heal], got %v", tags)
// 	}
// }
//
// func TestClassify_OrderingRule(t *testing.T) {
// 	// AoE damage: delivery (aoe) comes after effect family.
// 	// Tags should be [aoe] if no explicit effect family tag applies,
// 	// or [dot, aoe] if PoisonPower is set.
// 	sk := skill.New()
// 	// Zone > 1 cells
// 	sk.Targeting[property.Zone.String()] = def.MakeZoneProperty(nil, "Neighbours")
// 	sk.Effect.Properties = append(sk.Effect.Properties,
// 		defaultproperty.MakeIntProperty(property.PoisonPower, 5, property.Public, property.Skill),
// 	)
// 	tags := Classify(sk)
// 	// Expected order: [dot, aoe] — effect family before delivery
// 	if len(tags) < 2 || tags[0] != "dot" || tags[1] != "aoe" {
// 		t.Errorf("expected [dot, aoe], got %v", tags)
// 	}
// }
