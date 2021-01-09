package models

import "time"

// WebClientIdentity is ory kratos identity models/objects
type WebClientIdentity struct {
	ID                  string                        `json:"id,omitempty"`
	RecoveryAddresses   []*WebClientRecoveryAddress   `json:"recovery_addresses,omitempty"`
	SchemaID            string                        `json:"schema_id,omitempty"`
	SchemaURL           string                        `json:"schema_url,omitempty"`
	Traits              WebClientTraits               `json:"traits,omitempty"`
	VerifiableAddresses []*WebClientVerifiableAddress `json:"verifiable_addresses,omitempty"`
	Error               *WebClientError               `json:"error,omitempty"`
}

// WebClientRecoveryAddress is ory kratos recovery address on identity
type WebClientRecoveryAddress struct {
	ID    string `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
	Via   string `json:"via,omitempty"`
}

// WebClientTraits is our defined traits that used for identity and ORY Kratos logged in credentials
type WebClientTraits struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

// WebClientVerifiableAddress is ory kratos verifiable address data struct
type WebClientVerifiableAddress struct {
	ID         string    `json:"id,omitempty"`
	Status     string    `json:"status,omitempty"`
	Value      string    `json:"value,omitempty"`
	Verified   bool      `json:"verified,omitempty"`
	VerifiedAt time.Time `json:"verified_at,omitempty"`
	Via        string    `json:"via,omitempty"`
}

type WebClientError struct {
	Code    int32  `json:"code,omitempty"`
	Debug   string `json:"debug,omitempty"`
	Message string `json:"message,omitempty"`
	Reason  string `json:"reason,omitempty"`
	Request string `json:"request,omitempty"`
	Status  string `json:"status,omitempty"`
}
