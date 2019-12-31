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

// NewListSecretSearchResultsParams creates a new ListSecretSearchResultsParams object
// with the default values initialized.
func NewListSecretSearchResultsParams() *ListSecretSearchResultsParams {
	var ()
	return &ListSecretSearchResultsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSecretSearchResultsParamsWithTimeout creates a new ListSecretSearchResultsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSecretSearchResultsParamsWithTimeout(timeout time.Duration) *ListSecretSearchResultsParams {
	var ()
	return &ListSecretSearchResultsParams{

		timeout: timeout,
	}
}

// NewListSecretSearchResultsParamsWithContext creates a new ListSecretSearchResultsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSecretSearchResultsParamsWithContext(ctx context.Context) *ListSecretSearchResultsParams {
	var ()
	return &ListSecretSearchResultsParams{

		Context: ctx,
	}
}

// NewListSecretSearchResultsParamsWithHTTPClient creates a new ListSecretSearchResultsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSecretSearchResultsParamsWithHTTPClient(client *http.Client) *ListSecretSearchResultsParams {
	var ()
	return &ListSecretSearchResultsParams{
		HTTPClient: client,
	}
}

/*ListSecretSearchResultsParams contains all the parameters to send to the API endpoint
for the list secret search results operation typically these are written to a http.Request
*/
type ListSecretSearchResultsParams struct {

	/*ImageDigest*/
	ImageDigest string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list secret search results params
func (o *ListSecretSearchResultsParams) WithTimeout(timeout time.Duration) *ListSecretSearchResultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list secret search results params
func (o *ListSecretSearchResultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list secret search results params
func (o *ListSecretSearchResultsParams) WithContext(ctx context.Context) *ListSecretSearchResultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list secret search results params
func (o *ListSecretSearchResultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list secret search results params
func (o *ListSecretSearchResultsParams) WithHTTPClient(client *http.Client) *ListSecretSearchResultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list secret search results params
func (o *ListSecretSearchResultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImageDigest adds the imageDigest to the list secret search results params
func (o *ListSecretSearchResultsParams) WithImageDigest(imageDigest string) *ListSecretSearchResultsParams {
	o.SetImageDigest(imageDigest)
	return o
}

// SetImageDigest adds the imageDigest to the list secret search results params
func (o *ListSecretSearchResultsParams) SetImageDigest(imageDigest string) {
	o.ImageDigest = imageDigest
}

// WriteToRequest writes these params to a swagger request
func (o *ListSecretSearchResultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param imageDigest
	if err := r.SetPathParam("imageDigest", o.ImageDigest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}