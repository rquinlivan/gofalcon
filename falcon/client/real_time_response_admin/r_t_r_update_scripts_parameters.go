// Code generated by go-swagger; DO NOT EDIT.

package real_time_response_admin

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
	"github.com/go-openapi/swag"
)

// NewRTRUpdateScriptsParams creates a new RTRUpdateScriptsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRTRUpdateScriptsParams() *RTRUpdateScriptsParams {
	return &RTRUpdateScriptsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRTRUpdateScriptsParamsWithTimeout creates a new RTRUpdateScriptsParams object
// with the ability to set a timeout on a request.
func NewRTRUpdateScriptsParamsWithTimeout(timeout time.Duration) *RTRUpdateScriptsParams {
	return &RTRUpdateScriptsParams{
		timeout: timeout,
	}
}

// NewRTRUpdateScriptsParamsWithContext creates a new RTRUpdateScriptsParams object
// with the ability to set a context for a request.
func NewRTRUpdateScriptsParamsWithContext(ctx context.Context) *RTRUpdateScriptsParams {
	return &RTRUpdateScriptsParams{
		Context: ctx,
	}
}

// NewRTRUpdateScriptsParamsWithHTTPClient creates a new RTRUpdateScriptsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRTRUpdateScriptsParamsWithHTTPClient(client *http.Client) *RTRUpdateScriptsParams {
	return &RTRUpdateScriptsParams{
		HTTPClient: client,
	}
}

