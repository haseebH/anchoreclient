// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RegistryConfiguration A registry entry describing the endpoint and credentials for a registry to pull images from
// swagger:model RegistryConfiguration
type RegistryConfiguration struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// last upated
	// Format: date-time
	LastUpated strfmt.DateTime `json:"last_upated,omitempty"`

	// hostname:port string for accessing the registry, as would be used in a docker pull operation
	Registry string `json:"registry,omitempty"`

	// human readable name associated with registry record
	RegistryName string `json:"registry_name,omitempty"`

	// Type of registry
	RegistryType string `json:"registry_type,omitempty"`

	// Username portion of credential to use for this registry
	RegistryUser string `json:"registry_user,omitempty"`

	// Use TLS/SSL verification for the registry URL
	RegistryVerify bool `json:"registry_verify,omitempty"`

	// Engine user that owns this registry entry
	UserID string `json:"userId,omitempty"`
}

// Validate validates this registry configuration
func (m *RegistryConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistryConfiguration) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RegistryConfiguration) validateLastUpated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_upated", "body", "date-time", m.LastUpated.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegistryConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistryConfiguration) UnmarshalBinary(b []byte) error {
	var res RegistryConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
