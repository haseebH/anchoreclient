// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServiceVersion Version information for a service
// swagger:model ServiceVersion
type ServiceVersion struct {

	// api
	API *ServiceVersionAPI `json:"api,omitempty"`

	// db
	Db *ServiceVersionDb `json:"db,omitempty"`

	// service
	Service *ServiceVersionService `json:"service,omitempty"`
}

// Validate validates this service version
func (m *ServiceVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDb(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceVersion) validateAPI(formats strfmt.Registry) error {

	if swag.IsZero(m.API) { // not required
		return nil
	}

	if m.API != nil {
		if err := m.API.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("api")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceVersion) validateDb(formats strfmt.Registry) error {

	if swag.IsZero(m.Db) { // not required
		return nil
	}

	if m.Db != nil {
		if err := m.Db.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("db")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceVersion) validateService(formats strfmt.Registry) error {

	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceVersion) UnmarshalBinary(b []byte) error {
	var res ServiceVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceVersionAPI Api Version string
// swagger:model ServiceVersionAPI
type ServiceVersionAPI struct {

	// Semantic version of the api
	Version string `json:"version,omitempty"`
}

// Validate validates this service version API
func (m *ServiceVersionAPI) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceVersionAPI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceVersionAPI) UnmarshalBinary(b []byte) error {
	var res ServiceVersionAPI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceVersionDb service version db
// swagger:model ServiceVersionDb
type ServiceVersionDb struct {

	// Semantic version of the db schema
	SchemaVersion string `json:"schema_version,omitempty"`
}

// Validate validates this service version db
func (m *ServiceVersionDb) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceVersionDb) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceVersionDb) UnmarshalBinary(b []byte) error {
	var res ServiceVersionDb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceVersionService service version service
// swagger:model ServiceVersionService
type ServiceVersionService struct {

	// Semantic Version string of the service implementation
	Version string `json:"version,omitempty"`
}

// Validate validates this service version service
func (m *ServiceVersionService) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceVersionService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceVersionService) UnmarshalBinary(b []byte) error {
	var res ServiceVersionService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}