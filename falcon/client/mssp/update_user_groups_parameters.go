// Code generated by go-swagger; DO NOT EDIT.

package mssp

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewUpdateUserGroupsParams creates a new UpdateUserGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUserGroupsParams() *UpdateUserGroupsParams {
	return &UpdateUserGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserGroupsParamsWithTimeout creates a new UpdateUserGroupsParams object
// with the ability to set a timeout on a request.
func NewUpdateUserGroupsParamsWithTimeout(timeout time.Duration) *UpdateUserGroupsParams {
	return &UpdateUserGroupsParams{
		timeout: timeout,
	}
}

// NewUpdateUserGroupsParamsWithContext creates a new UpdateUserGroupsParams object
// with the ability to set a context for a request.
func NewUpdateUserGroupsParamsWithContext(ctx context.Context) *UpdateUserGroupsParams {
	return &UpdateUserGroupsParams{
		Context: ctx,
	}
}

// NewUpdateUserGroupsParamsWithHTTPClient creates a new UpdateUserGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUserGroupsParamsWithHTTPClient(client *http.Client) *UpdateUserGroupsParams {
	return &UpdateUserGroupsParams{
		HTTPClient: client,
	}
}

/*
UpdateUserGroupsParams contains all the parameters to send to the API endpoint

	for the update user groups operation.

	Typically these are written to a http.Request.
*/
type UpdateUserGroupsParams struct {

	// Body.
	Body *models.DomainUserGroupsRequestV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update user groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserGroupsParams) WithDefaults() *UpdateUserGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update user groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update user groups params
func (o *UpdateUserGroupsParams) WithTimeout(timeout time.Duration) *UpdateUserGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user groups params
func (o *UpdateUserGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user groups params
func (o *UpdateUserGroupsParams) WithContext(ctx context.Context) *UpdateUserGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user groups params
func (o *UpdateUserGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user groups params
func (o *UpdateUserGroupsParams) WithHTTPClient(client *http.Client) *UpdateUserGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user groups params
func (o *UpdateUserGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update user groups params
func (o *UpdateUserGroupsParams) WithBody(body *models.DomainUserGroupsRequestV1) *UpdateUserGroupsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update user groups params
func (o *UpdateUserGroupsParams) SetBody(body *models.DomainUserGroupsRequestV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
