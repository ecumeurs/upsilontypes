package skillgenerator

import (
	"strings"

	"github.com/ecumeurs/upsilontools/tools"
)

// Name returns a diegetic skill name (max 24 chars) for the given tags and grade.
// Template: [Modifier prefix] [Subject] [Suffix]
// @spec-link [[mech_skill_name_generation]]
func Name(primaryTag string, secondaryTags []string, grade string) string {
	prefix := pickPrefix(secondaryTags)
	subject := pickSubject(primaryTag)
	suffix := pickSuffix(grade)

	parts := make([]string, 0, 3)
	if prefix != "" {
		parts = append(parts, prefix)
	}
	if subject != "" {
		parts = append(parts, subject)
	}
	if suffix != "" {
		parts = append(parts, suffix)
	}

	name := strings.Join(parts, " ")
	if len(name) > 24 {
		name = name[:24]
	}
	return name
}

// pickPrefix selects a modifier prefix driven by the first matching secondary tag.
func pickPrefix(secondaryTags []string) string {
	prefixPools := map[string][]string{
		"dot":            {"Cinder", "Sludge", "Rot_", "Verm_"},
		"crit":           {"Razor", "Fang_", "Spike_"},
		"aoe":            {"Flux", "Cascade_", "Wave_"},
		"channeled":      {"Drift_", "Bleed_", "Slow_"},
		"buff":           {"Echo_", "Static", "Hex_"},
		"debuff":         {"Echo_", "Static", "Hex_"},
	}

	haxxorPool := []string{"", "Null_", "Void_", "Ghost_"}

	for _, tag := range secondaryTags {
		if pool, ok := prefixPools[tag]; ok {
			return pool[tools.RandomInt(0, len(pool))]
		}
	}
	return haxxorPool[tools.RandomInt(0, len(haxxorPool))]
}

// pickSubject selects the main noun from the primary tag's pool.
func pickSubject(primaryTag string) string {
	subjectPools := map[string][]string{
		"melee":    {"Strike", "Bash", "Cleaver", "Smash"},
		"ranged":   {"Bolt", "Lance", "Pulse", "Tracer"},
		"aoe":      {"Burst", "Field", "Storm", "Bloom"},
		"heal":     {"Mend", "Patch", "Pulse", "Suture"},
		"shield":   {"Bulwark", "Aegis", "Plate", "Shell"},
		"buff":     {"Flux", "Charge", "Surge", "Amp"},
		"debuff":   {"Wither", "Corrupt", "Drain", "Hex"},
		"dot":      {"Sear", "Fester", "Blight", "Venom"},
		"stun":     {"Jolt", "Stutter", "Lockdown", "Crack"},
		"trap":     {"Mine", "Snare", "Tripwire", "Hex"},
		"counter":  {"Riposte", "Ricochet", "Rebuke"},
		"reaction": {"Reflex", "Backlash"},
		"passive":  {"Aura", "Cycle", "Drift"},
		"mobility": {"Sprint", "Phase", "Vector"},
		"crit":     {"Pierce", "Fang", "Spike"},
		"channeled":{"Channel", "Sustain", "Draw"},
	}

	if pool, ok := subjectPools[primaryTag]; ok {
		return pool[tools.RandomInt(0, len(pool))]
	}
	return "Strike"
}

// pickSuffix returns a grade-flavored or haxxor suffix.
func pickSuffix(grade string) string {
	// 30% grade-flavored, 70% haxxor
	haxxorPool := []string{"_X", "v2", "_Z", "_Bot", "_666", "_Alpha"}
	gradeSuffixes := map[string]string{
		"I": "_I", "II": "_II", "III": "_III", "IV": "_IV", "V": "_V",
	}

	if tools.RandomInt(0, 100) < 30 {
		if s, ok := gradeSuffixes[grade]; ok {
			return s
		}
	}
	return haxxorPool[tools.RandomInt(0, len(haxxorPool))]
}
