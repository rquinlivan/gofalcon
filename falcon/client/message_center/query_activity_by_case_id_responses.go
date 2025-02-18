// Code generated by go-swagger; DO NOT EDIT.

package message_center

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

// QueryActivityByCaseIDReader is a Reader for the QueryActivityByCaseID structure.
type QueryActivityByCaseIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryActivityByCaseIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryActivityByCaseIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryActivityByCaseIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryActivityByCaseIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryActivityByCaseIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryActivityByCaseIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryActivityByCaseIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryActivityByCaseIDOK creates a QueryActivityByCaseIDOK with default headers values
func NewQueryActivityByCaseIDOK() *QueryActivityByCaseIDOK {
	return &QueryActivityByCaseIDOK{}
}

/*
QueryActivityByCaseIDOK describes a response with status code 200, with default header values.

OK
*/
type QueryActivityByCaseIDOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query activity by case Id o k response has a 2xx status code
func (o *QueryActivityByCaseIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query activity by case Id o k response has a 3xx status code
func (o *QueryActivityByCaseIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query activity by case Id o k response has a 4xx status code
func (o *QueryActivityByCaseIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query activity by case Id o k response has a 5xx status code
func (o *QueryActivityByCaseIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query activity by case Id o k response a status code equal to that given
func (o *QueryActivityByCaseIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query activity by case Id o k response
func (o *QueryActivityByCaseIDOK) Code() int {
	return 200
}

func (o *QueryActivityByCaseIDOK) Error() string {
	return fmt.Sprintf("[GET /message-center/queries/case-activities/v1][%d] queryActivityByCaseIdOK  %+v", 200, o.Payload)
}

func (o *QueryActivityByCaseIDOK) String() string {
	return fmt.Sprintf("[GET /message-center/queries/case-activities/v1][%d] queryActivityByCaseIdOK  %+v", 200, o.Payload)
}

func (o *QueryActivityByCaseIDOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryActivityByCaseIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryActivityByCaseIDBadRequest creates a QueryActivityByCaseIDBadRequest with default headers values
func NewQueryActivityByCaseIDBadRequest() *QueryActivityByCaseIDBadRequest {
	return &QueryActivityByCaseIDBadRequest{}
}

/*
QueryActivityByCaseIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryActivityByCaseIDBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this query activity by case Id bad request response has a 2xx status code
func (o *QueryActivityByCaseIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query activity by case Id bad request response has a 3xx status code
func (o *QueryActivityByCaseIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query activity by case Id bad request response has a 4xx status code
func (o *QueryActivityByCaseIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query activity by case Id bad request response has a 5xx status code
func (o *QueryActivityByCaseIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query activity by case Id bad request response a status code equal to that given
func (o *QueryActivityByCaseIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query activity by case Id bad request response
func (o *QueryActivityByCaseIDBadRequest) Code() int {
	return 400
}

func (o *QueryActivityByCaseIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /message-center/queries/case-activities/v1][%d] queryActivityByCaseIdBadRequest  %+v", 400, o.Payload)
}

func (o *QueryActivityByCaseIDBadRequest) String() string {
	return fmt.Sprintf("[GET /message-center/queries/case-activities/v1][%d] queryActivityByCaseIdBadRequest  %+v", 400, o.Payload)
}

func (o *QueryActivityByCaseIDBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryActivityByCaseIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryActivityByCaseIDForbidden creates a QueryActivityByCaseIDForbidden with default headers values
func NewQueryActivityByCaseIDForbidden() *QueryActivityByCaseIDForbidden {
	return &QueryActivityByCaseIDForbidden{}
}

/*
QueryActivityByCaseIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryActivityByCaseIDForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this query activity by case Id forbidden response has a 2xx status code
func (o *QueryActivityByCaseIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query activity by case Id forbidden response has a 3xx status code
func (o *QueryActivityByCaseIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query activity by case Id forbidden response has a 4xx status code
func (o *QueryActivityByCaseIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query activity by case Id forbidden response has a 5xx status code
func (o *QueryActivityByCaseIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query activity by case Id forbidden response a status code equal to that given
func (o *QueryActivityByCaseIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query activity by case Id forbidden response
func (o *QueryActivityByCaseIDForbidden) Code() int {
	return 403
}

func (o *QueryActivityByCaseIDForbidden) Error() string {
	return fmt.Sprintf("[GET /message-center/queries/case-activities/v1][%d] queryActivityByCaseIdForbidden  %+v", 403, o.Payload)
}

func (o *QueryActivityByCaseIDForbidden) String() string {
	return fmt.Sprintf("[GET /message-center/queries/case-activities/v1][%d] queryActivityByCaseIdForbidden  %+v", 403, o.Payload)
}

func (o *QueryActivityByCaseIDForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryActivityByCaseIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryActivityByCaseIDTooManyRequests creates a QueryActivityByCaseIDTooManyRequests with default headers values
func NewQueryActivityByCaseIDTooManyRequests() *QueryActivityByCaseIDTooManyRequests {
	return &QueryActivityByCaseIDTooManyRequests{}
}

/*
QueryActivityByCaseIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryActivityByCaseIDTooManyRequests struct {

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

// IsSuccess returns true when this query activity by case Id too many requests response has a 2xx status code
func (o *QueryActivityByCaseIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query activity by case Id too many requests response has a 3xx status code
func (o *QueryActivityByCaseIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query activity by case Id too many requests response has a 4xx status code
func (o *QueryActivityByCaseIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query activity by case Id too many requests response has a 5xx status code
func (o *QueryActivityByCaseIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query activity by case Id too many requests response a status code equal to that given
func (o *QueryActivityByCaseIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query activity by case Id too many requests response
func (o *QueryActivityByCaseIDTooManyRequests) Code() int {
	return 429
}

func (o *QueryActivityByCaseIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /message-center/queries/case-activities/v1][%d] queryActivityByCaseIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryActivityByCaseIDTooManyRequests) String() string {
	return fmt.Sprintf("[GET /message-center/queries/case-activities/v1][%d] queryActivityByCaseIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryActivityByCaseIDTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryActivityByCaseIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryActivityByCaseIDInternalServerError creates a QueryActivityByCaseIDInternalServerError with default headers values
func NewQueryActivityByCaseIDInternalServerError() *QueryActivityByCaseIDInternalServerError {
	return &QueryActivityByCaseIDInternalServerError{}
}

/*
QueryActivityByCaseIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryActivityByCaseIDInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this query activity by case Id internal server error response has a 2xx status code
func (o *QueryActivityByCaseIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query activity by case Id internal server error response has a 3xx status code
func (o *QueryActivityByCaseIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query activity by case Id internal server error response has a 4xx status code
func (o *QueryActivityByCaseIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query activity by case Id internal server error response has a 5xx status code
func (o *QueryActivityByCaseIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query activity by case Id internal server error response a status code equal to that given
func (o *QueryActivityByCaseIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query activity by case Id internal server error response
func (o *QueryActivityByCaseIDInternalServerError) Code() int {
	return 500
}

func (o *QueryActivityByCaseIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /message-center/queries/case-activities/v1][%d] queryActivityByCaseIdInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryActivityByCaseIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /message-center/queries/case-activities/v1][%d] queryActivityByCaseIdInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryActivityByCaseIDInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryActivityByCaseIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryActivityByCaseIDDefault creates a QueryActivityByCaseIDDefault with default headers values
func NewQueryActivityByCaseIDDefault(code int) *QueryActivityByCaseIDDefault {
	return &QueryActivityByCaseIDDefault{
		_statusCode: code,
	}
}

/*
QueryActivityByCaseIDDefault describes a response with status code -1, with default header values.

OK
*/
type QueryActivityByCaseIDDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query activity by case ID default response has a 2xx status code
func (o *QueryActivityByCaseIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query activity by case ID default response has a 3xx status code
func (o *QueryActivityByCaseIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query activity by case ID default response has a 4xx status code
func (o *QueryActivityByCaseIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query activity by case ID default response has a 5xx status code
func (o *QueryActivityByCaseIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query activity by case ID default response a status code equal to that given
func (o *QueryActivityByCaseIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the query activity by case ID default response
func (o *QueryActivityByCaseIDDefault) Code() int {
	return o._statusCode
}

func (o *QueryActivityByCaseIDDefault) Error() string {
	return fmt.Sprintf("[GET /message-center/queries/case-activities/v1][%d] QueryActivityByCaseID default  %+v", o._statusCode, o.Payload)
}

func (o *QueryActivityByCaseIDDefault) String() string {
	return fmt.Sprintf("[GET /message-center/queries/case-activities/v1][%d] QueryActivityByCaseID default  %+v", o._statusCode, o.Payload)
}

func (o *QueryActivityByCaseIDDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryActivityByCaseIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
