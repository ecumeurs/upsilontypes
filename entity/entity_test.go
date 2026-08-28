package entity

import (
	"testing"

	"github.com/ecumeurs/upsilontypes/property"
	"github.com/ecumeurs/upsilontypes/property/defaultproperty"
	"github.com/google/uuid"
)

func TestEntity(t *testing.T) {
	e := New()
	if e.ID == uuid.Nil {
		t.Error("New() should not return nil")
	}
}

func TestEntityGetPropertyWithoutBuffs(t *testing.T) {
	e := New()

	e.Properties[property.HP.String()] = defaultproperty.MakeIntCounterProperty(property.HP, 10, 10, property.Public, property.Character)

	if e.GetProperty(property.HP).Get().(int) != 10 {
		t.Error("GetProperty() should return 10")
	}
}

func TestEntityGetPropertyWithBuffs(t *testing.T) {
	e := New()

	e.Properties[property.HP.String()] = defaultproperty.MakeIntCounterProperty(property.HP, 10, 10, property.Public, property.Character)
	tmpBuff := property.MakeTemporaryProperties(10)
	tmpBuff.Properties[property.HP.String()] = defaultproperty.MakeIntCounterProperty(property.HP, 0, 10, property.Public, property.Character)

	e.RegisterBuff(tmpBuff)
	if e.GetProperty(property.HP).(*defaultproperty.DefaultIntCounterProperty).MaxValue != 20 {
		t.Error("GetProperty() should return 20")
	}
}

func TestEntityGetPropertyWithBuffsAndNegativeValue(t *testing.T) {
	e := New()

	e.Properties[property.HP.String()] = defaultproperty.MakeIntCounterProperty(property.HP, 10, 10, property.Public, property.Character)
	tmpBuff := property.MakeTemporaryProperties(10)
	tmpBuff.Properties[property.HP.String()] = defaultproperty.MakeIntCounterProperty(property.HP, 0, -5, property.Public, property.Character)

	e.RegisterBuff(tmpBuff)
	if e.GetProperty(property.HP).(*defaultproperty.DefaultIntCounterProperty).MaxValue != 5 {
		t.Error("GetProperty() should return 5")
	}
}

func TestBuffGetRemovedAfterTime(t *testing.T) {
	e := New()

	e.Properties[property.HP.String()] = defaultproperty.MakeIntCounterProperty(property.HP, 10, 10, property.Public, property.Character)
	tmpBuff := property.MakeTemporaryProperties(5)
	tmpBuff.Properties[property.HP.String()] = defaultproperty.MakeIntCounterProperty(property.HP, 0, 10, property.Public, property.Character)

	e.RegisterBuff(tmpBuff)

	for i := 0; i < 5; i++ {
		e.BuffTickDown()
	}

	if e.GetProperty(property.HP).(*defaultproperty.DefaultIntCounterProperty).MaxValue != 10 {
		t.Error("GetProperty() should return 10")
	}
}

// TestWriteIsolation_AdjustPropertyCValueDoesNotEscalateAndLeavesCleanBaseOnBuffRemoval
// is the ISS-144 write-isolation regression guard at the entity level. It
// exercises the full EXPECTATION section of
// rule_entity_property_write_isolation.atom.md directly against AdjustPropertyCValue:
//  1. a write derived from a composed read persists only the intended
//     base-level delta, with the buff's contribution absent from base;
//  2. repeated read-modify-write cycles that apply no new base-level delta
//     leave the base value unchanged (no escalation);
//  3. removing every buff leaves GetProperty returning exactly the last
//     correctly-isolated base value, with no buff contribution folded in.
func TestWriteIsolation_AdjustPropertyCValueDoesNotEscalateAndLeavesCleanBaseOnBuffRemoval(t *testing.T) {
	e := New()
	e.Properties[property.HP.String()] = defaultproperty.MakeIntCounterProperty(property.HP, 10, 10, property.Public, property.Character)

	buff := property.MakeTemporaryProperties(0)
	buff.Forever = true
	buff.Properties[property.HP.String()] = defaultproperty.MakeIntCounterProperty(property.HP, 10, 10, property.Public, property.Character)
	e.RegisterBuff(buff)

	// Composed HP is base(10) + buff(10) = 20/20.
	if hp := e.GetPropertyC(property.HP); hp.GetValue() != 20 || hp.GetMaxValue() != 20 {
		t.Fatalf("precondition failed: expected composed HP 20/20, got %d/%d", hp.GetValue(), hp.GetMaxValue())
	}

	// 1. A single delta-based write (-5) must persist only that delta to base.
	e.AdjustPropertyCValue(property.HP, -5)
	if base := e.GetBasePropertyC(property.HP); base.GetValue() != 5 {
		t.Fatalf("expected base HP value 5 after a -5 delta on a base of 10, got %d (buff contribution leaked into base)", base.GetValue())
	}
	if hp := e.GetPropertyC(property.HP); hp.GetValue() != 15 || hp.GetMaxValue() != 20 {
		t.Fatalf("expected composed HP 15/20 after the delta, got %d/%d", hp.GetValue(), hp.GetMaxValue())
	}

	// 2. Repeating a zero-delta write must not escalate the base or composed value.
	for i := 0; i < 3; i++ {
		e.AdjustPropertyCValue(property.HP, 0)
	}
	if base := e.GetBasePropertyC(property.HP); base.GetValue() != 5 {
		t.Fatalf("expected base HP to remain 5 across repeated zero-delta writes, got %d (escalation)", base.GetValue())
	}
	if hp := e.GetPropertyC(property.HP); hp.GetValue() != 15 || hp.GetMaxValue() != 20 {
		t.Fatalf("expected composed HP to remain 15/20 across repeated zero-delta writes, got %d/%d", hp.GetValue(), hp.GetMaxValue())
	}

	// 3. Removing the buff must leave GetProperty returning exactly the last
	// correctly-isolated base value — no buff residue folded into it.
	e.Buffs = nil
	if hp := e.GetPropertyC(property.HP); hp.GetValue() != 5 || hp.GetMaxValue() != 10 {
		t.Fatalf("expected composed HP to equal the clean base 5/10 after buff removal, got %d/%d (buff residue folded into base)", hp.GetValue(), hp.GetMaxValue())
	}
}

func TestRemoveBuffsByOrigin(t *testing.T) {
	e := New()
	origin1 := uuid.New()
	origin2 := uuid.New()

	buff1 := property.TemporaryProperties{Forever: true, OriginEntityID: origin1, Properties: make(map[string]property.Property)}
	buff2 := property.TemporaryProperties{Forever: true, OriginEntityID: origin2, Properties: make(map[string]property.Property)}
	buff3 := property.TemporaryProperties{Forever: true, OriginEntityID: origin1, Properties: make(map[string]property.Property)}

	e.RegisterBuff(buff1)
	e.RegisterBuff(buff2)
	e.RegisterBuff(buff3)

	if len(e.Buffs) != 3 {
		t.Errorf("Expected 3 buffs, got %d", len(e.Buffs))
	}

	e.RemoveBuffsByOrigin(origin1)

	if len(e.Buffs) != 1 {
		t.Errorf("Expected 1 buff after removal, got %d", len(e.Buffs))
	}

	if e.Buffs[0].OriginEntityID != origin2 {
		t.Error("Remaining buff should be origin2")
	}
}
