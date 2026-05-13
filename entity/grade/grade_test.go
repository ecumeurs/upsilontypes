package grade_test

// @test-link [[mechanic_ai_progression_matching]]

import (
	"testing"

	"github.com/ecumeurs/upsilontypes/entity/grade"
)

// TestGradeFromWins verifies the wins-to-grade mapping at all tier boundaries including the Grade V hard cap.
func TestGradeFromWins(t *testing.T) {
	cases := []struct {
		wins int
		want string
	}{
		// Grade I: 0–5
		{0, "I"}, {1, "I"}, {5, "I"},
		// Grade I+: 6–9
		{6, "I+"}, {9, "I+"},
		// Grade II: 10–15
		{10, "II"}, {15, "II"},
		// Grade II+: 16–19
		{16, "II+"}, {19, "II+"},
		// Grade III: 20–25
		{20, "III"}, {25, "III"},
		// Grade III+: 26–29
		{26, "III+"}, {29, "III+"},
		// Grade IV: 30–35
		{30, "IV"}, {35, "IV"},
		// Grade IV+: 36–39
		{36, "IV+"}, {39, "IV+"},
		// Grade V hard cap: 40 and above
		{40, "V"}, {41, "V"}, {100, "V"}, {9999, "V"},
		// Negative clamped to 0
		{-1, "I"}, {-100, "I"},
	}
	for _, tc := range cases {
		got := grade.GradeFromWins(tc.wins)
		if got != tc.want {
			t.Errorf("GradeFromWins(%d) = %q, want %q", tc.wins, got, tc.want)
		}
	}
}

// TestGradeIndex verifies that every canonical grade string maps to the correct zero-based index.
func TestGradeIndex(t *testing.T) {
	cases := []struct {
		g    string
		want int
	}{
		{"I", 0}, {"I+", 1}, {"II", 2}, {"II+", 3},
		{"III", 4}, {"III+", 5}, {"IV", 6}, {"IV+", 7}, {"V", 8},
	}
	for _, tc := range cases {
		got, err := grade.GradeIndex(tc.g)
		if err != nil {
			t.Errorf("GradeIndex(%q) unexpected error: %v", tc.g, err)
			continue
		}
		if got != tc.want {
			t.Errorf("GradeIndex(%q) = %d, want %d", tc.g, got, tc.want)
		}
	}

	_, err := grade.GradeIndex("VI")
	if err == nil {
		t.Error("GradeIndex(\"VI\") expected error for unknown grade")
	}
}

// TestCPForGrade verifies the CP formula yields 100+50*index for each grade.
func TestCPForGrade(t *testing.T) {
	cases := []struct {
		g    string
		want int
	}{
		{"I", 100}, {"I+", 150}, {"II", 200}, {"II+", 250},
		{"III", 300}, {"III+", 350}, {"IV", 400}, {"IV+", 450}, {"V", 500},
	}
	for _, tc := range cases {
		got := grade.CPForGrade(tc.g)
		if got != tc.want {
			t.Errorf("CPForGrade(%q) = %d, want %d", tc.g, got, tc.want)
		}
	}

	if cp := grade.CPForGrade("UNKNOWN"); cp != 0 {
		t.Errorf("CPForGrade(\"UNKNOWN\") = %d, want 0", cp)
	}
}

// TestGradeFromWinsRoundTrip checks that wins 0–50 all produce a CP value in the valid range [100, 500].
func TestGradeFromWinsRoundTrip(t *testing.T) {
	// GradeFromWins → CPForGrade should always give a valid non-zero result.
	for wins := 0; wins <= 50; wins++ {
		g := grade.GradeFromWins(wins)
		cp := grade.CPForGrade(g)
		if cp < 100 || cp > 500 {
			t.Errorf("wins=%d grade=%q CP=%d out of expected range [100,500]", wins, g, cp)
		}
	}
}
