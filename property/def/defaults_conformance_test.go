// Package def provides conformance and defect-repro tests for the property
// default-value constructors and their scope resolvers.
package def

import (
	"testing"

	"github.com/ecumeurs/upsilontypes/property"
)

// TestDefaults_MovementAbsenceDefaultIsThree reproduces the Movement defect
// recorded in architecture/property_key_vocabulary.md §3 row 2: the frozen
// vocabulary declares default-when-absent 3/3 for Movement (matching
// PropertiesForCharacter and every other code path), but the Movement()
// constructor returns 5/5.
// @test-link [[upsilontypes:module_property_key_registry]]
func TestDefaults_MovementAbsenceDefaultIsThree(t *testing.T) {
	mv := Movement()
	if mv == nil {
		t.Fatal("Movement() returned nil")
	}
	if mv.Value != 3 || mv.MaxValue != 3 {
		t.Errorf("Movement(): an entity created without an explicit Movement property "+
			"gets %d/%d tiles instead of the documented default of 3/3", mv.Value, mv.MaxValue)
	}
}

// keyConformanceCase pairs a registry key with the resolver call that should
// produce its default property, per the key's declared scope in
// architecture/property_key_vocabulary.md.
type keyConformanceCase struct {
	label   string
	key     property.Key
	resolve func() property.Property
}

// allKeyConformanceCases enumerates all 64 keys from the frozen vocabulary
// (architecture/property_key_vocabulary.md §3 Entity 20, §4 Skill 31, §5 Item
// 13), each routed to the resolver matching its declared scope. The four
// dual-scope keys (Accuracy, Dodge, CriticalChance, CriticalMultiplier) are
// routed through the Skill resolver, where they live today.
func allKeyConformanceCases() []keyConformanceCase {
	return []keyConformanceCase{
		// Entity scope (20)
		{"HP", property.HP, func() property.Property { return EntityProperty(property.HP) }},
		{"Movement", property.Movement, func() property.Property { return EntityProperty(property.Movement) }},
		{"SP", property.SP, func() property.Property { return EntityProperty(property.SP) }},
		{"MP", property.MP, func() property.Property { return EntityProperty(property.MP) }},
		{"Attack", property.Attack, func() property.Property { return EntityProperty(property.Attack) }},
		{"Defense", property.Defense, func() property.Property { return EntityProperty(property.Defense) }},
		{"JumpHeight", property.JumpHeight, func() property.Property { return EntityProperty(property.JumpHeight) }},
		{"TeamID", property.TeamID, func() property.Property { return EntityProperty(property.TeamID) }},
		{"IsDying", property.IsDying, func() property.Property { return EntityProperty(property.IsDying) }},
		{"HasMoved", property.HasMoved, func() property.Property { return EntityProperty(property.HasMoved) }},
		{"HasActed", property.HasActed, func() property.Property { return EntityProperty(property.HasActed) }},
		{"AttackRange", property.AttackRange, func() property.Property { return EntityProperty(property.AttackRange) }},
		{"EntityDuration", property.EntityDuration, func() property.Property { return EntityProperty(property.EntityDuration) }},
		{"ExpiresWithCaster", property.ExpiresWithCaster, func() property.Property { return EntityProperty(property.ExpiresWithCaster) }},
		{"WalkThrough", property.WalkThrough, func() property.Property { return EntityProperty(property.WalkThrough) }},
		{"Invisible", property.Invisible, func() property.Property { return EntityProperty(property.Invisible) }},
		{"AIArchetype", property.AIArchetype, func() property.Property { return EntityProperty(property.AIArchetype) }},
		{"Shield (Entity)", property.Shield, func() property.Property { return EntityProperty(property.Shield) }},
		{"Poison (Entity)", property.Poison, func() property.Property { return EntityProperty(property.Poison) }},
		{"Stun (Entity)", property.Stun, func() property.Property { return EntityProperty(property.Stun) }},

		// Skill scope (31), including the four dual-scope keys
		{"Behavior", property.Behavior, func() property.Property { return SkillProperty(property.Behavior) }},
		{"Range", property.Range, func() property.Property { return SkillProperty(property.Range) }},
		{"Zone", property.Zone, func() property.Property { return SkillProperty(property.Zone) }},
		{"TargetNumber", property.TargetNumber, func() property.Property { return SkillProperty(property.TargetNumber) }},
		{"Accuracy", property.Accuracy, func() property.Property { return SkillProperty(property.Accuracy) }},
		{"Dodge", property.Dodge, func() property.Property { return SkillProperty(property.Dodge) }},
		{"Parry", property.Parry, func() property.Property { return SkillProperty(property.Parry) }},
		{"TargetType", property.TargetType, func() property.Property { return SkillProperty(property.TargetType) }},
		{"TargetingMechanics", property.TargetingMechanics, func() property.Property { return SkillProperty(property.TargetingMechanics) }},
		{"DamageScale", property.DamageScale, func() property.Property { return SkillProperty(property.DamageScale) }},
		{"Heal", property.Heal, func() property.Property { return SkillProperty(property.Heal) }},
		{"ShieldPower", property.ShieldPower, func() property.Property { return SkillProperty(property.ShieldPower) }},
		{"StunPower", property.StunPower, func() property.Property { return SkillProperty(property.StunPower) }},
		{"StunChance", property.StunChance, func() property.Property { return SkillProperty(property.StunChance) }},
		{"CriticalChance", property.CriticalChance, func() property.Property { return SkillProperty(property.CriticalChance) }},
		{"CriticalMultiplier", property.CriticalMultiplier, func() property.Property { return SkillProperty(property.CriticalMultiplier) }},
		{"Duration", property.Duration, func() property.Property { return SkillProperty(property.Duration) }},
		{"PoisonPower", property.PoisonPower, func() property.Property { return SkillProperty(property.PoisonPower) }},
		{"PoisonChance", property.PoisonChance, func() property.Property { return SkillProperty(property.PoisonChance) }},
		{"Delay", property.Delay, func() property.Property { return SkillProperty(property.Delay) }},
		{"Channeling", property.Channeling, func() property.Property { return SkillProperty(property.Channeling) }},
		{"HPLeech", property.HPLeech, func() property.Property { return SkillProperty(property.HPLeech) }},
		{"MPLeech", property.MPLeech, func() property.Property { return SkillProperty(property.MPLeech) }},
		{"SPLeech", property.SPLeech, func() property.Property { return SkillProperty(property.SPLeech) }},
		{"MovementCost", property.MovementCost, func() property.Property { return SkillProperty(property.MovementCost) }},
		{"Cooldown", property.Cooldown, func() property.Property { return SkillProperty(property.Cooldown) }},
		{"RepositionSubject", property.RepositionSubject, func() property.Property { return SkillProperty(property.RepositionSubject) }},
		{"RepositionDistance", property.RepositionDistance, func() property.Property { return SkillProperty(property.RepositionDistance) }},
		{"TriggerType", property.TriggerType, func() property.Property { return SkillProperty(property.TriggerType) }},
		{"RemoveOnTrigger", property.RemoveOnTrigger, func() property.Property { return SkillProperty(property.RemoveOnTrigger) }},
		{"TriggerCount", property.TriggerCount, func() property.Property { return SkillProperty(property.TriggerCount) }},

		// Item scope (13)
		{"Durability", property.Durability, func() property.Property { return ItemProperty(property.Durability) }},
		{"Weight", property.Weight, func() property.Property { return ItemProperty(property.Weight) }},
		{"ItemType", property.ItemType, func() property.Property { return ItemProperty(property.ItemType) }},
		{"ArmorRating", property.ArmorRating, func() property.Property { return ItemProperty(property.ArmorRating) }},
		{"WeaponType", property.WeaponType, func() property.Property { return ItemProperty(property.WeaponType) }},
		{"ArmorType", property.ArmorType, func() property.Property { return ItemProperty(property.ArmorType) }},
		{"ToolType", property.ToolType, func() property.Property { return ItemProperty(property.ToolType) }},
		{"WeaponRange", property.WeaponRange, func() property.Property { return ItemProperty(property.WeaponRange) }},
		{"WeaponBaseDamage", property.WeaponBaseDamage, func() property.Property { return ItemProperty(property.WeaponBaseDamage) }},
		{"Stackable", property.Stackable, func() property.Property { return ItemProperty(property.Stackable) }},
		{"StackSize", property.StackSize, func() property.Property { return ItemProperty(property.StackSize) }},
		{"Effect", property.Effect, func() property.Property { return ItemProperty(property.Effect) }},
		{"ItemValue", property.ItemValue, func() property.Property { return ItemProperty(property.ItemValue) }},
	}
}

