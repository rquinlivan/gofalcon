// Code generated by go-swagger; DO NOT EDIT.

package scheduled_reports

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

// ScheduledReportsLaunchReader is a Reader for the ScheduledReportsLaunch structure.
type ScheduledReportsLaunchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScheduledReportsLaunchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScheduledReportsLaunchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewScheduledReportsLaunchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewScheduledReportsLaunchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewScheduledReportsLaunchTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewScheduledReportsLaunchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewScheduledReportsLaunchOK creates a ScheduledReportsLaunchOK with default headers values
func NewScheduledReportsLaunchOK() *ScheduledReportsLaunchOK {
	return &ScheduledReportsLaunchOK{}
}

/*
ScheduledReportsLaunchOK describes a response with status code 200, with default header values.

OK
*/
type ScheduledReportsLaunchOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIReportExecutionsResponseV1
}

func (o *ScheduledReportsLaunchOK) Error() string {
	return fmt.Sprintf("[POST /reports/entities/scheduled-reports/execution/v1][%d] scheduledReportsLaunchOK  %+v", 200, o.Payload)
}
func (o *ScheduledReportsLaunchOK) GetPayload() *models.APIReportExecutionsResponseV1 {
	return o.Payload
}

func (o *ScheduledReportsLaunchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIReportExecutionsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScheduledReportsLaunchBadRequest creates a ScheduledReportsLaunchBadRequest with default headers values
func NewScheduledReportsLaunchBadRequest() *ScheduledReportsLaunchBadRequest {
	return &ScheduledReportsLaunchBadRequest{}
}

/*
ScheduledReportsLaunchBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ScheduledReportsLaunchBadRequest struct {

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

func (o *ScheduledReportsLaunchBadRequest) Error() string {
	return fmt.Sprintf("[POST /reports/entities/scheduled-reports/execution/v1][%d] scheduledReportsLaunchBadRequest  %+v", 400, o.Payload)
}
func (o *ScheduledReportsLaunchBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ScheduledReportsLaunchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScheduledReportsLaunchForbidden creates a ScheduledReportsLaunchForbidden with default headers values
func NewScheduledReportsLaunchForbidden() *ScheduledReportsLaunchForbidden {
	return &ScheduledReportsLaunchForbidden{}
}

/*
ScheduledReportsLaunchForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ScheduledReportsLaunchForbidden struct {

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

func (o *ScheduledReportsLaunchForbidden) Error() string {
	return fmt.Sprintf("[POST /reports/entities/scheduled-reports/execution/v1][%d] scheduledReportsLaunchForbidden  %+v", 403, o.Payload)
}
func (o *ScheduledReportsLaunchForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ScheduledReportsLaunchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScheduledReportsLaunchTooManyRequests creates a ScheduledReportsLaunchTooManyRequests with default headers values
func NewScheduledReportsLaunchTooManyRequests() *ScheduledReportsLaunchTooManyRequests {
	return &ScheduledReportsLaunchTooManyRequests{}
}

/*
ScheduledReportsLaunchTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ScheduledReportsLaunchTooManyRequests struct {

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

func (o *ScheduledReportsLaunchTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /reports/entities/scheduled-reports/execution/v1][%d] scheduledReportsLaunchTooManyRequests  %+v", 429, o.Payload)
}
func (o *ScheduledReportsLaunchTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ScheduledReportsLaunchTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScheduledReportsLaunchDefault creates a ScheduledReportsLaunchDefault with default headers values
func NewScheduledReportsLaunchDefault(code int) *ScheduledReportsLaunchDefault {
	return &ScheduledReportsLaunchDefault{
		_statusCode: code,
	}
}

/*
ScheduledReportsLaunchDefault describes a response with status code -1, with default header values.

OK
*/
type ScheduledReportsLaunchDefault struct {
	_statusCode int

	Payload *models.APIReportExecutionsResponseV1
}

// Code gets the status code for the scheduled reports launch default response
func (o *ScheduledReportsLaunchDefault) Code() int {
	return o._statusCode
}

func (o *ScheduledReportsLaunchDefault) Error() string {
	return fmt.Sprintf("[POST /reports/entities/scheduled-reports/execution/v1][%d] scheduled-reports.launch default  %+v", o._statusCode, o.Payload)
}
func (o *ScheduledReportsLaunchDefault) GetPayload() *models.APIReportExecutionsResponseV1 {
	return o.Payload
}

func (o *ScheduledReportsLaunchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIReportExecutionsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
