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

// ResponsesIOARuleGroupV1 An IOA rule group which contains a set of IOA rules
//
// swagger:model responses.IOARuleGroupV1
type ResponsesIOARuleGroupV1 struct {

	// comment
	// Required: true
	Comment *string `json:"comment"`

	// The last attempted time CFS got this data on the rule group
	// Required: true
	// Format: date-time
	CommittedTimestamp *strfmt.DateTime `json:"committed_timestamp"`

	// The email of the user which created the rule group
	// Required: true
	CreatedBy *string `json:"created_by"`

	// The time at which the policy was created
	// Required: true
	// Format: date-time
	CreatedTimestamp *strfmt.DateTime `json:"created_timestamp"`

	// customer id
	// Required: true
	CustomerID *string `json:"customer_id"`

	// deleted
	// Required: true
	Deleted *bool `json:"deleted"`

	// An additional description of the group or the rules it contains
	// Required: true
	Description *string `json:"description"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// The identifier of this IOA rule group
	// Required: true
	ID *string `json:"id"`

	// The email of the user which last modified the rule group
	// Required: true
	ModifiedBy *string `json:"modified_by"`

	// The time at which the policy was last modified
	// Required: true
	// Format: date-time
	ModifiedTimestamp *strfmt.DateTime `json:"modified_timestamp"`

	// The name of the group
	// Required: true
	Name *string `json:"name"`

	// platform
	// Required: true
	Platform *string `json:"platform"`

	// rule ids
	// Required: true
	RuleIds []string `json:"rule_ids"`

	// version
	// Required: true
	Version *int64 `json:"version"`
}

// Validate validates this responses i o a rule group v1
func (m *ResponsesIOARuleGroupV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommittedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponsesIOARuleGroupV1) validateComment(formats strfmt.Registry) error {

	if err := validate.Required("comment", "body", m.Comment); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesIOARuleGroupV1) validateCommittedTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("committed_timestamp", "body", m.CommittedTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("committed_timestamp", "body", "date-time", m.CommittedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesIOARuleGroupV1) validateCreatedBy(formats strfmt.Registry) error {

	if err := validate.Required("created_by", "body", m.CreatedBy); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesIOARuleGroupV1) validateCreatedTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("created_timestamp", "body", m.CreatedTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("created_timestamp", "body", "date-time", m.CreatedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesIOARuleGroupV1) validateCustomerID(formats strfmt.Registry) error {

	if err := validate.Required("customer_id", "body", m.CustomerID); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesIOARuleGroupV1) validateDeleted(formats strfmt.Registry) error {

	if err := validate.Required("deleted", "body", m.Deleted); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesIOARuleGroupV1) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesIOARuleGroupV1) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesIOARuleGroupV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesIOARuleGroupV1) validateModifiedBy(formats strfmt.Registry) error {

	if err := validate.Required("modified_by", "body", m.ModifiedBy); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesIOARuleGroupV1) validateModifiedTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("modified_timestamp", "body", m.ModifiedTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("modified_timestamp", "body", "date-time", m.ModifiedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesIOARuleGroupV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesIOARuleGroupV1) validatePlatform(formats strfmt.Registry) error {

	if err := validate.Required("platform", "body", m.Platform); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesIOARuleGroupV1) validateRuleIds(formats strfmt.Registry) error {

	if err := validate.Required("rule_ids", "body", m.RuleIds); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesIOARuleGroupV1) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this responses i o a rule group v1 based on context it is used
func (m *ResponsesIOARuleGroupV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponsesIOARuleGroupV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponsesIOARuleGroupV1) UnmarshalBinary(b []byte) error {
	var res ResponsesIOARuleGroupV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
