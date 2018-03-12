package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/oci/core/models"
)

// TerminateInstanceReader is a Reader for the TerminateInstance structure.
type TerminateInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TerminateInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewTerminateInstanceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewTerminateInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTerminateInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewTerminateInstancePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewTerminateInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewTerminateInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTerminateInstanceNoContent creates a TerminateInstanceNoContent with default headers values
func NewTerminateInstanceNoContent() *TerminateInstanceNoContent {
	return &TerminateInstanceNoContent{}
}

/*TerminateInstanceNoContent handles this case with default header values.

The instance is being terminated.
*/
type TerminateInstanceNoContent struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string
}

func (o *TerminateInstanceNoContent) Error() string {
	return fmt.Sprintf("[DELETE /instances/{instanceId}][%d] terminateInstanceNoContent ", 204)
}

func (o *TerminateInstanceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	return nil
}

// NewTerminateInstanceUnauthorized creates a TerminateInstanceUnauthorized with default headers values
func NewTerminateInstanceUnauthorized() *TerminateInstanceUnauthorized {
	return &TerminateInstanceUnauthorized{}
}

/*TerminateInstanceUnauthorized handles this case with default header values.

Unauthorized
*/
type TerminateInstanceUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *TerminateInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /instances/{instanceId}][%d] terminateInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *TerminateInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTerminateInstanceNotFound creates a TerminateInstanceNotFound with default headers values
func NewTerminateInstanceNotFound() *TerminateInstanceNotFound {
	return &TerminateInstanceNotFound{}
}

/*TerminateInstanceNotFound handles this case with default header values.

Not Found
*/
type TerminateInstanceNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *TerminateInstanceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /instances/{instanceId}][%d] terminateInstanceNotFound  %+v", 404, o.Payload)
}

func (o *TerminateInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTerminateInstancePreconditionFailed creates a TerminateInstancePreconditionFailed with default headers values
func NewTerminateInstancePreconditionFailed() *TerminateInstancePreconditionFailed {
	return &TerminateInstancePreconditionFailed{}
}

/*TerminateInstancePreconditionFailed handles this case with default header values.

Precondition Failed
*/
type TerminateInstancePreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *TerminateInstancePreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /instances/{instanceId}][%d] terminateInstancePreconditionFailed  %+v", 412, o.Payload)
}

func (o *TerminateInstancePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTerminateInstanceInternalServerError creates a TerminateInstanceInternalServerError with default headers values
func NewTerminateInstanceInternalServerError() *TerminateInstanceInternalServerError {
	return &TerminateInstanceInternalServerError{}
}

/*TerminateInstanceInternalServerError handles this case with default header values.

Internal Server Error
*/
type TerminateInstanceInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *TerminateInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /instances/{instanceId}][%d] terminateInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *TerminateInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTerminateInstanceDefault creates a TerminateInstanceDefault with default headers values
func NewTerminateInstanceDefault(code int) *TerminateInstanceDefault {
	return &TerminateInstanceDefault{
		_statusCode: code,
	}
}

/*TerminateInstanceDefault handles this case with default header values.

An error has occurred.
*/
type TerminateInstanceDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the terminate instance default response
func (o *TerminateInstanceDefault) Code() int {
	return o._statusCode
}

func (o *TerminateInstanceDefault) Error() string {
	return fmt.Sprintf("[DELETE /instances/{instanceId}][%d] TerminateInstance default  %+v", o._statusCode, o.Payload)
}

func (o *TerminateInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}