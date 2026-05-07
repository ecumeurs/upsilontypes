---
id: contract_types_contract
status: STABLE
layer: BUSINESS
priority: 1
dependents: []
human_name: UpsilonTypes Contract
type: CONTRACT
version: 1.0
tags: [governance, contract, types]
parents:
  - [[shared:contract_upsilon_contract]]
---

# New Atom

## INTENT
Establish the stability and serialization requirements for the UpsilonTypes project.

## THE RULE / LOGIC
- **Stability:** Changes to core types must be backwards compatible or strictly versioned to avoid breaking the multi-stack integration.
- **Logic Isolation:** This project must contain only data structures and basic validation/serialization logic. No game mechanics.
- **Serialization:** Types must be compatible with both JSON and binary serialization formats.
- **Traceability:** Every field in core entities must be linked to its corresponding specification atom.

## TECHNICAL INTERFACE
- **Code Tag:** `@spec-link [[types_contract]]`
- **Related Atoms:** `[[shared:upsilon_contract]]`

## EXPECTATION
