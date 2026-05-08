---
id: mech_skill_generator_blueprint
human_name: "Skill Generator Blueprint DSL"
type: MECHANIC
layer: IMPLEMENTATION
version: 1.0
status: DRAFT
priority: 3
tags: [skills, generation, blueprint, dsl, iss-095]
parents:
  - [[mech_skill_generator_core]]
dependents: []
---

# Skill Generator Blueprint DSL

## INTENT
To provide a fluent, builder-style interface for constructing skills while encapsulating the complexity of property initialization and SW-budget tracking.

## THE RULE / LOGIC
- **Builder Pattern:** The `blueprint` struct tracks a `skill.Skill` instance.
- **Fluent Methods:** Provides methods like `addDamage()`, `setRange()`, `addStunChance()`, `addStunPower()`, `addPoisonChance()`, `addPoisonPower()` that return the blueprint for chaining.
- **Status Effect Pairing (ISS-095):** StunChance must always be paired with StunPower, and PoisonPower must always be paired with PoisonChance. Omitting either side renders the status effect mechanically dead in the effect applicator.
- **Internal Balancing:** Methods like `setRange()` and `addDamage()` handle the specific property keys and default values (e.g., setting LoS targeting by default).
- **PSW Reporting:** The `psw()` method provides real-time access to the skill's current Positive Skill Weight during the build process.
- **PoisonChance Clamping:** `addPoisonChance()` clamps input to [1, 100] range.

## TECHNICAL INTERFACE
- **Code Tag:** `@spec-link [[mech_skill_generator_blueprint]]`
- **File:** `upsilontypes/entity/skill/skillgenerator/blueprint.go`
- **Primary Struct:** `blueprint`

## EXPECTATION
- Isolated logic for skill property manipulation.
- Deterministic initialization of critical properties (Delay, Targeting).
