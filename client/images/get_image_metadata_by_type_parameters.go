// Code generated by go-swagger; DO NOT EDIT.

package images

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

// NewGetImageMetadataByTypeParams creates a new GetImageMetadataByTypeParams object
// with the default values initialized.
func NewGetImageMetadataByTypeParams() *GetImageMetadataByTypeParams {
	var ()
	return &GetImageMetadataByTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetImageMetadataByTypeParamsWithTimeout creates a new GetImageMetadataByTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetImageMetadataByTypeParamsWithTimeout(timeout time.Duration) *GetImageMetadataByTypeParams {
	var ()
	return &GetImageMetadataByTypeParams{

		timeout: timeout,
	}
}

// NewGetImageMetadataByTypeParamsWithContext creates a new GetImageMetadataByTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetImageMetadataByTypeParamsWithContext(ctx context.Context) *GetImageMetadataByTypeParams {
	var ()
	return &GetImageMetadataByTypeParams{

		Context: ctx,
	}
}

// NewGetImageMetadataByTypeParamsWithHTTPClient creates a new GetImageMetadataByTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetImageMetadataByTypeParamsWithHTTPClient(client *http.Client) *GetImageMetadataByTypeParams {
	var ()
	return &GetImageMetadataByTypeParams{
		HTTPClient: client,
	}
}

/*GetImageMetadataByTypeParams contains all the parameters to send to the API endpoint
for the get image metadata by type operation typically these are written to a http.Request
*/
type GetImageMetadataByTypeParams struct {

	/*ImageDigest*/
	ImageDigest string
	/*Mtype*/
	Mtype string
	/*XAnchoreAccount
	  An account name to change the resource scope of the request to that account, if permissions allow (admin only)

	*/
	XAnchoreAccount *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get image metadata by type params
func (o *GetImageMetadataByTypeParams) WithTimeout(timeout time.Duration) *GetImageMetadataByTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get image metadata by type params
func (o *GetImageMetadataByTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get image metadata by type params
func (o *GetImageMetadataByTypeParams) WithContext(ctx context.Context) *GetImageMetadataByTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get image metadata by type params
func (o *GetImageMetadataByTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get image metadata by type params
func (o *GetImageMetadataByTypeParams) WithHTTPClient(client *http.Client) *GetImageMetadataByTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get image metadata by type params
func (o *GetImageMetadataByTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImageDigest adds the imageDigest to the get image metadata by type params
func (o *GetImageMetadataByTypeParams) WithImageDigest(imageDigest string) *GetImageMetadataByTypeParams {
	o.SetImageDigest(imageDigest)
	return o
}

// SetImageDigest adds the imageDigest to the get image metadata by type params
func (o *GetImageMetadataByTypeParams) SetImageDigest(imageDigest string) {
	o.ImageDigest = imageDigest
}

// WithMtype adds the mtype to the get image metadata by type params
func (o *GetImageMetadataByTypeParams) WithMtype(mtype string) *GetImageMetadataByTypeParams {
	o.SetMtype(mtype)
	return o
}

// SetMtype adds the mtype to the get image metadata by type params
func (o *GetImageMetadataByTypeParams) SetMtype(mtype string) {
	o.Mtype = mtype
}

// WithXAnchoreAccount adds the xAnchoreAccount to the get image metadata by type params
func (o *GetImageMetadataByTypeParams) WithXAnchoreAccount(xAnchoreAccount *string) *GetImageMetadataByTypeParams {
	o.SetXAnchoreAccount(xAnchoreAccount)
	return o
}

// SetXAnchoreAccount adds the xAnchoreAccount to the get image metadata by type params
func (o *GetImageMetadataByTypeParams) SetXAnchoreAccount(xAnchoreAccount *string) {
	o.XAnchoreAccount = xAnchoreAccount
}

// WriteToRequest writes these params to a swagger request
func (o *GetImageMetadataByTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param imageDigest
	if err := r.SetPathParam("imageDigest", o.ImageDigest); err != nil {
		return err
	}

	// path param mtype
	if err := r.SetPathParam("mtype", o.Mtype); err != nil {
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
