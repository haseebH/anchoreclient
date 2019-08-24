// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

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

// NewListSubscriptionsParams creates a new ListSubscriptionsParams object
// with the default values initialized.
func NewListSubscriptionsParams() *ListSubscriptionsParams {
	var ()
	return &ListSubscriptionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSubscriptionsParamsWithTimeout creates a new ListSubscriptionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSubscriptionsParamsWithTimeout(timeout time.Duration) *ListSubscriptionsParams {
	var ()
	return &ListSubscriptionsParams{

		timeout: timeout,
	}
}

// NewListSubscriptionsParamsWithContext creates a new ListSubscriptionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSubscriptionsParamsWithContext(ctx context.Context) *ListSubscriptionsParams {
	var ()
	return &ListSubscriptionsParams{

		Context: ctx,
	}
}

// NewListSubscriptionsParamsWithHTTPClient creates a new ListSubscriptionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSubscriptionsParamsWithHTTPClient(client *http.Client) *ListSubscriptionsParams {
	var ()
	return &ListSubscriptionsParams{
		HTTPClient: client,
	}
}

/*ListSubscriptionsParams contains all the parameters to send to the API endpoint
for the list subscriptions operation typically these are written to a http.Request
*/
type ListSubscriptionsParams struct {

	/*SubscriptionKey
	  filter only subscriptions matching key

	*/
	SubscriptionKey *string
	/*SubscriptionType
	  filter only subscriptions matching type

	*/
	SubscriptionType *string
	/*XAnchoreAccount
	  An account name to change the resource scope of the request to that account, if permissions allow (admin only)

	*/
	XAnchoreAccount *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list subscriptions params
func (o *ListSubscriptionsParams) WithTimeout(timeout time.Duration) *ListSubscriptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list subscriptions params
func (o *ListSubscriptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list subscriptions params
func (o *ListSubscriptionsParams) WithContext(ctx context.Context) *ListSubscriptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list subscriptions params
func (o *ListSubscriptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list subscriptions params
func (o *ListSubscriptionsParams) WithHTTPClient(client *http.Client) *ListSubscriptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list subscriptions params
func (o *ListSubscriptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubscriptionKey adds the subscriptionKey to the list subscriptions params
func (o *ListSubscriptionsParams) WithSubscriptionKey(subscriptionKey *string) *ListSubscriptionsParams {
	o.SetSubscriptionKey(subscriptionKey)
	return o
}

// SetSubscriptionKey adds the subscriptionKey to the list subscriptions params
func (o *ListSubscriptionsParams) SetSubscriptionKey(subscriptionKey *string) {
	o.SubscriptionKey = subscriptionKey
}

// WithSubscriptionType adds the subscriptionType to the list subscriptions params
func (o *ListSubscriptionsParams) WithSubscriptionType(subscriptionType *string) *ListSubscriptionsParams {
	o.SetSubscriptionType(subscriptionType)
	return o
}

// SetSubscriptionType adds the subscriptionType to the list subscriptions params
func (o *ListSubscriptionsParams) SetSubscriptionType(subscriptionType *string) {
	o.SubscriptionType = subscriptionType
}

// WithXAnchoreAccount adds the xAnchoreAccount to the list subscriptions params
func (o *ListSubscriptionsParams) WithXAnchoreAccount(xAnchoreAccount *string) *ListSubscriptionsParams {
	o.SetXAnchoreAccount(xAnchoreAccount)
	return o
}

// SetXAnchoreAccount adds the xAnchoreAccount to the list subscriptions params
func (o *ListSubscriptionsParams) SetXAnchoreAccount(xAnchoreAccount *string) {
	o.XAnchoreAccount = xAnchoreAccount
}

// WriteToRequest writes these params to a swagger request
func (o *ListSubscriptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SubscriptionKey != nil {

		// query param subscription_key
		var qrSubscriptionKey string
		if o.SubscriptionKey != nil {
			qrSubscriptionKey = *o.SubscriptionKey
		}
		qSubscriptionKey := qrSubscriptionKey
		if qSubscriptionKey != "" {
			if err := r.SetQueryParam("subscription_key", qSubscriptionKey); err != nil {
				return err
			}
		}

	}

	if o.SubscriptionType != nil {

		// query param subscription_type
		var qrSubscriptionType string
		if o.SubscriptionType != nil {
			qrSubscriptionType = *o.SubscriptionType
		}
		qSubscriptionType := qrSubscriptionType
		if qSubscriptionType != "" {
			if err := r.SetQueryParam("subscription_type", qSubscriptionType); err != nil {
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