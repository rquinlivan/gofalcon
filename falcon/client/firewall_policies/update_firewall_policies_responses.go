// Code generated by go-swagger; DO NOT EDIT.

package firewall_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// UpdateFirewallPoliciesReader is a Reader for the UpdateFirewallPolicies structure.
type UpdateFirewallPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFirewallPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateFirewallPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateFirewallPoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateFirewallPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateFirewallPoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateFirewallPoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateFirewallPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateFirewallPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateFirewallPoliciesOK creates a UpdateFirewallPoliciesOK with default headers values
func NewUpdateFirewallPoliciesOK() *UpdateFirewallPoliciesOK {
	return &UpdateFirewallPoliciesOK{}
}

/*
UpdateFirewallPoliciesOK describes a response with status code 200, with default header values.

OK
*/
type UpdateFirewallPoliciesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesFirewallPoliciesV1
}

func (o *UpdateFirewallPoliciesOK) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/firewall/v1][%d] updateFirewallPoliciesOK  %+v", 200, o.Payload)
}
func (o *UpdateFirewallPoliciesOK) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *UpdateFirewallPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFirewallPoliciesBadRequest creates a UpdateFirewallPoliciesBadRequest with default headers values
func NewUpdateFirewallPoliciesBadRequest() *UpdateFirewallPoliciesBadRequest {
	return &UpdateFirewallPoliciesBadRequest{}
}

/*
UpdateFirewallPoliciesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateFirewallPoliciesBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesFirewallPoliciesV1
}

func (o *UpdateFirewallPoliciesBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/firewall/v1][%d] updateFirewallPoliciesBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateFirewallPoliciesBadRequest) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *UpdateFirewallPoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFirewallPoliciesForbidden creates a UpdateFirewallPoliciesForbidden with default headers values
func NewUpdateFirewallPoliciesForbidden() *UpdateFirewallPoliciesForbidden {
	return &UpdateFirewallPoliciesForbidden{}
}

/*
UpdateFirewallPoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateFirewallPoliciesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *UpdateFirewallPoliciesForbidden) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/firewall/v1][%d] updateFirewallPoliciesForbidden  %+v", 403, o.Payload)
}
func (o *UpdateFirewallPoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *UpdateFirewallPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFirewallPoliciesNotFound creates a UpdateFirewallPoliciesNotFound with default headers values
func NewUpdateFirewallPoliciesNotFound() *UpdateFirewallPoliciesNotFound {
	return &UpdateFirewallPoliciesNotFound{}
}

/*
UpdateFirewallPoliciesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateFirewallPoliciesNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesFirewallPoliciesV1
}

func (o *UpdateFirewallPoliciesNotFound) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/firewall/v1][%d] updateFirewallPoliciesNotFound  %+v", 404, o.Payload)
}
func (o *UpdateFirewallPoliciesNotFound) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *UpdateFirewallPoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFirewallPoliciesTooManyRequests creates a UpdateFirewallPoliciesTooManyRequests with default headers values
func NewUpdateFirewallPoliciesTooManyRequests() *UpdateFirewallPoliciesTooManyRequests {
	return &UpdateFirewallPoliciesTooManyRequests{}
}

/*
UpdateFirewallPoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateFirewallPoliciesTooManyRequests struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

func (o *UpdateFirewallPoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/firewall/v1][%d] updateFirewallPoliciesTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateFirewallPoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateFirewallPoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFirewallPoliciesInternalServerError creates a UpdateFirewallPoliciesInternalServerError with default headers values
func NewUpdateFirewallPoliciesInternalServerError() *UpdateFirewallPoliciesInternalServerError {
	return &UpdateFirewallPoliciesInternalServerError{}
}

/*
UpdateFirewallPoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateFirewallPoliciesInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesFirewallPoliciesV1
}

func (o *UpdateFirewallPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/firewall/v1][%d] updateFirewallPoliciesInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateFirewallPoliciesInternalServerError) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *UpdateFirewallPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFirewallPoliciesDefault creates a UpdateFirewallPoliciesDefault with default headers values
func NewUpdateFirewallPoliciesDefault(code int) *UpdateFirewallPoliciesDefault {
	return &UpdateFirewallPoliciesDefault{
		_statusCode: code,
	}
}

/*
UpdateFirewallPoliciesDefault describes a response with status code -1, with default header values.

OK
*/
type UpdateFirewallPoliciesDefault struct {
	_statusCode int

	Payload *models.ResponsesFirewallPoliciesV1
}

// Code gets the status code for the update firewall policies default response
func (o *UpdateFirewallPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *UpdateFirewallPoliciesDefault) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/firewall/v1][%d] updateFirewallPolicies default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateFirewallPoliciesDefault) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *UpdateFirewallPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
