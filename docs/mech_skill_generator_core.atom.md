---
id: mech_skill_generator_core
human_name: "Skill Generator Core Mechanic"
type: MECHANIC
layer: IMPLEMENTATION
version: 1.0
status: DRAFT
priority: 3
tags: [skills, generation, dispatcher]
parents:
  - [[module_skill_generator]]
dependents:
  - [[mech_skill_generator_blueprint]]
---

# Skill Generator Core Mechanic

## INTENT
To implement the dispatcher logic that selects primary producers, layers secondary properties, and ensures the final skill satisfies grade-specific PSW bands and net-zero SW balance.

## THE RULE / LOGIC
1. **Band Calculation:** Retrieve PSW range for the target grade from `gradeBand`.
2. **Primary Dispatch:** Randomly select an allowed primary producer (e.g., `melee`, `heal`).
3. **Secondary Layering:** Apply a secondary layer (e.g., `dot`, `aoe`) with grade-dependent probability.
4. **Cost Balancing:** Use `applyDelayCloser` to adjust the `Delay` property until the skill's net SW is exactly zero.
5. **Orchestration:** Pass the resulting skill to the classifier and name generator.

## TECHNICAL INTERFACE
- **Code Tag:** `@spec-link [[mech_skill_generator_core]]`
- **File:** `upsilontypes/entity/skill/skillgenerator/skillgenerator.go`
- **Primary Function:** `Generate(req GenerateRequest)`

## EXPECTATION
- The dispatcher always produces a skill within the requested grade's PSW band.
- Net SW is always zero after `applyDelayCloser`.
