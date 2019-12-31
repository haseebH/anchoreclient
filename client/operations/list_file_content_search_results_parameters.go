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

// NewListFileContentSearchResultsParams creates a new ListFileContentSearchResultsParams object
// with the default values initialized.
func NewListFileContentSearchResultsParams() *ListFileContentSearchResultsParams {
	var ()
	return &ListFileContentSearchResultsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListFileContentSearchResultsParamsWithTimeout creates a new ListFileContentSearchResultsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListFileContentSearchResultsParamsWithTimeout(timeout time.Duration) *ListFileContentSearchResultsParams {
	var ()
	return &ListFileContentSearchResultsParams{

		timeout: timeout,
	}
}

// NewListFileContentSearchResultsParamsWithContext creates a new ListFileContentSearchResultsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListFileContentSearchResultsParamsWithContext(ctx context.Context) *ListFileContentSearchResultsParams {
	var ()
	return &ListFileContentSearchResultsParams{

		Context: ctx,
	}
}

// NewListFileContentSearchResultsParamsWithHTTPClient creates a new ListFileContentSearchResultsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListFileContentSearchResultsParamsWithHTTPClient(client *http.Client) *ListFileContentSearchResultsParams {
	var ()
	return &ListFileContentSearchResultsParams{
		HTTPClient: client,
	}
}

/*ListFileContentSearchResultsParams contains all the parameters to send to the API endpoint
for the list file content search results operation typically these are written to a http.Request
*/
type ListFileContentSearchResultsParams struct {

	/*ImageDigest*/
	ImageDigest string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list file content search results params
func (o *ListFileContentSearchResultsParams) WithTimeout(timeout time.Duration) *ListFileContentSearchResultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list file content search results params
func (o *ListFileContentSearchResultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list file content search results params
func (o *ListFileContentSearchResultsParams) WithContext(ctx context.Context) *ListFileContentSearchResultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list file content search results params
func (o *ListFileContentSearchResultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list file content search results params
func (o *ListFileContentSearchResultsParams) WithHTTPClient(client *http.Client) *ListFileContentSearchResultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list file content search results params
func (o *ListFileContentSearchResultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImageDigest adds the imageDigest to the list file content search results params
func (o *ListFileContentSearchResultsParams) WithImageDigest(imageDigest string) *ListFileContentSearchResultsParams {
	o.SetImageDigest(imageDigest)
	return o
}

// SetImageDigest adds the imageDigest to the list file content search results params
func (o *ListFileContentSearchResultsParams) SetImageDigest(imageDigest string) {
	o.ImageDigest = imageDigest
}

// WriteToRequest writes these params to a swagger request
func (o *ListFileContentSearchResultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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