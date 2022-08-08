// Code generated by go-swagger; DO NOT EDIT.

package cloud_connect_aws

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

// DeleteAWSAccountsReader is a Reader for the DeleteAWSAccounts structure.
type DeleteAWSAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAWSAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAWSAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAWSAccountsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAWSAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteAWSAccountsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAWSAccountsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteAWSAccountsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAWSAccountsOK creates a DeleteAWSAccountsOK with default headers values
func NewDeleteAWSAccountsOK() *DeleteAWSAccountsOK {
	return &DeleteAWSAccountsOK{}
}

/*
DeleteAWSAccountsOK describes a response with status code 200, with default header values.

OK
*/
type DeleteAWSAccountsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ModelsBaseResponseV1
}

func (o *DeleteAWSAccountsOK) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-aws/entities/accounts/v1][%d] deleteAWSAccountsOK  %+v", 200, o.Payload)
}
func (o *DeleteAWSAccountsOK) GetPayload() *models.ModelsBaseResponseV1 {
	return o.Payload
}

func (o *DeleteAWSAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ModelsBaseResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAWSAccountsBadRequest creates a DeleteAWSAccountsBadRequest with default headers values
func NewDeleteAWSAccountsBadRequest() *DeleteAWSAccountsBadRequest {
	return &DeleteAWSAccountsBadRequest{}
}

/*
DeleteAWSAccountsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteAWSAccountsBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ModelsBaseResponseV1
}

func (o *DeleteAWSAccountsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-aws/entities/accounts/v1][%d] deleteAWSAccountsBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteAWSAccountsBadRequest) GetPayload() *models.ModelsBaseResponseV1 {
	return o.Payload
}

func (o *DeleteAWSAccountsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ModelsBaseResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAWSAccountsForbidden creates a DeleteAWSAccountsForbidden with default headers values
func NewDeleteAWSAccountsForbidden() *DeleteAWSAccountsForbidden {
	return &DeleteAWSAccountsForbidden{}
}

/*
DeleteAWSAccountsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteAWSAccountsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *DeleteAWSAccountsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-aws/entities/accounts/v1][%d] deleteAWSAccountsForbidden  %+v", 403, o.Payload)
}
func (o *DeleteAWSAccountsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteAWSAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteAWSAccountsTooManyRequests creates a DeleteAWSAccountsTooManyRequests with default headers values
func NewDeleteAWSAccountsTooManyRequests() *DeleteAWSAccountsTooManyRequests {
	return &DeleteAWSAccountsTooManyRequests{}
}

/*
DeleteAWSAccountsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteAWSAccountsTooManyRequests struct {

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

func (o *DeleteAWSAccountsTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-aws/entities/accounts/v1][%d] deleteAWSAccountsTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteAWSAccountsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteAWSAccountsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteAWSAccountsInternalServerError creates a DeleteAWSAccountsInternalServerError with default headers values
func NewDeleteAWSAccountsInternalServerError() *DeleteAWSAccountsInternalServerError {
	return &DeleteAWSAccountsInternalServerError{}
}

/*
DeleteAWSAccountsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteAWSAccountsInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ModelsBaseResponseV1
}

func (o *DeleteAWSAccountsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-aws/entities/accounts/v1][%d] deleteAWSAccountsInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteAWSAccountsInternalServerError) GetPayload() *models.ModelsBaseResponseV1 {
	return o.Payload
}

func (o *DeleteAWSAccountsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ModelsBaseResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAWSAccountsDefault creates a DeleteAWSAccountsDefault with default headers values
func NewDeleteAWSAccountsDefault(code int) *DeleteAWSAccountsDefault {
	return &DeleteAWSAccountsDefault{
		_statusCode: code,
	}
}

/*
DeleteAWSAccountsDefault describes a response with status code -1, with default header values.

OK
*/
type DeleteAWSAccountsDefault struct {
	_statusCode int

	Payload *models.ModelsBaseResponseV1
}

// Code gets the status code for the delete a w s accounts default response
func (o *DeleteAWSAccountsDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAWSAccountsDefault) Error() string {
	return fmt.Sprintf("[DELETE /cloud-connect-aws/entities/accounts/v1][%d] DeleteAWSAccounts default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteAWSAccountsDefault) GetPayload() *models.ModelsBaseResponseV1 {
	return o.Payload
}

func (o *DeleteAWSAccountsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsBaseResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
