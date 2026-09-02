package entity

import (
	"testing"

	"github.com/ecumeurs/upsilontypes/property"
	"github.com/ecumeurs/upsilontypes/property/defaultproperty"
	"github.com/google/uuid"
)

// TestGetBuffAttributionFor_Unbuffed asserts that a key with no contributing
// buff blocks still resolves the correct natural (unbuffed) base value and
// returns an empty, non-nil contribs slice rather than erroring or panicking.
//
// @test-link [[mechanic_buff_attribution_accessor]]
func TestGetBuffAttributionFor_Unbuffed(t *testing.T) {
	e := New()
	e.Properties[property.HP.String()] = defaultproperty.MakeIntCounterProperty(property.HP, 10, 10, property.Public, property.Character)

	base, contribs := e.GetBuffAttributionFor(property.HP)

	if base.(*defaultproperty.DefaultIntCounterProperty).Value != 10 {
		t.Fatalf("expected base HP value 10, got %d", base.(*defaultproperty.DefaultIntCounterProperty).Value)
	}
	if contribs == nil {
		t.Fatal("expected contribs to be an empty non-nil slice, got nil")
	}
	if len(contribs) != 0 {
		t.Fatalf("expected 0 contribs for an unbuffed key, got %d", len(contribs))
	}
}

// TestGetBuffAttributionFor_TwoDistinctBlocks is the centrepiece case that
// GetBuffsFor cannot express: two separate buff blocks, from two distinct
// origins, both contributing to the same key. GetBuffAttributionFor must
// report each contribution against its own owning block.
//
// @test-link [[mechanic_buff_attribution_accessor]]
func TestGetBuffAttributionFor_TwoDistinctBlocks(t *testing.T) {
	e := New()
	e.Properties[property.HP.String()] = defaultproperty.MakeIntCounterProperty(property.HP, 10, 10, property.Public, property.Character)

	origin1 := uuid.New()
	origin2 := uuid.New()

	buff1 := property.MakeTemporaryProperties(5)
	buff1.OriginEntityID = origin1
	buff1.Properties[property.HP.String()] = defaultproperty.MakeIntCounterProperty(property.HP, 3, 3, property.Public, property.Character)
	e.RegisterBuff(buff1)

	buff2 := property.MakeTemporaryProperties(0)
	buff2.Forever = true
	buff2.OriginEntityID = origin2
	buff2.Properties[property.HP.String()] = defaultproperty.MakeIntCounterProperty(property.HP, 7, 7, property.Public, property.Character)
	e.RegisterBuff(buff2)

	base, contribs := e.GetBuffAttributionFor(property.HP)

	if base.(*defaultproperty.DefaultIntCounterProperty).Value != 10 {
		t.Fatalf("expected base HP value 10, got %d", base.(*defaultproperty.DefaultIntCounterProperty).Value)
	}
	if len(contribs) != 2 {
		t.Fatalf("expected 2 contribs, got %d", len(contribs))
	}

	if contribs[0].Buff.OriginEntityID != origin1 {
		t.Errorf("expected contribs[0] to originate from origin1, got %s", contribs[0].Buff.OriginEntityID)
	}
	if contribs[0].Value.(*defaultproperty.DefaultIntCounterProperty).Value != 3 {
		t.Errorf("expected contribs[0] value 3, got %d", contribs[0].Value.(*defaultproperty.DefaultIntCounterProperty).Value)
	}

	if contribs[1].Buff.OriginEntityID != origin2 {
		t.Errorf("expected contribs[1] to originate from origin2, got %s", contribs[1].Buff.OriginEntityID)
	}
	if contribs[1].Value.(*defaultproperty.DefaultIntCounterProperty).Value != 7 {
		t.Errorf("expected contribs[1] value 7, got %d", contribs[1].Value.(*defaultproperty.DefaultIntCounterProperty).Value)
	}
}

// TestGetBuffAttributionFor_BlockNotCarryingKeyExcluded asserts that a buff
// block which does not carry the requested key is excluded from contribs,
// even though it exists on the entity and contributes to a different key.
//
// @test-link [[mechanic_buff_attribution_accessor]]
func TestGetBuffAttributionFor_BlockNotCarryingKeyExcluded(t *testing.T) {
	e := New()
	e.Properties[property.HP.String()] = defaultproperty.MakeIntCounterProperty(property.HP, 10, 10, property.Public, property.Character)
	e.Properties[property.Defense.String()] = defaultproperty.MakeIntProperty(property.Defense, 5, property.Public, property.Character)

	buff := property.MakeTemporaryProperties(5)
	buff.Properties[property.Defense.String()] = defaultproperty.MakeIntProperty(property.Defense, 2, property.Public, property.Character)
	e.RegisterBuff(buff)

	_, contribs := e.GetBuffAttributionFor(property.HP)

	if len(contribs) != 0 {
		t.Fatalf("expected buff block not carrying HP to be excluded from HP contribs, got %d", len(contribs))
	}
}

// TestGetBuffAttributionFor_ConsistentWithGetProperty asserts that composing
// the returned contributions' Values over base reproduces exactly what
// GetProperty(name) returns. A divergence here is a real finding about the
// composition path, not a test bug.
//
// @test-link [[mechanic_buff_attribution_accessor]]
func TestGetBuffAttributionFor_ConsistentWithGetProperty(t *testing.T) {
	e := New()
	e.Properties[property.HP.String()] = defaultproperty.MakeIntCounterProperty(property.HP, 10, 10, property.Public, property.Character)

	buff1 := property.MakeTemporaryProperties(5)
	buff1.Properties[property.HP.String()] = defaultproperty.MakeIntCounterProperty(property.HP, 3, 3, property.Public, property.Character)
	e.RegisterBuff(buff1)

	buff2 := property.MakeTemporaryProperties(0)
	buff2.Forever = true
	buff2.Properties[property.HP.String()] = defaultproperty.MakeIntCounterProperty(property.HP, 7, 7, property.Public, property.Character)
	e.RegisterBuff(buff2)

	base, contribs := e.GetBuffAttributionFor(property.HP)

	composed := base
	for _, c := range contribs {
		composed = composed.ApplyBuff(c.Value)
	}

	want := e.GetProperty(property.HP)

	composedIC := composed.(*defaultproperty.DefaultIntCounterProperty)
	wantIC := want.(*defaultproperty.DefaultIntCounterProperty)

	if composedIC.Value != wantIC.Value || composedIC.MaxValue != wantIC.MaxValue {
		t.Fatalf("composing GetBuffAttributionFor's contributions over base (%d/%d) diverges from GetProperty (%d/%d)",
			composedIC.Value, composedIC.MaxValue, wantIC.Value, wantIC.MaxValue)
	}
}
