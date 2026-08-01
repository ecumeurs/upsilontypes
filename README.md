# UpsilonTypes

**UpsilonTypes** is the central repository for core data structures, enums, and shared types used across the Upsilon ecosystem.

This subproject serves as the "Single Source of Truth" for all entities and properties, ensuring that the **Battle Engine**, **API**, **UI**, and **CLI** operate on identical data definitions.

## 🏗️ Architectural Role

- **Decoupling**: By isolating type definitions from business logic (mechanics/ruler), we prevent circular dependencies between the engine and its consumers.
- **Traceability**: Every type defined here is linked to an **Atomic Documentation (ATD)** atom, ensuring that implementation exactly matches the architectural specification.
- **Stability**: As a leaf-level package, `upsilontypes` is designed for high stability and low volatility.

## 📂 Package Structure

- **`entity/`**: Core game units and their state.
  - `Entity`: The base unit for characters, monsters, traps, area effects and other battle-turn actors; carries position, properties, buffs and skills.
  - `entity/skill/`: `Skill` — data structures for active/passive/reaction/counter/trap abilities (with `skillgenerator/` and `skillweight/` subpackages for procedural generation and balancing).
  - `entity/grade/`: `Grade` — entity power-tier classification.
  - `entity/entitygenerator/`: procedural entity generation.
- **`property/`**: Enumerations and metadata for combat arithmetic.
  - `EntityProperties`, `SkillProperties`, `ItemProperties`: string-enum lists of engine-recognized properties (HP, Attack, etc.) for each subject kind.
  - `TemporaryProperties`: structures for temporary stat modifications (buffs/debuffs), with duration tick-down.
  - `TriggerTypeValue`: enumeration of event hooks (OnEnter, OnStep, OnTurn, etc.) for positional effects.
  - `property/def/` and `property/defaultproperty/`: default property constructors/implementations for entities, skills and items.
  - `property/effect/`: `Effect` — definition of positional and temporal status effects.
- **`authv1/`**: Plain wire DTOs for the `upsilonauth` service boundary — `User`, `Token`, introspection requests/responses, and game-service registration contracts (`RegisterServiceRequest`, `Registration`).
- **`economyv1/`**: Plain wire DTOs for the `upsiloneconomy` service boundary — `Wallet`, batch wallet lookups, idempotent `AwardRequest`/`AwardResponse`, and shop catalog types.
- **`seedids/`**: Deterministic UUIDv5 derivation (`Account`, `ShopItem`) so independently-seeded services (auth, economy, hub) agree on well-known ids without cross-querying each other.

## 🚀 Usage

To use these types in your Go project:

```go
import (
    "github.com/ecumeurs/upsilontypes/entity"
    "github.com/ecumeurs/upsilontypes/property"
    "github.com/ecumeurs/upsilontypes/authv1"
    "github.com/ecumeurs/upsilontypes/economyv1"
)
```

## 🔗 ATD Traceability

This project strictly follows the ATD ruleset. All core structs and enums are annotated with `@spec-link` tags pointing to their respective atoms in the documentation, which are housed directly in `upsilontypes/docs/`.

Example:
```go
// @spec-link [[upsiloneconomy:contract_economy_service]]
type AwardRequest struct {
    IdempotencyKey string
    PlayerID       uuid.UUID
    Amount         int64
    // ...
}
```

## 🛠️ Development

When adding new types or properties:
1. Ensure a corresponding **DRAFT** or **STABLE** atom exists in `upsilontypes/docs/`.
2. Add the type with the appropriate `@spec-link` annotation.
3. Run `atd_weave` to update the dependency graph across the workspace.
4. Verify implementation coverage using `atd_trace`.

---
*Part of the [UpsilonBattle](https://github.com/ecumeurs/upsilon-hub) ecosystem.*
