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

// BatchGetCmdReader is a Reader for the BatchGetCmd structure.
type BatchGetCmdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BatchGetCmdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewBatchGetCmdCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBatchGetCmdBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBatchGetCmdForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewBatchGetCmdTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBatchGetCmdInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBatchGetCmdCreated creates a BatchGetCmdCreated with default headers values
func NewBatchGetCmdCreated() *BatchGetCmdCreated {
	return &BatchGetCmdCreated{}
}

/*
BatchGetCmdCreated describes a response with status code 201, with default header values.

Created
*/
type BatchGetCmdCreated struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainBatchGetCommandResponse
}

// IsSuccess returns true when this batch get cmd created response has a 2xx status code
func (o *BatchGetCmdCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this batch get cmd created response has a 3xx status code
func (o *BatchGetCmdCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this batch get cmd created response has a 4xx status code
func (o *BatchGetCmdCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this batch get cmd created response has a 5xx status code
func (o *BatchGetCmdCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this batch get cmd created response a status code equal to that given
func (o *BatchGetCmdCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the batch get cmd created response
func (o *BatchGetCmdCreated) Code() int {
	return 201
}

func (o *BatchGetCmdCreated) Error() string {
	return fmt.Sprintf("[POST /real-time-response/combined/batch-get-command/v1][%d] batchGetCmdCreated  %+v", 201, o.Payload)
}

func (o *BatchGetCmdCreated) String() string {
	return fmt.Sprintf("[POST /real-time-response/combined/batch-get-command/v1][%d] batchGetCmdCreated  %+v", 201, o.Payload)
}

func (o *BatchGetCmdCreated) GetPayload() *models.DomainBatchGetCommandResponse {
	return o.Payload
}

func (o *BatchGetCmdCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.DomainBatchGetCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBatchGetCmdBadRequest creates a BatchGetCmdBadRequest with default headers values
func NewBatchGetCmdBadRequest() *BatchGetCmdBadRequest {
	return &BatchGetCmdBadRequest{}
}

/*
BatchGetCmdBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BatchGetCmdBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this batch get cmd bad request response has a 2xx status code
func (o *BatchGetCmdBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this batch get cmd bad request response has a 3xx status code
func (o *BatchGetCmdBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this batch get cmd bad request response has a 4xx status code
func (o *BatchGetCmdBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this batch get cmd bad request response has a 5xx status code
func (o *BatchGetCmdBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this batch get cmd bad request response a status code equal to that given
func (o *BatchGetCmdBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the batch get cmd bad request response
func (o *BatchGetCmdBadRequest) Code() int {
	return 400
}

func (o *BatchGetCmdBadRequest) Error() string {
	return fmt.Sprintf("[POST /real-time-response/combined/batch-get-command/v1][%d] batchGetCmdBadRequest  %+v", 400, o.Payload)
}

func (o *BatchGetCmdBadRequest) String() string {
	return fmt.Sprintf("[POST /real-time-response/combined/batch-get-command/v1][%d] batchGetCmdBadRequest  %+v", 400, o.Payload)
}

func (o *BatchGetCmdBadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *BatchGetCmdBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

// NewBatchGetCmdForbidden creates a BatchGetCmdForbidden with default headers values
func NewBatchGetCmdForbidden() *BatchGetCmdForbidden {
	return &BatchGetCmdForbidden{}
}

/*
BatchGetCmdForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BatchGetCmdForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this batch get cmd forbidden response has a 2xx status code
func (o *BatchGetCmdForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this batch get cmd forbidden response has a 3xx status code
func (o *BatchGetCmdForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this batch get cmd forbidden response has a 4xx status code
func (o *BatchGetCmdForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this batch get cmd forbidden response has a 5xx status code
func (o *BatchGetCmdForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this batch get cmd forbidden response a status code equal to that given
func (o *BatchGetCmdForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the batch get cmd forbidden response
func (o *BatchGetCmdForbidden) Code() int {
	return 403
}

func (o *BatchGetCmdForbidden) Error() string {
	return fmt.Sprintf("[POST /real-time-response/combined/batch-get-command/v1][%d] batchGetCmdForbidden  %+v", 403, o.Payload)
}

func (o *BatchGetCmdForbidden) String() string {
	return fmt.Sprintf("[POST /real-time-response/combined/batch-get-command/v1][%d] batchGetCmdForbidden  %+v", 403, o.Payload)
}

func (o *BatchGetCmdForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *BatchGetCmdForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

// NewBatchGetCmdTooManyRequests creates a BatchGetCmdTooManyRequests with default headers values
func NewBatchGetCmdTooManyRequests() *BatchGetCmdTooManyRequests {
	return &BatchGetCmdTooManyRequests{}
}

/*
BatchGetCmdTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type BatchGetCmdTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

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

// IsSuccess returns true when this batch get cmd too many requests response has a 2xx status code
func (o *BatchGetCmdTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this batch get cmd too many requests response has a 3xx status code
func (o *BatchGetCmdTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this batch get cmd too many requests response has a 4xx status code
func (o *BatchGetCmdTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this batch get cmd too many requests response has a 5xx status code
func (o *BatchGetCmdTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this batch get cmd too many requests response a status code equal to that given
func (o *BatchGetCmdTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the batch get cmd too many requests response
func (o *BatchGetCmdTooManyRequests) Code() int {
	return 429
}

func (o *BatchGetCmdTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /real-time-response/combined/batch-get-command/v1][%d] batchGetCmdTooManyRequests  %+v", 429, o.Payload)
}

func (o *BatchGetCmdTooManyRequests) String() string {
	return fmt.Sprintf("[POST /real-time-response/combined/batch-get-command/v1][%d] batchGetCmdTooManyRequests  %+v", 429, o.Payload)
}

func (o *BatchGetCmdTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *BatchGetCmdTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

// NewBatchGetCmdInternalServerError creates a BatchGetCmdInternalServerError with default headers values
func NewBatchGetCmdInternalServerError() *BatchGetCmdInternalServerError {
	return &BatchGetCmdInternalServerError{}
}

/*
BatchGetCmdInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type BatchGetCmdInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this batch get cmd internal server error response has a 2xx status code
func (o *BatchGetCmdInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this batch get cmd internal server error response has a 3xx status code
func (o *BatchGetCmdInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this batch get cmd internal server error response has a 4xx status code
func (o *BatchGetCmdInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this batch get cmd internal server error response has a 5xx status code
func (o *BatchGetCmdInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this batch get cmd internal server error response a status code equal to that given
func (o *BatchGetCmdInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the batch get cmd internal server error response
func (o *BatchGetCmdInternalServerError) Code() int {
	return 500
}

func (o *BatchGetCmdInternalServerError) Error() string {
	return fmt.Sprintf("[POST /real-time-response/combined/batch-get-command/v1][%d] batchGetCmdInternalServerError  %+v", 500, o.Payload)
}

func (o *BatchGetCmdInternalServerError) String() string {
	return fmt.Sprintf("[POST /real-time-response/combined/batch-get-command/v1][%d] batchGetCmdInternalServerError  %+v", 500, o.Payload)
}

func (o *BatchGetCmdInternalServerError) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *BatchGetCmdInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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
