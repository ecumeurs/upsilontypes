package skillgenerator

import (
	"math/rand"
	"testing"

	"github.com/ecumeurs/upsilontypes/entity/skill"
	"github.com/ecumeurs/upsilontypes/entity/skill/skillweight"
	"github.com/ecumeurs/upsilontypes/property"
	"github.com/ecumeurs/upsilontypes/property/def"
	"github.com/ecumeurs/upsilontools/tools"
)

// @test-link [[mech_skill_generator_blueprint]]
// @test-link [[mechanic_effect_stun]]
// @test-link [[mechanic_effect_poison]]

// resetRand restores the default random function after deterministic tests.
func resetRand() {
	tools.TesterRand(rand.Intn)
}

// assertBalanced checks that a skill has netSW=0 (after applyDelayCloser) and PSW within the given band.
func assertBalanced(t *testing.T, name string, sk skill.Skill, minPSW, maxPSW int) {
	t.Helper()
	// Producers don't call applyDelayCloser — only Generate() does.
	// Apply it here so we can verify the skill balances to netSW=0.
	applyDelayCloser(&sk)
	pSW, _, netSW := skillweight.Calculate(&sk)
	if netSW != 0 {
		t.Errorf("%s: netSW=%d (expected 0), pSW=%d", name, netSW, pSW)
	}
	if minPSW > 0 && pSW < minPSW {
		t.Errorf("%s: PSW=%d below minimum %d", name, pSW, minPSW)
	}
	if maxPSW > 0 && pSW > maxPSW {
		t.Errorf("%s: PSW=%d above maximum %d", name, pSW, maxPSW)
	}
}

// --- Melee ---

func TestProduceMelee_DamageOnly(t *testing.T) {
	// Force the random path that skips stun (random >= 50)
	tools.TesterRand(func(n int) int { return n - 1 })
	defer resetRand()

	sk := produceMelee(100)
	assertBalanced(t, "melee-dmg", sk, 0, 0)

	rng := sk.GetPropertyC(property.Range).GetMaxValue()
	if rng != 1 {
		t.Errorf("melee range: expected 1, got %d", rng)
	}

	dmg := sk.GetPropertyI(property.Damage).I()
	if dmg <= 0 {
		t.Errorf("melee damage: expected > 0, got %d", dmg)
	}

	tt := sk.GetProperty(property.TargetType).Get().(string)
	if tt != string(def.TargetTypeEnemyOnly) {
		t.Errorf("melee target type: expected EnemyOnly, got %s", tt)
	}
}

func TestProduceMelee_WithStunPairing(t *testing.T) {
	// Force the stun variant path (random < 50 for stun decision, targetPSW >= 80)
	tools.TesterRand(func(n int) int { return 0 })
	defer resetRand()

	sk := produceMelee(120)
	assertBalanced(t, "melee-stun", sk, 0, 0)

	stunChance := sk.GetPropertyI(property.StunChance).I()
	stunPower := sk.GetPropertyI(property.StunPower).I()

	if stunChance <= 0 {
		t.Errorf("melee stun variant: expected StunChance > 0, got %d", stunChance)
	}
	// ISS-095: StunPower MUST be paired with StunChance
	if stunPower <= 0 {
		t.Errorf("melee stun variant: StunPower must be > 0 when StunChance=%d (ISS-095 pairing fix)", stunChance)
	}
}

// --- Ranged ---

func TestProduceRanged_RangeAndDamage(t *testing.T) {
	tools.TesterRand(func(n int) int { return 0 })
	defer resetRand()

	sk := produceRanged(100)
	assertBalanced(t, "ranged", sk, 0, 0)

	rng := sk.GetPropertyC(property.Range).GetMaxValue()
	if rng < 2 || rng > 4 {
		t.Errorf("ranged range: expected [2,4], got %d", rng)
	}

	dmg := sk.GetPropertyI(property.Damage).I()
	if dmg <= 0 {
		t.Errorf("ranged damage: expected > 0, got %d", dmg)
	}

	tt := sk.GetProperty(property.TargetType).Get().(string)
	if tt != string(def.TargetTypeEnemyOnly) {
		t.Errorf("ranged target type: expected EnemyOnly, got %s", tt)
	}
}

// --- AoE ---

