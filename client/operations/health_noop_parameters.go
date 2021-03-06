// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewHealthNoopParams creates a new HealthNoopParams object
// with the default values initialized.
func NewHealthNoopParams() *HealthNoopParams {

	return &HealthNoopParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHealthNoopParamsWithTimeout creates a new HealthNoopParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHealthNoopParamsWithTimeout(timeout time.Duration) *HealthNoopParams {

	return &HealthNoopParams{

		timeout: timeout,
	}
}

// NewHealthNoopParamsWithContext creates a new HealthNoopParams object
// with the default values initialized, and the ability to set a context for a request
func NewHealthNoopParamsWithContext(ctx context.Context) *HealthNoopParams {

	return &HealthNoopParams{

		Context: ctx,
	}
}

// NewHealthNoopParamsWithHTTPClient creates a new HealthNoopParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHealthNoopParamsWithHTTPClient(client *http.Client) *HealthNoopParams {

	return &HealthNoopParams{
		HTTPClient: client,
	}
}

/*HealthNoopParams contains all the parameters to send to the API endpoint
for the health noop operation typically these are written to a http.Request
*/
type HealthNoopParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the health noop params
func (o *HealthNoopParams) WithTimeout(timeout time.Duration) *HealthNoopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the health noop params
func (o *HealthNoopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the health noop params
func (o *HealthNoopParams) WithContext(ctx context.Context) *HealthNoopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the health noop params
func (o *HealthNoopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the health noop params
func (o *HealthNoopParams) WithHTTPClient(client *http.Client) *HealthNoopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the health noop params
func (o *HealthNoopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HealthNoopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
