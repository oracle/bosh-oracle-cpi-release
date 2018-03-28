// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CrossConnectStatus The status of the cross-connect.
// swagger:model CrossConnectStatus
type CrossConnectStatus struct {

	// The OCID of the cross-connect.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	CrossConnectID *string `json:"crossConnectId"`

	// Whether Oracle's side of the interface is up or down.
	InterfaceState string `json:"interfaceState,omitempty"`

	// The light level of the cross-connect (in dBm).
	//
	// Example: `14.0`
	//
	// Maximum: 50
	// Minimum: -50
	LightLevelIndBm *float32 `json:"lightLevelIndBm,omitempty"`

	// Status indicator corresponding to the light level.
	//
	//   * **NO_LIGHT:** No measurable light
	//
	//   * **LOW_WARN:** There's measurable light but it's too low
	//
	//   * **HIGH_WARN:** Light level is too high
	//
	//   * **BAD:** There's measurable light but the signal-to-noise ratio is bad
	//
	//   * **GOOD:** Good light level
	//
	LightLevelIndicator string `json:"lightLevelIndicator,omitempty"`
}

// Validate validates this cross connect status
func (m *CrossConnectStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCrossConnectID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInterfaceState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLightLevelIndBm(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLightLevelIndicator(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CrossConnectStatus) validateCrossConnectID(formats strfmt.Registry) error {

	if err := validate.Required("crossConnectId", "body", m.CrossConnectID); err != nil {
		return err
	}

	if err := validate.MinLength("crossConnectId", "body", string(*m.CrossConnectID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("crossConnectId", "body", string(*m.CrossConnectID), 255); err != nil {
		return err
	}

	return nil
}

var crossConnectStatusTypeInterfaceStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UP","DOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		crossConnectStatusTypeInterfaceStatePropEnum = append(crossConnectStatusTypeInterfaceStatePropEnum, v)
	}
}

const (

	// CrossConnectStatusInterfaceStateUP captures enum value "UP"
	CrossConnectStatusInterfaceStateUP string = "UP"

	// CrossConnectStatusInterfaceStateDOWN captures enum value "DOWN"
	CrossConnectStatusInterfaceStateDOWN string = "DOWN"
)

// prop value enum
func (m *CrossConnectStatus) validateInterfaceStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, crossConnectStatusTypeInterfaceStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CrossConnectStatus) validateInterfaceState(formats strfmt.Registry) error {

	if swag.IsZero(m.InterfaceState) { // not required
		return nil
	}

	// value enum
	if err := m.validateInterfaceStateEnum("interfaceState", "body", m.InterfaceState); err != nil {
		return err
	}

	return nil
}

func (m *CrossConnectStatus) validateLightLevelIndBm(formats strfmt.Registry) error {

	if swag.IsZero(m.LightLevelIndBm) { // not required
		return nil
	}

	if err := validate.Minimum("lightLevelIndBm", "body", float64(*m.LightLevelIndBm), -50, false); err != nil {
		return err
	}

	if err := validate.Maximum("lightLevelIndBm", "body", float64(*m.LightLevelIndBm), 50, false); err != nil {
		return err
	}

	return nil
}

var crossConnectStatusTypeLightLevelIndicatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NO_LIGHT","LOW_WARN","HIGH_WARN","BAD","GOOD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		crossConnectStatusTypeLightLevelIndicatorPropEnum = append(crossConnectStatusTypeLightLevelIndicatorPropEnum, v)
	}
}

const (

	// CrossConnectStatusLightLevelIndicatorNOLIGHT captures enum value "NO_LIGHT"
	CrossConnectStatusLightLevelIndicatorNOLIGHT string = "NO_LIGHT"

	// CrossConnectStatusLightLevelIndicatorLOWWARN captures enum value "LOW_WARN"
	CrossConnectStatusLightLevelIndicatorLOWWARN string = "LOW_WARN"

	// CrossConnectStatusLightLevelIndicatorHIGHWARN captures enum value "HIGH_WARN"
	CrossConnectStatusLightLevelIndicatorHIGHWARN string = "HIGH_WARN"

	// CrossConnectStatusLightLevelIndicatorBAD captures enum value "BAD"
	CrossConnectStatusLightLevelIndicatorBAD string = "BAD"

	// CrossConnectStatusLightLevelIndicatorGOOD captures enum value "GOOD"
	CrossConnectStatusLightLevelIndicatorGOOD string = "GOOD"
)

// prop value enum
func (m *CrossConnectStatus) validateLightLevelIndicatorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, crossConnectStatusTypeLightLevelIndicatorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CrossConnectStatus) validateLightLevelIndicator(formats strfmt.Registry) error {

	if swag.IsZero(m.LightLevelIndicator) { // not required
		return nil
	}

	// value enum
	if err := m.validateLightLevelIndicatorEnum("lightLevelIndicator", "body", m.LightLevelIndicator); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CrossConnectStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CrossConnectStatus) UnmarshalBinary(b []byte) error {
	var res CrossConnectStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
