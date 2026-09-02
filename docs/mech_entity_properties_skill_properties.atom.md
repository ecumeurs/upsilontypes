---
id: mech_entity_properties_skill_properties
human_name: Skill Properties Mechanic
type: MECHANIC
layer: IMPLEMENTATION
version: 1.0
status: DRAFT
priority: 5
tags: superseded,property
parents:
  - [[mech_entity_properties]]
dependents: []
---
# Skill Properties Mechanic

## INTENT
To strictly catalog dynamic properties dictating combat arithmetic for Skills

## THE RULE / LOGIC
**SUPERSEDED 2026-08-30 by `[[upsilontypes:module_property_key_registry]]`** (Property Key Space Unification round), via its parent `[[upsilontypes:mech_entity_properties]]` (also superseded the same day). Original content, retained for history: "Skills map to dynamic properties dictating combat arithmetic." The `SkillProperties` named-string type this atom catalogs is deleted by the registry decision; skill-scoped keys become ordinary `property.Key` entries carrying `Scope: Skill` metadata in the unified registry instead. Note: unlike its sibling `[[upsilontypes:mech_entity_properties_item_properties]]`, this atom carries one real `@spec-link` to `upsilontypes/property/property.go` (`atd trace` shows total_code_files: 1) — a live code link this cleanup pass has not removed; the link is now stale against the superseded content and should be reconciled (removed or re-pointed) during the Workflow B pass that follows implementation, not silently dropped here. Demoted to DRAFT pending removal/re-scope in a future ATD cleanup pass rather than deleted outright.

## TECHNICAL INTERFACE (The Bridge)
- **Code Tag:** `@spec-link [[mech_entity_properties_skill_properties]]`
