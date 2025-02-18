// Code generated by go-swagger; DO NOT EDIT.

package falconx_sandbox

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

// GetArtifactsReader is a Reader for the GetArtifacts structure.
type GetArtifactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArtifactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewGetArtifactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArtifactsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArtifactsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetArtifactsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArtifactsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArtifactsBadRequest creates a GetArtifactsBadRequest with default headers values
func NewGetArtifactsBadRequest() *GetArtifactsBadRequest {
	return &GetArtifactsBadRequest{}
}

/*
GetArtifactsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetArtifactsBadRequest struct {

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

// IsSuccess returns true when this get artifacts bad request response has a 2xx status code
func (o *GetArtifactsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get artifacts bad request response has a 3xx status code
func (o *GetArtifactsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get artifacts bad request response has a 4xx status code
func (o *GetArtifactsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get artifacts bad request response has a 5xx status code
func (o *GetArtifactsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get artifacts bad request response a status code equal to that given
func (o *GetArtifactsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get artifacts bad request response
func (o *GetArtifactsBadRequest) Code() int {
	return 400
}

func (o *GetArtifactsBadRequest) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/artifacts/v1][%d] getArtifactsBadRequest  %+v", 400, o.Payload)
}

func (o *GetArtifactsBadRequest) String() string {
	return fmt.Sprintf("[GET /falconx/entities/artifacts/v1][%d] getArtifactsBadRequest  %+v", 400, o.Payload)
}

func (o *GetArtifactsBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetArtifactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetArtifactsForbidden creates a GetArtifactsForbidden with default headers values
func NewGetArtifactsForbidden() *GetArtifactsForbidden {
	return &GetArtifactsForbidden{}
}

/*
GetArtifactsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetArtifactsForbidden struct {

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

// IsSuccess returns true when this get artifacts forbidden response has a 2xx status code
func (o *GetArtifactsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get artifacts forbidden response has a 3xx status code
func (o *GetArtifactsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get artifacts forbidden response has a 4xx status code
func (o *GetArtifactsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get artifacts forbidden response has a 5xx status code
func (o *GetArtifactsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get artifacts forbidden response a status code equal to that given
func (o *GetArtifactsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get artifacts forbidden response
func (o *GetArtifactsForbidden) Code() int {
	return 403
}

func (o *GetArtifactsForbidden) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/artifacts/v1][%d] getArtifactsForbidden  %+v", 403, o.Payload)
}

func (o *GetArtifactsForbidden) String() string {
	return fmt.Sprintf("[GET /falconx/entities/artifacts/v1][%d] getArtifactsForbidden  %+v", 403, o.Payload)
}

func (o *GetArtifactsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetArtifactsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetArtifactsNotFound creates a GetArtifactsNotFound with default headers values
func NewGetArtifactsNotFound() *GetArtifactsNotFound {
	return &GetArtifactsNotFound{}
}

/*
GetArtifactsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetArtifactsNotFound struct {

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

// IsSuccess returns true when this get artifacts not found response has a 2xx status code
func (o *GetArtifactsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get artifacts not found response has a 3xx status code
func (o *GetArtifactsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get artifacts not found response has a 4xx status code
func (o *GetArtifactsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get artifacts not found response has a 5xx status code
func (o *GetArtifactsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get artifacts not found response a status code equal to that given
func (o *GetArtifactsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get artifacts not found response
func (o *GetArtifactsNotFound) Code() int {
	return 404
}

func (o *GetArtifactsNotFound) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/artifacts/v1][%d] getArtifactsNotFound  %+v", 404, o.Payload)
}

func (o *GetArtifactsNotFound) String() string {
	return fmt.Sprintf("[GET /falconx/entities/artifacts/v1][%d] getArtifactsNotFound  %+v", 404, o.Payload)
}

func (o *GetArtifactsNotFound) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetArtifactsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetArtifactsTooManyRequests creates a GetArtifactsTooManyRequests with default headers values
func NewGetArtifactsTooManyRequests() *GetArtifactsTooManyRequests {
	return &GetArtifactsTooManyRequests{}
}

/*
GetArtifactsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetArtifactsTooManyRequests struct {

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

// IsSuccess returns true when this get artifacts too many requests response has a 2xx status code
func (o *GetArtifactsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get artifacts too many requests response has a 3xx status code
func (o *GetArtifactsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get artifacts too many requests response has a 4xx status code
func (o *GetArtifactsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get artifacts too many requests response has a 5xx status code
func (o *GetArtifactsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get artifacts too many requests response a status code equal to that given
func (o *GetArtifactsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get artifacts too many requests response
func (o *GetArtifactsTooManyRequests) Code() int {
	return 429
}

func (o *GetArtifactsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/artifacts/v1][%d] getArtifactsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArtifactsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /falconx/entities/artifacts/v1][%d] getArtifactsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArtifactsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetArtifactsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetArtifactsInternalServerError creates a GetArtifactsInternalServerError with default headers values
func NewGetArtifactsInternalServerError() *GetArtifactsInternalServerError {
	return &GetArtifactsInternalServerError{}
}

/*
GetArtifactsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetArtifactsInternalServerError struct {

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

// IsSuccess returns true when this get artifacts internal server error response has a 2xx status code
func (o *GetArtifactsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get artifacts internal server error response has a 3xx status code
func (o *GetArtifactsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get artifacts internal server error response has a 4xx status code
func (o *GetArtifactsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get artifacts internal server error response has a 5xx status code
func (o *GetArtifactsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get artifacts internal server error response a status code equal to that given
func (o *GetArtifactsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get artifacts internal server error response
func (o *GetArtifactsInternalServerError) Code() int {
	return 500
}

func (o *GetArtifactsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/artifacts/v1][%d] getArtifactsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArtifactsInternalServerError) String() string {
	return fmt.Sprintf("[GET /falconx/entities/artifacts/v1][%d] getArtifactsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArtifactsInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetArtifactsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
