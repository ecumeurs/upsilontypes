---
id: rule_character_create_character
human_name: Character creation base allocation
type: RULE
layer: ARCHITECTURE
version: 2.0
status: STABLE
priority: 5
tags: [character, creation]
parents:
  - [[entity_character]]
dependents: []
---
# Character creation base allocation

## INTENT
Define the V2 base attributes (x10 baseline), the initial 100 CP point-buy allocation, and the single grade-gated initial skill roll granted to new characters at creation.

## THE RULE / LOGIC
- **Base Attributes (V2 x10 Baseline):** Every character starts with:
  - HP: 30
  - Attack: 10
  - Defense: 5
  - Movement: 3
- **Initial Allocation (Point-Buy):** Characters are initialized with 100 Character Points (CP) available to spend on top of their base attributes. The schema tracks only a `spent_cp` column (default 0) — there is no separate stored CP-pool/remaining-CP column; a character's remaining CP is always the implicit value 100 minus `spent_cp`, computed rather than stored.
- **Unspent CP:** Upon creation `spent_cp` is 0, so the full 100 CP remain unspent (100 − 0), leaving the character "ready for progression" via the progression allocation rules.
- **Initial Skill Selection:** Character creation also carries a single grade-gated skill roll (Grade I-II, at the caller's requested grade) — not a choice of 3. A single skill is generated and acquired directly onto the character, and the character's `roulette_used` flag (default false) is marked consumed by that roll (see `upsilonbattle:mech_skill_selection_progression`).

## TECHNICAL INTERFACE (The Bridge)
- **Code Tag:** `@spec-link [[rule_character_create_character]]`
- **Test Names:** `TestInitialStatConsistency`
