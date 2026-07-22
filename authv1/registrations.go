// Service-registration wire contracts: the game-agnostic account model's
// record of which platform services (games) an account has enrolled in.
// Games own the enrollment act (they create their game-local state, e.g.
// battle's roster) and then record it here; auth only keeps the registry.
package authv1

import (
	"time"

	"github.com/google/uuid"
)

// ServiceTactical is the service key of the battle game (upsilonhub, future
// upsilontactical). Every game claims exactly one stable key; keys are part
// of the wire contract and never renamed.
const ServiceTactical = "tactical"

// RegisterServiceRequest records one account's enrollment in a service.
// Idempotent: re-registering an already-registered service is a no-op
// success, so game enroll endpoints can safely retry.
//
// @spec-link [[upsilonauth:contract_auth_service]]
type RegisterServiceRequest struct {
	Service string `json:"service"`
}

// Registration is one enrolled service on an account.
//
// @spec-link [[upsilonauth:contract_auth_service]]
type Registration struct {
	UserID       uuid.UUID `json:"user_id"`
	Service      string    `json:"service"`
	RegisteredAt time.Time `json:"registered_at"`
}

// RegistrationsResponse lists an account's enrolled services.
type RegistrationsResponse struct {
	Registrations []Registration `json:"registrations"`
}
