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

// RegistryTagSource An image reference using a tag in a registry, this is the most common source type.
// swagger:model RegistryTagSource
type RegistryTagSource struct {

	// Base64 encoded content of the dockerfile used to build the image, if available.
	// Pattern: ^[a-zA-Z0-9+/=]+$
	Dockerfile string `json:"dockerfile,omitempty"`

	// A docker pull string (e.g. docker.io/nginx:latest, or docker.io/nginx@sha256:abd) to retrieve the image
	// Required: true
	Pullstring *string `json:"pullstring"`
}

// Validate validates this registry tag source
func (m *RegistryTagSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDockerfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePullstring(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistryTagSource) validateDockerfile(formats strfmt.Registry) error {

	if swag.IsZero(m.Dockerfile) { // not required
		return nil
	}

	if err := validate.Pattern("dockerfile", "body", string(m.Dockerfile), `^[a-zA-Z0-9+/=]+$`); err != nil {
		return err
	}

	return nil
}

func (m *RegistryTagSource) validatePullstring(formats strfmt.Registry) error {

	if err := validate.Required("pullstring", "body", m.Pullstring); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegistryTagSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistryTagSource) UnmarshalBinary(b []byte) error {
	var res RegistryTagSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
