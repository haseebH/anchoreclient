// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CVSSV2Scores c v s s v2 scores
// swagger:model CVSSV2Scores
type CVSSV2Scores struct {

	// base score
	BaseScore *float64 `json:"base_score,omitempty"`

	// exploitability score
	ExploitabilityScore *float64 `json:"exploitability_score,omitempty"`

	// impact score
	ImpactScore *float64 `json:"impact_score,omitempty"`
}

// Validate validates this c v s s v2 scores
func (m *CVSSV2Scores) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CVSSV2Scores) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CVSSV2Scores) UnmarshalBinary(b []byte) error {
	var res CVSSV2Scores
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
