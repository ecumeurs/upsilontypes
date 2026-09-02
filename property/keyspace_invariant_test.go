package property

import "testing"

// TestKeySpace_IdentifierMatchesStringValue is a hand-written, red-first specification for the
// Property Key Space Unification round (architecture/property_key_vocabulary.md §1, §7).
//
// §1 states the target invariant: constant identifier == string value, for all 64 keys. §7 lists
// the 7-entry rename map that gets the current 64 keys there. Until the unified property.Key +
// def registry lands, this table is hand-maintained against the current flat property.Key
// declarations in propertyenum.go (formerly split across the now-removed three-namespace,
// Entity-/Skill-/Item-scoped type aliases). Once the rename lands, this table's left column
// moves to the new identifiers and the whole test collapses into a walk over the future registry
// (asserting entry.Key == property.Key(<identifier>) for every entry) — it is deliberately a
// hand-written table only until then.
//
// @test-link [[upsilontypes:module_property_key_registry]]
func TestKeySpace_IdentifierMatchesStringValue(t *testing.T) {
	type row struct {
		identifier    string // the Go constant's own spelling, for failure messages
		currentValue  string // PropertyToString(constant) today
		expectedFinal string // the string value this key must have once the rename lands (§7)
	}

	rows := []row{
		// --- Entity scope (20) ---
		{"HP", PropertyToString(HP), "HP"},
		{"Movement", PropertyToString(Movement), "Movement"},
		{"SP", PropertyToString(SP), "SP"},
		{"MP", PropertyToString(MP), "MP"},
		{"Attack", PropertyToString(Attack), "Attack"},
		{"Defense", PropertyToString(Defense), "Defense"},
		{"JumpHeight", PropertyToString(JumpHeight), "JumpHeight"},
		{"TeamID", PropertyToString(TeamID), "TeamID"},
		{"IsDying", PropertyToString(IsDying), "IsDying"},
		{"HasMoved", PropertyToString(HasMoved), "HasMoved"},
		{"HasActed", PropertyToString(HasActed), "HasActed"},
		{"AttackRange", PropertyToString(AttackRange), "AttackRange"},
		{"EntityDuration", PropertyToString(EntityDuration), "EntityDuration"},
		{"ExpiresWithCaster", PropertyToString(ExpiresWithCaster), "ExpiresWithCaster"},
		{"WalkThrough", PropertyToString(WalkThrough), "WalkThrough"},
		{"Invisible", PropertyToString(Invisible), "Invisible"},
		{"AIArchetype", PropertyToString(AIArchetype), "AIArchetype"},
		// entity Shield/Poison/Stun keep their bare names; the skill-side counterparts rename below.
		{"Shield (entity)", PropertyToString(Shield), "Shield"},
		{"Poison (entity)", PropertyToString(Poison), "Poison"},
		{"Stun (entity)", PropertyToString(Stun), "Stun"},

		// --- Skill scope: targeting (9) ---
		{"Behavior", PropertyToString(Behavior), "Behavior"},
		{"Range", PropertyToString(Range), "Range"},
		{"Zone", PropertyToString(Zone), "Zone"},
		{"TargetNumber", PropertyToString(TargetNumber), "TargetNumber"},
		{"Accuracy", PropertyToString(Accuracy), "Accuracy"},
		{"Dodge", PropertyToString(Dodge), "Dodge"},
		{"Parry", PropertyToString(Parry), "Parry"},
		{"TargetType", PropertyToString(TargetType), "TargetType"},
		{"TargetingMechanics", PropertyToString(TargetingMechanics), "TargetingMechanics"},

		// --- Skill scope: effect (10) ---
		// RENAME: Damage currently stringifies "Damage"; §7 renames it to DamageScale/"DamageScale".
		{"DamageScale", PropertyToString(DamageScale), "DamageScale"},
		{"Heal", PropertyToString(Heal), "Heal"},
		// RENAME (ISS-147 family): ShieldPower currently stringifies "Shield", colliding with the
		// entity-scope Shield key above. §7 renames the value to "ShieldPower".
		{"ShieldPower", PropertyToString(ShieldPower), "ShieldPower"},
		// RENAME (ISS-147 family): StunPower currently stringifies "Stun", colliding with the
		// entity-scope Stun key above. §7 renames the value to "StunPower".
		{"StunPower", PropertyToString(StunPower), "StunPower"},
		{"StunChance", PropertyToString(StunChance), "StunChance"},
		{"CriticalChance", PropertyToString(CriticalChance), "CriticalChance"},
		{"CriticalMultiplier", PropertyToString(CriticalMultiplier), "CriticalMultiplier"},
		{"Duration", PropertyToString(Duration), "Duration"},
		// RENAME — ISS-147 itself: PoisonPower currently stringifies "Poison", colliding with the
		// entity-scope Poison key above. §7 renames the value to "PoisonPower".
		{"PoisonPower", PropertyToString(PoisonPower), "PoisonPower"},
		{"PoisonChance", PropertyToString(PoisonChance), "PoisonChance"},

		// --- Skill scope: cost (7) ---
		{"Delay", PropertyToString(Delay), "Delay"},
		{"Channeling", PropertyToString(Channeling), "Channeling"},
		{"HPLeech", PropertyToString(HPLeech), "HPLeech"},
		{"MPLeech", PropertyToString(MPLeech), "MPLeech"},
		{"SPLeech", PropertyToString(SPLeech), "SPLeech"},
		// RENAME: MvtCost currently stringifies "MvtCost"; §7 renames it to MovementCost/"MovementCost".
		{"MovementCost", PropertyToString(MovementCost), "MovementCost"},
		{"Cooldown", PropertyToString(Cooldown), "Cooldown"},

		// --- Skill scope: reposition (2) ---
		{"RepositionSubject", PropertyToString(RepositionSubject), "RepositionSubject"},
		{"RepositionDistance", PropertyToString(RepositionDistance), "RepositionDistance"},

		// --- Skill scope: trigger (3) ---
		{"TriggerType", PropertyToString(TriggerType), "TriggerType"},
		{"RemoveOnTrigger", PropertyToString(RemoveOnTrigger), "RemoveOnTrigger"},
		{"TriggerCount", PropertyToString(TriggerCount), "TriggerCount"},

		// --- Item scope (13) ---
		{"Durability", PropertyToString(Durability), "Durability"},
		{"Weight", PropertyToString(Weight), "Weight"},
		{"ItemType", PropertyToString(ItemType), "ItemType"},
		// RENAME: ArmorRating currently stringifies "Armor"; §7 renames the value to "ArmorRating"
		// to restore the §1 invariant.
		{"ArmorRating", PropertyToString(ArmorRating), "ArmorRating"},
		{"WeaponType", PropertyToString(WeaponType), "WeaponType"},
		{"ArmorType", PropertyToString(ArmorType), "ArmorType"},
		{"ToolType", PropertyToString(ToolType), "ToolType"},
		{"WeaponRange", PropertyToString(WeaponRange), "WeaponRange"},
		{"WeaponBaseDamage", PropertyToString(WeaponBaseDamage), "WeaponBaseDamage"},
		{"Stackable", PropertyToString(Stackable), "Stackable"},
		{"StackSize", PropertyToString(StackSize), "StackSize"},
		{"Effect", PropertyToString(Effect), "Effect"},
		// RENAME: Value currently stringifies "Value"; §7 renames it to ItemValue/"ItemValue"
		// (too generic once the three namespaces merge into one flat key space).
		{"ItemValue", PropertyToString(ItemValue), "ItemValue"},
	}

	if got, want := len(rows), 64; got != want {
		t.Fatalf("keyspace table has %d rows, want %d (the frozen 64 keys)", got, want)
	}

	for _, r := range rows {
		r := r
		t.Run(r.identifier, func(t *testing.T) {
			// Assertion 1: today's string value must equal the final (post-rename) value.
			// This is red on exactly the 7 keys in the §7 rename map; green on the other 57.
			if r.currentValue != r.expectedFinal {
				t.Errorf("%s: current string value %q does not yet match the target final value %q "+
					"(pending rename per architecture/property_key_vocabulary.md §7)",
					r.identifier, r.currentValue, r.expectedFinal)
			}

			// Assertion 2: the §1 invariant. Once the rename lands, the identifier's own spelling
			// (stripped of any disambiguating suffix used only in this table, e.g. "(entity)") must
			// equal the final string value — documenting *why* the 7 renames exist.
			bareIdentifier := r.identifier
			if idx := indexOfSpace(bareIdentifier); idx != -1 {
				bareIdentifier = bareIdentifier[:idx]
			}
			if bareIdentifier != r.expectedFinal {
				t.Errorf("%s: §1 invariant violated — identifier %q does not match target final value %q",
					r.identifier, bareIdentifier, r.expectedFinal)
			}
		})
	}
}

// indexOfSpace returns the index of the first space in s, or -1 if none. Local helper kept tiny
// and dependency-free since this file is deliberately throwaway scaffolding (see file comment).
func indexOfSpace(s string) int {
	for i, c := range s {
		if c == ' ' {
			return i
		}
	}
	return -1
}
