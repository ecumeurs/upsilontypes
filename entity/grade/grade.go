// Package grade defines the AI grade system used for archetype progression.
//
// Grades run from I to V (hard cap). Each grade represents a bracket of
// player wins that determines the AI's CP budget at character creation.
//
// Grade table:
//
//	I    →  0– 5 wins  (idx 0, 100 CP)
//	I+   →  6– 9 wins  (idx 1, 150 CP)
//	II   → 10–15 wins  (idx 2, 200 CP)
//	II+  → 16–19 wins  (idx 3, 250 CP)
//	III  → 20–25 wins  (idx 4, 300 CP)
//	III+ → 26–29 wins  (idx 5, 350 CP)
//	IV   → 30–35 wins  (idx 6, 400 CP)
//	IV+  → 36–39 wins  (idx 7, 450 CP)
//	V    → 40+  wins   (idx 8, 500 CP)  ← hard cap, no V+
//
// @spec-link [[mechanic_ai_progression_matching]]
package grade

import "fmt"

// grades is the ordered list of grade strings, index 0..8.
var grades = []string{"I", "I+", "II", "II+", "III", "III+", "IV", "IV+", "V"}

// GradeFromWins derives the AI grade string from a player's total win count.
// Wins < 0 are clamped to 0. Wins ≥ 40 return "V" (hard cap).
func GradeFromWins(wins int) string {
	if wins < 0 {
		wins = 0
	}
	if wins >= 40 {
		return "V"
	}
	main := wins / 10    // 0..3 → base grade I/II/III/IV
	plus := wins%10 >= 6 // X6-X9 → "+" variant
	if plus {
		return grades[main*2+1]
	}
	return grades[main*2]
}

// GradeIndex returns the zero-based index (0..8) for a grade string.
// Returns an error for unknown grades.
func GradeIndex(g string) (int, error) {
	for i, v := range grades {
		if v == g {
			return i, nil
		}
	}
	return 0, fmt.Errorf("grade: unknown grade %q", g)
}

// MustGradeIndex is like GradeIndex but panics on unknown grades.
// Suitable for use with compile-time-known grade strings.
func MustGradeIndex(g string) int {
	i, err := GradeIndex(g)
	if err != nil {
		panic(err)
	}
	return i
}

// CPForGrade returns the CP budget for a given grade string.
// Formula: 100 + 50 × GradeIndex.
// Returns 0 for unknown grades.
func CPForGrade(g string) int {
	i, err := GradeIndex(g)
	if err != nil {
		return 0
	}
	return 100 + 50*i
}

// ValidGrade reports whether g is a known grade string.
func ValidGrade(g string) bool {
	_, err := GradeIndex(g)
	return err == nil
}

// AllGrades returns the ordered slice of all grade strings.
func AllGrades() []string {
	cp := make([]string, len(grades))
	copy(cp, grades)
	return cp
}
