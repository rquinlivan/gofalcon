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

// GetFirewallPoliciesReader is a Reader for the GetFirewallPolicies structure.
type GetFirewallPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFirewallPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFirewallPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetFirewallPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFirewallPoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFirewallPoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFirewallPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetFirewallPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFirewallPoliciesOK creates a GetFirewallPoliciesOK with default headers values
func NewGetFirewallPoliciesOK() *GetFirewallPoliciesOK {
	return &GetFirewallPoliciesOK{}
}

/*
GetFirewallPoliciesOK describes a response with status code 200, with default header values.

OK
*/
type GetFirewallPoliciesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesFirewallPoliciesV1
}

func (o *GetFirewallPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /policy/entities/firewall/v1][%d] getFirewallPoliciesOK  %+v", 200, o.Payload)
}
func (o *GetFirewallPoliciesOK) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *GetFirewallPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFirewallPoliciesForbidden creates a GetFirewallPoliciesForbidden with default headers values
func NewGetFirewallPoliciesForbidden() *GetFirewallPoliciesForbidden {
	return &GetFirewallPoliciesForbidden{}
}

/*
GetFirewallPoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetFirewallPoliciesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *GetFirewallPoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/entities/firewall/v1][%d] getFirewallPoliciesForbidden  %+v", 403, o.Payload)
}
func (o *GetFirewallPoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetFirewallPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFirewallPoliciesNotFound creates a GetFirewallPoliciesNotFound with default headers values
func NewGetFirewallPoliciesNotFound() *GetFirewallPoliciesNotFound {
	return &GetFirewallPoliciesNotFound{}
}

/*
GetFirewallPoliciesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetFirewallPoliciesNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesFirewallPoliciesV1
}

func (o *GetFirewallPoliciesNotFound) Error() string {
	return fmt.Sprintf("[GET /policy/entities/firewall/v1][%d] getFirewallPoliciesNotFound  %+v", 404, o.Payload)
}
func (o *GetFirewallPoliciesNotFound) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *GetFirewallPoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFirewallPoliciesTooManyRequests creates a GetFirewallPoliciesTooManyRequests with default headers values
func NewGetFirewallPoliciesTooManyRequests() *GetFirewallPoliciesTooManyRequests {
	return &GetFirewallPoliciesTooManyRequests{}
}

/*
GetFirewallPoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetFirewallPoliciesTooManyRequests struct {

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

func (o *GetFirewallPoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/entities/firewall/v1][%d] getFirewallPoliciesTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetFirewallPoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetFirewallPoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFirewallPoliciesInternalServerError creates a GetFirewallPoliciesInternalServerError with default headers values
func NewGetFirewallPoliciesInternalServerError() *GetFirewallPoliciesInternalServerError {
	return &GetFirewallPoliciesInternalServerError{}
}

/*
GetFirewallPoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetFirewallPoliciesInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesFirewallPoliciesV1
}

func (o *GetFirewallPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/entities/firewall/v1][%d] getFirewallPoliciesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetFirewallPoliciesInternalServerError) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *GetFirewallPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFirewallPoliciesDefault creates a GetFirewallPoliciesDefault with default headers values
func NewGetFirewallPoliciesDefault(code int) *GetFirewallPoliciesDefault {
	return &GetFirewallPoliciesDefault{
		_statusCode: code,
	}
}

/*
GetFirewallPoliciesDefault describes a response with status code -1, with default header values.

OK
*/
type GetFirewallPoliciesDefault struct {
	_statusCode int

	Payload *models.ResponsesFirewallPoliciesV1
}

// Code gets the status code for the get firewall policies default response
func (o *GetFirewallPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *GetFirewallPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /policy/entities/firewall/v1][%d] getFirewallPolicies default  %+v", o._statusCode, o.Payload)
}
func (o *GetFirewallPoliciesDefault) GetPayload() *models.ResponsesFirewallPoliciesV1 {
	return o.Payload
}

func (o *GetFirewallPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesFirewallPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
