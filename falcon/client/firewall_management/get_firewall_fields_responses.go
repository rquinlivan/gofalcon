// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

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

// GetFirewallFieldsReader is a Reader for the GetFirewallFields structure.
type GetFirewallFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFirewallFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFirewallFieldsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetFirewallFieldsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFirewallFieldsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetFirewallFieldsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFirewallFieldsOK creates a GetFirewallFieldsOK with default headers values
func NewGetFirewallFieldsOK() *GetFirewallFieldsOK {
	return &GetFirewallFieldsOK{}
}

/*
GetFirewallFieldsOK describes a response with status code 200, with default header values.

OK
*/
type GetFirewallFieldsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrAPIFirewallFieldsResponse
}

func (o *GetFirewallFieldsOK) Error() string {
	return fmt.Sprintf("[GET /fwmgr/entities/firewall-fields/v1][%d] getFirewallFieldsOK  %+v", 200, o.Payload)
}
func (o *GetFirewallFieldsOK) GetPayload() *models.FwmgrAPIFirewallFieldsResponse {
	return o.Payload
}

func (o *GetFirewallFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FwmgrAPIFirewallFieldsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFirewallFieldsForbidden creates a GetFirewallFieldsForbidden with default headers values
func NewGetFirewallFieldsForbidden() *GetFirewallFieldsForbidden {
	return &GetFirewallFieldsForbidden{}
}

/*
GetFirewallFieldsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetFirewallFieldsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *GetFirewallFieldsForbidden) Error() string {
	return fmt.Sprintf("[GET /fwmgr/entities/firewall-fields/v1][%d] getFirewallFieldsForbidden  %+v", 403, o.Payload)
}
func (o *GetFirewallFieldsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetFirewallFieldsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFirewallFieldsTooManyRequests creates a GetFirewallFieldsTooManyRequests with default headers values
func NewGetFirewallFieldsTooManyRequests() *GetFirewallFieldsTooManyRequests {
	return &GetFirewallFieldsTooManyRequests{}
}

/*
GetFirewallFieldsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetFirewallFieldsTooManyRequests struct {

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

func (o *GetFirewallFieldsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /fwmgr/entities/firewall-fields/v1][%d] getFirewallFieldsTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetFirewallFieldsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetFirewallFieldsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFirewallFieldsDefault creates a GetFirewallFieldsDefault with default headers values
func NewGetFirewallFieldsDefault(code int) *GetFirewallFieldsDefault {
	return &GetFirewallFieldsDefault{
		_statusCode: code,
	}
}

/*
GetFirewallFieldsDefault describes a response with status code -1, with default header values.

OK
*/
type GetFirewallFieldsDefault struct {
	_statusCode int

	Payload *models.FwmgrAPIFirewallFieldsResponse
}

// Code gets the status code for the get firewall fields default response
func (o *GetFirewallFieldsDefault) Code() int {
	return o._statusCode
}

func (o *GetFirewallFieldsDefault) Error() string {
	return fmt.Sprintf("[GET /fwmgr/entities/firewall-fields/v1][%d] get-firewall-fields default  %+v", o._statusCode, o.Payload)
}
func (o *GetFirewallFieldsDefault) GetPayload() *models.FwmgrAPIFirewallFieldsResponse {
	return o.Payload
}

func (o *GetFirewallFieldsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FwmgrAPIFirewallFieldsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
