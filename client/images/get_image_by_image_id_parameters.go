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

// NewGetImageByImageIDParams creates a new GetImageByImageIDParams object
// with the default values initialized.
func NewGetImageByImageIDParams() *GetImageByImageIDParams {
	var ()
	return &GetImageByImageIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetImageByImageIDParamsWithTimeout creates a new GetImageByImageIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetImageByImageIDParamsWithTimeout(timeout time.Duration) *GetImageByImageIDParams {
	var ()
	return &GetImageByImageIDParams{

		timeout: timeout,
	}
}

// NewGetImageByImageIDParamsWithContext creates a new GetImageByImageIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetImageByImageIDParamsWithContext(ctx context.Context) *GetImageByImageIDParams {
	var ()
	return &GetImageByImageIDParams{

		Context: ctx,
	}
}

// NewGetImageByImageIDParamsWithHTTPClient creates a new GetImageByImageIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetImageByImageIDParamsWithHTTPClient(client *http.Client) *GetImageByImageIDParams {
	var ()
	return &GetImageByImageIDParams{
		HTTPClient: client,
	}
}

/*GetImageByImageIDParams contains all the parameters to send to the API endpoint
for the get image by image Id operation typically these are written to a http.Request
*/
type GetImageByImageIDParams struct {

	/*ImageID*/
	ImageID string
	/*XAnchoreAccount
	  An account name to change the resource scope of the request to that account, if permissions allow (admin only)

	*/
	XAnchoreAccount *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get image by image Id params
func (o *GetImageByImageIDParams) WithTimeout(timeout time.Duration) *GetImageByImageIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get image by image Id params
func (o *GetImageByImageIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get image by image Id params
func (o *GetImageByImageIDParams) WithContext(ctx context.Context) *GetImageByImageIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get image by image Id params
func (o *GetImageByImageIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get image by image Id params
func (o *GetImageByImageIDParams) WithHTTPClient(client *http.Client) *GetImageByImageIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get image by image Id params
func (o *GetImageByImageIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImageID adds the imageID to the get image by image Id params
func (o *GetImageByImageIDParams) WithImageID(imageID string) *GetImageByImageIDParams {
	o.SetImageID(imageID)
	return o
}

// SetImageID adds the imageId to the get image by image Id params
func (o *GetImageByImageIDParams) SetImageID(imageID string) {
	o.ImageID = imageID
}

// WithXAnchoreAccount adds the xAnchoreAccount to the get image by image Id params
func (o *GetImageByImageIDParams) WithXAnchoreAccount(xAnchoreAccount *string) *GetImageByImageIDParams {
	o.SetXAnchoreAccount(xAnchoreAccount)
	return o
}

// SetXAnchoreAccount adds the xAnchoreAccount to the get image by image Id params
func (o *GetImageByImageIDParams) SetXAnchoreAccount(xAnchoreAccount *string) {
	o.XAnchoreAccount = xAnchoreAccount
}

// WriteToRequest writes these params to a swagger request
func (o *GetImageByImageIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param imageId
	if err := r.SetPathParam("imageId", o.ImageID); err != nil {
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
