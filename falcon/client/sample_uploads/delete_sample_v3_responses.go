// Code generated by go-swagger; DO NOT EDIT.

package sample_uploads

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

// DeleteSampleV3Reader is a Reader for the DeleteSampleV3 structure.
type DeleteSampleV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSampleV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSampleV3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSampleV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteSampleV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSampleV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteSampleV3TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteSampleV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteSampleV3Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSampleV3OK creates a DeleteSampleV3OK with default headers values
func NewDeleteSampleV3OK() *DeleteSampleV3OK {
	return &DeleteSampleV3OK{}
}

/*
DeleteSampleV3OK describes a response with status code 200, with default header values.

OK
*/
type DeleteSampleV3OK struct {

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

// IsSuccess returns true when this delete sample v3 o k response has a 2xx status code
func (o *DeleteSampleV3OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete sample v3 o k response has a 3xx status code
func (o *DeleteSampleV3OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete sample v3 o k response has a 4xx status code
func (o *DeleteSampleV3OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete sample v3 o k response has a 5xx status code
func (o *DeleteSampleV3OK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete sample v3 o k response a status code equal to that given
func (o *DeleteSampleV3OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete sample v3 o k response
func (o *DeleteSampleV3OK) Code() int {
	return 200
}

func (o *DeleteSampleV3OK) Error() string {
	return fmt.Sprintf("[DELETE /samples/entities/samples/v3][%d] deleteSampleV3OK  %+v", 200, o.Payload)
}

func (o *DeleteSampleV3OK) String() string {
	return fmt.Sprintf("[DELETE /samples/entities/samples/v3][%d] deleteSampleV3OK  %+v", 200, o.Payload)
}

func (o *DeleteSampleV3OK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeleteSampleV3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSampleV3BadRequest creates a DeleteSampleV3BadRequest with default headers values
func NewDeleteSampleV3BadRequest() *DeleteSampleV3BadRequest {
	return &DeleteSampleV3BadRequest{}
}

/*
DeleteSampleV3BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteSampleV3BadRequest struct {

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

// IsSuccess returns true when this delete sample v3 bad request response has a 2xx status code
func (o *DeleteSampleV3BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete sample v3 bad request response has a 3xx status code
func (o *DeleteSampleV3BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete sample v3 bad request response has a 4xx status code
func (o *DeleteSampleV3BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete sample v3 bad request response has a 5xx status code
func (o *DeleteSampleV3BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete sample v3 bad request response a status code equal to that given
func (o *DeleteSampleV3BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete sample v3 bad request response
func (o *DeleteSampleV3BadRequest) Code() int {
	return 400
}

func (o *DeleteSampleV3BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /samples/entities/samples/v3][%d] deleteSampleV3BadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSampleV3BadRequest) String() string {
	return fmt.Sprintf("[DELETE /samples/entities/samples/v3][%d] deleteSampleV3BadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSampleV3BadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeleteSampleV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSampleV3Forbidden creates a DeleteSampleV3Forbidden with default headers values
func NewDeleteSampleV3Forbidden() *DeleteSampleV3Forbidden {
	return &DeleteSampleV3Forbidden{}
}

/*
DeleteSampleV3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteSampleV3Forbidden struct {

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

// IsSuccess returns true when this delete sample v3 forbidden response has a 2xx status code
func (o *DeleteSampleV3Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete sample v3 forbidden response has a 3xx status code
func (o *DeleteSampleV3Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete sample v3 forbidden response has a 4xx status code
func (o *DeleteSampleV3Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete sample v3 forbidden response has a 5xx status code
func (o *DeleteSampleV3Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete sample v3 forbidden response a status code equal to that given
func (o *DeleteSampleV3Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete sample v3 forbidden response
func (o *DeleteSampleV3Forbidden) Code() int {
	return 403
}

func (o *DeleteSampleV3Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /samples/entities/samples/v3][%d] deleteSampleV3Forbidden  %+v", 403, o.Payload)
}

func (o *DeleteSampleV3Forbidden) String() string {
	return fmt.Sprintf("[DELETE /samples/entities/samples/v3][%d] deleteSampleV3Forbidden  %+v", 403, o.Payload)
}

func (o *DeleteSampleV3Forbidden) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeleteSampleV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSampleV3NotFound creates a DeleteSampleV3NotFound with default headers values
func NewDeleteSampleV3NotFound() *DeleteSampleV3NotFound {
	return &DeleteSampleV3NotFound{}
}

/*
DeleteSampleV3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteSampleV3NotFound struct {

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

// IsSuccess returns true when this delete sample v3 not found response has a 2xx status code
func (o *DeleteSampleV3NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete sample v3 not found response has a 3xx status code
func (o *DeleteSampleV3NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete sample v3 not found response has a 4xx status code
func (o *DeleteSampleV3NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete sample v3 not found response has a 5xx status code
func (o *DeleteSampleV3NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete sample v3 not found response a status code equal to that given
func (o *DeleteSampleV3NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete sample v3 not found response
func (o *DeleteSampleV3NotFound) Code() int {
	return 404
}

func (o *DeleteSampleV3NotFound) Error() string {
	return fmt.Sprintf("[DELETE /samples/entities/samples/v3][%d] deleteSampleV3NotFound  %+v", 404, o.Payload)
}

func (o *DeleteSampleV3NotFound) String() string {
	return fmt.Sprintf("[DELETE /samples/entities/samples/v3][%d] deleteSampleV3NotFound  %+v", 404, o.Payload)
}

func (o *DeleteSampleV3NotFound) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeleteSampleV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSampleV3TooManyRequests creates a DeleteSampleV3TooManyRequests with default headers values
func NewDeleteSampleV3TooManyRequests() *DeleteSampleV3TooManyRequests {
	return &DeleteSampleV3TooManyRequests{}
}

/*
DeleteSampleV3TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteSampleV3TooManyRequests struct {

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

// IsSuccess returns true when this delete sample v3 too many requests response has a 2xx status code
func (o *DeleteSampleV3TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete sample v3 too many requests response has a 3xx status code
func (o *DeleteSampleV3TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete sample v3 too many requests response has a 4xx status code
func (o *DeleteSampleV3TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete sample v3 too many requests response has a 5xx status code
func (o *DeleteSampleV3TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete sample v3 too many requests response a status code equal to that given
func (o *DeleteSampleV3TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete sample v3 too many requests response
func (o *DeleteSampleV3TooManyRequests) Code() int {
	return 429
}

func (o *DeleteSampleV3TooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /samples/entities/samples/v3][%d] deleteSampleV3TooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteSampleV3TooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /samples/entities/samples/v3][%d] deleteSampleV3TooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteSampleV3TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteSampleV3TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSampleV3InternalServerError creates a DeleteSampleV3InternalServerError with default headers values
func NewDeleteSampleV3InternalServerError() *DeleteSampleV3InternalServerError {
	return &DeleteSampleV3InternalServerError{}
}

/*
DeleteSampleV3InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteSampleV3InternalServerError struct {

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

// IsSuccess returns true when this delete sample v3 internal server error response has a 2xx status code
func (o *DeleteSampleV3InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete sample v3 internal server error response has a 3xx status code
func (o *DeleteSampleV3InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete sample v3 internal server error response has a 4xx status code
func (o *DeleteSampleV3InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete sample v3 internal server error response has a 5xx status code
func (o *DeleteSampleV3InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete sample v3 internal server error response a status code equal to that given
func (o *DeleteSampleV3InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete sample v3 internal server error response
func (o *DeleteSampleV3InternalServerError) Code() int {
	return 500
}

func (o *DeleteSampleV3InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /samples/entities/samples/v3][%d] deleteSampleV3InternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSampleV3InternalServerError) String() string {
	return fmt.Sprintf("[DELETE /samples/entities/samples/v3][%d] deleteSampleV3InternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSampleV3InternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeleteSampleV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSampleV3Default creates a DeleteSampleV3Default with default headers values
func NewDeleteSampleV3Default(code int) *DeleteSampleV3Default {
	return &DeleteSampleV3Default{
		_statusCode: code,
	}
}

/*
DeleteSampleV3Default describes a response with status code -1, with default header values.

OK
*/
type DeleteSampleV3Default struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this delete sample v3 default response has a 2xx status code
func (o *DeleteSampleV3Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete sample v3 default response has a 3xx status code
func (o *DeleteSampleV3Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete sample v3 default response has a 4xx status code
func (o *DeleteSampleV3Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete sample v3 default response has a 5xx status code
func (o *DeleteSampleV3Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete sample v3 default response a status code equal to that given
func (o *DeleteSampleV3Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete sample v3 default response
func (o *DeleteSampleV3Default) Code() int {
	return o._statusCode
}

func (o *DeleteSampleV3Default) Error() string {
	return fmt.Sprintf("[DELETE /samples/entities/samples/v3][%d] DeleteSampleV3 default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSampleV3Default) String() string {
	return fmt.Sprintf("[DELETE /samples/entities/samples/v3][%d] DeleteSampleV3 default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSampleV3Default) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeleteSampleV3Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
