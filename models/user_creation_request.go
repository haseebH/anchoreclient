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

// UserCreationRequest A payload for creating a new user, includes the username and password in a single request
// swagger:model UserCreationRequest
type UserCreationRequest struct {

	// The initial password for the user, must be at least 6 characters, up to 128
	// Required: true
	// Pattern: .{6,128}$
	Password *string `json:"password"`

	// The username to create
	// Required: true
	// Pattern: ^[a-zA-Z0-9][a-zA-Z0-9_-]{1,126}[a-zA-Z0-9]$
	Username *string `json:"username"`
}

// Validate validates this user creation request
func (m *UserCreationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserCreationRequest) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.Pattern("password", "body", string(*m.Password), `.{6,128}$`); err != nil {
		return err
	}

	return nil
}

func (m *UserCreationRequest) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.Pattern("username", "body", string(*m.Username), `^[a-zA-Z0-9][a-zA-Z0-9_-]{1,126}[a-zA-Z0-9]$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserCreationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserCreationRequest) UnmarshalBinary(b []byte) error {
	var res UserCreationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}