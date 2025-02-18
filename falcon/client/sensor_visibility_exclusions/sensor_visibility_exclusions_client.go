// Code generated by go-swagger; DO NOT EDIT.

package sensor_visibility_exclusions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new sensor visibility exclusions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sensor visibility exclusions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateSVExclusionsV1(params *CreateSVExclusionsV1Params, opts ...ClientOption) (*CreateSVExclusionsV1OK, error)

	DeleteSensorVisibilityExclusionsV1(params *DeleteSensorVisibilityExclusionsV1Params, opts ...ClientOption) (*DeleteSensorVisibilityExclusionsV1OK, error)

	GetSensorVisibilityExclusionsV1(params *GetSensorVisibilityExclusionsV1Params, opts ...ClientOption) (*GetSensorVisibilityExclusionsV1OK, error)

	QuerySensorVisibilityExclusionsV1(params *QuerySensorVisibilityExclusionsV1Params, opts ...ClientOption) (*QuerySensorVisibilityExclusionsV1OK, error)

	UpdateSensorVisibilityExclusionsV1(params *UpdateSensorVisibilityExclusionsV1Params, opts ...ClientOption) (*UpdateSensorVisibilityExclusionsV1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateSVExclusionsV1 creates the sensor visibility exclusions
*/
func (a *Client) CreateSVExclusionsV1(params *CreateSVExclusionsV1Params, opts ...ClientOption) (*CreateSVExclusionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSVExclusionsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "createSVExclusionsV1",
		Method:             "POST",
		PathPattern:        "/policy/entities/sv-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSVExclusionsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateSVExclusionsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateSVExclusionsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteSensorVisibilityExclusionsV1 deletes the sensor visibility exclusions by id
*/
func (a *Client) DeleteSensorVisibilityExclusionsV1(params *DeleteSensorVisibilityExclusionsV1Params, opts ...ClientOption) (*DeleteSensorVisibilityExclusionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSensorVisibilityExclusionsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSensorVisibilityExclusionsV1",
		Method:             "DELETE",
		PathPattern:        "/policy/entities/sv-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSensorVisibilityExclusionsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteSensorVisibilityExclusionsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteSensorVisibilityExclusionsV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSensorVisibilityExclusionsV1 gets a set of sensor visibility exclusions by specifying their i ds
*/
func (a *Client) GetSensorVisibilityExclusionsV1(params *GetSensorVisibilityExclusionsV1Params, opts ...ClientOption) (*GetSensorVisibilityExclusionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSensorVisibilityExclusionsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSensorVisibilityExclusionsV1",
		Method:             "GET",
		PathPattern:        "/policy/entities/sv-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSensorVisibilityExclusionsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSensorVisibilityExclusionsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSensorVisibilityExclusionsV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QuerySensorVisibilityExclusionsV1 searches for sensor visibility exclusions
*/
func (a *Client) QuerySensorVisibilityExclusionsV1(params *QuerySensorVisibilityExclusionsV1Params, opts ...ClientOption) (*QuerySensorVisibilityExclusionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuerySensorVisibilityExclusionsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "querySensorVisibilityExclusionsV1",
		Method:             "GET",
		PathPattern:        "/policy/queries/sv-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QuerySensorVisibilityExclusionsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QuerySensorVisibilityExclusionsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QuerySensorVisibilityExclusionsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateSensorVisibilityExclusionsV1 updates the sensor visibility exclusions
*/
func (a *Client) UpdateSensorVisibilityExclusionsV1(params *UpdateSensorVisibilityExclusionsV1Params, opts ...ClientOption) (*UpdateSensorVisibilityExclusionsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSensorVisibilityExclusionsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateSensorVisibilityExclusionsV1",
		Method:             "PATCH",
		PathPattern:        "/policy/entities/sv-exclusions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSensorVisibilityExclusionsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateSensorVisibilityExclusionsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateSensorVisibilityExclusionsV1Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
