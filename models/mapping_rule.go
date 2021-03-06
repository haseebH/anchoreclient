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

// MappingRule mapping rule
// swagger:model MappingRule
type MappingRule struct {

	// id
	ID string `json:"id,omitempty"`

	// image
	// Required: true
	Image *ImageRef `json:"image"`

	// name
	// Required: true
	Name *string `json:"name"`

	// Optional single policy to evalute, if set will override any value in policy_ids, for backwards compatibility. Generally, policy_ids should be used even with a array of length 1.
	PolicyID string `json:"policy_id,omitempty"`

	// List of policyIds to evaluate in order, to completion
	PolicyIds []string `json:"policy_ids"`

	// registry
	// Required: true
	Registry *string `json:"registry"`

	// repository
	// Required: true
	Repository *string `json:"repository"`

	// whitelist ids
	WhitelistIds []string `json:"whitelist_ids"`
}

// Validate validates this mapping rule
func (m *MappingRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MappingRule) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *MappingRule) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *MappingRule) validateRegistry(formats strfmt.Registry) error {

	if err := validate.Required("registry", "body", m.Registry); err != nil {
		return err
	}

	return nil
}

func (m *MappingRule) validateRepository(formats strfmt.Registry) error {

	if err := validate.Required("repository", "body", m.Repository); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MappingRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MappingRule) UnmarshalBinary(b []byte) error {
	var res MappingRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
