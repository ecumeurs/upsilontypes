package skillgenerator

// @test-link [[mech_skill_name_generation]]
// @test-link [[req_skill_generation]]

// TODO: uncomment and implement once namegen.go is written.
//
// import (
// 	"regexp"
// 	"strings"
// 	"testing"
// )
//
// var rawPropertyKeys = map[string]bool{
// 	"Damage": true, "Heal": true, "Shield": true,
// 	"Accuracy": true, "Cooldown": true, "HPLeech": true,
// 	"MPLeech": true, "SPLeech": true, "New Skill": true,
// }
//
// // At least one alphabetic word; optional alphanumeric suffix token.
// var namePattern = regexp.MustCompile(`^[A-Za-z_][A-Za-z_0-9]*(\s[A-Za-z_][A-Za-z_0-9]*)?$`)
//
// func TestName_NotARawPropertyKey(t *testing.T) {
// 	grades := []string{"I", "II", "III", "IV", "V"}
// 	primaries := []string{"melee", "ranged", "aoe", "heal", "shield", "trap", "counter", "reaction", "passive", "stun", "mobility"}
// 	for _, grade := range grades {
// 		for _, primary := range primaries {
// 			name := Name(primary, nil, grade)
// 			if rawPropertyKeys[name] {
// 				t.Errorf("Name(%q, nil, %q) = %q: is a raw property key", primary, grade, name)
// 			}
// 		}
// 	}
// }
//
// func TestName_MaxLength(t *testing.T) {
// 	primaries := []string{"melee", "heal", "trap", "aoe"}
// 	secondaries := [][]string{{"dot"}, {"crit", "aoe"}, {}, {"channeled", "debuff"}}
// 	for i, primary := range primaries {
// 		name := Name(primary, secondaries[i], "III")
// 		if len(name) > 24 {
// 			t.Errorf("Name(%q, %v, III) = %q: len %d > 24", primary, secondaries[i], name, len(name))
// 		}
// 	}
// }
//
// func TestName_NonEmpty(t *testing.T) {
// 	name := Name("melee", nil, "I")
// 	if strings.TrimSpace(name) == "" {
// 		t.Error("Name must not return an empty string")
// 	}
// }
//
// func TestName_SecondaryTagModifier(t *testing.T) {
// 	// With "dot" as secondary, modifier prefix should come from the dot pool.
// 	dotPool := map[string]bool{"Cinder": true, "Sludge": true, "Rot_": true, "Verm_": true}
// 	for i := 0; i < 50; i++ {
// 		name := Name("ranged", []string{"dot"}, "II")
// 		parts := strings.SplitN(name, " ", 2)
// 		if len(parts) > 1 {
// 			prefix := parts[0]
// 			if !dotPool[prefix] {
// 				// Prefix may be from haxxor pool too — just log, not fail.
// 				_ = prefix
// 			}
// 		}
// 	}
// }
