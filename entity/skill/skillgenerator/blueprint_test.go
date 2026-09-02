package skillgenerator

import (
	"testing"

	"github.com/ecumeurs/upsilontypes/property"
)

// @test-link [[mech_skill_generator_blueprint]]

func TestBlueprint_SetEffectProperty_AddsNew(t *testing.T) {
	bp := newBlueprint()
	bp.addDamage(50)

	dmg := bp.sk.GetPropertyI(property.DamageScale).I()
	if dmg != 50 {
		t.Errorf("expected Damage=50, got %d", dmg)
	}
}

func TestBlueprint_SetEffectProperty_ReplacesExisting(t *testing.T) {
	bp := newBlueprint()
	// newBlueprint initializes Damage to 0
	bp.addDamage(80)

	dmg := bp.sk.GetPropertyI(property.DamageScale).I()
	// Since Damage is a non-counter (IntProperty), setEffectProperty accumulates:
	// existing 0 + 80 = 80
	if dmg != 80 {
		t.Errorf("expected Damage=80, got %d", dmg)
	}
}

func TestBlueprint_RangeMinimum(t *testing.T) {
	bp := newBlueprint()
	bp.setRange(0)

	rng := bp.sk.GetPropertyC(property.Range).GetMaxValue()
	if rng < 1 {
		t.Errorf("setRange(0) should clamp to 1, got %d", rng)
	}
}

func TestBlueprint_StunPower_Clamped(t *testing.T) {
	bp := newBlueprint()
	bp.addStunPower(0) // should be no-op
	sp := bp.sk.GetPropertyI(property.StunPower).I()
	if sp != 0 {
		t.Errorf("addStunPower(0) should be no-op, got %d", sp)
	}

	bp.addStunPower(5)
	sp = bp.sk.GetPropertyI(property.StunPower).I()
	if sp != 5 {
		t.Errorf("expected StunPower=5, got %d", sp)
	}
}

func TestBlueprint_PoisonChance_Clamped(t *testing.T) {
	bp := newBlueprint()
	bp.addPoisonChance(0) // should be no-op
	pc := bp.sk.GetPropertyI(property.PoisonChance).I()
	if pc != 0 {
		t.Errorf("addPoisonChance(0) should be no-op, got %d", pc)
	}

	bp.addPoisonChance(150) // should clamp to 100
	pc = bp.sk.GetPropertyI(property.PoisonChance).I()
	if pc != 100 {
		t.Errorf("addPoisonChance(150) should clamp to 100, got %d", pc)
	}
}
