// Code generated by go-swagger; DO NOT EDIT.

package sensor_update_policies

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

// GetSensorUpdatePoliciesReader is a Reader for the GetSensorUpdatePolicies structure.
type GetSensorUpdatePoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSensorUpdatePoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSensorUpdatePoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetSensorUpdatePoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSensorUpdatePoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSensorUpdatePoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSensorUpdatePoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSensorUpdatePoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSensorUpdatePoliciesOK creates a GetSensorUpdatePoliciesOK with default headers values
func NewGetSensorUpdatePoliciesOK() *GetSensorUpdatePoliciesOK {
	return &GetSensorUpdatePoliciesOK{}
}

/*
GetSensorUpdatePoliciesOK describes a response with status code 200, with default header values.

OK
*/
type GetSensorUpdatePoliciesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV1
}

func (o *GetSensorUpdatePoliciesOK) Error() string {
	return fmt.Sprintf("[GET /policy/entities/sensor-update/v1][%d] getSensorUpdatePoliciesOK  %+v", 200, o.Payload)
}
func (o *GetSensorUpdatePoliciesOK) GetPayload() *models.ResponsesSensorUpdatePoliciesV1 {
	return o.Payload
}

func (o *GetSensorUpdatePoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSensorUpdatePoliciesForbidden creates a GetSensorUpdatePoliciesForbidden with default headers values
func NewGetSensorUpdatePoliciesForbidden() *GetSensorUpdatePoliciesForbidden {
	return &GetSensorUpdatePoliciesForbidden{}
}

/*
GetSensorUpdatePoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSensorUpdatePoliciesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *GetSensorUpdatePoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/entities/sensor-update/v1][%d] getSensorUpdatePoliciesForbidden  %+v", 403, o.Payload)
}
func (o *GetSensorUpdatePoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetSensorUpdatePoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSensorUpdatePoliciesNotFound creates a GetSensorUpdatePoliciesNotFound with default headers values
func NewGetSensorUpdatePoliciesNotFound() *GetSensorUpdatePoliciesNotFound {
	return &GetSensorUpdatePoliciesNotFound{}
}

/*
GetSensorUpdatePoliciesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetSensorUpdatePoliciesNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV1
}

func (o *GetSensorUpdatePoliciesNotFound) Error() string {
	return fmt.Sprintf("[GET /policy/entities/sensor-update/v1][%d] getSensorUpdatePoliciesNotFound  %+v", 404, o.Payload)
}
func (o *GetSensorUpdatePoliciesNotFound) GetPayload() *models.ResponsesSensorUpdatePoliciesV1 {
	return o.Payload
}

func (o *GetSensorUpdatePoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSensorUpdatePoliciesTooManyRequests creates a GetSensorUpdatePoliciesTooManyRequests with default headers values
func NewGetSensorUpdatePoliciesTooManyRequests() *GetSensorUpdatePoliciesTooManyRequests {
	return &GetSensorUpdatePoliciesTooManyRequests{}
}

/*
GetSensorUpdatePoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetSensorUpdatePoliciesTooManyRequests struct {

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

func (o *GetSensorUpdatePoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/entities/sensor-update/v1][%d] getSensorUpdatePoliciesTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetSensorUpdatePoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSensorUpdatePoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSensorUpdatePoliciesInternalServerError creates a GetSensorUpdatePoliciesInternalServerError with default headers values
func NewGetSensorUpdatePoliciesInternalServerError() *GetSensorUpdatePoliciesInternalServerError {
	return &GetSensorUpdatePoliciesInternalServerError{}
}

/*
GetSensorUpdatePoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSensorUpdatePoliciesInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV1
}

func (o *GetSensorUpdatePoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/entities/sensor-update/v1][%d] getSensorUpdatePoliciesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSensorUpdatePoliciesInternalServerError) GetPayload() *models.ResponsesSensorUpdatePoliciesV1 {
	return o.Payload
}

func (o *GetSensorUpdatePoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSensorUpdatePoliciesDefault creates a GetSensorUpdatePoliciesDefault with default headers values
func NewGetSensorUpdatePoliciesDefault(code int) *GetSensorUpdatePoliciesDefault {
	return &GetSensorUpdatePoliciesDefault{
		_statusCode: code,
	}
}

/*
GetSensorUpdatePoliciesDefault describes a response with status code -1, with default header values.

OK
*/
type GetSensorUpdatePoliciesDefault struct {
	_statusCode int

	Payload *models.ResponsesSensorUpdatePoliciesV1
}

// Code gets the status code for the get sensor update policies default response
func (o *GetSensorUpdatePoliciesDefault) Code() int {
	return o._statusCode
}

func (o *GetSensorUpdatePoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /policy/entities/sensor-update/v1][%d] getSensorUpdatePolicies default  %+v", o._statusCode, o.Payload)
}
func (o *GetSensorUpdatePoliciesDefault) GetPayload() *models.ResponsesSensorUpdatePoliciesV1 {
	return o.Payload
}

func (o *GetSensorUpdatePoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
