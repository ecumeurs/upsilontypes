---
id: mech_entity_properties_item_properties
human_name: Item Properties Mechanic
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
# Item Properties Mechanic

## INTENT
To strictly catalog static properties utilized when acting as weapons or armor gear for Items

## THE RULE / LOGIC
**SUPERSEDED 2026-08-30 by `[[upsilontypes:module_property_key_registry]]`** (Property Key Space Unification round), via its parent `[[upsilontypes:mech_entity_properties]]` (also superseded the same day). Original content, retained for history: "Items carry static properties utilized when acting as weapons or armor gear." The `ItemProperties` named-string type this atom catalogs is deleted by the registry decision; item-scoped keys become ordinary `property.Key` entries carrying `Scope: Item` metadata in the unified registry instead. Demoted to DRAFT pending removal/re-scope in a future ATD cleanup pass rather than deleted outright.

## TECHNICAL INTERFACE (The Bridge)
- **Code Tag:** `@spec-link [[mech_entity_properties_item_properties]]`
