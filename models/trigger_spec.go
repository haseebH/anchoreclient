// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TriggerSpec Definition of a trigger and its parameters
// swagger:model TriggerSpec
type TriggerSpec struct {

	// Trigger description for what it tests and when it will fire during evaluation
	Description string `json:"description,omitempty"`

	// Name of the trigger as it would appear in a policy document
	Name string `json:"name,omitempty"`

	// The list of parameters that are valid for this trigger
	Parameters []*TriggerParamSpec `json:"parameters"`

	// State of the trigger
	// Enum: [active deprecated eol]
	State string `json:"state,omitempty"`

	// The name of another trigger that supercedes this on functionally if this is deprecated
	SupercededBy *string `json:"superceded_by,omitempty"`
}

// Validate validates this trigger spec
func (m *TriggerSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TriggerSpec) validateParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var triggerSpecTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","deprecated","eol"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		triggerSpecTypeStatePropEnum = append(triggerSpecTypeStatePropEnum, v)
	}
}

const (

	// TriggerSpecStateActive captures enum value "active"
	TriggerSpecStateActive string = "active"

	// TriggerSpecStateDeprecated captures enum value "deprecated"
	TriggerSpecStateDeprecated string = "deprecated"

	// TriggerSpecStateEol captures enum value "eol"
	TriggerSpecStateEol string = "eol"
)

// prop value enum
func (m *TriggerSpec) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, triggerSpecTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TriggerSpec) validateState(formats strfmt.Registry) error {

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
func (m *TriggerSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TriggerSpec) UnmarshalBinary(b []byte) error {
	var res TriggerSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}