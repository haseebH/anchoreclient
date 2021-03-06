// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteEventParams creates a new DeleteEventParams object
// with the default values initialized.
func NewDeleteEventParams() *DeleteEventParams {
	var ()
	return &DeleteEventParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEventParamsWithTimeout creates a new DeleteEventParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteEventParamsWithTimeout(timeout time.Duration) *DeleteEventParams {
	var ()
	return &DeleteEventParams{

		timeout: timeout,
	}
}

// NewDeleteEventParamsWithContext creates a new DeleteEventParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteEventParamsWithContext(ctx context.Context) *DeleteEventParams {
	var ()
	return &DeleteEventParams{

		Context: ctx,
	}
}

// NewDeleteEventParamsWithHTTPClient creates a new DeleteEventParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteEventParamsWithHTTPClient(client *http.Client) *DeleteEventParams {
	var ()
	return &DeleteEventParams{
		HTTPClient: client,
	}
}

/*DeleteEventParams contains all the parameters to send to the API endpoint
for the delete event operation typically these are written to a http.Request
*/
type DeleteEventParams struct {

	/*EventID
	  Event ID of the event to be deleted

	*/
	EventID string
	/*XAnchoreAccount
	  An account name to change the resource scope of the request to that account, if permissions allow (admin only)

	*/
	XAnchoreAccount *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete event params
func (o *DeleteEventParams) WithTimeout(timeout time.Duration) *DeleteEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete event params
func (o *DeleteEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete event params
func (o *DeleteEventParams) WithContext(ctx context.Context) *DeleteEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete event params
func (o *DeleteEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete event params
func (o *DeleteEventParams) WithHTTPClient(client *http.Client) *DeleteEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete event params
func (o *DeleteEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEventID adds the eventID to the delete event params
func (o *DeleteEventParams) WithEventID(eventID string) *DeleteEventParams {
	o.SetEventID(eventID)
	return o
}

// SetEventID adds the eventId to the delete event params
func (o *DeleteEventParams) SetEventID(eventID string) {
	o.EventID = eventID
}

// WithXAnchoreAccount adds the xAnchoreAccount to the delete event params
func (o *DeleteEventParams) WithXAnchoreAccount(xAnchoreAccount *string) *DeleteEventParams {
	o.SetXAnchoreAccount(xAnchoreAccount)
	return o
}

// SetXAnchoreAccount adds the xAnchoreAccount to the delete event params
func (o *DeleteEventParams) SetXAnchoreAccount(xAnchoreAccount *string) {
	o.XAnchoreAccount = xAnchoreAccount
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param eventId
	if err := r.SetPathParam("eventId", o.EventID); err != nil {
		return err
	}

	if o.XAnchoreAccount != nil {

		// header param x-anchore-account
		if err := r.SetHeaderParam("x-anchore-account", *o.XAnchoreAccount); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
