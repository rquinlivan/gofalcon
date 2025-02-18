// Code generated by go-swagger; DO NOT EDIT.

package real_time_response_admin

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

// RTRCreatePutFilesReader is a Reader for the RTRCreatePutFiles structure.
type RTRCreatePutFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRCreatePutFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRTRCreatePutFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRCreatePutFilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRCreatePutFilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRCreatePutFilesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRTRCreatePutFilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRTRCreatePutFilesOK creates a RTRCreatePutFilesOK with default headers values
func NewRTRCreatePutFilesOK() *RTRCreatePutFilesOK {
	return &RTRCreatePutFilesOK{}
}

/*
RTRCreatePutFilesOK describes a response with status code 200, with default header values.

OK
*/
type RTRCreatePutFilesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r create put files o k response has a 2xx status code
func (o *RTRCreatePutFilesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this r t r create put files o k response has a 3xx status code
func (o *RTRCreatePutFilesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r create put files o k response has a 4xx status code
func (o *RTRCreatePutFilesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r create put files o k response has a 5xx status code
func (o *RTRCreatePutFilesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r create put files o k response a status code equal to that given
func (o *RTRCreatePutFilesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the r t r create put files o k response
func (o *RTRCreatePutFilesOK) Code() int {
	return 200
}

func (o *RTRCreatePutFilesOK) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/put-files/v1][%d] rTRCreatePutFilesOK  %+v", 200, o.Payload)
}

func (o *RTRCreatePutFilesOK) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/put-files/v1][%d] rTRCreatePutFilesOK  %+v", 200, o.Payload)
}

func (o *RTRCreatePutFilesOK) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRCreatePutFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRCreatePutFilesBadRequest creates a RTRCreatePutFilesBadRequest with default headers values
func NewRTRCreatePutFilesBadRequest() *RTRCreatePutFilesBadRequest {
	return &RTRCreatePutFilesBadRequest{}
}

/*
RTRCreatePutFilesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RTRCreatePutFilesBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r create put files bad request response has a 2xx status code
func (o *RTRCreatePutFilesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r create put files bad request response has a 3xx status code
func (o *RTRCreatePutFilesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r create put files bad request response has a 4xx status code
func (o *RTRCreatePutFilesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r create put files bad request response has a 5xx status code
func (o *RTRCreatePutFilesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r create put files bad request response a status code equal to that given
func (o *RTRCreatePutFilesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the r t r create put files bad request response
func (o *RTRCreatePutFilesBadRequest) Code() int {
	return 400
}

func (o *RTRCreatePutFilesBadRequest) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/put-files/v1][%d] rTRCreatePutFilesBadRequest  %+v", 400, o.Payload)
}

func (o *RTRCreatePutFilesBadRequest) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/put-files/v1][%d] rTRCreatePutFilesBadRequest  %+v", 400, o.Payload)
}

func (o *RTRCreatePutFilesBadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRCreatePutFilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRCreatePutFilesForbidden creates a RTRCreatePutFilesForbidden with default headers values
func NewRTRCreatePutFilesForbidden() *RTRCreatePutFilesForbidden {
	return &RTRCreatePutFilesForbidden{}
}

/*
RTRCreatePutFilesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRCreatePutFilesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r create put files forbidden response has a 2xx status code
func (o *RTRCreatePutFilesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r create put files forbidden response has a 3xx status code
func (o *RTRCreatePutFilesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r create put files forbidden response has a 4xx status code
func (o *RTRCreatePutFilesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r create put files forbidden response has a 5xx status code
func (o *RTRCreatePutFilesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r create put files forbidden response a status code equal to that given
func (o *RTRCreatePutFilesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the r t r create put files forbidden response
func (o *RTRCreatePutFilesForbidden) Code() int {
	return 403
}

func (o *RTRCreatePutFilesForbidden) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/put-files/v1][%d] rTRCreatePutFilesForbidden  %+v", 403, o.Payload)
}

func (o *RTRCreatePutFilesForbidden) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/put-files/v1][%d] rTRCreatePutFilesForbidden  %+v", 403, o.Payload)
}

func (o *RTRCreatePutFilesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRCreatePutFilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRCreatePutFilesTooManyRequests creates a RTRCreatePutFilesTooManyRequests with default headers values
func NewRTRCreatePutFilesTooManyRequests() *RTRCreatePutFilesTooManyRequests {
	return &RTRCreatePutFilesTooManyRequests{}
}

/*
RTRCreatePutFilesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRCreatePutFilesTooManyRequests struct {

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

// IsSuccess returns true when this r t r create put files too many requests response has a 2xx status code
func (o *RTRCreatePutFilesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r create put files too many requests response has a 3xx status code
func (o *RTRCreatePutFilesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r create put files too many requests response has a 4xx status code
func (o *RTRCreatePutFilesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r create put files too many requests response has a 5xx status code
func (o *RTRCreatePutFilesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r create put files too many requests response a status code equal to that given
func (o *RTRCreatePutFilesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the r t r create put files too many requests response
func (o *RTRCreatePutFilesTooManyRequests) Code() int {
	return 429
}

func (o *RTRCreatePutFilesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/put-files/v1][%d] rTRCreatePutFilesTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRCreatePutFilesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/put-files/v1][%d] rTRCreatePutFilesTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRCreatePutFilesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRCreatePutFilesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRCreatePutFilesDefault creates a RTRCreatePutFilesDefault with default headers values
func NewRTRCreatePutFilesDefault(code int) *RTRCreatePutFilesDefault {
	return &RTRCreatePutFilesDefault{
		_statusCode: code,
	}
}

/*
RTRCreatePutFilesDefault describes a response with status code -1, with default header values.

OK
*/
type RTRCreatePutFilesDefault struct {
	_statusCode int

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r create put files default response has a 2xx status code
func (o *RTRCreatePutFilesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this r t r create put files default response has a 3xx status code
func (o *RTRCreatePutFilesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this r t r create put files default response has a 4xx status code
func (o *RTRCreatePutFilesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this r t r create put files default response has a 5xx status code
func (o *RTRCreatePutFilesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this r t r create put files default response a status code equal to that given
func (o *RTRCreatePutFilesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the r t r create put files default response
func (o *RTRCreatePutFilesDefault) Code() int {
	return o._statusCode
}

func (o *RTRCreatePutFilesDefault) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/put-files/v1][%d] RTR-CreatePut-Files default  %+v", o._statusCode, o.Payload)
}

func (o *RTRCreatePutFilesDefault) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/put-files/v1][%d] RTR-CreatePut-Files default  %+v", o._statusCode, o.Payload)
}

func (o *RTRCreatePutFilesDefault) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRCreatePutFilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
