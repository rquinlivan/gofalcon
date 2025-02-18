// Code generated by go-swagger; DO NOT EDIT.

package zero_trust_assessment

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

// GetAssessmentV1Reader is a Reader for the GetAssessmentV1 structure.
type GetAssessmentV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAssessmentV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAssessmentV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAssessmentV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAssessmentV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAssessmentV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAssessmentV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAssessmentV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAssessmentV1OK creates a GetAssessmentV1OK with default headers values
func NewGetAssessmentV1OK() *GetAssessmentV1OK {
	return &GetAssessmentV1OK{}
}

/*
GetAssessmentV1OK describes a response with status code 200, with default header values.

OK
*/
type GetAssessmentV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAssessmentsResponse
}

// IsSuccess returns true when this get assessment v1 o k response has a 2xx status code
func (o *GetAssessmentV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get assessment v1 o k response has a 3xx status code
func (o *GetAssessmentV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assessment v1 o k response has a 4xx status code
func (o *GetAssessmentV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get assessment v1 o k response has a 5xx status code
func (o *GetAssessmentV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get assessment v1 o k response a status code equal to that given
func (o *GetAssessmentV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get assessment v1 o k response
func (o *GetAssessmentV1OK) Code() int {
	return 200
}

func (o *GetAssessmentV1OK) Error() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/entities/assessments/v1][%d] getAssessmentV1OK  %+v", 200, o.Payload)
}

func (o *GetAssessmentV1OK) String() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/entities/assessments/v1][%d] getAssessmentV1OK  %+v", 200, o.Payload)
}

func (o *GetAssessmentV1OK) GetPayload() *models.DomainAssessmentsResponse {
	return o.Payload
}

func (o *GetAssessmentV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAssessmentsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssessmentV1BadRequest creates a GetAssessmentV1BadRequest with default headers values
func NewGetAssessmentV1BadRequest() *GetAssessmentV1BadRequest {
	return &GetAssessmentV1BadRequest{}
}

/*
GetAssessmentV1BadRequest describes a response with status code 400, with default header values.

Number of agent IDs exceeds the limit of 1000.
*/
type GetAssessmentV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAssessmentsResponse
}

// IsSuccess returns true when this get assessment v1 bad request response has a 2xx status code
func (o *GetAssessmentV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get assessment v1 bad request response has a 3xx status code
func (o *GetAssessmentV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assessment v1 bad request response has a 4xx status code
func (o *GetAssessmentV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get assessment v1 bad request response has a 5xx status code
func (o *GetAssessmentV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get assessment v1 bad request response a status code equal to that given
func (o *GetAssessmentV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get assessment v1 bad request response
func (o *GetAssessmentV1BadRequest) Code() int {
	return 400
}

func (o *GetAssessmentV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/entities/assessments/v1][%d] getAssessmentV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetAssessmentV1BadRequest) String() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/entities/assessments/v1][%d] getAssessmentV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetAssessmentV1BadRequest) GetPayload() *models.DomainAssessmentsResponse {
	return o.Payload
}

func (o *GetAssessmentV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAssessmentsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssessmentV1Forbidden creates a GetAssessmentV1Forbidden with default headers values
func NewGetAssessmentV1Forbidden() *GetAssessmentV1Forbidden {
	return &GetAssessmentV1Forbidden{}
}

/*
GetAssessmentV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAssessmentV1Forbidden struct {

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

// IsSuccess returns true when this get assessment v1 forbidden response has a 2xx status code
func (o *GetAssessmentV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get assessment v1 forbidden response has a 3xx status code
func (o *GetAssessmentV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assessment v1 forbidden response has a 4xx status code
func (o *GetAssessmentV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get assessment v1 forbidden response has a 5xx status code
func (o *GetAssessmentV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get assessment v1 forbidden response a status code equal to that given
func (o *GetAssessmentV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get assessment v1 forbidden response
func (o *GetAssessmentV1Forbidden) Code() int {
	return 403
}

func (o *GetAssessmentV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/entities/assessments/v1][%d] getAssessmentV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetAssessmentV1Forbidden) String() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/entities/assessments/v1][%d] getAssessmentV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetAssessmentV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetAssessmentV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAssessmentV1NotFound creates a GetAssessmentV1NotFound with default headers values
func NewGetAssessmentV1NotFound() *GetAssessmentV1NotFound {
	return &GetAssessmentV1NotFound{}
}

/*
GetAssessmentV1NotFound describes a response with status code 404, with default header values.

One or more of the specified agent IDs is not found.
*/
type GetAssessmentV1NotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAssessmentsResponse
}

// IsSuccess returns true when this get assessment v1 not found response has a 2xx status code
func (o *GetAssessmentV1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get assessment v1 not found response has a 3xx status code
func (o *GetAssessmentV1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assessment v1 not found response has a 4xx status code
func (o *GetAssessmentV1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get assessment v1 not found response has a 5xx status code
func (o *GetAssessmentV1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get assessment v1 not found response a status code equal to that given
func (o *GetAssessmentV1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get assessment v1 not found response
func (o *GetAssessmentV1NotFound) Code() int {
	return 404
}

func (o *GetAssessmentV1NotFound) Error() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/entities/assessments/v1][%d] getAssessmentV1NotFound  %+v", 404, o.Payload)
}

func (o *GetAssessmentV1NotFound) String() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/entities/assessments/v1][%d] getAssessmentV1NotFound  %+v", 404, o.Payload)
}

func (o *GetAssessmentV1NotFound) GetPayload() *models.DomainAssessmentsResponse {
	return o.Payload
}

func (o *GetAssessmentV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAssessmentsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssessmentV1TooManyRequests creates a GetAssessmentV1TooManyRequests with default headers values
func NewGetAssessmentV1TooManyRequests() *GetAssessmentV1TooManyRequests {
	return &GetAssessmentV1TooManyRequests{}
}

/*
GetAssessmentV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAssessmentV1TooManyRequests struct {

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

// IsSuccess returns true when this get assessment v1 too many requests response has a 2xx status code
func (o *GetAssessmentV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get assessment v1 too many requests response has a 3xx status code
func (o *GetAssessmentV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get assessment v1 too many requests response has a 4xx status code
func (o *GetAssessmentV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get assessment v1 too many requests response has a 5xx status code
func (o *GetAssessmentV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get assessment v1 too many requests response a status code equal to that given
func (o *GetAssessmentV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get assessment v1 too many requests response
func (o *GetAssessmentV1TooManyRequests) Code() int {
	return 429
}

func (o *GetAssessmentV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/entities/assessments/v1][%d] getAssessmentV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAssessmentV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/entities/assessments/v1][%d] getAssessmentV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAssessmentV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetAssessmentV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAssessmentV1Default creates a GetAssessmentV1Default with default headers values
func NewGetAssessmentV1Default(code int) *GetAssessmentV1Default {
	return &GetAssessmentV1Default{
		_statusCode: code,
	}
}

/*
GetAssessmentV1Default describes a response with status code -1, with default header values.

OK
*/
type GetAssessmentV1Default struct {
	_statusCode int

	Payload *models.DomainAssessmentsResponse
}

// IsSuccess returns true when this get assessment v1 default response has a 2xx status code
func (o *GetAssessmentV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get assessment v1 default response has a 3xx status code
func (o *GetAssessmentV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get assessment v1 default response has a 4xx status code
func (o *GetAssessmentV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get assessment v1 default response has a 5xx status code
func (o *GetAssessmentV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get assessment v1 default response a status code equal to that given
func (o *GetAssessmentV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get assessment v1 default response
func (o *GetAssessmentV1Default) Code() int {
	return o._statusCode
}

func (o *GetAssessmentV1Default) Error() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/entities/assessments/v1][%d] getAssessmentV1 default  %+v", o._statusCode, o.Payload)
}

func (o *GetAssessmentV1Default) String() string {
	return fmt.Sprintf("[GET /zero-trust-assessment/entities/assessments/v1][%d] getAssessmentV1 default  %+v", o._statusCode, o.Payload)
}

func (o *GetAssessmentV1Default) GetPayload() *models.DomainAssessmentsResponse {
	return o.Payload
}

func (o *GetAssessmentV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainAssessmentsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
