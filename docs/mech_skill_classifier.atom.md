---
id: mech_skill_classifier
human_name: "Skill Classifier Mechanic"
type: MECHANIC
layer: IMPLEMENTATION
version: 1.0
status: DRAFT
priority: 3
tags: [skills, classification, tags]
parents:
  - [[module_skill_generator]]
dependents: []
---

# Skill Classifier Mechanic

## INTENT
To infer a set of ordered semantic tags from a skill's final properties, which are then used to drive diegetic naming and icon selection.

## THE RULE / LOGIC
Classification follows a 4-layer priority loop:
1. **Behavior:** `trap`, `counter`, `reaction`, `passive` (derived from `sk.Behavior`).
2. **Effect Family:** `heal`, `shield`, `dot`, `stun`, `buff`, `debuff` (based on non-zero property values).
3. **Delivery:** `aoe` (zone size > 1), `ranged` (range >= 2), `melee` (range == 1).
4. **Modifiers:** `crit` (chance >= 25%), `channeled` (channeling > 0), `instant` (delay <= 100), `mobility` (self-buff movement).

## TECHNICAL INTERFACE
- **Code Tag:** `@spec-link [[mech_skill_classifier]]`
- **File:** `upsilontypes/entity/skill/skillgenerator/classifier.go`
- **Primary Function:** `Classify(sk skill.Skill) []string`

## EXPECTATION
- Every skill receives at least one tag.
- Tags are ordered correctly to reflect behavior priority.
- Fallback to `melee` if no other tags are inferred.
