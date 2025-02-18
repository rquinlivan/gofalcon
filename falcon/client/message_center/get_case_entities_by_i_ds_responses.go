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

// GetCaseEntitiesByIDsReader is a Reader for the GetCaseEntitiesByIDs structure.
type GetCaseEntitiesByIDsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCaseEntitiesByIDsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCaseEntitiesByIDsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCaseEntitiesByIDsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCaseEntitiesByIDsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCaseEntitiesByIDsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCaseEntitiesByIDsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCaseEntitiesByIDsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCaseEntitiesByIDsOK creates a GetCaseEntitiesByIDsOK with default headers values
func NewGetCaseEntitiesByIDsOK() *GetCaseEntitiesByIDsOK {
	return &GetCaseEntitiesByIDsOK{}
}

/*
GetCaseEntitiesByIDsOK describes a response with status code 200, with default header values.

OK
*/
type GetCaseEntitiesByIDsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIMessageCenterCasesResponse
}

// IsSuccess returns true when this get case entities by i ds o k response has a 2xx status code
func (o *GetCaseEntitiesByIDsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get case entities by i ds o k response has a 3xx status code
func (o *GetCaseEntitiesByIDsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get case entities by i ds o k response has a 4xx status code
func (o *GetCaseEntitiesByIDsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get case entities by i ds o k response has a 5xx status code
func (o *GetCaseEntitiesByIDsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get case entities by i ds o k response a status code equal to that given
func (o *GetCaseEntitiesByIDsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get case entities by i ds o k response
func (o *GetCaseEntitiesByIDsOK) Code() int {
	return 200
}

func (o *GetCaseEntitiesByIDsOK) Error() string {
	return fmt.Sprintf("[POST /message-center/entities/cases/GET/v1][%d] getCaseEntitiesByIDsOK  %+v", 200, o.Payload)
}

func (o *GetCaseEntitiesByIDsOK) String() string {
	return fmt.Sprintf("[POST /message-center/entities/cases/GET/v1][%d] getCaseEntitiesByIDsOK  %+v", 200, o.Payload)
}

func (o *GetCaseEntitiesByIDsOK) GetPayload() *models.APIMessageCenterCasesResponse {
	return o.Payload
}

func (o *GetCaseEntitiesByIDsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIMessageCenterCasesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCaseEntitiesByIDsBadRequest creates a GetCaseEntitiesByIDsBadRequest with default headers values
func NewGetCaseEntitiesByIDsBadRequest() *GetCaseEntitiesByIDsBadRequest {
	return &GetCaseEntitiesByIDsBadRequest{}
}

/*
GetCaseEntitiesByIDsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCaseEntitiesByIDsBadRequest struct {

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

// IsSuccess returns true when this get case entities by i ds bad request response has a 2xx status code
func (o *GetCaseEntitiesByIDsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get case entities by i ds bad request response has a 3xx status code
func (o *GetCaseEntitiesByIDsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get case entities by i ds bad request response has a 4xx status code
func (o *GetCaseEntitiesByIDsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get case entities by i ds bad request response has a 5xx status code
func (o *GetCaseEntitiesByIDsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get case entities by i ds bad request response a status code equal to that given
func (o *GetCaseEntitiesByIDsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get case entities by i ds bad request response
func (o *GetCaseEntitiesByIDsBadRequest) Code() int {
	return 400
}

func (o *GetCaseEntitiesByIDsBadRequest) Error() string {
	return fmt.Sprintf("[POST /message-center/entities/cases/GET/v1][%d] getCaseEntitiesByIDsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCaseEntitiesByIDsBadRequest) String() string {
	return fmt.Sprintf("[POST /message-center/entities/cases/GET/v1][%d] getCaseEntitiesByIDsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCaseEntitiesByIDsBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCaseEntitiesByIDsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCaseEntitiesByIDsForbidden creates a GetCaseEntitiesByIDsForbidden with default headers values
func NewGetCaseEntitiesByIDsForbidden() *GetCaseEntitiesByIDsForbidden {
	return &GetCaseEntitiesByIDsForbidden{}
}

/*
GetCaseEntitiesByIDsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCaseEntitiesByIDsForbidden struct {

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

// IsSuccess returns true when this get case entities by i ds forbidden response has a 2xx status code
func (o *GetCaseEntitiesByIDsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get case entities by i ds forbidden response has a 3xx status code
func (o *GetCaseEntitiesByIDsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get case entities by i ds forbidden response has a 4xx status code
func (o *GetCaseEntitiesByIDsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get case entities by i ds forbidden response has a 5xx status code
func (o *GetCaseEntitiesByIDsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get case entities by i ds forbidden response a status code equal to that given
func (o *GetCaseEntitiesByIDsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get case entities by i ds forbidden response
func (o *GetCaseEntitiesByIDsForbidden) Code() int {
	return 403
}

func (o *GetCaseEntitiesByIDsForbidden) Error() string {
	return fmt.Sprintf("[POST /message-center/entities/cases/GET/v1][%d] getCaseEntitiesByIDsForbidden  %+v", 403, o.Payload)
}

func (o *GetCaseEntitiesByIDsForbidden) String() string {
	return fmt.Sprintf("[POST /message-center/entities/cases/GET/v1][%d] getCaseEntitiesByIDsForbidden  %+v", 403, o.Payload)
}

func (o *GetCaseEntitiesByIDsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCaseEntitiesByIDsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCaseEntitiesByIDsTooManyRequests creates a GetCaseEntitiesByIDsTooManyRequests with default headers values
func NewGetCaseEntitiesByIDsTooManyRequests() *GetCaseEntitiesByIDsTooManyRequests {
	return &GetCaseEntitiesByIDsTooManyRequests{}
}

/*
GetCaseEntitiesByIDsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCaseEntitiesByIDsTooManyRequests struct {

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

// IsSuccess returns true when this get case entities by i ds too many requests response has a 2xx status code
func (o *GetCaseEntitiesByIDsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get case entities by i ds too many requests response has a 3xx status code
func (o *GetCaseEntitiesByIDsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get case entities by i ds too many requests response has a 4xx status code
func (o *GetCaseEntitiesByIDsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get case entities by i ds too many requests response has a 5xx status code
func (o *GetCaseEntitiesByIDsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get case entities by i ds too many requests response a status code equal to that given
func (o *GetCaseEntitiesByIDsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get case entities by i ds too many requests response
func (o *GetCaseEntitiesByIDsTooManyRequests) Code() int {
	return 429
}

func (o *GetCaseEntitiesByIDsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /message-center/entities/cases/GET/v1][%d] getCaseEntitiesByIDsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCaseEntitiesByIDsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /message-center/entities/cases/GET/v1][%d] getCaseEntitiesByIDsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCaseEntitiesByIDsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCaseEntitiesByIDsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCaseEntitiesByIDsInternalServerError creates a GetCaseEntitiesByIDsInternalServerError with default headers values
func NewGetCaseEntitiesByIDsInternalServerError() *GetCaseEntitiesByIDsInternalServerError {
	return &GetCaseEntitiesByIDsInternalServerError{}
}

/*
GetCaseEntitiesByIDsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetCaseEntitiesByIDsInternalServerError struct {

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

// IsSuccess returns true when this get case entities by i ds internal server error response has a 2xx status code
func (o *GetCaseEntitiesByIDsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get case entities by i ds internal server error response has a 3xx status code
func (o *GetCaseEntitiesByIDsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get case entities by i ds internal server error response has a 4xx status code
func (o *GetCaseEntitiesByIDsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get case entities by i ds internal server error response has a 5xx status code
func (o *GetCaseEntitiesByIDsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get case entities by i ds internal server error response a status code equal to that given
func (o *GetCaseEntitiesByIDsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get case entities by i ds internal server error response
func (o *GetCaseEntitiesByIDsInternalServerError) Code() int {
	return 500
}

func (o *GetCaseEntitiesByIDsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /message-center/entities/cases/GET/v1][%d] getCaseEntitiesByIDsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCaseEntitiesByIDsInternalServerError) String() string {
	return fmt.Sprintf("[POST /message-center/entities/cases/GET/v1][%d] getCaseEntitiesByIDsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCaseEntitiesByIDsInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCaseEntitiesByIDsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCaseEntitiesByIDsDefault creates a GetCaseEntitiesByIDsDefault with default headers values
func NewGetCaseEntitiesByIDsDefault(code int) *GetCaseEntitiesByIDsDefault {
	return &GetCaseEntitiesByIDsDefault{
		_statusCode: code,
	}
}

/*
GetCaseEntitiesByIDsDefault describes a response with status code -1, with default header values.

OK
*/
type GetCaseEntitiesByIDsDefault struct {
	_statusCode int

	Payload *models.APIMessageCenterCasesResponse
}

// IsSuccess returns true when this get case entities by i ds default response has a 2xx status code
func (o *GetCaseEntitiesByIDsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get case entities by i ds default response has a 3xx status code
func (o *GetCaseEntitiesByIDsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get case entities by i ds default response has a 4xx status code
func (o *GetCaseEntitiesByIDsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get case entities by i ds default response has a 5xx status code
func (o *GetCaseEntitiesByIDsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get case entities by i ds default response a status code equal to that given
func (o *GetCaseEntitiesByIDsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get case entities by i ds default response
func (o *GetCaseEntitiesByIDsDefault) Code() int {
	return o._statusCode
}

func (o *GetCaseEntitiesByIDsDefault) Error() string {
	return fmt.Sprintf("[POST /message-center/entities/cases/GET/v1][%d] GetCaseEntitiesByIDs default  %+v", o._statusCode, o.Payload)
}

func (o *GetCaseEntitiesByIDsDefault) String() string {
	return fmt.Sprintf("[POST /message-center/entities/cases/GET/v1][%d] GetCaseEntitiesByIDs default  %+v", o._statusCode, o.Payload)
}

func (o *GetCaseEntitiesByIDsDefault) GetPayload() *models.APIMessageCenterCasesResponse {
	return o.Payload
}

func (o *GetCaseEntitiesByIDsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageCenterCasesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
