// Code generated by go-swagger; DO NOT EDIT.

package virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "oracle/oci/core/models"
)

// ListAllowedPeerRegionsForRemotePeeringReader is a Reader for the ListAllowedPeerRegionsForRemotePeering structure.
type ListAllowedPeerRegionsForRemotePeeringReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllowedPeerRegionsForRemotePeeringReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAllowedPeerRegionsForRemotePeeringOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListAllowedPeerRegionsForRemotePeeringUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListAllowedPeerRegionsForRemotePeeringInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListAllowedPeerRegionsForRemotePeeringDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAllowedPeerRegionsForRemotePeeringOK creates a ListAllowedPeerRegionsForRemotePeeringOK with default headers values
func NewListAllowedPeerRegionsForRemotePeeringOK() *ListAllowedPeerRegionsForRemotePeeringOK {
	return &ListAllowedPeerRegionsForRemotePeeringOK{}
}

/*ListAllowedPeerRegionsForRemotePeeringOK handles this case with default header values.

The list is retrieved.
*/
type ListAllowedPeerRegionsForRemotePeeringOK struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload []*models.PeerRegionForRemotePeering
}

func (o *ListAllowedPeerRegionsForRemotePeeringOK) Error() string {
	return fmt.Sprintf("[GET /allowedPeerRegionsForRemotePeering][%d] listAllowedPeerRegionsForRemotePeeringOK  %+v", 200, o.Payload)
}

func (o *ListAllowedPeerRegionsForRemotePeeringOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllowedPeerRegionsForRemotePeeringUnauthorized creates a ListAllowedPeerRegionsForRemotePeeringUnauthorized with default headers values
func NewListAllowedPeerRegionsForRemotePeeringUnauthorized() *ListAllowedPeerRegionsForRemotePeeringUnauthorized {
	return &ListAllowedPeerRegionsForRemotePeeringUnauthorized{}
}

/*ListAllowedPeerRegionsForRemotePeeringUnauthorized handles this case with default header values.

Unauthorized
*/
type ListAllowedPeerRegionsForRemotePeeringUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListAllowedPeerRegionsForRemotePeeringUnauthorized) Error() string {
	return fmt.Sprintf("[GET /allowedPeerRegionsForRemotePeering][%d] listAllowedPeerRegionsForRemotePeeringUnauthorized  %+v", 401, o.Payload)
}

func (o *ListAllowedPeerRegionsForRemotePeeringUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllowedPeerRegionsForRemotePeeringInternalServerError creates a ListAllowedPeerRegionsForRemotePeeringInternalServerError with default headers values
func NewListAllowedPeerRegionsForRemotePeeringInternalServerError() *ListAllowedPeerRegionsForRemotePeeringInternalServerError {
	return &ListAllowedPeerRegionsForRemotePeeringInternalServerError{}
}

/*ListAllowedPeerRegionsForRemotePeeringInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListAllowedPeerRegionsForRemotePeeringInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListAllowedPeerRegionsForRemotePeeringInternalServerError) Error() string {
	return fmt.Sprintf("[GET /allowedPeerRegionsForRemotePeering][%d] listAllowedPeerRegionsForRemotePeeringInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAllowedPeerRegionsForRemotePeeringInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllowedPeerRegionsForRemotePeeringDefault creates a ListAllowedPeerRegionsForRemotePeeringDefault with default headers values
func NewListAllowedPeerRegionsForRemotePeeringDefault(code int) *ListAllowedPeerRegionsForRemotePeeringDefault {
	return &ListAllowedPeerRegionsForRemotePeeringDefault{
		_statusCode: code,
	}
}

/*ListAllowedPeerRegionsForRemotePeeringDefault handles this case with default header values.

An error has occurred.
*/
type ListAllowedPeerRegionsForRemotePeeringDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the list allowed peer regions for remote peering default response
func (o *ListAllowedPeerRegionsForRemotePeeringDefault) Code() int {
	return o._statusCode
}

func (o *ListAllowedPeerRegionsForRemotePeeringDefault) Error() string {
	return fmt.Sprintf("[GET /allowedPeerRegionsForRemotePeering][%d] ListAllowedPeerRegionsForRemotePeering default  %+v", o._statusCode, o.Payload)
}

func (o *ListAllowedPeerRegionsForRemotePeeringDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
