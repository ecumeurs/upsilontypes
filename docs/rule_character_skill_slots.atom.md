---
id: rule_character_skill_slots
status: STABLE
parents:
  - [[entity_character_skill_inventory]]
dependents:
  - [[upsilonapi:api_character_skill_inventory]]
  - [[entity_character_skill_inventory]]
priority: 5
layer: ARCHITECTURE
version: 2.0
tags: ["skills", "progression", "equipment"]
---

# New Atom

## INTENT
To establish the character skill slot system where characters gain a new skill slot every 10 levels, starting with 1 base slot at level 1. This defines the maximum number of skills a character can have equipped for battle.

## THE RULE / LOGIC
**Slot Formula:** `skill_slots = min(5, 1 + intdiv(player.total_wins, 10))`

Implemented as a computed accessor on the `Character` model. `player.total_wins` is the owning `Player` record's win counter — not a character-level stat.

**Slot Progression (by player wins):**
- 0–9 wins: 1 slot (base)
- 10–19 wins: 2 slots
- 20–29 wins: 3 slots
- 30–39 wins: 4 slots
- 40+ wins: 5 slots (hard cap)

**Invariants:**
- Characters cannot equip more skills than `skill_slots`.
- Attempting to equip beyond the limit returns 422 with reason `slot_full`.
- Unequipping a skill frees its slot immediately.
- Inventory accumulation (rolling skills) is uncapped — players may own more skills than slots.
- Only equipped skills are sent to the engine at arena initialization.

## TECHNICAL INTERFACE
- **Code Tag:** `@spec-link [[rule_character_skill_slots]]`
- **Accessor:** `Character::getSkillSlotsAttribute()` in `App\Models\Character`
- **Enforcement:** `SkillService::equip()` — `lockForUpdate()->get(['id'])->count()` vs `skill_slots`.

## EXPECTATION
- Fresh character (0 wins) has `skill_slots = 1`.
- After 10 player wins, `skill_slots` becomes 2.
- Equipping beyond limit returns 422.
- `skill_slots` field is exposed in `CharacterResource`.