// TestDefaults_ResolvedNameMatchesRegistryKey guards against the Name()
// desync trap recorded in architecture/property_key_vocabulary.md §8.1:
// Entity.UpdateProperty keys storage by the resolved property's own
// self-reported Name(GameMaster), not the key it was constructed from. If a
// resolver's constructed default ever disagrees with the registry key it was
// asked to build (e.g. a future string-value rename touches one side but not
// the other), a write under the registry key would silently land under a
// different storage key with no visible symptom. property.GameMaster is used
// deliberately: it is exactly the information level UpdateProperty passes,
// and it is the highest level, so this test never trips the legitimate
// low-info-level "" returned by EffectProperty.Name below Analyser.
//
// All 64 keys are expected to resolve and match today except TargetNumber:
// def.SkillProperty has no case for it (a live, documented gap — see
// architecture/property_key_vocabulary.md §4a row 24 and §8.2), so it
// resolves nil and is expected to fail here until that gap is closed.
// @test-link [[upsilontypes:module_property_key_registry]]
func TestDefaults_ResolvedNameMatchesRegistryKey(t *testing.T) {
	cases := allKeyConformanceCases()
	if len(cases) != 64 {
		t.Fatalf("expected 64 registry keys under test, got %d", len(cases))
	}

	for _, c := range cases {
		c := c
		t.Run(c.label, func(t *testing.T) {
			resolved := c.resolve()
			if resolved == nil {
				t.Errorf("resolver for key %q returned nil property", c.label)
				return
			}

			wantName := property.PropertyToString(c.key)
			gotName := resolved.Name(property.GameMaster)
			if gotName != wantName {
				t.Errorf("resolved property Name(GameMaster) = %q, want %q (registry key %q)",
					gotName, wantName, c.label)
			}
		})
	}
}
