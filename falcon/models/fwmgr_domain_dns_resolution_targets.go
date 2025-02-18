// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FwmgrDomainDNSResolutionTargets fwmgr domain DNS resolution targets
//
// swagger:model fwmgr.domain.DNSResolutionTargets
type FwmgrDomainDNSResolutionTargets struct {

	// targets
	// Required: true
	Targets []*FwmgrDomainDNSTarget `json:"targets"`
}

// Validate validates this fwmgr domain DNS resolution targets
func (m *FwmgrDomainDNSResolutionTargets) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTargets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrDomainDNSResolutionTargets) validateTargets(formats strfmt.Registry) error {

	if err := validate.Required("targets", "body", m.Targets); err != nil {
		return err
	}

	for i := 0; i < len(m.Targets); i++ {
		if swag.IsZero(m.Targets[i]) { // not required
			continue
		}

		if m.Targets[i] != nil {
			if err := m.Targets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this fwmgr domain DNS resolution targets based on the context it is used
func (m *FwmgrDomainDNSResolutionTargets) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTargets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrDomainDNSResolutionTargets) contextValidateTargets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Targets); i++ {

		if m.Targets[i] != nil {
			if err := m.Targets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("targets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FwmgrDomainDNSResolutionTargets) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FwmgrDomainDNSResolutionTargets) UnmarshalBinary(b []byte) error {
	var res FwmgrDomainDNSResolutionTargets
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
