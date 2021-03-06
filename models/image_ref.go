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

// ImageRef A reference to an image
// swagger:model ImageRef
type ImageRef struct {

	// type
	// Required: true
	// Enum: [tag digest id]
	Type interface{} `json:"type"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this image ref
func (m *ImageRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var imageRefTypeTypePropEnum []interface{}

func init() {
	var res []interface{}
	if err := json.Unmarshal([]byte(`["tag","digest","id"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		imageRefTypeTypePropEnum = append(imageRefTypeTypePropEnum, v)
	}
}

// prop value enum
func (m *ImageRef) validateTypeEnum(path, location string, value interface{}) error {
	if err := validate.Enum(path, location, value, imageRefTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ImageRef) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ImageRef) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageRef) UnmarshalBinary(b []byte) error {
	var res ImageRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
