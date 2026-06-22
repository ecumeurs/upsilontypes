---
id: vision_types_vision
status: STABLE
type: VISION
tags: [governance, vision, types]
parents:
  - [[shared:vision_upsilon_vision]]
human_name: UpsilonTypes Vision
layer: BUSINESS
version: 1.0
priority: 1
dependents: []
---

# UpsilonTypes Vision

## INTENT
Define the vision for UpsilonTypes as the unified domain model authority for the entire ecosystem.

## THE RULE / LOGIC
- **Single Source of Truth:** Provide common type definitions, interfaces, and enums shared across all Go modules.
- **Interoperability:** Ensure seamless serialization and deserialization of game state between the engine, API, and frontend.
- **Domain Consistency:** Enforce strict typing for critical game entities (Characters, Skills, Boards) to prevent cross-module logic errors.

## TECHNICAL INTERFACE
- **Code Tag:** `@spec-link [[vision_types_vision]]`
- **Related Atoms:** `[[entity_character]]`, `[[shared:vision_upsilon_vision]]`

## EXPECTATION
