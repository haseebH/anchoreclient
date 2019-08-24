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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteImageByImageIDParams creates a new DeleteImageByImageIDParams object
// with the default values initialized.
func NewDeleteImageByImageIDParams() *DeleteImageByImageIDParams {
	var ()
	return &DeleteImageByImageIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteImageByImageIDParamsWithTimeout creates a new DeleteImageByImageIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteImageByImageIDParamsWithTimeout(timeout time.Duration) *DeleteImageByImageIDParams {
	var ()
	return &DeleteImageByImageIDParams{

		timeout: timeout,
	}
}

// NewDeleteImageByImageIDParamsWithContext creates a new DeleteImageByImageIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteImageByImageIDParamsWithContext(ctx context.Context) *DeleteImageByImageIDParams {
	var ()
	return &DeleteImageByImageIDParams{

		Context: ctx,
	}
}

// NewDeleteImageByImageIDParamsWithHTTPClient creates a new DeleteImageByImageIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteImageByImageIDParamsWithHTTPClient(client *http.Client) *DeleteImageByImageIDParams {
	var ()
	return &DeleteImageByImageIDParams{
		HTTPClient: client,
	}
}

/*DeleteImageByImageIDParams contains all the parameters to send to the API endpoint
for the delete image by image Id operation typically these are written to a http.Request
*/
type DeleteImageByImageIDParams struct {

	/*Force*/
	Force *bool
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

// WithTimeout adds the timeout to the delete image by image Id params
func (o *DeleteImageByImageIDParams) WithTimeout(timeout time.Duration) *DeleteImageByImageIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete image by image Id params
func (o *DeleteImageByImageIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete image by image Id params
func (o *DeleteImageByImageIDParams) WithContext(ctx context.Context) *DeleteImageByImageIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete image by image Id params
func (o *DeleteImageByImageIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete image by image Id params
func (o *DeleteImageByImageIDParams) WithHTTPClient(client *http.Client) *DeleteImageByImageIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete image by image Id params
func (o *DeleteImageByImageIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the delete image by image Id params
func (o *DeleteImageByImageIDParams) WithForce(force *bool) *DeleteImageByImageIDParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the delete image by image Id params
func (o *DeleteImageByImageIDParams) SetForce(force *bool) {
	o.Force = force
}

// WithImageID adds the imageID to the delete image by image Id params
func (o *DeleteImageByImageIDParams) WithImageID(imageID string) *DeleteImageByImageIDParams {
	o.SetImageID(imageID)
	return o
}

// SetImageID adds the imageId to the delete image by image Id params
func (o *DeleteImageByImageIDParams) SetImageID(imageID string) {
	o.ImageID = imageID
}

// WithXAnchoreAccount adds the xAnchoreAccount to the delete image by image Id params
func (o *DeleteImageByImageIDParams) WithXAnchoreAccount(xAnchoreAccount *string) *DeleteImageByImageIDParams {
	o.SetXAnchoreAccount(xAnchoreAccount)
	return o
}

// SetXAnchoreAccount adds the xAnchoreAccount to the delete image by image Id params
func (o *DeleteImageByImageIDParams) SetXAnchoreAccount(xAnchoreAccount *string) {
	o.XAnchoreAccount = xAnchoreAccount
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteImageByImageIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param force
		var qrForce bool
		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {
			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}

	}

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