// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DomainUser domain user
//
// swagger:model domain.User
type DomainUser struct {

	// cid
	Cid string `json:"cid,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this domain user
func (m *DomainUser) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this domain user based on context it is used
func (m *DomainUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainUser) UnmarshalBinary(b []byte) error {
	var res DomainUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
