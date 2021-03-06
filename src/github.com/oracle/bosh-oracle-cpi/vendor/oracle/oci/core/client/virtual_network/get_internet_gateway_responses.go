package virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/oci/core/models"
)

// GetInternetGatewayReader is a Reader for the GetInternetGateway structure.
type GetInternetGatewayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInternetGatewayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInternetGatewayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetInternetGatewayUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetInternetGatewayNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetInternetGatewayInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetInternetGatewayDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetInternetGatewayOK creates a GetInternetGatewayOK with default headers values
func NewGetInternetGatewayOK() *GetInternetGatewayOK {
	return &GetInternetGatewayOK{}
}

/*GetInternetGatewayOK handles this case with default header values.

The Internet Gateway was retrieved.
*/
type GetInternetGatewayOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.InternetGateway
}

func (o *GetInternetGatewayOK) Error() string {
	return fmt.Sprintf("[GET /internetGateways/{igId}][%d] getInternetGatewayOK  %+v", 200, o.Payload)
}

func (o *GetInternetGatewayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.InternetGateway)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInternetGatewayUnauthorized creates a GetInternetGatewayUnauthorized with default headers values
func NewGetInternetGatewayUnauthorized() *GetInternetGatewayUnauthorized {
	return &GetInternetGatewayUnauthorized{}
}

/*GetInternetGatewayUnauthorized handles this case with default header values.

Unauthorized
*/
type GetInternetGatewayUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetInternetGatewayUnauthorized) Error() string {
	return fmt.Sprintf("[GET /internetGateways/{igId}][%d] getInternetGatewayUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInternetGatewayUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInternetGatewayNotFound creates a GetInternetGatewayNotFound with default headers values
func NewGetInternetGatewayNotFound() *GetInternetGatewayNotFound {
	return &GetInternetGatewayNotFound{}
}

/*GetInternetGatewayNotFound handles this case with default header values.

Not Found
*/
type GetInternetGatewayNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetInternetGatewayNotFound) Error() string {
	return fmt.Sprintf("[GET /internetGateways/{igId}][%d] getInternetGatewayNotFound  %+v", 404, o.Payload)
}

func (o *GetInternetGatewayNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInternetGatewayInternalServerError creates a GetInternetGatewayInternalServerError with default headers values
func NewGetInternetGatewayInternalServerError() *GetInternetGatewayInternalServerError {
	return &GetInternetGatewayInternalServerError{}
}

/*GetInternetGatewayInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetInternetGatewayInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetInternetGatewayInternalServerError) Error() string {
	return fmt.Sprintf("[GET /internetGateways/{igId}][%d] getInternetGatewayInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInternetGatewayInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInternetGatewayDefault creates a GetInternetGatewayDefault with default headers values
func NewGetInternetGatewayDefault(code int) *GetInternetGatewayDefault {
	return &GetInternetGatewayDefault{
		_statusCode: code,
	}
}

/*GetInternetGatewayDefault handles this case with default header values.

An error has occurred.
*/
type GetInternetGatewayDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the get internet gateway default response
func (o *GetInternetGatewayDefault) Code() int {
	return o._statusCode
}

func (o *GetInternetGatewayDefault) Error() string {
	return fmt.Sprintf("[GET /internetGateways/{igId}][%d] GetInternetGateway default  %+v", o._statusCode, o.Payload)
}

func (o *GetInternetGatewayDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
