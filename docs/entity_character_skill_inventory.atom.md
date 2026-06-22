---
id: entity_character_skill_inventory
human_name: "Character Skill Inventory"
status: STABLE
tags: ["skills", "inventory", "progression"]
parents:
  - [[entity_skill_template]]
  - [[rule_character_skill_slots]]
dependents:
  - [[upsilonapi:api_character_skill_inventory]]
  - [[rule_character_skill_slots]]
priority: 5
layer: ARCHITECTURE
version: 2.0
type: ENTITY
---

# Character Skill Inventory

## INTENT
To establish a character skill inventory system where players can store, manage, and access all acquired skills regardless of current equipment status. This enables roguelike-style skill collection and swapping.

## THE RULE / LOGIC
**`character_skills` table fields:**
- **id:** UUID primary key.
- **character_id:** FK → `characters.id`.
- **skill_template_id:** FK → `skill_templates.id` (informational only — snapshot is the authoritative copy).
- **source:** VARCHAR — `'roll'` for player-acquired skills; reserved for future sources.
- **instance_data:** JSON — frozen snapshot of the skill template row at acquisition time. Engine reads from this, not from the live template.
- **equipped:** Boolean (default false). True when the skill occupies an active slot.
- **acquired_at / equipped_at:** Timestamps.

**Inventory vs Equipped distinction:**
- Inventory capacity is unlimited; players accumulate skills over time.
- Equipped count is bounded by `character.skill_slots` (see `[[rule_character_skill_slots]]`).
- Equipping a skill sets `equipped=true`; unequipping sets it back to `false`.

**Acquisition methods:** Roll lottery (sole method in V2.0). Reforging and shop-bought skill inventory deferred to future versions.

## TECHNICAL INTERFACE
- **Code Tag:** `@spec-link [[entity_character_skill_inventory]]`
- **Laravel Model:** `App\Models\CharacterSkill`
- **Migration:** `*_create_character_skills_table.php`
- **Resource:** `App\Http\Resources\CharacterSkillResource`
- **Relationship:** `Character` hasMany `CharacterSkill`; scope `equippedSkills()` filters `equipped=true`.

## EXPECTATION
- `character_skills.instance_data` is set at creation and never mutated.
- `equipped` flag is the sole source of truth for slot occupancy.
- Deleting or updating a `skill_template` row does not retroactively change existing `instance_data` entries.
