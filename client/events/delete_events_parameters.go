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

// NewDeleteEventsParams creates a new DeleteEventsParams object
// with the default values initialized.
func NewDeleteEventsParams() *DeleteEventsParams {
	var ()
	return &DeleteEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEventsParamsWithTimeout creates a new DeleteEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteEventsParamsWithTimeout(timeout time.Duration) *DeleteEventsParams {
	var ()
	return &DeleteEventsParams{

		timeout: timeout,
	}
}

// NewDeleteEventsParamsWithContext creates a new DeleteEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteEventsParamsWithContext(ctx context.Context) *DeleteEventsParams {
	var ()
	return &DeleteEventsParams{

		Context: ctx,
	}
}

// NewDeleteEventsParamsWithHTTPClient creates a new DeleteEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteEventsParamsWithHTTPClient(client *http.Client) *DeleteEventsParams {
	var ()
	return &DeleteEventsParams{
		HTTPClient: client,
	}
}

/*DeleteEventsParams contains all the parameters to send to the API endpoint
for the delete events operation typically these are written to a http.Request
*/
type DeleteEventsParams struct {

	/*Before
	  Delete events that occurred before the timestamp

	*/
	Before *string
	/*Level
	  Delete events that match the level - INFO or ERROR

	*/
	Level *string
	/*Since
	  Delete events that occurred after the timestamp

	*/
	Since *string
	/*XAnchoreAccount
	  An account name to change the resource scope of the request to that account, if permissions allow (admin only)

	*/
	XAnchoreAccount *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete events params
func (o *DeleteEventsParams) WithTimeout(timeout time.Duration) *DeleteEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete events params
func (o *DeleteEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete events params
func (o *DeleteEventsParams) WithContext(ctx context.Context) *DeleteEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete events params
func (o *DeleteEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete events params
func (o *DeleteEventsParams) WithHTTPClient(client *http.Client) *DeleteEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete events params
func (o *DeleteEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBefore adds the before to the delete events params
func (o *DeleteEventsParams) WithBefore(before *string) *DeleteEventsParams {
	o.SetBefore(before)
	return o
}

// SetBefore adds the before to the delete events params
func (o *DeleteEventsParams) SetBefore(before *string) {
	o.Before = before
}

// WithLevel adds the level to the delete events params
func (o *DeleteEventsParams) WithLevel(level *string) *DeleteEventsParams {
	o.SetLevel(level)
	return o
}

// SetLevel adds the level to the delete events params
func (o *DeleteEventsParams) SetLevel(level *string) {
	o.Level = level
}

// WithSince adds the since to the delete events params
func (o *DeleteEventsParams) WithSince(since *string) *DeleteEventsParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the delete events params
func (o *DeleteEventsParams) SetSince(since *string) {
	o.Since = since
}

// WithXAnchoreAccount adds the xAnchoreAccount to the delete events params
func (o *DeleteEventsParams) WithXAnchoreAccount(xAnchoreAccount *string) *DeleteEventsParams {
	o.SetXAnchoreAccount(xAnchoreAccount)
	return o
}

// SetXAnchoreAccount adds the xAnchoreAccount to the delete events params
func (o *DeleteEventsParams) SetXAnchoreAccount(xAnchoreAccount *string) {
	o.XAnchoreAccount = xAnchoreAccount
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Before != nil {

		// query param before
		var qrBefore string
		if o.Before != nil {
			qrBefore = *o.Before
		}
		qBefore := qrBefore
		if qBefore != "" {
			if err := r.SetQueryParam("before", qBefore); err != nil {
				return err
			}
		}

	}

	if o.Level != nil {

		// query param level
		var qrLevel string
		if o.Level != nil {
			qrLevel = *o.Level
		}
		qLevel := qrLevel
		if qLevel != "" {
			if err := r.SetQueryParam("level", qLevel); err != nil {
				return err
			}
		}

	}

	if o.Since != nil {

		// query param since
		var qrSince string
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := qrSince
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

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
