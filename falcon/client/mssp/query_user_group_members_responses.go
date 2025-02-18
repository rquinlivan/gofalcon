// Code generated by go-swagger; DO NOT EDIT.

package mssp

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

// QueryUserGroupMembersReader is a Reader for the QueryUserGroupMembers structure.
type QueryUserGroupMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryUserGroupMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryUserGroupMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryUserGroupMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryUserGroupMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryUserGroupMembersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryUserGroupMembersOK creates a QueryUserGroupMembersOK with default headers values
func NewQueryUserGroupMembersOK() *QueryUserGroupMembersOK {
	return &QueryUserGroupMembersOK{}
}

/*
QueryUserGroupMembersOK describes a response with status code 200, with default header values.

OK
*/
type QueryUserGroupMembersOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query user group members o k response has a 2xx status code
func (o *QueryUserGroupMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query user group members o k response has a 3xx status code
func (o *QueryUserGroupMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query user group members o k response has a 4xx status code
func (o *QueryUserGroupMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query user group members o k response has a 5xx status code
func (o *QueryUserGroupMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query user group members o k response a status code equal to that given
func (o *QueryUserGroupMembersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query user group members o k response
func (o *QueryUserGroupMembersOK) Code() int {
	return 200
}

func (o *QueryUserGroupMembersOK) Error() string {
	return fmt.Sprintf("[GET /mssp/queries/user-group-members/v1][%d] queryUserGroupMembersOK  %+v", 200, o.Payload)
}

func (o *QueryUserGroupMembersOK) String() string {
	return fmt.Sprintf("[GET /mssp/queries/user-group-members/v1][%d] queryUserGroupMembersOK  %+v", 200, o.Payload)
}

func (o *QueryUserGroupMembersOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryUserGroupMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryUserGroupMembersForbidden creates a QueryUserGroupMembersForbidden with default headers values
func NewQueryUserGroupMembersForbidden() *QueryUserGroupMembersForbidden {
	return &QueryUserGroupMembersForbidden{}
}

/*
QueryUserGroupMembersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryUserGroupMembersForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this query user group members forbidden response has a 2xx status code
func (o *QueryUserGroupMembersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query user group members forbidden response has a 3xx status code
func (o *QueryUserGroupMembersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query user group members forbidden response has a 4xx status code
func (o *QueryUserGroupMembersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query user group members forbidden response has a 5xx status code
func (o *QueryUserGroupMembersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query user group members forbidden response a status code equal to that given
func (o *QueryUserGroupMembersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query user group members forbidden response
func (o *QueryUserGroupMembersForbidden) Code() int {
	return 403
}

func (o *QueryUserGroupMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /mssp/queries/user-group-members/v1][%d] queryUserGroupMembersForbidden  %+v", 403, o.Payload)
}

func (o *QueryUserGroupMembersForbidden) String() string {
	return fmt.Sprintf("[GET /mssp/queries/user-group-members/v1][%d] queryUserGroupMembersForbidden  %+v", 403, o.Payload)
}

func (o *QueryUserGroupMembersForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryUserGroupMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryUserGroupMembersTooManyRequests creates a QueryUserGroupMembersTooManyRequests with default headers values
func NewQueryUserGroupMembersTooManyRequests() *QueryUserGroupMembersTooManyRequests {
	return &QueryUserGroupMembersTooManyRequests{}
}

/*
QueryUserGroupMembersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryUserGroupMembersTooManyRequests struct {

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

// IsSuccess returns true when this query user group members too many requests response has a 2xx status code
func (o *QueryUserGroupMembersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query user group members too many requests response has a 3xx status code
func (o *QueryUserGroupMembersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query user group members too many requests response has a 4xx status code
func (o *QueryUserGroupMembersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query user group members too many requests response has a 5xx status code
func (o *QueryUserGroupMembersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query user group members too many requests response a status code equal to that given
func (o *QueryUserGroupMembersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query user group members too many requests response
func (o *QueryUserGroupMembersTooManyRequests) Code() int {
	return 429
}

func (o *QueryUserGroupMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mssp/queries/user-group-members/v1][%d] queryUserGroupMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryUserGroupMembersTooManyRequests) String() string {
	return fmt.Sprintf("[GET /mssp/queries/user-group-members/v1][%d] queryUserGroupMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryUserGroupMembersTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryUserGroupMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryUserGroupMembersDefault creates a QueryUserGroupMembersDefault with default headers values
func NewQueryUserGroupMembersDefault(code int) *QueryUserGroupMembersDefault {
	return &QueryUserGroupMembersDefault{
		_statusCode: code,
	}
}

/*
QueryUserGroupMembersDefault describes a response with status code -1, with default header values.

OK
*/
type QueryUserGroupMembersDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query user group members default response has a 2xx status code
func (o *QueryUserGroupMembersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query user group members default response has a 3xx status code
func (o *QueryUserGroupMembersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query user group members default response has a 4xx status code
func (o *QueryUserGroupMembersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query user group members default response has a 5xx status code
func (o *QueryUserGroupMembersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query user group members default response a status code equal to that given
func (o *QueryUserGroupMembersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the query user group members default response
func (o *QueryUserGroupMembersDefault) Code() int {
	return o._statusCode
}

func (o *QueryUserGroupMembersDefault) Error() string {
	return fmt.Sprintf("[GET /mssp/queries/user-group-members/v1][%d] queryUserGroupMembers default  %+v", o._statusCode, o.Payload)
}

func (o *QueryUserGroupMembersDefault) String() string {
	return fmt.Sprintf("[GET /mssp/queries/user-group-members/v1][%d] queryUserGroupMembers default  %+v", o._statusCode, o.Payload)
}

func (o *QueryUserGroupMembersDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryUserGroupMembersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
