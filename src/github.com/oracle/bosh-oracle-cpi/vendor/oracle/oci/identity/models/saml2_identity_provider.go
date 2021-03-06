package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Saml2IdentityProvider A special type of [IdentityProvider](#/en/identity/20160918/IdentityProvider/) that
// supports the SAML 2.0 protocol. For more information, see
// [Identity Providers and Federation](/Content/Identity/Concepts/federation.htm).
//
// swagger:model Saml2IdentityProvider
type Saml2IdentityProvider struct {
	compartmentIdField *string

	descriptionField *string

	idField *string

	inactiveStatusField int64

	lifecycleStateField *string

	nameField *string

	productTypeField *string

	timeCreatedField *strfmt.DateTime

	// The URL for retrieving the identity provider's metadata, which
	// contains information required for federating.
	//
	// Required: true
	// Max Length: 512
	// Min Length: 1
	MetadataURL *string `json:"metadataUrl"`

	// The URL to redirect federated users to for authentication with the
	// identity provider.
	//
	// Required: true
	// Max Length: 512
	// Min Length: 1
	RedirectURL *string `json:"redirectUrl"`

	// The identity provider's signing certificate used by the IAM Service
	// to validate the SAML2 token.
	//
	// Required: true
	// Max Length: 10000
	// Min Length: 1
	SigningCertificate *string `json:"signingCertificate"`
}

func (m *Saml2IdentityProvider) CompartmentID() *string {
	return m.compartmentIdField
}
func (m *Saml2IdentityProvider) SetCompartmentID(val *string) {
	m.compartmentIdField = val
}

func (m *Saml2IdentityProvider) Description() *string {
	return m.descriptionField
}
func (m *Saml2IdentityProvider) SetDescription(val *string) {
	m.descriptionField = val
}

func (m *Saml2IdentityProvider) ID() *string {
	return m.idField
}
func (m *Saml2IdentityProvider) SetID(val *string) {
	m.idField = val
}

func (m *Saml2IdentityProvider) InactiveStatus() int64 {
	return m.inactiveStatusField
}
func (m *Saml2IdentityProvider) SetInactiveStatus(val int64) {
	m.inactiveStatusField = val
}

func (m *Saml2IdentityProvider) LifecycleState() *string {
	return m.lifecycleStateField
}
func (m *Saml2IdentityProvider) SetLifecycleState(val *string) {
	m.lifecycleStateField = val
}

func (m *Saml2IdentityProvider) Name() *string {
	return m.nameField
}
func (m *Saml2IdentityProvider) SetName(val *string) {
	m.nameField = val
}

func (m *Saml2IdentityProvider) ProductType() *string {
	return m.productTypeField
}
func (m *Saml2IdentityProvider) SetProductType(val *string) {
	m.productTypeField = val
}

func (m *Saml2IdentityProvider) Protocol() string {
	return "Saml2IdentityProvider"
}
func (m *Saml2IdentityProvider) SetProtocol(val string) {

}

func (m *Saml2IdentityProvider) TimeCreated() *strfmt.DateTime {
	return m.timeCreatedField
}
func (m *Saml2IdentityProvider) SetTimeCreated(val *strfmt.DateTime) {
	m.timeCreatedField = val
}

// UnmarshalJSON unmarshals this polymorphic type from a JSON structure
func (m *Saml2IdentityProvider) UnmarshalJSON(raw []byte) error {
	var data struct {
		CompartmentID *string `json:"compartmentId"`

		Description *string `json:"description"`

		ID *string `json:"id"`

		InactiveStatus int64 `json:"inactiveStatus,omitempty"`

		LifecycleState *string `json:"lifecycleState"`

		Name *string `json:"name"`

		ProductType *string `json:"productType"`

		Protocol string `json:"protocol"`

		TimeCreated *strfmt.DateTime `json:"timeCreated"`

		// The URL for retrieving the identity provider's metadata, which
		// contains information required for federating.
		//
		// Required: true
		// Max Length: 512
		// Min Length: 1
		MetadataURL *string `json:"metadataUrl"`

		// The URL to redirect federated users to for authentication with the
		// identity provider.
		//
		// Required: true
		// Max Length: 512
		// Min Length: 1
		RedirectURL *string `json:"redirectUrl"`

		// The identity provider's signing certificate used by the IAM Service
		// to validate the SAML2 token.
		//
		// Required: true
		// Max Length: 10000
		// Min Length: 1
		SigningCertificate *string `json:"signingCertificate"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	m.compartmentIdField = data.CompartmentID
	m.descriptionField = data.Description
	m.idField = data.ID
	m.inactiveStatusField = data.InactiveStatus
	m.lifecycleStateField = data.LifecycleState
	m.nameField = data.Name
	m.productTypeField = data.ProductType
	m.timeCreatedField = data.TimeCreated
	m.MetadataURL = data.MetadataURL
	m.RedirectURL = data.RedirectURL
	m.SigningCertificate = data.SigningCertificate

	return nil
}

// MarshalJSON marshals this polymorphic type to a JSON structure
func (m Saml2IdentityProvider) MarshalJSON() ([]byte, error) {
	var data struct {
		CompartmentID *string `json:"compartmentId"`

		Description *string `json:"description"`

		ID *string `json:"id"`

		InactiveStatus int64 `json:"inactiveStatus,omitempty"`

		LifecycleState *string `json:"lifecycleState"`

		Name *string `json:"name"`

		ProductType *string `json:"productType"`

		Protocol string `json:"protocol"`

		TimeCreated *strfmt.DateTime `json:"timeCreated"`

		// The URL for retrieving the identity provider's metadata, which
		// contains information required for federating.
		//
		// Required: true
		// Max Length: 512
		// Min Length: 1
		MetadataURL *string `json:"metadataUrl"`

		// The URL to redirect federated users to for authentication with the
		// identity provider.
		//
		// Required: true
		// Max Length: 512
		// Min Length: 1
		RedirectURL *string `json:"redirectUrl"`

		// The identity provider's signing certificate used by the IAM Service
		// to validate the SAML2 token.
		//
		// Required: true
		// Max Length: 10000
		// Min Length: 1
		SigningCertificate *string `json:"signingCertificate"`
	}

	data.CompartmentID = m.compartmentIdField
	data.Description = m.descriptionField
	data.ID = m.idField
	data.InactiveStatus = m.inactiveStatusField
	data.LifecycleState = m.lifecycleStateField
	data.Name = m.nameField
	data.ProductType = m.productTypeField
	data.TimeCreated = m.timeCreatedField
	data.MetadataURL = m.MetadataURL
	data.RedirectURL = m.RedirectURL
	data.SigningCertificate = m.SigningCertificate
	data.Protocol = "Saml2IdentityProvider"
	return json.Marshal(data)
}

// Validate validates this saml2 identity provider
func (m *Saml2IdentityProvider) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompartmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLifecycleState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadataURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirectURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSigningCertificate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Saml2IdentityProvider) validateCompartmentID(formats strfmt.Registry) error {

	if err := validate.Required("compartmentId", "body", m.CompartmentID()); err != nil {
		return err
	}

	return nil
}

func (m *Saml2IdentityProvider) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description()); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description()), 400); err != nil {
		return err
	}

	return nil
}

func (m *Saml2IdentityProvider) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID()); err != nil {
		return err
	}

	return nil
}

var saml2IdentityProviderTypeLifecycleStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATING","ACTIVE","INACTIVE","DELETING","DELETED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		saml2IdentityProviderTypeLifecycleStatePropEnum = append(saml2IdentityProviderTypeLifecycleStatePropEnum, v)
	}
}

// property enum
func (m *Saml2IdentityProvider) validateLifecycleStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, saml2IdentityProviderTypeLifecycleStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Saml2IdentityProvider) validateLifecycleState(formats strfmt.Registry) error {

	if err := validate.Required("lifecycleState", "body", m.LifecycleState()); err != nil {
		return err
	}

	if err := validate.MinLength("lifecycleState", "body", string(*m.LifecycleState()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("lifecycleState", "body", string(*m.LifecycleState()), 64); err != nil {
		return err
	}

	// value enum
	if err := m.validateLifecycleStateEnum("lifecycleState", "body", *m.LifecycleState()); err != nil {
		return err
	}

	return nil
}

func (m *Saml2IdentityProvider) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name()), 100); err != nil {
		return err
	}

	return nil
}

func (m *Saml2IdentityProvider) validateProductType(formats strfmt.Registry) error {

	if err := validate.Required("productType", "body", m.ProductType()); err != nil {
		return err
	}

	if err := validate.MinLength("productType", "body", string(*m.ProductType()), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("productType", "body", string(*m.ProductType()), 64); err != nil {
		return err
	}

	return nil
}

func (m *Saml2IdentityProvider) validateTimeCreated(formats strfmt.Registry) error {

	if err := validate.Required("timeCreated", "body", m.TimeCreated()); err != nil {
		return err
	}

	return nil
}

func (m *Saml2IdentityProvider) validateMetadataURL(formats strfmt.Registry) error {

	if err := validate.Required("metadataUrl", "body", m.MetadataURL); err != nil {
		return err
	}

	if err := validate.MinLength("metadataUrl", "body", string(*m.MetadataURL), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("metadataUrl", "body", string(*m.MetadataURL), 512); err != nil {
		return err
	}

	return nil
}

func (m *Saml2IdentityProvider) validateRedirectURL(formats strfmt.Registry) error {

	if err := validate.Required("redirectUrl", "body", m.RedirectURL); err != nil {
		return err
	}

	if err := validate.MinLength("redirectUrl", "body", string(*m.RedirectURL), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("redirectUrl", "body", string(*m.RedirectURL), 512); err != nil {
		return err
	}

	return nil
}

func (m *Saml2IdentityProvider) validateSigningCertificate(formats strfmt.Registry) error {

	if err := validate.Required("signingCertificate", "body", m.SigningCertificate); err != nil {
		return err
	}

	if err := validate.MinLength("signingCertificate", "body", string(*m.SigningCertificate), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("signingCertificate", "body", string(*m.SigningCertificate), 10000); err != nil {
		return err
	}

	return nil
}
