// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// NewGetPolicyParams creates a new GetPolicyParams object
// with the default values initialized.
func NewGetPolicyParams() *GetPolicyParams {
	var ()
	return &GetPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPolicyParamsWithTimeout creates a new GetPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPolicyParamsWithTimeout(timeout time.Duration) *GetPolicyParams {
	var ()
	return &GetPolicyParams{

		timeout: timeout,
	}
}

// NewGetPolicyParamsWithContext creates a new GetPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPolicyParamsWithContext(ctx context.Context) *GetPolicyParams {
	var ()
	return &GetPolicyParams{

		Context: ctx,
	}
}

// NewGetPolicyParamsWithHTTPClient creates a new GetPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPolicyParamsWithHTTPClient(client *http.Client) *GetPolicyParams {
	var ()
	return &GetPolicyParams{
		HTTPClient: client,
	}
}

/*GetPolicyParams contains all the parameters to send to the API endpoint
for the get policy operation typically these are written to a http.Request
*/
type GetPolicyParams struct {

	/*Detail
	  Include policy bundle detail in the form of the full bundle content for each entry

	*/
	Detail *bool
	/*PolicyID*/
	PolicyID string
	/*XAnchoreAccount
	  An account name to change the resource scope of the request to that account, if permissions allow (admin only)

	*/
	XAnchoreAccount *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get policy params
func (o *GetPolicyParams) WithTimeout(timeout time.Duration) *GetPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get policy params
func (o *GetPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get policy params
func (o *GetPolicyParams) WithContext(ctx context.Context) *GetPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get policy params
func (o *GetPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get policy params
func (o *GetPolicyParams) WithHTTPClient(client *http.Client) *GetPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get policy params
func (o *GetPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDetail adds the detail to the get policy params
func (o *GetPolicyParams) WithDetail(detail *bool) *GetPolicyParams {
	o.SetDetail(detail)
	return o
}

// SetDetail adds the detail to the get policy params
func (o *GetPolicyParams) SetDetail(detail *bool) {
	o.Detail = detail
}

// WithPolicyID adds the policyID to the get policy params
func (o *GetPolicyParams) WithPolicyID(policyID string) *GetPolicyParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the get policy params
func (o *GetPolicyParams) SetPolicyID(policyID string) {
	o.PolicyID = policyID
}

// WithXAnchoreAccount adds the xAnchoreAccount to the get policy params
func (o *GetPolicyParams) WithXAnchoreAccount(xAnchoreAccount *string) *GetPolicyParams {
	o.SetXAnchoreAccount(xAnchoreAccount)
	return o
}

// SetXAnchoreAccount adds the xAnchoreAccount to the get policy params
func (o *GetPolicyParams) SetXAnchoreAccount(xAnchoreAccount *string) {
	o.XAnchoreAccount = xAnchoreAccount
}

// WriteToRequest writes these params to a swagger request
func (o *GetPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Detail != nil {

		// query param detail
		var qrDetail bool
		if o.Detail != nil {
			qrDetail = *o.Detail
		}
		qDetail := swag.FormatBool(qrDetail)
		if qDetail != "" {
			if err := r.SetQueryParam("detail", qDetail); err != nil {
				return err
			}
		}

	}

	// path param policyId
	if err := r.SetPathParam("policyId", o.PolicyID); err != nil {
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