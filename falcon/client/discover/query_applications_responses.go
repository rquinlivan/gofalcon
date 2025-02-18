// Code generated by go-swagger; DO NOT EDIT.

package discover

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

// QueryApplicationsReader is a Reader for the QueryApplications structure.
type QueryApplicationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryApplicationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryApplicationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryApplicationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryApplicationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryApplicationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryApplicationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryApplicationsOK creates a QueryApplicationsOK with default headers values
func NewQueryApplicationsOK() *QueryApplicationsOK {
	return &QueryApplicationsOK{}
}

/*
QueryApplicationsOK describes a response with status code 200, with default header values.

OK
*/
type QueryApplicationsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this query applications o k response has a 2xx status code
func (o *QueryApplicationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query applications o k response has a 3xx status code
func (o *QueryApplicationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query applications o k response has a 4xx status code
func (o *QueryApplicationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query applications o k response has a 5xx status code
func (o *QueryApplicationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query applications o k response a status code equal to that given
func (o *QueryApplicationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query applications o k response
func (o *QueryApplicationsOK) Code() int {
	return 200
}

func (o *QueryApplicationsOK) Error() string {
	return fmt.Sprintf("[GET /discover/queries/applications/v1][%d] queryApplicationsOK  %+v", 200, o.Payload)
}

func (o *QueryApplicationsOK) String() string {
	return fmt.Sprintf("[GET /discover/queries/applications/v1][%d] queryApplicationsOK  %+v", 200, o.Payload)
}

func (o *QueryApplicationsOK) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *QueryApplicationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryApplicationsBadRequest creates a QueryApplicationsBadRequest with default headers values
func NewQueryApplicationsBadRequest() *QueryApplicationsBadRequest {
	return &QueryApplicationsBadRequest{}
}

/*
QueryApplicationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryApplicationsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this query applications bad request response has a 2xx status code
func (o *QueryApplicationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query applications bad request response has a 3xx status code
func (o *QueryApplicationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query applications bad request response has a 4xx status code
func (o *QueryApplicationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query applications bad request response has a 5xx status code
func (o *QueryApplicationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query applications bad request response a status code equal to that given
func (o *QueryApplicationsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query applications bad request response
func (o *QueryApplicationsBadRequest) Code() int {
	return 400
}

func (o *QueryApplicationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /discover/queries/applications/v1][%d] queryApplicationsBadRequest  %+v", 400, o.Payload)
}

func (o *QueryApplicationsBadRequest) String() string {
	return fmt.Sprintf("[GET /discover/queries/applications/v1][%d] queryApplicationsBadRequest  %+v", 400, o.Payload)
}

func (o *QueryApplicationsBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *QueryApplicationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryApplicationsForbidden creates a QueryApplicationsForbidden with default headers values
func NewQueryApplicationsForbidden() *QueryApplicationsForbidden {
	return &QueryApplicationsForbidden{}
}

/*
QueryApplicationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryApplicationsForbidden struct {

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

// IsSuccess returns true when this query applications forbidden response has a 2xx status code
func (o *QueryApplicationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query applications forbidden response has a 3xx status code
func (o *QueryApplicationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query applications forbidden response has a 4xx status code
func (o *QueryApplicationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query applications forbidden response has a 5xx status code
func (o *QueryApplicationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query applications forbidden response a status code equal to that given
func (o *QueryApplicationsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query applications forbidden response
func (o *QueryApplicationsForbidden) Code() int {
	return 403
}

func (o *QueryApplicationsForbidden) Error() string {
	return fmt.Sprintf("[GET /discover/queries/applications/v1][%d] queryApplicationsForbidden  %+v", 403, o.Payload)
}

func (o *QueryApplicationsForbidden) String() string {
	return fmt.Sprintf("[GET /discover/queries/applications/v1][%d] queryApplicationsForbidden  %+v", 403, o.Payload)
}

func (o *QueryApplicationsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryApplicationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryApplicationsTooManyRequests creates a QueryApplicationsTooManyRequests with default headers values
func NewQueryApplicationsTooManyRequests() *QueryApplicationsTooManyRequests {
	return &QueryApplicationsTooManyRequests{}
}

/*
QueryApplicationsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryApplicationsTooManyRequests struct {

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

// IsSuccess returns true when this query applications too many requests response has a 2xx status code
func (o *QueryApplicationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query applications too many requests response has a 3xx status code
func (o *QueryApplicationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query applications too many requests response has a 4xx status code
func (o *QueryApplicationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query applications too many requests response has a 5xx status code
func (o *QueryApplicationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query applications too many requests response a status code equal to that given
func (o *QueryApplicationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query applications too many requests response
func (o *QueryApplicationsTooManyRequests) Code() int {
	return 429
}

func (o *QueryApplicationsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /discover/queries/applications/v1][%d] queryApplicationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryApplicationsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /discover/queries/applications/v1][%d] queryApplicationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryApplicationsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryApplicationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryApplicationsInternalServerError creates a QueryApplicationsInternalServerError with default headers values
func NewQueryApplicationsInternalServerError() *QueryApplicationsInternalServerError {
	return &QueryApplicationsInternalServerError{}
}

/*
QueryApplicationsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryApplicationsInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this query applications internal server error response has a 2xx status code
func (o *QueryApplicationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query applications internal server error response has a 3xx status code
func (o *QueryApplicationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query applications internal server error response has a 4xx status code
func (o *QueryApplicationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query applications internal server error response has a 5xx status code
func (o *QueryApplicationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query applications internal server error response a status code equal to that given
func (o *QueryApplicationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query applications internal server error response
func (o *QueryApplicationsInternalServerError) Code() int {
	return 500
}

func (o *QueryApplicationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /discover/queries/applications/v1][%d] queryApplicationsInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryApplicationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /discover/queries/applications/v1][%d] queryApplicationsInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryApplicationsInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *QueryApplicationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
