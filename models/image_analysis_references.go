// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
)

// ImageAnalysisReferences List of image digests to archive
// swagger:model ImageAnalysisReferences
type ImageAnalysisReferences []string

// Validate validates this image analysis references
func (m ImageAnalysisReferences) Validate(formats strfmt.Registry) error {
	return nil
}
