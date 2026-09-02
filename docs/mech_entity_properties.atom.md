---
id: mech_entity_properties
human_name: Entity Property Declarations Mechanic
type: MODULE
layer: ARCHITECTURE
version: 1.0
status: DRAFT
priority: 5
tags: superseded,property
parents:
  - [[upsilonapi:domain_credit_economy]]
  - [[upsilonapi:domain_skill_system]]
dependents:
  - [[mech_entity_properties_item_properties]]
  - [[mech_entity_properties_skill_properties]]
---
# Entity Property Declarations Mechanic

## INTENT
To aggregate the constituent rules of Entity Property Declarations Mechanic.

## THE RULE / LOGIC
**SUPERSEDED 2026-08-30 by `[[upsilontypes:module_property_key_registry]]`** (Property Key Space Unification round). Original content, retained for history: "Defines the properties for Items and Skills in the game engine." This described the three independent namespaces `EntityProperties`/`SkillProperties`/`ItemProperties` in `upsilontypes/property/propertyenum.go` — three named `string` types sharing no common ancestor but `interface{}`. That three-way split is exactly what the unified `property.Key` + `def` registry (`[[upsilontypes:module_property_key_registry]]`) replaces. This atom no longer describes what `upsilontypes/property` does once the registry lands; demoted to DRAFT pending removal/re-scope in a future ATD cleanup pass rather than deleted outright.

## TECHNICAL INTERFACE (The Bridge)
- **Code Tag:** `@spec-link [[mech_entity_properties]]`
