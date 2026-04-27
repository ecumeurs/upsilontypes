# Upsilon Types

**Upsilon Types** is the central repository for core data structures, enums, and shared types used across the Upsilon ecosystem.

This subproject serves as the "Single Source of Truth" for all entities and properties, ensuring that the **Battle Engine**, **API**, **UI**, and **CLI** operate on identical data definitions.

## 🏗️ Architectural Role

- **Decoupling**: By isolating type definitions from business logic (mechanics/ruler), we prevent circular dependencies between the engine and its consumers.
- **Traceability**: Every type defined here is linked to an **Atomic Documentation (ATD)** atom, ensuring that implementation exactly matches the architectural specification.
- **Stability**: As a leaf-level package, `upsilon-types` is designed for high stability and low volatility.

## 📂 Package Structure

- **`entity/`**: Core game units and their state.
  - `Entity`: The base unit for characters, summons, and temporary effects.
  - `Skill`: Data structures for active and passive abilities.
  - `Effect`: Definition of positional and temporal status effects.
- **`property/`**: Enumerations and metadata for combat arithmetic.
  - `PropertyEnum`: Unified list of all engine-recognized properties (HP, Attack, etc.).
  - `Buff`: Structures for temporary stat modifications.
  - `TriggerType`: Enumeration of event hooks (OnStep, OnTurn, etc.).

## 🚀 Usage

To use these types in your Go project:

```go
import (
    "github.com/ecumeurs/upsilon-types/entity"
    "github.com/ecumeurs/upsilon-types/property"
)
```

## 🔗 ATD Traceability

This project strictly follows the ATD ruleset. All core structs and enums are annotated with `@spec-link` tags pointing to their respective atoms in the documentation.

Example:
```go
// @spec-link [[entity_character]]
type Character struct {
    ID   uuid.UUID
    Name string
    // ...
}
```

## 🛠️ Development

When adding new types or properties:
1. Ensure a corresponding **DRAFT** or **STABLE** atom exists in the documentation.
2. Add the type with the appropriate `@spec-link` annotation.
3. Run `atd_weave` to update the dependency graph.
4. Verify implementation coverage using `atd_trace`.

---
*Part of the [UpsilonBattle](https://github.com/ecumeurs/upsilon-hub) ecosystem.*
