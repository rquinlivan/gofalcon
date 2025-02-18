// Code generated by go-swagger; DO NOT EDIT.

package hosts

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

// QueryDevicesByFilterReader is a Reader for the QueryDevicesByFilter structure.
type QueryDevicesByFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryDevicesByFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryDevicesByFilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryDevicesByFilterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryDevicesByFilterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryDevicesByFilterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryDevicesByFilterOK creates a QueryDevicesByFilterOK with default headers values
func NewQueryDevicesByFilterOK() *QueryDevicesByFilterOK {
	return &QueryDevicesByFilterOK{}
}

/*
QueryDevicesByFilterOK describes a response with status code 200, with default header values.

OK
*/
type QueryDevicesByFilterOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query devices by filter o k response has a 2xx status code
func (o *QueryDevicesByFilterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query devices by filter o k response has a 3xx status code
func (o *QueryDevicesByFilterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query devices by filter o k response has a 4xx status code
func (o *QueryDevicesByFilterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query devices by filter o k response has a 5xx status code
func (o *QueryDevicesByFilterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query devices by filter o k response a status code equal to that given
func (o *QueryDevicesByFilterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query devices by filter o k response
func (o *QueryDevicesByFilterOK) Code() int {
	return 200
}

func (o *QueryDevicesByFilterOK) Error() string {
	return fmt.Sprintf("[GET /devices/queries/devices/v1][%d] queryDevicesByFilterOK  %+v", 200, o.Payload)
}

func (o *QueryDevicesByFilterOK) String() string {
	return fmt.Sprintf("[GET /devices/queries/devices/v1][%d] queryDevicesByFilterOK  %+v", 200, o.Payload)
}

func (o *QueryDevicesByFilterOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryDevicesByFilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryDevicesByFilterForbidden creates a QueryDevicesByFilterForbidden with default headers values
func NewQueryDevicesByFilterForbidden() *QueryDevicesByFilterForbidden {
	return &QueryDevicesByFilterForbidden{}
}

/*
QueryDevicesByFilterForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryDevicesByFilterForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this query devices by filter forbidden response has a 2xx status code
func (o *QueryDevicesByFilterForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query devices by filter forbidden response has a 3xx status code
func (o *QueryDevicesByFilterForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query devices by filter forbidden response has a 4xx status code
func (o *QueryDevicesByFilterForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query devices by filter forbidden response has a 5xx status code
func (o *QueryDevicesByFilterForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query devices by filter forbidden response a status code equal to that given
func (o *QueryDevicesByFilterForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query devices by filter forbidden response
func (o *QueryDevicesByFilterForbidden) Code() int {
	return 403
}

func (o *QueryDevicesByFilterForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/queries/devices/v1][%d] queryDevicesByFilterForbidden  %+v", 403, o.Payload)
}

func (o *QueryDevicesByFilterForbidden) String() string {
	return fmt.Sprintf("[GET /devices/queries/devices/v1][%d] queryDevicesByFilterForbidden  %+v", 403, o.Payload)
}

func (o *QueryDevicesByFilterForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryDevicesByFilterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryDevicesByFilterTooManyRequests creates a QueryDevicesByFilterTooManyRequests with default headers values
func NewQueryDevicesByFilterTooManyRequests() *QueryDevicesByFilterTooManyRequests {
	return &QueryDevicesByFilterTooManyRequests{}
}

/*
QueryDevicesByFilterTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryDevicesByFilterTooManyRequests struct {

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

// IsSuccess returns true when this query devices by filter too many requests response has a 2xx status code
func (o *QueryDevicesByFilterTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query devices by filter too many requests response has a 3xx status code
func (o *QueryDevicesByFilterTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query devices by filter too many requests response has a 4xx status code
func (o *QueryDevicesByFilterTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query devices by filter too many requests response has a 5xx status code
func (o *QueryDevicesByFilterTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query devices by filter too many requests response a status code equal to that given
func (o *QueryDevicesByFilterTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query devices by filter too many requests response
func (o *QueryDevicesByFilterTooManyRequests) Code() int {
	return 429
}

func (o *QueryDevicesByFilterTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /devices/queries/devices/v1][%d] queryDevicesByFilterTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryDevicesByFilterTooManyRequests) String() string {
	return fmt.Sprintf("[GET /devices/queries/devices/v1][%d] queryDevicesByFilterTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryDevicesByFilterTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryDevicesByFilterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryDevicesByFilterDefault creates a QueryDevicesByFilterDefault with default headers values
func NewQueryDevicesByFilterDefault(code int) *QueryDevicesByFilterDefault {
	return &QueryDevicesByFilterDefault{
		_statusCode: code,
	}
}

/*
QueryDevicesByFilterDefault describes a response with status code -1, with default header values.

OK
*/
type QueryDevicesByFilterDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query devices by filter default response has a 2xx status code
func (o *QueryDevicesByFilterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query devices by filter default response has a 3xx status code
func (o *QueryDevicesByFilterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query devices by filter default response has a 4xx status code
func (o *QueryDevicesByFilterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query devices by filter default response has a 5xx status code
func (o *QueryDevicesByFilterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query devices by filter default response a status code equal to that given
func (o *QueryDevicesByFilterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the query devices by filter default response
func (o *QueryDevicesByFilterDefault) Code() int {
	return o._statusCode
}

func (o *QueryDevicesByFilterDefault) Error() string {
	return fmt.Sprintf("[GET /devices/queries/devices/v1][%d] QueryDevicesByFilter default  %+v", o._statusCode, o.Payload)
}

func (o *QueryDevicesByFilterDefault) String() string {
	return fmt.Sprintf("[GET /devices/queries/devices/v1][%d] QueryDevicesByFilter default  %+v", o._statusCode, o.Payload)
}

func (o *QueryDevicesByFilterDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryDevicesByFilterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
