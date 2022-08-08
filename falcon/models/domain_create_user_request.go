// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DomainCreateUserRequest domain create user request
//
// swagger:model domain.CreateUserRequest
type DomainCreateUserRequest struct {

	// cid
	Cid string `json:"cid,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this domain create user request
func (m *DomainCreateUserRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this domain create user request based on context it is used
func (m *DomainCreateUserRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainCreateUserRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainCreateUserRequest) UnmarshalBinary(b []byte) error {
	var res DomainCreateUserRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
