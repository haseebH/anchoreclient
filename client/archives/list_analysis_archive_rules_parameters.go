// Code generated by go-swagger; DO NOT EDIT.

package archives

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListAnalysisArchiveRulesParams creates a new ListAnalysisArchiveRulesParams object
// with the default values initialized.
func NewListAnalysisArchiveRulesParams() *ListAnalysisArchiveRulesParams {
	var ()
	return &ListAnalysisArchiveRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAnalysisArchiveRulesParamsWithTimeout creates a new ListAnalysisArchiveRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAnalysisArchiveRulesParamsWithTimeout(timeout time.Duration) *ListAnalysisArchiveRulesParams {
	var ()
	return &ListAnalysisArchiveRulesParams{

		timeout: timeout,
	}
}

// NewListAnalysisArchiveRulesParamsWithContext creates a new ListAnalysisArchiveRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAnalysisArchiveRulesParamsWithContext(ctx context.Context) *ListAnalysisArchiveRulesParams {
	var ()
	return &ListAnalysisArchiveRulesParams{

		Context: ctx,
	}
}

// NewListAnalysisArchiveRulesParamsWithHTTPClient creates a new ListAnalysisArchiveRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAnalysisArchiveRulesParamsWithHTTPClient(client *http.Client) *ListAnalysisArchiveRulesParams {
	var ()
	return &ListAnalysisArchiveRulesParams{
		HTTPClient: client,
	}
}

/*ListAnalysisArchiveRulesParams contains all the parameters to send to the API endpoint
for the list analysis archive rules operation typically these are written to a http.Request
*/
type ListAnalysisArchiveRulesParams struct {

	/*SystemGlobal
	  If true include system global rules (owned by admin) even for non-admin users. Defaults to true if not set. Can be set to false to exclude globals

	*/
	SystemGlobal *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list analysis archive rules params
func (o *ListAnalysisArchiveRulesParams) WithTimeout(timeout time.Duration) *ListAnalysisArchiveRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list analysis archive rules params
func (o *ListAnalysisArchiveRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list analysis archive rules params
func (o *ListAnalysisArchiveRulesParams) WithContext(ctx context.Context) *ListAnalysisArchiveRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list analysis archive rules params
func (o *ListAnalysisArchiveRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list analysis archive rules params
func (o *ListAnalysisArchiveRulesParams) WithHTTPClient(client *http.Client) *ListAnalysisArchiveRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list analysis archive rules params
func (o *ListAnalysisArchiveRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSystemGlobal adds the systemGlobal to the list analysis archive rules params
func (o *ListAnalysisArchiveRulesParams) WithSystemGlobal(systemGlobal *bool) *ListAnalysisArchiveRulesParams {
	o.SetSystemGlobal(systemGlobal)
	return o
}

// SetSystemGlobal adds the systemGlobal to the list analysis archive rules params
func (o *ListAnalysisArchiveRulesParams) SetSystemGlobal(systemGlobal *bool) {
	o.SystemGlobal = systemGlobal
}

// WriteToRequest writes these params to a swagger request
func (o *ListAnalysisArchiveRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SystemGlobal != nil {

		// query param system_global
		var qrSystemGlobal bool
		if o.SystemGlobal != nil {
			qrSystemGlobal = *o.SystemGlobal
		}
		qSystemGlobal := swag.FormatBool(qrSystemGlobal)
		if qSystemGlobal != "" {
			if err := r.SetQueryParam("system_global", qSystemGlobal); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
