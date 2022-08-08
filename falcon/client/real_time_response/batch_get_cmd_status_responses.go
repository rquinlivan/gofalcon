// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

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

// BatchGetCmdStatusReader is a Reader for the BatchGetCmdStatus structure.
type BatchGetCmdStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BatchGetCmdStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBatchGetCmdStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBatchGetCmdStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBatchGetCmdStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBatchGetCmdStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewBatchGetCmdStatusTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBatchGetCmdStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBatchGetCmdStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBatchGetCmdStatusOK creates a BatchGetCmdStatusOK with default headers values
func NewBatchGetCmdStatusOK() *BatchGetCmdStatusOK {
	return &BatchGetCmdStatusOK{}
}

/*
BatchGetCmdStatusOK describes a response with status code 200, with default header values.

OK
*/
type BatchGetCmdStatusOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainBatchGetCmdStatusResponse
}

func (o *BatchGetCmdStatusOK) Error() string {
	return fmt.Sprintf("[GET /real-time-response/combined/batch-get-command/v1][%d] batchGetCmdStatusOK  %+v", 200, o.Payload)
}
func (o *BatchGetCmdStatusOK) GetPayload() *models.DomainBatchGetCmdStatusResponse {
	return o.Payload
}

func (o *BatchGetCmdStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainBatchGetCmdStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBatchGetCmdStatusBadRequest creates a BatchGetCmdStatusBadRequest with default headers values
func NewBatchGetCmdStatusBadRequest() *BatchGetCmdStatusBadRequest {
	return &BatchGetCmdStatusBadRequest{}
}

/*
BatchGetCmdStatusBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BatchGetCmdStatusBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

func (o *BatchGetCmdStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /real-time-response/combined/batch-get-command/v1][%d] batchGetCmdStatusBadRequest  %+v", 400, o.Payload)
}
func (o *BatchGetCmdStatusBadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *BatchGetCmdStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBatchGetCmdStatusForbidden creates a BatchGetCmdStatusForbidden with default headers values
func NewBatchGetCmdStatusForbidden() *BatchGetCmdStatusForbidden {
	return &BatchGetCmdStatusForbidden{}
}

/*
BatchGetCmdStatusForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BatchGetCmdStatusForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *BatchGetCmdStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /real-time-response/combined/batch-get-command/v1][%d] batchGetCmdStatusForbidden  %+v", 403, o.Payload)
}
func (o *BatchGetCmdStatusForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *BatchGetCmdStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewBatchGetCmdStatusNotFound creates a BatchGetCmdStatusNotFound with default headers values
func NewBatchGetCmdStatusNotFound() *BatchGetCmdStatusNotFound {
	return &BatchGetCmdStatusNotFound{}
}

/*
BatchGetCmdStatusNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BatchGetCmdStatusNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

func (o *BatchGetCmdStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /real-time-response/combined/batch-get-command/v1][%d] batchGetCmdStatusNotFound  %+v", 404, o.Payload)
}
func (o *BatchGetCmdStatusNotFound) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *BatchGetCmdStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBatchGetCmdStatusTooManyRequests creates a BatchGetCmdStatusTooManyRequests with default headers values
func NewBatchGetCmdStatusTooManyRequests() *BatchGetCmdStatusTooManyRequests {
	return &BatchGetCmdStatusTooManyRequests{}
}

/*
BatchGetCmdStatusTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type BatchGetCmdStatusTooManyRequests struct {

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

func (o *BatchGetCmdStatusTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /real-time-response/combined/batch-get-command/v1][%d] batchGetCmdStatusTooManyRequests  %+v", 429, o.Payload)
}
func (o *BatchGetCmdStatusTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *BatchGetCmdStatusTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewBatchGetCmdStatusInternalServerError creates a BatchGetCmdStatusInternalServerError with default headers values
func NewBatchGetCmdStatusInternalServerError() *BatchGetCmdStatusInternalServerError {
	return &BatchGetCmdStatusInternalServerError{}
}

/*
BatchGetCmdStatusInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type BatchGetCmdStatusInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

func (o *BatchGetCmdStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /real-time-response/combined/batch-get-command/v1][%d] batchGetCmdStatusInternalServerError  %+v", 500, o.Payload)
}
func (o *BatchGetCmdStatusInternalServerError) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *BatchGetCmdStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBatchGetCmdStatusDefault creates a BatchGetCmdStatusDefault with default headers values
func NewBatchGetCmdStatusDefault(code int) *BatchGetCmdStatusDefault {
	return &BatchGetCmdStatusDefault{
		_statusCode: code,
	}
}

/*
BatchGetCmdStatusDefault describes a response with status code -1, with default header values.

OK
*/
type BatchGetCmdStatusDefault struct {
	_statusCode int

	Payload *models.DomainBatchGetCmdStatusResponse
}

// Code gets the status code for the batch get cmd status default response
func (o *BatchGetCmdStatusDefault) Code() int {
	return o._statusCode
}

func (o *BatchGetCmdStatusDefault) Error() string {
	return fmt.Sprintf("[GET /real-time-response/combined/batch-get-command/v1][%d] BatchGetCmdStatus default  %+v", o._statusCode, o.Payload)
}
func (o *BatchGetCmdStatusDefault) GetPayload() *models.DomainBatchGetCmdStatusResponse {
	return o.Payload
}

func (o *BatchGetCmdStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainBatchGetCmdStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
