// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"
)

// ImageSourceViaObjectStorageTupleDetailsAllOf1 image source via object storage tuple details all of1
// swagger:discriminator imageSourceViaObjectStorageTupleDetailsAllOf1 objectStorageTuple
type ImageSourceViaObjectStorageTupleDetailsAllOf1 interface {
	runtime.Validatable

	// The Object Storage bucket for the image.
	// Required: true
	BucketName() *string
	SetBucketName(*string)

	// The Object Storage namespace for the image.
	// Required: true
	NamespaceName() *string
	SetNamespaceName(*string)

	// The Object Storage name for the image.
	// Required: true
	ObjectName() *string
	SetObjectName(*string)
}

type imageSourceViaObjectStorageTupleDetailsAllOf1 struct {
	bucketNameField *string

	namespaceNameField *string

	objectNameField *string
}

func (m *imageSourceViaObjectStorageTupleDetailsAllOf1) BucketName() *string {
	return m.bucketNameField
}
func (m *imageSourceViaObjectStorageTupleDetailsAllOf1) SetBucketName(val *string) {
	m.bucketNameField = val
}

func (m *imageSourceViaObjectStorageTupleDetailsAllOf1) NamespaceName() *string {
	return m.namespaceNameField
}
func (m *imageSourceViaObjectStorageTupleDetailsAllOf1) SetNamespaceName(val *string) {
	m.namespaceNameField = val
}

func (m *imageSourceViaObjectStorageTupleDetailsAllOf1) ObjectName() *string {
	return m.objectNameField
}
func (m *imageSourceViaObjectStorageTupleDetailsAllOf1) SetObjectName(val *string) {
	m.objectNameField = val
}

// Validate validates this image source via object storage tuple details all of1
func (m *imageSourceViaObjectStorageTupleDetailsAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBucketName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNamespaceName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateObjectName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *imageSourceViaObjectStorageTupleDetailsAllOf1) validateBucketName(formats strfmt.Registry) error {

	if err := validate.Required("bucketName", "body", m.BucketName()); err != nil {
		return err
	}

	return nil
}

func (m *imageSourceViaObjectStorageTupleDetailsAllOf1) validateNamespaceName(formats strfmt.Registry) error {

	if err := validate.Required("namespaceName", "body", m.NamespaceName()); err != nil {
		return err
	}

	return nil
}

func (m *imageSourceViaObjectStorageTupleDetailsAllOf1) validateObjectName(formats strfmt.Registry) error {

	if err := validate.Required("objectName", "body", m.ObjectName()); err != nil {
		return err
	}

	return nil
}

// UnmarshalImageSourceViaObjectStorageTupleDetailsAllOf1Slice unmarshals polymorphic slices of ImageSourceViaObjectStorageTupleDetailsAllOf1
func UnmarshalImageSourceViaObjectStorageTupleDetailsAllOf1Slice(reader io.Reader, consumer runtime.Consumer) ([]ImageSourceViaObjectStorageTupleDetailsAllOf1, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []ImageSourceViaObjectStorageTupleDetailsAllOf1
	for _, element := range elements {
		obj, err := unmarshalImageSourceViaObjectStorageTupleDetailsAllOf1(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalImageSourceViaObjectStorageTupleDetailsAllOf1 unmarshals polymorphic ImageSourceViaObjectStorageTupleDetailsAllOf1
func UnmarshalImageSourceViaObjectStorageTupleDetailsAllOf1(reader io.Reader, consumer runtime.Consumer) (ImageSourceViaObjectStorageTupleDetailsAllOf1, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalImageSourceViaObjectStorageTupleDetailsAllOf1(data, consumer)
}

func unmarshalImageSourceViaObjectStorageTupleDetailsAllOf1(data []byte, consumer runtime.Consumer) (ImageSourceViaObjectStorageTupleDetailsAllOf1, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the objectStorageTuple property.
	var getType struct {
		ObjectStorageTuple string `json:"objectStorageTuple"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("objectStorageTuple", "body", getType.ObjectStorageTuple); err != nil {
		return nil, err
	}

	// The value of objectStorageTuple is used to determine which type to create and unmarshal the data into
	switch getType.ObjectStorageTuple {
	case "ImageSourceViaObjectStorageTupleDetails":
		var result ImageSourceViaObjectStorageTupleDetails
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "imageSourceViaObjectStorageTupleDetailsAllOf1":
		var result imageSourceViaObjectStorageTupleDetailsAllOf1
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid objectStorageTuple value: %q", getType.ObjectStorageTuple)

}
