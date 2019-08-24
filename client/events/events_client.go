// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new events API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for events API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteEvent deletes event

Delete an event by its event ID
*/
func (a *Client) DeleteEvent(params *DeleteEventParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete_event",
		Method:             "DELETE",
		PathPattern:        "/events/{eventId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteEventReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteEventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_event: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteEvents deletes events

Delete all or a subset of events filtered using the optional query parameters
*/
func (a *Client) DeleteEvents(params *DeleteEventsParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete_events",
		Method:             "DELETE",
		PathPattern:        "/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteEventsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_events: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEvent gets event

Lookup an event by its event ID
*/
func (a *Client) GetEvent(params *GetEventParams, authInfo runtime.ClientAuthInfoWriter) (*GetEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_event",
		Method:             "GET",
		PathPattern:        "/events/{eventId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEventReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_event: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListEvents lists events

Returns a paginated list of events in the descending order of their occurrence. Optional query parameters may be used for filtering results
*/
func (a *Client) ListEvents(params *ListEventsParams, authInfo runtime.ClientAuthInfoWriter) (*ListEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "list_events",
		Method:             "GET",
		PathPattern:        "/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListEventsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list_events: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}