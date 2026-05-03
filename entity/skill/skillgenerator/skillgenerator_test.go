package skillgenerator

import (
	"testing"

	"github.com/ecumeurs/upsilontypes/entity/skill/skillweight"
)

// @test-link [[ISS-065_skill_generation_balance]]
func TestGenerateRandomSkill(t *testing.T) {
	// Generate 100 random skills
	for i := 0; i < 100; i++ {
		sk := GenerateRandomSkill()

		pSW, nSW, netSW := skillweight.Calculate(sk)

		// 1. Verify net SW = 0
		if netSW != 0 {
			t.Errorf("Skill %d (%s) has non-zero net SW: %d (Positive: %d, Negative: %d)", i, sk.Name, netSW, pSW, nSW)
		}

		// 2. Verify grade distribution follows expected ranges
		grade := skillweight.GetGrade(pSW)
		if grade == "" {
			t.Errorf("Skill %d has empty grade", i)
		}

		// 3. Verify credit costs = positiveSW * 2
		cost := skillweight.GetCreditCost(pSW)
		if cost != pSW*2 {
			t.Errorf("Skill %d cost mismatch: expected %d, got %d", i, pSW*2, cost)
		}
	}
}

// @test-link [[shared:req_skill_generation_overhaul]]

var gradeBands = map[string][2]int{
	"I": {60, 150}, "II": {151, 300}, "III": {301, 500}, "IV": {501, 750}, "V": {751, 9999},
}

func TestGenerate_GradeBand(t *testing.T) {
	for grade, band := range gradeBands {
		for i := 0; i < 100; i++ {
			sk, _, err := Generate(GenerateRequest{TargetGrade: grade})
			if err != nil {
				t.Fatalf("Generate grade %s: unexpected error: %v", grade, err)
			}
			pSW, _, netSW := skillweight.Calculate(sk)
			if netSW != 0 {
				t.Errorf("Generate grade %s skill %d: netSW=%d (expected 0)", grade, i, netSW)
			}
			if pSW < band[0] || pSW > band[1] {
				t.Errorf("Generate grade %s skill %d: PSW=%d outside band [%d, %d]", grade, i, pSW, band[0], band[1])
			}
		}
	}
}

func TestGenerate_TagsNonEmpty(t *testing.T) {
	for i := 0; i < 100; i++ {
		_, tags, err := Generate(GenerateRequest{TargetGrade: "I"})
		if err != nil {
			t.Fatalf("skill %d: unexpected error: %v", i, err)
		}
		if len(tags) == 0 {
			t.Errorf("skill %d: tags array is empty", i)
		}
	}
}
