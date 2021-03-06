// Code generated by go-swagger; DO NOT EDIT.

package system

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

// NewGetServiceDetailParams creates a new GetServiceDetailParams object
// with the default values initialized.
func NewGetServiceDetailParams() *GetServiceDetailParams {

	return &GetServiceDetailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceDetailParamsWithTimeout creates a new GetServiceDetailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServiceDetailParamsWithTimeout(timeout time.Duration) *GetServiceDetailParams {

	return &GetServiceDetailParams{

		timeout: timeout,
	}
}

// NewGetServiceDetailParamsWithContext creates a new GetServiceDetailParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServiceDetailParamsWithContext(ctx context.Context) *GetServiceDetailParams {

	return &GetServiceDetailParams{

		Context: ctx,
	}
}

// NewGetServiceDetailParamsWithHTTPClient creates a new GetServiceDetailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServiceDetailParamsWithHTTPClient(client *http.Client) *GetServiceDetailParams {

	return &GetServiceDetailParams{
		HTTPClient: client,
	}
}

/*GetServiceDetailParams contains all the parameters to send to the API endpoint
for the get service detail operation typically these are written to a http.Request
*/
type GetServiceDetailParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get service detail params
func (o *GetServiceDetailParams) WithTimeout(timeout time.Duration) *GetServiceDetailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service detail params
func (o *GetServiceDetailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service detail params
func (o *GetServiceDetailParams) WithContext(ctx context.Context) *GetServiceDetailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service detail params
func (o *GetServiceDetailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service detail params
func (o *GetServiceDetailParams) WithHTTPClient(client *http.Client) *GetServiceDetailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service detail params
func (o *GetServiceDetailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceDetailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
