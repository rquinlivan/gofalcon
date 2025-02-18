// Code generated by go-swagger; DO NOT EDIT.

package ioa_exclusions

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

// GetIOAExclusionsV1Reader is a Reader for the GetIOAExclusionsV1 structure.
type GetIOAExclusionsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIOAExclusionsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIOAExclusionsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIOAExclusionsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIOAExclusionsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIOAExclusionsV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIOAExclusionsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIOAExclusionsV1OK creates a GetIOAExclusionsV1OK with default headers values
func NewGetIOAExclusionsV1OK() *GetIOAExclusionsV1OK {
	return &GetIOAExclusionsV1OK{}
}

/*
GetIOAExclusionsV1OK describes a response with status code 200, with default header values.

OK
*/
type GetIOAExclusionsV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesIoaExclusionRespV1
}

// IsSuccess returns true when this get i o a exclusions v1 o k response has a 2xx status code
func (o *GetIOAExclusionsV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get i o a exclusions v1 o k response has a 3xx status code
func (o *GetIOAExclusionsV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get i o a exclusions v1 o k response has a 4xx status code
func (o *GetIOAExclusionsV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get i o a exclusions v1 o k response has a 5xx status code
func (o *GetIOAExclusionsV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get i o a exclusions v1 o k response a status code equal to that given
func (o *GetIOAExclusionsV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get i o a exclusions v1 o k response
func (o *GetIOAExclusionsV1OK) Code() int {
	return 200
}

func (o *GetIOAExclusionsV1OK) Error() string {
	return fmt.Sprintf("[GET /policy/entities/ioa-exclusions/v1][%d] getIOAExclusionsV1OK  %+v", 200, o.Payload)
}

func (o *GetIOAExclusionsV1OK) String() string {
	return fmt.Sprintf("[GET /policy/entities/ioa-exclusions/v1][%d] getIOAExclusionsV1OK  %+v", 200, o.Payload)
}

func (o *GetIOAExclusionsV1OK) GetPayload() *models.ResponsesIoaExclusionRespV1 {
	return o.Payload
}

func (o *GetIOAExclusionsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesIoaExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIOAExclusionsV1BadRequest creates a GetIOAExclusionsV1BadRequest with default headers values
func NewGetIOAExclusionsV1BadRequest() *GetIOAExclusionsV1BadRequest {
	return &GetIOAExclusionsV1BadRequest{}
}

/*
GetIOAExclusionsV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetIOAExclusionsV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesIoaExclusionRespV1
}

// IsSuccess returns true when this get i o a exclusions v1 bad request response has a 2xx status code
func (o *GetIOAExclusionsV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get i o a exclusions v1 bad request response has a 3xx status code
func (o *GetIOAExclusionsV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get i o a exclusions v1 bad request response has a 4xx status code
func (o *GetIOAExclusionsV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get i o a exclusions v1 bad request response has a 5xx status code
func (o *GetIOAExclusionsV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get i o a exclusions v1 bad request response a status code equal to that given
func (o *GetIOAExclusionsV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get i o a exclusions v1 bad request response
func (o *GetIOAExclusionsV1BadRequest) Code() int {
	return 400
}

func (o *GetIOAExclusionsV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/entities/ioa-exclusions/v1][%d] getIOAExclusionsV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetIOAExclusionsV1BadRequest) String() string {
	return fmt.Sprintf("[GET /policy/entities/ioa-exclusions/v1][%d] getIOAExclusionsV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetIOAExclusionsV1BadRequest) GetPayload() *models.ResponsesIoaExclusionRespV1 {
	return o.Payload
}

func (o *GetIOAExclusionsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesIoaExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIOAExclusionsV1Forbidden creates a GetIOAExclusionsV1Forbidden with default headers values
func NewGetIOAExclusionsV1Forbidden() *GetIOAExclusionsV1Forbidden {
	return &GetIOAExclusionsV1Forbidden{}
}

/*
GetIOAExclusionsV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetIOAExclusionsV1Forbidden struct {

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

// IsSuccess returns true when this get i o a exclusions v1 forbidden response has a 2xx status code
func (o *GetIOAExclusionsV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get i o a exclusions v1 forbidden response has a 3xx status code
func (o *GetIOAExclusionsV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get i o a exclusions v1 forbidden response has a 4xx status code
func (o *GetIOAExclusionsV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get i o a exclusions v1 forbidden response has a 5xx status code
func (o *GetIOAExclusionsV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get i o a exclusions v1 forbidden response a status code equal to that given
func (o *GetIOAExclusionsV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get i o a exclusions v1 forbidden response
func (o *GetIOAExclusionsV1Forbidden) Code() int {
	return 403
}

func (o *GetIOAExclusionsV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /policy/entities/ioa-exclusions/v1][%d] getIOAExclusionsV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetIOAExclusionsV1Forbidden) String() string {
	return fmt.Sprintf("[GET /policy/entities/ioa-exclusions/v1][%d] getIOAExclusionsV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetIOAExclusionsV1Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetIOAExclusionsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIOAExclusionsV1TooManyRequests creates a GetIOAExclusionsV1TooManyRequests with default headers values
func NewGetIOAExclusionsV1TooManyRequests() *GetIOAExclusionsV1TooManyRequests {
	return &GetIOAExclusionsV1TooManyRequests{}
}

/*
GetIOAExclusionsV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetIOAExclusionsV1TooManyRequests struct {

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

// IsSuccess returns true when this get i o a exclusions v1 too many requests response has a 2xx status code
func (o *GetIOAExclusionsV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get i o a exclusions v1 too many requests response has a 3xx status code
func (o *GetIOAExclusionsV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get i o a exclusions v1 too many requests response has a 4xx status code
func (o *GetIOAExclusionsV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get i o a exclusions v1 too many requests response has a 5xx status code
func (o *GetIOAExclusionsV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get i o a exclusions v1 too many requests response a status code equal to that given
func (o *GetIOAExclusionsV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get i o a exclusions v1 too many requests response
func (o *GetIOAExclusionsV1TooManyRequests) Code() int {
	return 429
}

func (o *GetIOAExclusionsV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/entities/ioa-exclusions/v1][%d] getIOAExclusionsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIOAExclusionsV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /policy/entities/ioa-exclusions/v1][%d] getIOAExclusionsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIOAExclusionsV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetIOAExclusionsV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIOAExclusionsV1InternalServerError creates a GetIOAExclusionsV1InternalServerError with default headers values
func NewGetIOAExclusionsV1InternalServerError() *GetIOAExclusionsV1InternalServerError {
	return &GetIOAExclusionsV1InternalServerError{}
}

/*
GetIOAExclusionsV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetIOAExclusionsV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesIoaExclusionRespV1
}

// IsSuccess returns true when this get i o a exclusions v1 internal server error response has a 2xx status code
func (o *GetIOAExclusionsV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get i o a exclusions v1 internal server error response has a 3xx status code
func (o *GetIOAExclusionsV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get i o a exclusions v1 internal server error response has a 4xx status code
func (o *GetIOAExclusionsV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get i o a exclusions v1 internal server error response has a 5xx status code
func (o *GetIOAExclusionsV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get i o a exclusions v1 internal server error response a status code equal to that given
func (o *GetIOAExclusionsV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get i o a exclusions v1 internal server error response
func (o *GetIOAExclusionsV1InternalServerError) Code() int {
	return 500
}

func (o *GetIOAExclusionsV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/entities/ioa-exclusions/v1][%d] getIOAExclusionsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetIOAExclusionsV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /policy/entities/ioa-exclusions/v1][%d] getIOAExclusionsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetIOAExclusionsV1InternalServerError) GetPayload() *models.ResponsesIoaExclusionRespV1 {
	return o.Payload
}

func (o *GetIOAExclusionsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesIoaExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
