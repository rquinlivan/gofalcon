// Code generated by go-swagger; DO NOT EDIT.

package falcon_complete_dashboard

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

// AggregateBlockListReader is a Reader for the AggregateBlockList structure.
type AggregateBlockListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregateBlockListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregateBlockListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAggregateBlockListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAggregateBlockListTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAggregateBlockListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAggregateBlockListOK creates a AggregateBlockListOK with default headers values
func NewAggregateBlockListOK() *AggregateBlockListOK {
	return &AggregateBlockListOK{}
}

/*
AggregateBlockListOK describes a response with status code 200, with default header values.

OK
*/
type AggregateBlockListOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaAggregatesResponse
}

func (o *AggregateBlockListOK) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/blocklist/GET/v1][%d] aggregateBlockListOK  %+v", 200, o.Payload)
}
func (o *AggregateBlockListOK) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *AggregateBlockListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateBlockListForbidden creates a AggregateBlockListForbidden with default headers values
func NewAggregateBlockListForbidden() *AggregateBlockListForbidden {
	return &AggregateBlockListForbidden{}
}

/*
AggregateBlockListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AggregateBlockListForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *AggregateBlockListForbidden) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/blocklist/GET/v1][%d] aggregateBlockListForbidden  %+v", 403, o.Payload)
}
func (o *AggregateBlockListForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateBlockListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateBlockListTooManyRequests creates a AggregateBlockListTooManyRequests with default headers values
func NewAggregateBlockListTooManyRequests() *AggregateBlockListTooManyRequests {
	return &AggregateBlockListTooManyRequests{}
}

/*
AggregateBlockListTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AggregateBlockListTooManyRequests struct {

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

func (o *AggregateBlockListTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/blocklist/GET/v1][%d] aggregateBlockListTooManyRequests  %+v", 429, o.Payload)
}
func (o *AggregateBlockListTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateBlockListTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateBlockListDefault creates a AggregateBlockListDefault with default headers values
func NewAggregateBlockListDefault(code int) *AggregateBlockListDefault {
	return &AggregateBlockListDefault{
		_statusCode: code,
	}
}

/*
AggregateBlockListDefault describes a response with status code -1, with default header values.

OK
*/
type AggregateBlockListDefault struct {
	_statusCode int

	Payload *models.MsaAggregatesResponse
}

// Code gets the status code for the aggregate block list default response
func (o *AggregateBlockListDefault) Code() int {
	return o._statusCode
}

func (o *AggregateBlockListDefault) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/blocklist/GET/v1][%d] AggregateBlockList default  %+v", o._statusCode, o.Payload)
}
func (o *AggregateBlockListDefault) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *AggregateBlockListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
