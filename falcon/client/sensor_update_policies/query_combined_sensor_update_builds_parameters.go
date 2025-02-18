// Code generated by go-swagger; DO NOT EDIT.

package sensor_update_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewQueryCombinedSensorUpdateBuildsParams creates a new QueryCombinedSensorUpdateBuildsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQueryCombinedSensorUpdateBuildsParams() *QueryCombinedSensorUpdateBuildsParams {
	return &QueryCombinedSensorUpdateBuildsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQueryCombinedSensorUpdateBuildsParamsWithTimeout creates a new QueryCombinedSensorUpdateBuildsParams object
// with the ability to set a timeout on a request.
func NewQueryCombinedSensorUpdateBuildsParamsWithTimeout(timeout time.Duration) *QueryCombinedSensorUpdateBuildsParams {
	return &QueryCombinedSensorUpdateBuildsParams{
		timeout: timeout,
	}
}

// NewQueryCombinedSensorUpdateBuildsParamsWithContext creates a new QueryCombinedSensorUpdateBuildsParams object
// with the ability to set a context for a request.
func NewQueryCombinedSensorUpdateBuildsParamsWithContext(ctx context.Context) *QueryCombinedSensorUpdateBuildsParams {
	return &QueryCombinedSensorUpdateBuildsParams{
		Context: ctx,
	}
}

// NewQueryCombinedSensorUpdateBuildsParamsWithHTTPClient creates a new QueryCombinedSensorUpdateBuildsParams object
// with the ability to set a custom HTTPClient for a request.
func NewQueryCombinedSensorUpdateBuildsParamsWithHTTPClient(client *http.Client) *QueryCombinedSensorUpdateBuildsParams {
	return &QueryCombinedSensorUpdateBuildsParams{
		HTTPClient: client,
	}
}

/*
QueryCombinedSensorUpdateBuildsParams contains all the parameters to send to the API endpoint

	for the query combined sensor update builds operation.

	Typically these are written to a http.Request.
*/
type QueryCombinedSensorUpdateBuildsParams struct {

	/* Platform.

	   The platform to return builds for
	*/
	Platform *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the query combined sensor update builds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryCombinedSensorUpdateBuildsParams) WithDefaults() *QueryCombinedSensorUpdateBuildsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the query combined sensor update builds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QueryCombinedSensorUpdateBuildsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the query combined sensor update builds params
func (o *QueryCombinedSensorUpdateBuildsParams) WithTimeout(timeout time.Duration) *QueryCombinedSensorUpdateBuildsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query combined sensor update builds params
func (o *QueryCombinedSensorUpdateBuildsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query combined sensor update builds params
func (o *QueryCombinedSensorUpdateBuildsParams) WithContext(ctx context.Context) *QueryCombinedSensorUpdateBuildsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query combined sensor update builds params
func (o *QueryCombinedSensorUpdateBuildsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query combined sensor update builds params
func (o *QueryCombinedSensorUpdateBuildsParams) WithHTTPClient(client *http.Client) *QueryCombinedSensorUpdateBuildsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query combined sensor update builds params
func (o *QueryCombinedSensorUpdateBuildsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPlatform adds the platform to the query combined sensor update builds params
func (o *QueryCombinedSensorUpdateBuildsParams) WithPlatform(platform *string) *QueryCombinedSensorUpdateBuildsParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the query combined sensor update builds params
func (o *QueryCombinedSensorUpdateBuildsParams) SetPlatform(platform *string) {
	o.Platform = platform
}

// WriteToRequest writes these params to a swagger request
func (o *QueryCombinedSensorUpdateBuildsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Platform != nil {

		// query param platform
		var qrPlatform string

		if o.Platform != nil {
			qrPlatform = *o.Platform
		}
		qPlatform := qrPlatform
		if qPlatform != "" {

			if err := r.SetQueryParam("platform", qPlatform); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
