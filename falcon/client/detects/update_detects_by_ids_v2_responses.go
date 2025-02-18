// Code generated by go-swagger; DO NOT EDIT.

package detects

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

// UpdateDetectsByIdsV2Reader is a Reader for the UpdateDetectsByIdsV2 structure.
type UpdateDetectsByIdsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDetectsByIdsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDetectsByIdsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateDetectsByIdsV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateDetectsByIdsV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateDetectsByIdsV2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateDetectsByIdsV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateDetectsByIdsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDetectsByIdsV2OK creates a UpdateDetectsByIdsV2OK with default headers values
func NewUpdateDetectsByIdsV2OK() *UpdateDetectsByIdsV2OK {
	return &UpdateDetectsByIdsV2OK{}
}

/*
UpdateDetectsByIdsV2OK describes a response with status code 200, with default header values.

OK
*/
type UpdateDetectsByIdsV2OK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this update detects by ids v2 o k response has a 2xx status code
func (o *UpdateDetectsByIdsV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update detects by ids v2 o k response has a 3xx status code
func (o *UpdateDetectsByIdsV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update detects by ids v2 o k response has a 4xx status code
func (o *UpdateDetectsByIdsV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update detects by ids v2 o k response has a 5xx status code
func (o *UpdateDetectsByIdsV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this update detects by ids v2 o k response a status code equal to that given
func (o *UpdateDetectsByIdsV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update detects by ids v2 o k response
func (o *UpdateDetectsByIdsV2OK) Code() int {
	return 200
}

func (o *UpdateDetectsByIdsV2OK) Error() string {
	return fmt.Sprintf("[PATCH /detects/entities/detects/v2][%d] updateDetectsByIdsV2OK  %+v", 200, o.Payload)
}

func (o *UpdateDetectsByIdsV2OK) String() string {
	return fmt.Sprintf("[PATCH /detects/entities/detects/v2][%d] updateDetectsByIdsV2OK  %+v", 200, o.Payload)
}

func (o *UpdateDetectsByIdsV2OK) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateDetectsByIdsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateDetectsByIdsV2BadRequest creates a UpdateDetectsByIdsV2BadRequest with default headers values
func NewUpdateDetectsByIdsV2BadRequest() *UpdateDetectsByIdsV2BadRequest {
	return &UpdateDetectsByIdsV2BadRequest{}
}

/*
UpdateDetectsByIdsV2BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateDetectsByIdsV2BadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this update detects by ids v2 bad request response has a 2xx status code
func (o *UpdateDetectsByIdsV2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update detects by ids v2 bad request response has a 3xx status code
func (o *UpdateDetectsByIdsV2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update detects by ids v2 bad request response has a 4xx status code
func (o *UpdateDetectsByIdsV2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update detects by ids v2 bad request response has a 5xx status code
func (o *UpdateDetectsByIdsV2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update detects by ids v2 bad request response a status code equal to that given
func (o *UpdateDetectsByIdsV2BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update detects by ids v2 bad request response
func (o *UpdateDetectsByIdsV2BadRequest) Code() int {
	return 400
}

func (o *UpdateDetectsByIdsV2BadRequest) Error() string {
	return fmt.Sprintf("[PATCH /detects/entities/detects/v2][%d] updateDetectsByIdsV2BadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDetectsByIdsV2BadRequest) String() string {
	return fmt.Sprintf("[PATCH /detects/entities/detects/v2][%d] updateDetectsByIdsV2BadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDetectsByIdsV2BadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateDetectsByIdsV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateDetectsByIdsV2Forbidden creates a UpdateDetectsByIdsV2Forbidden with default headers values
func NewUpdateDetectsByIdsV2Forbidden() *UpdateDetectsByIdsV2Forbidden {
	return &UpdateDetectsByIdsV2Forbidden{}
}

/*
UpdateDetectsByIdsV2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateDetectsByIdsV2Forbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this update detects by ids v2 forbidden response has a 2xx status code
func (o *UpdateDetectsByIdsV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update detects by ids v2 forbidden response has a 3xx status code
func (o *UpdateDetectsByIdsV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update detects by ids v2 forbidden response has a 4xx status code
func (o *UpdateDetectsByIdsV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update detects by ids v2 forbidden response has a 5xx status code
func (o *UpdateDetectsByIdsV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update detects by ids v2 forbidden response a status code equal to that given
func (o *UpdateDetectsByIdsV2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update detects by ids v2 forbidden response
func (o *UpdateDetectsByIdsV2Forbidden) Code() int {
	return 403
}

func (o *UpdateDetectsByIdsV2Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /detects/entities/detects/v2][%d] updateDetectsByIdsV2Forbidden  %+v", 403, o.Payload)
}

func (o *UpdateDetectsByIdsV2Forbidden) String() string {
	return fmt.Sprintf("[PATCH /detects/entities/detects/v2][%d] updateDetectsByIdsV2Forbidden  %+v", 403, o.Payload)
}

func (o *UpdateDetectsByIdsV2Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateDetectsByIdsV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateDetectsByIdsV2TooManyRequests creates a UpdateDetectsByIdsV2TooManyRequests with default headers values
func NewUpdateDetectsByIdsV2TooManyRequests() *UpdateDetectsByIdsV2TooManyRequests {
	return &UpdateDetectsByIdsV2TooManyRequests{}
}

/*
UpdateDetectsByIdsV2TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateDetectsByIdsV2TooManyRequests struct {

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

// IsSuccess returns true when this update detects by ids v2 too many requests response has a 2xx status code
func (o *UpdateDetectsByIdsV2TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update detects by ids v2 too many requests response has a 3xx status code
func (o *UpdateDetectsByIdsV2TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update detects by ids v2 too many requests response has a 4xx status code
func (o *UpdateDetectsByIdsV2TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update detects by ids v2 too many requests response has a 5xx status code
func (o *UpdateDetectsByIdsV2TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update detects by ids v2 too many requests response a status code equal to that given
func (o *UpdateDetectsByIdsV2TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the update detects by ids v2 too many requests response
func (o *UpdateDetectsByIdsV2TooManyRequests) Code() int {
	return 429
}

func (o *UpdateDetectsByIdsV2TooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /detects/entities/detects/v2][%d] updateDetectsByIdsV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateDetectsByIdsV2TooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /detects/entities/detects/v2][%d] updateDetectsByIdsV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateDetectsByIdsV2TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateDetectsByIdsV2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateDetectsByIdsV2InternalServerError creates a UpdateDetectsByIdsV2InternalServerError with default headers values
func NewUpdateDetectsByIdsV2InternalServerError() *UpdateDetectsByIdsV2InternalServerError {
	return &UpdateDetectsByIdsV2InternalServerError{}
}

/*
UpdateDetectsByIdsV2InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateDetectsByIdsV2InternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this update detects by ids v2 internal server error response has a 2xx status code
func (o *UpdateDetectsByIdsV2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update detects by ids v2 internal server error response has a 3xx status code
func (o *UpdateDetectsByIdsV2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update detects by ids v2 internal server error response has a 4xx status code
func (o *UpdateDetectsByIdsV2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update detects by ids v2 internal server error response has a 5xx status code
func (o *UpdateDetectsByIdsV2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update detects by ids v2 internal server error response a status code equal to that given
func (o *UpdateDetectsByIdsV2InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update detects by ids v2 internal server error response
func (o *UpdateDetectsByIdsV2InternalServerError) Code() int {
	return 500
}

func (o *UpdateDetectsByIdsV2InternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /detects/entities/detects/v2][%d] updateDetectsByIdsV2InternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateDetectsByIdsV2InternalServerError) String() string {
	return fmt.Sprintf("[PATCH /detects/entities/detects/v2][%d] updateDetectsByIdsV2InternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateDetectsByIdsV2InternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateDetectsByIdsV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateDetectsByIdsV2Default creates a UpdateDetectsByIdsV2Default with default headers values
func NewUpdateDetectsByIdsV2Default(code int) *UpdateDetectsByIdsV2Default {
	return &UpdateDetectsByIdsV2Default{
		_statusCode: code,
	}
}

/*
UpdateDetectsByIdsV2Default describes a response with status code -1, with default header values.

OK
*/
type UpdateDetectsByIdsV2Default struct {
	_statusCode int

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this update detects by ids v2 default response has a 2xx status code
func (o *UpdateDetectsByIdsV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update detects by ids v2 default response has a 3xx status code
func (o *UpdateDetectsByIdsV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update detects by ids v2 default response has a 4xx status code
func (o *UpdateDetectsByIdsV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update detects by ids v2 default response has a 5xx status code
func (o *UpdateDetectsByIdsV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update detects by ids v2 default response a status code equal to that given
func (o *UpdateDetectsByIdsV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update detects by ids v2 default response
func (o *UpdateDetectsByIdsV2Default) Code() int {
	return o._statusCode
}

func (o *UpdateDetectsByIdsV2Default) Error() string {
	return fmt.Sprintf("[PATCH /detects/entities/detects/v2][%d] UpdateDetectsByIdsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDetectsByIdsV2Default) String() string {
	return fmt.Sprintf("[PATCH /detects/entities/detects/v2][%d] UpdateDetectsByIdsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDetectsByIdsV2Default) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateDetectsByIdsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