/*
RTRUpdateScriptsParams contains all the parameters to send to the API endpoint

	for the r t r update scripts operation.

	Typically these are written to a http.Request.
*/
type RTRUpdateScriptsParams struct {

	/* CommentsForAuditLog.

	   The audit log comment
	*/
	CommentsForAuditLog *string

	/* Content.

	   The script text that you want to use to upload
	*/
	Content *string

	/* Description.

	   File description
	*/
	Description *string

	/* File.

	   custom-script file to upload.  These should be powershell scripts.
	*/
	File runtime.NamedReadCloser

	/* ID.

	   ID to update
	*/
	ID string

	/* Name.

	   File name (if different than actual file name)
	*/
	Name *string

	/* PermissionType.

	    Permission for the custom-script. Valid permission values:
	- `private`, usable by only the user who uploaded it
	- `group`, usable by all RTR Admins
	- `public`, usable by all active-responders and RTR admins

	    Default: "none"
	*/
	PermissionType *string

	/* Platform.

	   Platforms for the file. Currently supports: windows, mac,
	*/
	Platform []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the r t r update scripts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RTRUpdateScriptsParams) WithDefaults() *RTRUpdateScriptsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the r t r update scripts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RTRUpdateScriptsParams) SetDefaults() {
	var (
		permissionTypeDefault = string("none")
	)

	val := RTRUpdateScriptsParams{
		PermissionType: &permissionTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the r t r update scripts params
func (o *RTRUpdateScriptsParams) WithTimeout(timeout time.Duration) *RTRUpdateScriptsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the r t r update scripts params
func (o *RTRUpdateScriptsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the r t r update scripts params
func (o *RTRUpdateScriptsParams) WithContext(ctx context.Context) *RTRUpdateScriptsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the r t r update scripts params
func (o *RTRUpdateScriptsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the r t r update scripts params
func (o *RTRUpdateScriptsParams) WithHTTPClient(client *http.Client) *RTRUpdateScriptsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the r t r update scripts params
func (o *RTRUpdateScriptsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommentsForAuditLog adds the commentsForAuditLog to the r t r update scripts params
func (o *RTRUpdateScriptsParams) WithCommentsForAuditLog(commentsForAuditLog *string) *RTRUpdateScriptsParams {
	o.SetCommentsForAuditLog(commentsForAuditLog)
	return o
}

// SetCommentsForAuditLog adds the commentsForAuditLog to the r t r update scripts params
func (o *RTRUpdateScriptsParams) SetCommentsForAuditLog(commentsForAuditLog *string) {
	o.CommentsForAuditLog = commentsForAuditLog
}

// WithContent adds the content to the r t r update scripts params
func (o *RTRUpdateScriptsParams) WithContent(content *string) *RTRUpdateScriptsParams {
	o.SetContent(content)
	return o
}

// SetContent adds the content to the r t r update scripts params
func (o *RTRUpdateScriptsParams) SetContent(content *string) {
	o.Content = content
}

// WithDescription adds the description to the r t r update scripts params
func (o *RTRUpdateScriptsParams) WithDescription(description *string) *RTRUpdateScriptsParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the r t r update scripts params
func (o *RTRUpdateScriptsParams) SetDescription(description *string) {
	o.Description = description
}

// WithFile adds the file to the r t r update scripts params
func (o *RTRUpdateScriptsParams) WithFile(file runtime.NamedReadCloser) *RTRUpdateScriptsParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the r t r update scripts params
func (o *RTRUpdateScriptsParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithID adds the id to the r t r update scripts params
func (o *RTRUpdateScriptsParams) WithID(id string) *RTRUpdateScriptsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the r t r update scripts params
func (o *RTRUpdateScriptsParams) SetID(id string) {
	o.ID = id
}

// WithName adds the name to the r t r update scripts params
func (o *RTRUpdateScriptsParams) WithName(name *string) *RTRUpdateScriptsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the r t r update scripts params
func (o *RTRUpdateScriptsParams) SetName(name *string) {
	o.Name = name
}

// WithPermissionType adds the permissionType to the r t r update scripts params
func (o *RTRUpdateScriptsParams) WithPermissionType(permissionType *string) *RTRUpdateScriptsParams {
	o.SetPermissionType(permissionType)
	return o
}

// SetPermissionType adds the permissionType to the r t r update scripts params
func (o *RTRUpdateScriptsParams) SetPermissionType(permissionType *string) {
	o.PermissionType = permissionType
}

// WithPlatform adds the platform to the r t r update scripts params
func (o *RTRUpdateScriptsParams) WithPlatform(platform []string) *RTRUpdateScriptsParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the r t r update scripts params
func (o *RTRUpdateScriptsParams) SetPlatform(platform []string) {
	o.Platform = platform
}

// WriteToRequest writes these params to a swagger request
func (o *RTRUpdateScriptsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CommentsForAuditLog != nil {

		// form param comments_for_audit_log
		var frCommentsForAuditLog string
		if o.CommentsForAuditLog != nil {
			frCommentsForAuditLog = *o.CommentsForAuditLog
		}
		fCommentsForAuditLog := frCommentsForAuditLog
		if fCommentsForAuditLog != "" {
			if err := r.SetFormParam("comments_for_audit_log", fCommentsForAuditLog); err != nil {
				return err
			}
		}
	}

	if o.Content != nil {

		// form param content
		var frContent string
		if o.Content != nil {
			frContent = *o.Content
		}
		fContent := frContent
		if fContent != "" {
			if err := r.SetFormParam("content", fContent); err != nil {
				return err
			}
		}
	}

	if o.Description != nil {

		// form param description
		var frDescription string
		if o.Description != nil {
			frDescription = *o.Description
		}
		fDescription := frDescription
		if fDescription != "" {
			if err := r.SetFormParam("description", fDescription); err != nil {
				return err
			}
		}
	}

	if o.File != nil {

		if o.File != nil {
			// form file param file
			if err := r.SetFileParam("file", o.File); err != nil {
				return err
			}
		}
	}

	// form param id
	frID := o.ID
	fID := frID
	if fID != "" {
		if err := r.SetFormParam("id", fID); err != nil {
			return err
		}
	}

	if o.Name != nil {

		// form param name
		var frName string
		if o.Name != nil {
			frName = *o.Name
		}
		fName := frName
		if fName != "" {
			if err := r.SetFormParam("name", fName); err != nil {
				return err
			}
		}
	}

	if o.PermissionType != nil {

		// form param permission_type
		var frPermissionType string
		if o.PermissionType != nil {
			frPermissionType = *o.PermissionType
		}
		fPermissionType := frPermissionType
		if fPermissionType != "" {
			if err := r.SetFormParam("permission_type", fPermissionType); err != nil {
				return err
			}
		}
	}

	if o.Platform != nil {

		// binding items for platform
		joinedPlatform := o.bindParamPlatform(reg)

		// form array param platform
		if err := r.SetFormParam("platform", joinedPlatform...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamRTRUpdateScripts binds the parameter platform
func (o *RTRUpdateScriptsParams) bindParamPlatform(formats strfmt.Registry) []string {
	platformIR := o.Platform

	var platformIC []string
	for _, platformIIR := range platformIR { // explode []string

		platformIIV := platformIIR // string as string
		platformIC = append(platformIC, platformIIV)
	}

	// items.CollectionFormat: "multi"
	platformIS := swag.JoinByFormat(platformIC, "multi")

	return platformIS
}
