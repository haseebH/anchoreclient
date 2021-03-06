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

// EventResponse A record of occurance of an asynchronous event triggered either by system or by user activity
// swagger:model EventResponse
type EventResponse struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// event
	Event *EventResponseEvent `json:"event,omitempty"`

	// generated uuid
	GeneratedUUID string `json:"generated_uuid,omitempty"`
}

// Validate validates this event response
func (m *EventResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventResponse) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EventResponse) validateEvent(formats strfmt.Registry) error {

	if swag.IsZero(m.Event) { // not required
		return nil
	}

	if m.Event != nil {
		if err := m.Event.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EventResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventResponse) UnmarshalBinary(b []byte) error {
	var res EventResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EventResponseEvent event response event
// swagger:model EventResponseEvent
type EventResponseEvent struct {

	// category
	Category string `json:"category,omitempty"`

	// details
	Details interface{} `json:"details,omitempty"`

	// level
	Level string `json:"level,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// resource
	Resource *EventResponseEventResource `json:"resource,omitempty"`

	// source
	Source *EventResponseEventSource `json:"source,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this event response event
func (m *EventResponseEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventResponseEvent) validateResource(formats strfmt.Registry) error {

	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event" + "." + "resource")
			}
			return err
		}
	}

	return nil
}

func (m *EventResponseEvent) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event" + "." + "source")
			}
			return err
		}
	}

	return nil
}

func (m *EventResponseEvent) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("event"+"."+"timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EventResponseEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventResponseEvent) UnmarshalBinary(b []byte) error {
	var res EventResponseEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EventResponseEventResource event response event resource
// swagger:model EventResponseEventResource
type EventResponseEventResource struct {

	// id
	ID string `json:"id,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this event response event resource
func (m *EventResponseEventResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventResponseEventResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventResponseEventResource) UnmarshalBinary(b []byte) error {
	var res EventResponseEventResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EventResponseEventSource event response event source
// swagger:model EventResponseEventSource
type EventResponseEventSource struct {

	// base url
	BaseURL string `json:"base_url,omitempty"`

	// hostid
	Hostid string `json:"hostid,omitempty"`

	// request id
	RequestID string `json:"request_id,omitempty"`

	// servicename
	Servicename string `json:"servicename,omitempty"`
}

// Validate validates this event response event source
func (m *EventResponseEventSource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventResponseEventSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventResponseEventSource) UnmarshalBinary(b []byte) error {
	var res EventResponseEventSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
