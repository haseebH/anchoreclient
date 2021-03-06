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

// ImageSelectionRule image selection rule
// swagger:model ImageSelectionRule
type ImageSelectionRule struct {

	// id
	ID string `json:"id,omitempty"`

	// image
	// Required: true
	Image *ImageRef `json:"image"`

	// name
	// Required: true
	Name *string `json:"name"`

	// registry
	// Required: true
	Registry *string `json:"registry"`

	// repository
	// Required: true
	Repository *string `json:"repository"`
}

// Validate validates this image selection rule
func (m *ImageSelectionRule) Validate(formats strfmt.Registry) error {
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

func (m *ImageSelectionRule) validateImage(formats strfmt.Registry) error {

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

func (m *ImageSelectionRule) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ImageSelectionRule) validateRegistry(formats strfmt.Registry) error {

	if err := validate.Required("registry", "body", m.Registry); err != nil {
		return err
	}

	return nil
}

func (m *ImageSelectionRule) validateRepository(formats strfmt.Registry) error {

	if err := validate.Required("repository", "body", m.Repository); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageSelectionRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageSelectionRule) UnmarshalBinary(b []byte) error {
	var res ImageSelectionRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
