// Code generated by go-swagger; DO NOT EDIT.

package events

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

// NewListEventsParams creates a new ListEventsParams object
// with the default values initialized.
func NewListEventsParams() *ListEventsParams {
	var (
		limitDefault = int64(100)
		pageDefault  = int64(1)
	)
	return &ListEventsParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListEventsParamsWithTimeout creates a new ListEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListEventsParamsWithTimeout(timeout time.Duration) *ListEventsParams {
	var (
		limitDefault = int64(100)
		pageDefault  = int64(1)
	)
	return &ListEventsParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: timeout,
	}
}

// NewListEventsParamsWithContext creates a new ListEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListEventsParamsWithContext(ctx context.Context) *ListEventsParams {
	var (
		limitDefault = int64(100)
		pageDefault  = int64(1)
	)
	return &ListEventsParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		Context: ctx,
	}
}

// NewListEventsParamsWithHTTPClient creates a new ListEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListEventsParamsWithHTTPClient(client *http.Client) *ListEventsParams {
	var (
		limitDefault = int64(100)
		pageDefault  = int64(1)
	)
	return &ListEventsParams{
		Limit:      &limitDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*ListEventsParams contains all the parameters to send to the API endpoint
for the list events operation typically these are written to a http.Request
*/
type ListEventsParams struct {

	/*Before
	  Return events that occurred before the timestamp

	*/
	Before *string
	/*Level
	  Filter events by the level - INFO or ERROR

	*/
	Level *string
	/*Limit
	  Number of events in the result set. Defaults to 100 if left empty

	*/
	Limit *int64
	/*Page
	  Pagination controls - return the nth page of results. Defaults to first page if left empty

	*/
	Page *int64
	/*ResourceID
	  Filter events by the id of the resource

	*/
	ResourceID *string
	/*ResourceType
	  Filter events by the type of resource - tag, imageDigest, repository etc

	*/
	ResourceType *string
	/*Since
	  Return events that occurred after the timestamp

	*/
	Since *string
	/*SourceHostid
	  Filter events by the originating host ID

	*/
	SourceHostid *string
	/*SourceServicename
	  Filter events by the originating service

	*/
	SourceServicename *string
	/*XAnchoreAccount
	  An account name to change the resource scope of the request to that account, if permissions allow (admin only)

	*/
	XAnchoreAccount *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list events params
func (o *ListEventsParams) WithTimeout(timeout time.Duration) *ListEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list events params
func (o *ListEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list events params
func (o *ListEventsParams) WithContext(ctx context.Context) *ListEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list events params
func (o *ListEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list events params
func (o *ListEventsParams) WithHTTPClient(client *http.Client) *ListEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list events params
func (o *ListEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBefore adds the before to the list events params
func (o *ListEventsParams) WithBefore(before *string) *ListEventsParams {
	o.SetBefore(before)
	return o
}

// SetBefore adds the before to the list events params
func (o *ListEventsParams) SetBefore(before *string) {
	o.Before = before
}

// WithLevel adds the level to the list events params
func (o *ListEventsParams) WithLevel(level *string) *ListEventsParams {
	o.SetLevel(level)
	return o
}

// SetLevel adds the level to the list events params
func (o *ListEventsParams) SetLevel(level *string) {
	o.Level = level
}

// WithLimit adds the limit to the list events params
func (o *ListEventsParams) WithLimit(limit *int64) *ListEventsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list events params
func (o *ListEventsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the list events params
func (o *ListEventsParams) WithPage(page *int64) *ListEventsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list events params
func (o *ListEventsParams) SetPage(page *int64) {
	o.Page = page
}

// WithResourceID adds the resourceID to the list events params
func (o *ListEventsParams) WithResourceID(resourceID *string) *ListEventsParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the list events params
func (o *ListEventsParams) SetResourceID(resourceID *string) {
	o.ResourceID = resourceID
}

// WithResourceType adds the resourceType to the list events params
func (o *ListEventsParams) WithResourceType(resourceType *string) *ListEventsParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the list events params
func (o *ListEventsParams) SetResourceType(resourceType *string) {
	o.ResourceType = resourceType
}

// WithSince adds the since to the list events params
func (o *ListEventsParams) WithSince(since *string) *ListEventsParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the list events params
func (o *ListEventsParams) SetSince(since *string) {
	o.Since = since
}

// WithSourceHostid adds the sourceHostid to the list events params
func (o *ListEventsParams) WithSourceHostid(sourceHostid *string) *ListEventsParams {
	o.SetSourceHostid(sourceHostid)
	return o
}

// SetSourceHostid adds the sourceHostid to the list events params
func (o *ListEventsParams) SetSourceHostid(sourceHostid *string) {
	o.SourceHostid = sourceHostid
}

// WithSourceServicename adds the sourceServicename to the list events params
func (o *ListEventsParams) WithSourceServicename(sourceServicename *string) *ListEventsParams {
	o.SetSourceServicename(sourceServicename)
	return o
}

// SetSourceServicename adds the sourceServicename to the list events params
func (o *ListEventsParams) SetSourceServicename(sourceServicename *string) {
	o.SourceServicename = sourceServicename
}

// WithXAnchoreAccount adds the xAnchoreAccount to the list events params
func (o *ListEventsParams) WithXAnchoreAccount(xAnchoreAccount *string) *ListEventsParams {
	o.SetXAnchoreAccount(xAnchoreAccount)
	return o
}

// SetXAnchoreAccount adds the xAnchoreAccount to the list events params
func (o *ListEventsParams) SetXAnchoreAccount(xAnchoreAccount *string) {
	o.XAnchoreAccount = xAnchoreAccount
}

// WriteToRequest writes these params to a swagger request
func (o *ListEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Before != nil {

		// query param before
		var qrBefore string
		if o.Before != nil {
			qrBefore = *o.Before
		}
		qBefore := qrBefore
		if qBefore != "" {
			if err := r.SetQueryParam("before", qBefore); err != nil {
				return err
			}
		}

	}

	if o.Level != nil {

		// query param level
		var qrLevel string
		if o.Level != nil {
			qrLevel = *o.Level
		}
		qLevel := qrLevel
		if qLevel != "" {
			if err := r.SetQueryParam("level", qLevel); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.ResourceID != nil {

		// query param resource_id
		var qrResourceID string
		if o.ResourceID != nil {
			qrResourceID = *o.ResourceID
		}
		qResourceID := qrResourceID
		if qResourceID != "" {
			if err := r.SetQueryParam("resource_id", qResourceID); err != nil {
				return err
			}
		}

	}

	if o.ResourceType != nil {

		// query param resource_type
		var qrResourceType string
		if o.ResourceType != nil {
			qrResourceType = *o.ResourceType
		}
		qResourceType := qrResourceType
		if qResourceType != "" {
			if err := r.SetQueryParam("resource_type", qResourceType); err != nil {
				return err
			}
		}

	}

	if o.Since != nil {

		// query param since
		var qrSince string
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := qrSince
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

	}

	if o.SourceHostid != nil {

		// query param source_hostid
		var qrSourceHostid string
		if o.SourceHostid != nil {
			qrSourceHostid = *o.SourceHostid
		}
		qSourceHostid := qrSourceHostid
		if qSourceHostid != "" {
			if err := r.SetQueryParam("source_hostid", qSourceHostid); err != nil {
				return err
			}
		}

	}

	if o.SourceServicename != nil {

		// query param source_servicename
		var qrSourceServicename string
		if o.SourceServicename != nil {
			qrSourceServicename = *o.SourceServicename
		}
		qSourceServicename := qrSourceServicename
		if qSourceServicename != "" {
			if err := r.SetQueryParam("source_servicename", qSourceServicename); err != nil {
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
