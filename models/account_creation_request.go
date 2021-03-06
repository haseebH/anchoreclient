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

// AccountCreationRequest An account to create/add to the system. If already exists will return 400.
// swagger:model AccountCreationRequest
type AccountCreationRequest struct {

	// An optional email to associate with the account for contact purposes
	// Pattern: [a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?
	Email string `json:"email,omitempty"`

	// The account name to use. This will identify the account and must be globally unique in the system.
	// Required: true
	// Pattern: ^[a-zA-Z0-9][a-zA-Z0-9@.!#$+-=^_~;]{1,126}[a-zA-Z0-9]$
	Name *string `json:"name"`
}

// Validate validates this account creation request
func (m *AccountCreationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountCreationRequest) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.Pattern("email", "body", string(m.Email), `[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?`); err != nil {
		return err
	}

	return nil
}

func (m *AccountCreationRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `^[a-zA-Z0-9][a-zA-Z0-9@.!#$+-=^_~;]{1,126}[a-zA-Z0-9]$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountCreationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountCreationRequest) UnmarshalBinary(b []byte) error {
	var res AccountCreationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
