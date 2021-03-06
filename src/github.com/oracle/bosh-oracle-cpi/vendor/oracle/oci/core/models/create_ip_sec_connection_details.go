package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateIPSecConnectionDetails create IP sec connection details
// swagger:model CreateIPSecConnectionDetails
type CreateIPSecConnectionDetails struct {

	// The OCID of the compartment to contain the IPSec connection.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	CompartmentID *string `json:"compartmentId"`

	// The OCID of the CPE.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	CpeID *string `json:"cpeId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	// Max Length: 255
	// Min Length: 1
	DisplayName string `json:"displayName,omitempty"`

	// The OCID of the DRG.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	DrgID *string `json:"drgId"`

	// Static routes to the CPE. At least one route must be included. The CIDR must not be a
	// multicast address or class E address.
	//
	// Example: `10.0.1.0/24`
	//
	// Required: true
	// Max Length: 10
	// Min Length: 1
	// Unique: true
	StaticRoutes []string `json:"staticRoutes"`
}

// Validate validates this create IP sec connection details
func (m *CreateIPSecConnectionDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompartmentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCpeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDrgID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStaticRoutes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateIPSecConnectionDetails) validateCompartmentID(formats strfmt.Registry) error {

	if err := validate.Required("compartmentId", "body", m.CompartmentID); err != nil {
		return err
	}

	if err := validate.MinLength("compartmentId", "body", string(*m.CompartmentID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("compartmentId", "body", string(*m.CompartmentID), 255); err != nil {
		return err
	}

	return nil
}

func (m *CreateIPSecConnectionDetails) validateCpeID(formats strfmt.Registry) error {

	if err := validate.Required("cpeId", "body", m.CpeID); err != nil {
		return err
	}

	if err := validate.MinLength("cpeId", "body", string(*m.CpeID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("cpeId", "body", string(*m.CpeID), 255); err != nil {
		return err
	}

	return nil
}

func (m *CreateIPSecConnectionDetails) validateDisplayName(formats strfmt.Registry) error {

	if swag.IsZero(m.DisplayName) { // not required
		return nil
	}

	if err := validate.MinLength("displayName", "body", string(m.DisplayName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("displayName", "body", string(m.DisplayName), 255); err != nil {
		return err
	}

	return nil
}

func (m *CreateIPSecConnectionDetails) validateDrgID(formats strfmt.Registry) error {

	if err := validate.Required("drgId", "body", m.DrgID); err != nil {
		return err
	}

	if err := validate.MinLength("drgId", "body", string(*m.DrgID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("drgId", "body", string(*m.DrgID), 255); err != nil {
		return err
	}

	return nil
}

func (m *CreateIPSecConnectionDetails) validateStaticRoutes(formats strfmt.Registry) error {

	if err := validate.Required("staticRoutes", "body", m.StaticRoutes); err != nil {
		return err
	}

	if err := validate.UniqueItems("staticRoutes", "body", m.StaticRoutes); err != nil {
		return err
	}

	return nil
}
