---
id: module_skill_generator
human_name: "Skill Generator Module"
type: MODULE
layer: ARCHITECTURE
version: 1.0
status: DRAFT
priority: 3
tags: [skills, generation, module]
parents:
  - [[shared:req_skill_generation]]
dependents:
  - [[api_skill_generation]]
  - [[mech_skill_classifier]]
  - [[mech_skill_generator_core]]
---

# Skill Generator Module

## INTENT
To group and orchestrate the procedural generation, classification, and naming of skills within the Upsilon Hub ecosystem, ensuring grade-aware and tag-driven outputs.

## THE RULE / LOGIC
The module acts as a high-level container for the following sub-components:
- **API Interface:** Defines how external services (like the character roll) request new skills.
- **Generator Core:** Handles the primary dispatching and budget logic.
- **Classifier:** Infers semantic tags from the final skill properties.
- **Name Generator:** Produces diegetic names based on the inferred tags.

## TECHNICAL INTERFACE
- **Module Path:** `upsilontypes/entity/skill/skillgenerator/`
- **Primary Entrance:** `Generate(GenerateRequest)`

## EXPECTATION
- Cohesive management of the skill generation pipeline.
- All sub-components are linked and traceable back to this module.
