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

// CreateVcnReader is a Reader for the CreateVcn structure.
type CreateVcnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVcnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateVcnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateVcnBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateVcnUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateVcnNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateVcnConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateVcnInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateVcnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateVcnOK creates a CreateVcnOK with default headers values
func NewCreateVcnOK() *CreateVcnOK {
	return &CreateVcnOK{}
}

/*CreateVcnOK handles this case with default header values.

The VCN was created.
*/
type CreateVcnOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Vcn
}

func (o *CreateVcnOK) Error() string {
	return fmt.Sprintf("[POST /vcns][%d] createVcnOK  %+v", 200, o.Payload)
}

func (o *CreateVcnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Vcn)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVcnBadRequest creates a CreateVcnBadRequest with default headers values
func NewCreateVcnBadRequest() *CreateVcnBadRequest {
	return &CreateVcnBadRequest{}
}

/*CreateVcnBadRequest handles this case with default header values.

Bad Request
*/
type CreateVcnBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateVcnBadRequest) Error() string {
	return fmt.Sprintf("[POST /vcns][%d] createVcnBadRequest  %+v", 400, o.Payload)
}

func (o *CreateVcnBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVcnUnauthorized creates a CreateVcnUnauthorized with default headers values
func NewCreateVcnUnauthorized() *CreateVcnUnauthorized {
	return &CreateVcnUnauthorized{}
}

/*CreateVcnUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateVcnUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateVcnUnauthorized) Error() string {
	return fmt.Sprintf("[POST /vcns][%d] createVcnUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateVcnUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVcnNotFound creates a CreateVcnNotFound with default headers values
func NewCreateVcnNotFound() *CreateVcnNotFound {
	return &CreateVcnNotFound{}
}

/*CreateVcnNotFound handles this case with default header values.

Not Found
*/
type CreateVcnNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateVcnNotFound) Error() string {
	return fmt.Sprintf("[POST /vcns][%d] createVcnNotFound  %+v", 404, o.Payload)
}

func (o *CreateVcnNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVcnConflict creates a CreateVcnConflict with default headers values
func NewCreateVcnConflict() *CreateVcnConflict {
	return &CreateVcnConflict{}
}

/*CreateVcnConflict handles this case with default header values.

Conflict
*/
type CreateVcnConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateVcnConflict) Error() string {
	return fmt.Sprintf("[POST /vcns][%d] createVcnConflict  %+v", 409, o.Payload)
}

func (o *CreateVcnConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVcnInternalServerError creates a CreateVcnInternalServerError with default headers values
func NewCreateVcnInternalServerError() *CreateVcnInternalServerError {
	return &CreateVcnInternalServerError{}
}

/*CreateVcnInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateVcnInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreateVcnInternalServerError) Error() string {
	return fmt.Sprintf("[POST /vcns][%d] createVcnInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateVcnInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVcnDefault creates a CreateVcnDefault with default headers values
func NewCreateVcnDefault(code int) *CreateVcnDefault {
	return &CreateVcnDefault{
		_statusCode: code,
	}
}

/*CreateVcnDefault handles this case with default header values.

An error has occurred.
*/
type CreateVcnDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the create vcn default response
func (o *CreateVcnDefault) Code() int {
	return o._statusCode
}

func (o *CreateVcnDefault) Error() string {
	return fmt.Sprintf("[POST /vcns][%d] CreateVcn default  %+v", o._statusCode, o.Payload)
}

func (o *CreateVcnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
