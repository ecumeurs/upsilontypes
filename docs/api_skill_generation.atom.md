---
id: api_skill_generation
human_name: "Skill Generation API"
type: API
layer: ARCHITECTURE
version: 1.0
status: DRAFT
priority: 3
tags: [skills, api, generation]
parents:
  - [[module_skill_generator]]
dependents: []
---

# Skill Generation API

## INTENT
To define the formal contract for requesting procedurally generated skills, allowing callers to specify grade targets and tag constraints.

## THE RULE / LOGIC
The API supports the following structures:

**GenerateRequest:**
- `TargetGrade` (string): Grade I-V.
- `AllowedTags` ([]string): Inclusive filter for primary producers.
- `ForbidTags` ([]string): Exclusive filter.

**GenerateResponse:**
- `Skill` (skill.Skill): The resulting skill object.
- `Tags` ([]string): The ordered tags inferred by the classifier.

## TECHNICAL INTERFACE
- **Code Tag:** `@spec-link [[api_skill_generation]]`
- **Request Struct:** `GenerateRequest` in `upsilontypes/entity/skill/skillgenerator/skillgenerator.go`
- **Bridge Endpoint:** `POST /api/v1/profile/character/{id}/skills/roll`

## EXPECTATION
- Clear separation between request parameters and generated output.
- Support for grade-aware generation.
