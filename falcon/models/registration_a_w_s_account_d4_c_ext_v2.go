// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RegistrationAWSAccountD4CExtV2 registration a w s account d4 c ext v2
//
// swagger:model registration.AWSAccountD4CExtV2
type RegistrationAWSAccountD4CExtV2 struct {

	// account id
	// Required: true
	AccountID *string `json:"account_id"`

	// account type
	AccountType string `json:"account_type,omitempty"`

	// cloudtrail region
	// Required: true
	CloudtrailRegion *string `json:"cloudtrail_region"`

	// is master
	IsMaster bool `json:"is_master,omitempty"`

	// organization id
	// Required: true
	OrganizationID *string `json:"organization_id"`
}

// Validate validates this registration a w s account d4 c ext v2
func (m *RegistrationAWSAccountD4CExtV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudtrailRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistrationAWSAccountD4CExtV2) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("account_id", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *RegistrationAWSAccountD4CExtV2) validateCloudtrailRegion(formats strfmt.Registry) error {

	if err := validate.Required("cloudtrail_region", "body", m.CloudtrailRegion); err != nil {
		return err
	}

	return nil
}

func (m *RegistrationAWSAccountD4CExtV2) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organization_id", "body", m.OrganizationID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this registration a w s account d4 c ext v2 based on context it is used
func (m *RegistrationAWSAccountD4CExtV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegistrationAWSAccountD4CExtV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistrationAWSAccountD4CExtV2) UnmarshalBinary(b []byte) error {
	var res RegistrationAWSAccountD4CExtV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