func TestProduceAOE_ZoneAndDamage(t *testing.T) {
	sk := produceAOE(150)
	assertBalanced(t, "aoe", sk, 0, 0)

	zoneProp, ok := sk.Targeting[property.Zone.String()]
	if !ok {
		t.Fatal("aoe: missing zone property")
	}
	zp, ok := zoneProp.(*def.ZoneProperty)
	if !ok {
		t.Fatal("aoe: zone is not ZoneProperty")
	}
	if len(zp.ZonePattern) < 2 {
		t.Errorf("aoe: expected zone cells >= 2, got %d", len(zp.ZonePattern))
	}

	tt := sk.GetProperty(property.TargetType).Get().(string)
	if tt != string(def.TargetTypeEnemyOnly) {
		t.Errorf("aoe target type: expected EnemyOnly, got %s", tt)
	}
}

// --- Heal ---

func TestProduceHeal_FriendOnly(t *testing.T) {
	sk := produceHeal(100)
	assertBalanced(t, "heal", sk, 0, 0)

	heal := sk.GetPropertyI(property.Heal).I()
	if heal <= 0 {
		t.Errorf("heal: expected Heal > 0, got %d", heal)
	}

	tt := sk.GetProperty(property.TargetType).Get().(string)
	if tt != string(def.TargetTypeFriendOnly) {
		t.Errorf("heal target type: expected FriendOnly, got %s", tt)
	}
}

// --- Shield ---

func TestProduceShield_ShieldPower(t *testing.T) {
	tools.TesterRand(func(n int) int { return 0 })
	defer resetRand()

	sk := produceShield(100)
	assertBalanced(t, "shield", sk, 0, 0)

	sp := sk.GetPropertyI(property.ShieldPower).I()
	if sp <= 0 {
		t.Errorf("shield: expected ShieldPower > 0, got %d", sp)
	}

	tt := sk.GetProperty(property.TargetType).Get().(string)
	if tt != string(def.TargetTypeSelf) && tt != string(def.TargetTypeFriendOnly) {
		t.Errorf("shield target type: expected Self or FriendOnly, got %s", tt)
	}
}

// --- Buff ---

func TestProduceBuff_CritDuration(t *testing.T) {
	sk := produceBuff(120)
	assertBalanced(t, "buff", sk, 0, 0)

	dur := sk.GetPropertyC(property.Duration).GetMaxValue()
	if dur != 3 {
		t.Errorf("buff: expected Duration=3, got %d", dur)
	}

	crit := sk.GetPropertyI(property.CriticalChance).I()
	if crit <= 0 {
		t.Errorf("buff: expected CritChance > 0, got %d", crit)
	}

	tt := sk.GetProperty(property.TargetType).Get().(string)
	if tt != string(def.TargetTypeSelf) {
		t.Errorf("buff target type: expected Self, got %s", tt)
	}
}

// --- Debuff ---

func TestProduceDebuff_PoisonDurationPairing(t *testing.T) {
	sk := produceDebuff(120)
	assertBalanced(t, "debuff", sk, 0, 0)

	dur := sk.GetPropertyC(property.Duration).GetMaxValue()
	if dur != 3 {
		t.Errorf("debuff: expected Duration=3, got %d", dur)
	}

	pp := sk.GetPropertyI(property.PoisonPower).I()
	if pp <= 0 {
		t.Errorf("debuff: expected PoisonPower > 0, got %d", pp)
	}

	// ISS-095: PoisonChance MUST be paired with PoisonPower
	pc := sk.GetPropertyI(property.PoisonChance).I()
	if pc <= 0 {
		t.Errorf("debuff: PoisonChance must be > 0 when PoisonPower=%d (ISS-095 pairing fix)", pp)
	}

	tt := sk.GetProperty(property.TargetType).Get().(string)
	if tt != string(def.TargetTypeEnemyOnly) {
		t.Errorf("debuff target type: expected EnemyOnly, got %s", tt)
	}
}

// --- DOT ---

func TestProduceDot_PoisonPairing(t *testing.T) {
	sk := produceDot(100)
	assertBalanced(t, "dot", sk, 0, 0)

	pp := sk.GetPropertyI(property.PoisonPower).I()
	if pp <= 0 {
		t.Errorf("dot: expected PoisonPower > 0, got %d", pp)
	}

	// ISS-095: PoisonChance MUST be paired with PoisonPower
	pc := sk.GetPropertyI(property.PoisonChance).I()
	if pc <= 0 {
		t.Errorf("dot: PoisonChance must be > 0 when PoisonPower=%d (ISS-095 pairing fix)", pp)
	}

	tt := sk.GetProperty(property.TargetType).Get().(string)
	if tt != string(def.TargetTypeEnemyOnly) {
		t.Errorf("dot target type: expected EnemyOnly, got %s", tt)
	}
}

