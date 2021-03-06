// Code generated by go-swagger; DO NOT EDIT.

package registries

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

	models "github.com/haseebh/anchoreclient/models"
)

// NewCreateRegistryParams creates a new CreateRegistryParams object
// with the default values initialized.
func NewCreateRegistryParams() *CreateRegistryParams {
	var ()
	return &CreateRegistryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRegistryParamsWithTimeout creates a new CreateRegistryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRegistryParamsWithTimeout(timeout time.Duration) *CreateRegistryParams {
	var ()
	return &CreateRegistryParams{

		timeout: timeout,
	}
}

// NewCreateRegistryParamsWithContext creates a new CreateRegistryParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRegistryParamsWithContext(ctx context.Context) *CreateRegistryParams {
	var ()
	return &CreateRegistryParams{

		Context: ctx,
	}
}

// NewCreateRegistryParamsWithHTTPClient creates a new CreateRegistryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRegistryParamsWithHTTPClient(client *http.Client) *CreateRegistryParams {
	var ()
	return &CreateRegistryParams{
		HTTPClient: client,
	}
}

/*CreateRegistryParams contains all the parameters to send to the API endpoint
for the create registry operation typically these are written to a http.Request
*/
type CreateRegistryParams struct {

	/*Registrydata*/
	Registrydata *models.RegistryConfigurationRequest
	/*Validate
	  flag to determine whether or not to validate registry/credential at registry add time

	*/
	Validate *bool
	/*XAnchoreAccount
	  An account name to change the resource scope of the request to that account, if permissions allow (admin only)

	*/
	XAnchoreAccount *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create registry params
func (o *CreateRegistryParams) WithTimeout(timeout time.Duration) *CreateRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create registry params
func (o *CreateRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create registry params
func (o *CreateRegistryParams) WithContext(ctx context.Context) *CreateRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create registry params
func (o *CreateRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create registry params
func (o *CreateRegistryParams) WithHTTPClient(client *http.Client) *CreateRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create registry params
func (o *CreateRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegistrydata adds the registrydata to the create registry params
func (o *CreateRegistryParams) WithRegistrydata(registrydata *models.RegistryConfigurationRequest) *CreateRegistryParams {
	o.SetRegistrydata(registrydata)
	return o
}

// SetRegistrydata adds the registrydata to the create registry params
func (o *CreateRegistryParams) SetRegistrydata(registrydata *models.RegistryConfigurationRequest) {
	o.Registrydata = registrydata
}

// WithValidate adds the validate to the create registry params
func (o *CreateRegistryParams) WithValidate(validate *bool) *CreateRegistryParams {
	o.SetValidate(validate)
	return o
}

// SetValidate adds the validate to the create registry params
func (o *CreateRegistryParams) SetValidate(validate *bool) {
	o.Validate = validate
}

// WithXAnchoreAccount adds the xAnchoreAccount to the create registry params
func (o *CreateRegistryParams) WithXAnchoreAccount(xAnchoreAccount *string) *CreateRegistryParams {
	o.SetXAnchoreAccount(xAnchoreAccount)
	return o
}

// SetXAnchoreAccount adds the xAnchoreAccount to the create registry params
func (o *CreateRegistryParams) SetXAnchoreAccount(xAnchoreAccount *string) {
	o.XAnchoreAccount = xAnchoreAccount
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Registrydata != nil {
		if err := r.SetBodyParam(o.Registrydata); err != nil {
			return err
		}
	}

	if o.Validate != nil {

		// query param validate
		var qrValidate bool
		if o.Validate != nil {
			qrValidate = *o.Validate
		}
		qValidate := swag.FormatBool(qrValidate)
		if qValidate != "" {
			if err := r.SetQueryParam("validate", qValidate); err != nil {
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
