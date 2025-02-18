// Code generated by go-swagger; DO NOT EDIT.

package device_control_policies

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

// GetDeviceControlPoliciesReader is a Reader for the GetDeviceControlPolicies structure.
type GetDeviceControlPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceControlPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceControlPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetDeviceControlPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeviceControlPoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDeviceControlPoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDeviceControlPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeviceControlPoliciesOK creates a GetDeviceControlPoliciesOK with default headers values
func NewGetDeviceControlPoliciesOK() *GetDeviceControlPoliciesOK {
	return &GetDeviceControlPoliciesOK{}
}

/*
GetDeviceControlPoliciesOK describes a response with status code 200, with default header values.

OK
*/
type GetDeviceControlPoliciesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesDeviceControlPoliciesV1
}

// IsSuccess returns true when this get device control policies o k response has a 2xx status code
func (o *GetDeviceControlPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device control policies o k response has a 3xx status code
func (o *GetDeviceControlPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device control policies o k response has a 4xx status code
func (o *GetDeviceControlPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device control policies o k response has a 5xx status code
func (o *GetDeviceControlPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device control policies o k response a status code equal to that given
func (o *GetDeviceControlPoliciesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get device control policies o k response
func (o *GetDeviceControlPoliciesOK) Code() int {
	return 200
}

func (o *GetDeviceControlPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /policy/entities/device-control/v1][%d] getDeviceControlPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetDeviceControlPoliciesOK) String() string {
	return fmt.Sprintf("[GET /policy/entities/device-control/v1][%d] getDeviceControlPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetDeviceControlPoliciesOK) GetPayload() *models.ResponsesDeviceControlPoliciesV1 {
	return o.Payload
}

func (o *GetDeviceControlPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesDeviceControlPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceControlPoliciesForbidden creates a GetDeviceControlPoliciesForbidden with default headers values
func NewGetDeviceControlPoliciesForbidden() *GetDeviceControlPoliciesForbidden {
	return &GetDeviceControlPoliciesForbidden{}
}

/*
GetDeviceControlPoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDeviceControlPoliciesForbidden struct {

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

// IsSuccess returns true when this get device control policies forbidden response has a 2xx status code
func (o *GetDeviceControlPoliciesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get device control policies forbidden response has a 3xx status code
func (o *GetDeviceControlPoliciesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device control policies forbidden response has a 4xx status code
func (o *GetDeviceControlPoliciesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get device control policies forbidden response has a 5xx status code
func (o *GetDeviceControlPoliciesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get device control policies forbidden response a status code equal to that given
func (o *GetDeviceControlPoliciesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get device control policies forbidden response
func (o *GetDeviceControlPoliciesForbidden) Code() int {
	return 403
}

func (o *GetDeviceControlPoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/entities/device-control/v1][%d] getDeviceControlPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *GetDeviceControlPoliciesForbidden) String() string {
	return fmt.Sprintf("[GET /policy/entities/device-control/v1][%d] getDeviceControlPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *GetDeviceControlPoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetDeviceControlPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeviceControlPoliciesNotFound creates a GetDeviceControlPoliciesNotFound with default headers values
func NewGetDeviceControlPoliciesNotFound() *GetDeviceControlPoliciesNotFound {
	return &GetDeviceControlPoliciesNotFound{}
}

/*
GetDeviceControlPoliciesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDeviceControlPoliciesNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesDeviceControlPoliciesV1
}

// IsSuccess returns true when this get device control policies not found response has a 2xx status code
func (o *GetDeviceControlPoliciesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get device control policies not found response has a 3xx status code
func (o *GetDeviceControlPoliciesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device control policies not found response has a 4xx status code
func (o *GetDeviceControlPoliciesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get device control policies not found response has a 5xx status code
func (o *GetDeviceControlPoliciesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get device control policies not found response a status code equal to that given
func (o *GetDeviceControlPoliciesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get device control policies not found response
func (o *GetDeviceControlPoliciesNotFound) Code() int {
	return 404
}

func (o *GetDeviceControlPoliciesNotFound) Error() string {
	return fmt.Sprintf("[GET /policy/entities/device-control/v1][%d] getDeviceControlPoliciesNotFound  %+v", 404, o.Payload)
}

func (o *GetDeviceControlPoliciesNotFound) String() string {
	return fmt.Sprintf("[GET /policy/entities/device-control/v1][%d] getDeviceControlPoliciesNotFound  %+v", 404, o.Payload)
}

func (o *GetDeviceControlPoliciesNotFound) GetPayload() *models.ResponsesDeviceControlPoliciesV1 {
	return o.Payload
}

func (o *GetDeviceControlPoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesDeviceControlPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceControlPoliciesTooManyRequests creates a GetDeviceControlPoliciesTooManyRequests with default headers values
func NewGetDeviceControlPoliciesTooManyRequests() *GetDeviceControlPoliciesTooManyRequests {
	return &GetDeviceControlPoliciesTooManyRequests{}
}

/*
GetDeviceControlPoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetDeviceControlPoliciesTooManyRequests struct {

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

// IsSuccess returns true when this get device control policies too many requests response has a 2xx status code
func (o *GetDeviceControlPoliciesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get device control policies too many requests response has a 3xx status code
func (o *GetDeviceControlPoliciesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device control policies too many requests response has a 4xx status code
func (o *GetDeviceControlPoliciesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get device control policies too many requests response has a 5xx status code
func (o *GetDeviceControlPoliciesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get device control policies too many requests response a status code equal to that given
func (o *GetDeviceControlPoliciesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get device control policies too many requests response
func (o *GetDeviceControlPoliciesTooManyRequests) Code() int {
	return 429
}

func (o *GetDeviceControlPoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/entities/device-control/v1][%d] getDeviceControlPoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDeviceControlPoliciesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /policy/entities/device-control/v1][%d] getDeviceControlPoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDeviceControlPoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetDeviceControlPoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeviceControlPoliciesInternalServerError creates a GetDeviceControlPoliciesInternalServerError with default headers values
func NewGetDeviceControlPoliciesInternalServerError() *GetDeviceControlPoliciesInternalServerError {
	return &GetDeviceControlPoliciesInternalServerError{}
}

/*
GetDeviceControlPoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDeviceControlPoliciesInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesDeviceControlPoliciesV1
}

// IsSuccess returns true when this get device control policies internal server error response has a 2xx status code
func (o *GetDeviceControlPoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get device control policies internal server error response has a 3xx status code
func (o *GetDeviceControlPoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device control policies internal server error response has a 4xx status code
func (o *GetDeviceControlPoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device control policies internal server error response has a 5xx status code
func (o *GetDeviceControlPoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get device control policies internal server error response a status code equal to that given
func (o *GetDeviceControlPoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get device control policies internal server error response
func (o *GetDeviceControlPoliciesInternalServerError) Code() int {
	return 500
}

func (o *GetDeviceControlPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/entities/device-control/v1][%d] getDeviceControlPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDeviceControlPoliciesInternalServerError) String() string {
	return fmt.Sprintf("[GET /policy/entities/device-control/v1][%d] getDeviceControlPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDeviceControlPoliciesInternalServerError) GetPayload() *models.ResponsesDeviceControlPoliciesV1 {
	return o.Payload
}

func (o *GetDeviceControlPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesDeviceControlPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
