// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PaginationProperties Properties for common pagination handling to be included in any wrapping object that needs pagination elements
// swagger:model PaginationProperties
type PaginationProperties struct {

	// True if additional pages exist (page + 1) or False if this is the last page
	NextPage string `json:"next_page,omitempty"`

	// The page number returned (should match the requested page query string param)
	Page string `json:"page,omitempty"`

	// The number of items sent in this response
	ReturnedCount int64 `json:"returned_count,omitempty"`
}

// Validate validates this pagination properties
func (m *PaginationProperties) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaginationProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaginationProperties) UnmarshalBinary(b []byte) error {
	var res PaginationProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
