// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountStatus A summary of account status
// swagger:model AccountStatus
type AccountStatus struct {

	// The status of the account
	// Enum: [enabled disabled]
	State string `json:"state,omitempty"`
}

// Validate validates this account status
func (m *AccountStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var accountStatusTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountStatusTypeStatePropEnum = append(accountStatusTypeStatePropEnum, v)
	}
}

const (

	// AccountStatusStateEnabled captures enum value "enabled"
	AccountStatusStateEnabled string = "enabled"

	// AccountStatusStateDisabled captures enum value "disabled"
	AccountStatusStateDisabled string = "disabled"
)

// prop value enum
func (m *AccountStatus) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, accountStatusTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountStatus) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountStatus) UnmarshalBinary(b []byte) error {
	var res AccountStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