// --- Stun ---

func TestProduceStun_StunPairing(t *testing.T) {
	sk := produceStun(100)
	assertBalanced(t, "stun", sk, 0, 0)

	sc := sk.GetPropertyI(property.StunChance).I()
	if sc <= 0 {
		t.Errorf("stun: expected StunChance > 0, got %d", sc)
	}

	// ISS-095: StunPower MUST be paired with StunChance
	sp := sk.GetPropertyI(property.StunPower).I()
	if sp <= 0 {
		t.Errorf("stun: StunPower must be > 0 when StunChance=%d (ISS-095 pairing fix)", sc)
	}

	tt := sk.GetProperty(property.TargetType).Get().(string)
	if tt != string(def.TargetTypeEnemyOnly) {
		t.Errorf("stun target type: expected EnemyOnly, got %s", tt)
	}
}

// --- Trap ---

func TestProduceTrap_TileBehavior(t *testing.T) {
	sk := produceTrap(100)
	assertBalanced(t, "trap", sk, 0, 0)

	bh := sk.Behavior.Get().(string)
	if bh != string(def.BehaviorTypeTrap) {
		t.Errorf("trap behavior: expected Trap, got %s", bh)
	}

	tt := sk.GetProperty(property.TargetType).Get().(string)
	if tt != string(def.TargetTypeTile) {
		t.Errorf("trap target type: expected Tile, got %s", tt)
	}
}

// --- Passive ---

func TestProducePassive_PassiveBehavior(t *testing.T) {
	tools.TesterRand(func(n int) int { return 0 })
	defer resetRand()

	sk := producePassive(120)
	assertBalanced(t, "passive", sk, 0, 0)

	bh := sk.Behavior.Get().(string)
	if bh != string(def.BehaviorTypePassive) {
		t.Errorf("passive behavior: expected Passive, got %s", bh)
	}

	tt := sk.GetProperty(property.TargetType).Get().(string)
	if tt != string(def.TargetTypeSelf) {
		t.Errorf("passive target type: expected Self, got %s", tt)
	}
}

// --- Disabled Producers ---

func TestDisabledProducers_NotInAllTags(t *testing.T) {
	disabled := []string{"counter", "reaction", "mobility"}
	for _, tag := range disabled {
		for _, active := range allProducerTags {
			if active == tag {
				t.Errorf("disabled producer %q should not be in allProducerTags (ISS-095 #6)", tag)
			}
		}
		if _, ok := producers[tag]; ok {
			t.Errorf("disabled producer %q should not be in producers map (ISS-095 #6)", tag)
		}
	}
}

// --- Secondary Layer Pairing ---

func TestLayerStun_PairsStunPower(t *testing.T) {
	bp := newBlueprint()
	bp.addDamage(100)
	sk := bp.build()

	layerStun(&sk, 60)

	sc := sk.GetPropertyI(property.StunChance).I()
	sp := sk.GetPropertyI(property.StunPower).I()

	if sc <= 0 {
		t.Errorf("layerStun: expected StunChance > 0, got %d", sc)
	}
	if sp <= 0 {
		t.Errorf("layerStun: StunPower must be > 0 when StunChance=%d (ISS-095 pairing fix)", sc)
	}
}

func TestLayerDot_PairsPoisonChance(t *testing.T) {
	bp := newBlueprint()
	bp.addDamage(100)
	sk := bp.build()

	layerDot(&sk, 60)

	pp := sk.GetPropertyI(property.PoisonPower).I()
	pc := sk.GetPropertyI(property.PoisonChance).I()

	if pp <= 0 {
		t.Errorf("layerDot: expected PoisonPower > 0, got %d", pp)
	}
	if pc <= 0 {
		t.Errorf("layerDot: PoisonChance must be > 0 when PoisonPower=%d (ISS-095 pairing fix)", pp)
	}
}

func TestLayerDebuff_PairsPoisonChance(t *testing.T) {
	bp := newBlueprint()
	bp.setTargetType(def.TargetTypeEnemyOnly)
	bp.addDamage(100)
	sk := bp.build()

	layerDebuff(&sk, 60)

	pp := sk.GetPropertyI(property.PoisonPower).I()
	pc := sk.GetPropertyI(property.PoisonChance).I()

	if pp <= 0 {
		t.Errorf("layerDebuff: expected PoisonPower > 0, got %d", pp)
	}
	if pc <= 0 {
		t.Errorf("layerDebuff: PoisonChance must be > 0 when PoisonPower=%d (ISS-095 pairing fix)", pp)
	}
}
