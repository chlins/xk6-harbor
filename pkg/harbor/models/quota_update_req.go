// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QuotaUpdateReq quota update req
//
// swagger:model QuotaUpdateReq
type QuotaUpdateReq struct {

	// The new hard limits for the quota
	Hard ResourceList `json:"hard,omitempty" js:"hard"`
}

// Validate validates this quota update req
func (m *QuotaUpdateReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaUpdateReq) validateHard(formats strfmt.Registry) error {
	if swag.IsZero(m.Hard) { // not required
		return nil
	}

	if m.Hard != nil {
		if err := m.Hard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hard")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this quota update req based on the context it is used
func (m *QuotaUpdateReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaUpdateReq) contextValidateHard(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Hard.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hard")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuotaUpdateReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaUpdateReq) UnmarshalBinary(b []byte) error {
	var res QuotaUpdateReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
