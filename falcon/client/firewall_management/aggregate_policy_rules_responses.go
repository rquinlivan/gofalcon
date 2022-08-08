// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

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

// AggregatePolicyRulesReader is a Reader for the AggregatePolicyRules structure.
type AggregatePolicyRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregatePolicyRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregatePolicyRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAggregatePolicyRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAggregatePolicyRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAggregatePolicyRulesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAggregatePolicyRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAggregatePolicyRulesOK creates a AggregatePolicyRulesOK with default headers values
func NewAggregatePolicyRulesOK() *AggregatePolicyRulesOK {
	return &AggregatePolicyRulesOK{}
}

/*
AggregatePolicyRulesOK describes a response with status code 200, with default header values.

OK
*/
type AggregatePolicyRulesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrAPIAggregatesResponse
}

func (o *AggregatePolicyRulesOK) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/policy-rules/GET/v1][%d] aggregatePolicyRulesOK  %+v", 200, o.Payload)
}
func (o *AggregatePolicyRulesOK) GetPayload() *models.FwmgrAPIAggregatesResponse {
	return o.Payload
}

func (o *AggregatePolicyRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FwmgrAPIAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregatePolicyRulesBadRequest creates a AggregatePolicyRulesBadRequest with default headers values
func NewAggregatePolicyRulesBadRequest() *AggregatePolicyRulesBadRequest {
	return &AggregatePolicyRulesBadRequest{}
}

/*
AggregatePolicyRulesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AggregatePolicyRulesBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrMsaReplyMetaOnly
}

func (o *AggregatePolicyRulesBadRequest) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/policy-rules/GET/v1][%d] aggregatePolicyRulesBadRequest  %+v", 400, o.Payload)
}
func (o *AggregatePolicyRulesBadRequest) GetPayload() *models.FwmgrMsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregatePolicyRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FwmgrMsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregatePolicyRulesForbidden creates a AggregatePolicyRulesForbidden with default headers values
func NewAggregatePolicyRulesForbidden() *AggregatePolicyRulesForbidden {
	return &AggregatePolicyRulesForbidden{}
}

/*
AggregatePolicyRulesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AggregatePolicyRulesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *AggregatePolicyRulesForbidden) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/policy-rules/GET/v1][%d] aggregatePolicyRulesForbidden  %+v", 403, o.Payload)
}
func (o *AggregatePolicyRulesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregatePolicyRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregatePolicyRulesTooManyRequests creates a AggregatePolicyRulesTooManyRequests with default headers values
func NewAggregatePolicyRulesTooManyRequests() *AggregatePolicyRulesTooManyRequests {
	return &AggregatePolicyRulesTooManyRequests{}
}

/*
AggregatePolicyRulesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AggregatePolicyRulesTooManyRequests struct {

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

func (o *AggregatePolicyRulesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/policy-rules/GET/v1][%d] aggregatePolicyRulesTooManyRequests  %+v", 429, o.Payload)
}
func (o *AggregatePolicyRulesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregatePolicyRulesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregatePolicyRulesDefault creates a AggregatePolicyRulesDefault with default headers values
func NewAggregatePolicyRulesDefault(code int) *AggregatePolicyRulesDefault {
	return &AggregatePolicyRulesDefault{
		_statusCode: code,
	}
}

/*
AggregatePolicyRulesDefault describes a response with status code -1, with default header values.

OK
*/
type AggregatePolicyRulesDefault struct {
	_statusCode int

	Payload *models.FwmgrAPIAggregatesResponse
}

// Code gets the status code for the aggregate policy rules default response
func (o *AggregatePolicyRulesDefault) Code() int {
	return o._statusCode
}

func (o *AggregatePolicyRulesDefault) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/policy-rules/GET/v1][%d] aggregate-policy-rules default  %+v", o._statusCode, o.Payload)
}
func (o *AggregatePolicyRulesDefault) GetPayload() *models.FwmgrAPIAggregatesResponse {
	return o.Payload
}

func (o *AggregatePolicyRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FwmgrAPIAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
