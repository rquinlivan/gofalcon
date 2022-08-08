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

// RTRGetPutFilesV2Reader is a Reader for the RTRGetPutFilesV2 structure.
type RTRGetPutFilesV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRGetPutFilesV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRTRGetPutFilesV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRGetPutFilesV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRGetPutFilesV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRTRGetPutFilesV2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRGetPutFilesV2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRTRGetPutFilesV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRTRGetPutFilesV2OK creates a RTRGetPutFilesV2OK with default headers values
func NewRTRGetPutFilesV2OK() *RTRGetPutFilesV2OK {
	return &RTRGetPutFilesV2OK{}
}

/*
RTRGetPutFilesV2OK describes a response with status code 200, with default header values.

OK
*/
type RTRGetPutFilesV2OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainMsaPFResponseV2
}

func (o *RTRGetPutFilesV2OK) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/put-files/v2][%d] rTRGetPutFilesV2OK  %+v", 200, o.Payload)
}
func (o *RTRGetPutFilesV2OK) GetPayload() *models.DomainMsaPFResponseV2 {
	return o.Payload
}

func (o *RTRGetPutFilesV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainMsaPFResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRGetPutFilesV2BadRequest creates a RTRGetPutFilesV2BadRequest with default headers values
func NewRTRGetPutFilesV2BadRequest() *RTRGetPutFilesV2BadRequest {
	return &RTRGetPutFilesV2BadRequest{}
}

/*
RTRGetPutFilesV2BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RTRGetPutFilesV2BadRequest struct {

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

func (o *RTRGetPutFilesV2BadRequest) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/put-files/v2][%d] rTRGetPutFilesV2BadRequest  %+v", 400, o.Payload)
}
func (o *RTRGetPutFilesV2BadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRGetPutFilesV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRGetPutFilesV2Forbidden creates a RTRGetPutFilesV2Forbidden with default headers values
func NewRTRGetPutFilesV2Forbidden() *RTRGetPutFilesV2Forbidden {
	return &RTRGetPutFilesV2Forbidden{}
}

/*
RTRGetPutFilesV2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRGetPutFilesV2Forbidden struct {

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

func (o *RTRGetPutFilesV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/put-files/v2][%d] rTRGetPutFilesV2Forbidden  %+v", 403, o.Payload)
}
func (o *RTRGetPutFilesV2Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRGetPutFilesV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRGetPutFilesV2NotFound creates a RTRGetPutFilesV2NotFound with default headers values
func NewRTRGetPutFilesV2NotFound() *RTRGetPutFilesV2NotFound {
	return &RTRGetPutFilesV2NotFound{}
}

/*
RTRGetPutFilesV2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type RTRGetPutFilesV2NotFound struct {

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

func (o *RTRGetPutFilesV2NotFound) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/put-files/v2][%d] rTRGetPutFilesV2NotFound  %+v", 404, o.Payload)
}
func (o *RTRGetPutFilesV2NotFound) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRGetPutFilesV2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRGetPutFilesV2TooManyRequests creates a RTRGetPutFilesV2TooManyRequests with default headers values
func NewRTRGetPutFilesV2TooManyRequests() *RTRGetPutFilesV2TooManyRequests {
	return &RTRGetPutFilesV2TooManyRequests{}
}

/*
RTRGetPutFilesV2TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRGetPutFilesV2TooManyRequests struct {

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

func (o *RTRGetPutFilesV2TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/put-files/v2][%d] rTRGetPutFilesV2TooManyRequests  %+v", 429, o.Payload)
}
func (o *RTRGetPutFilesV2TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRGetPutFilesV2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRGetPutFilesV2Default creates a RTRGetPutFilesV2Default with default headers values
func NewRTRGetPutFilesV2Default(code int) *RTRGetPutFilesV2Default {
	return &RTRGetPutFilesV2Default{
		_statusCode: code,
	}
}

/*
RTRGetPutFilesV2Default describes a response with status code -1, with default header values.

OK
*/
type RTRGetPutFilesV2Default struct {
	_statusCode int

	Payload *models.DomainMsaPFResponseV2
}

// Code gets the status code for the r t r get put files v2 default response
func (o *RTRGetPutFilesV2Default) Code() int {
	return o._statusCode
}

func (o *RTRGetPutFilesV2Default) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/put-files/v2][%d] RTR-GetPut-FilesV2 default  %+v", o._statusCode, o.Payload)
}
func (o *RTRGetPutFilesV2Default) GetPayload() *models.DomainMsaPFResponseV2 {
	return o.Payload
}

func (o *RTRGetPutFilesV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainMsaPFResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
