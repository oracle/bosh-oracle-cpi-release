// Code generated by go-swagger; DO NOT EDIT.

package virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetVirtualCircuitParams creates a new GetVirtualCircuitParams object
// with the default values initialized.
func NewGetVirtualCircuitParams() *GetVirtualCircuitParams {
	var ()
	return &GetVirtualCircuitParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVirtualCircuitParamsWithTimeout creates a new GetVirtualCircuitParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVirtualCircuitParamsWithTimeout(timeout time.Duration) *GetVirtualCircuitParams {
	var ()
	return &GetVirtualCircuitParams{

		timeout: timeout,
	}
}

// NewGetVirtualCircuitParamsWithContext creates a new GetVirtualCircuitParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVirtualCircuitParamsWithContext(ctx context.Context) *GetVirtualCircuitParams {
	var ()
	return &GetVirtualCircuitParams{

		Context: ctx,
	}
}

// NewGetVirtualCircuitParamsWithHTTPClient creates a new GetVirtualCircuitParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVirtualCircuitParamsWithHTTPClient(client *http.Client) *GetVirtualCircuitParams {
	var ()
	return &GetVirtualCircuitParams{
		HTTPClient: client,
	}
}

/*GetVirtualCircuitParams contains all the parameters to send to the API endpoint
for the get virtual circuit operation typically these are written to a http.Request
*/
type GetVirtualCircuitParams struct {

	/*VirtualCircuitID
	  The OCID of the virtual circuit.

	*/
	VirtualCircuitID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get virtual circuit params
func (o *GetVirtualCircuitParams) WithTimeout(timeout time.Duration) *GetVirtualCircuitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get virtual circuit params
func (o *GetVirtualCircuitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get virtual circuit params
func (o *GetVirtualCircuitParams) WithContext(ctx context.Context) *GetVirtualCircuitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get virtual circuit params
func (o *GetVirtualCircuitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get virtual circuit params
func (o *GetVirtualCircuitParams) WithHTTPClient(client *http.Client) *GetVirtualCircuitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get virtual circuit params
func (o *GetVirtualCircuitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVirtualCircuitID adds the virtualCircuitID to the get virtual circuit params
func (o *GetVirtualCircuitParams) WithVirtualCircuitID(virtualCircuitID string) *GetVirtualCircuitParams {
	o.SetVirtualCircuitID(virtualCircuitID)
	return o
}

// SetVirtualCircuitID adds the virtualCircuitId to the get virtual circuit params
func (o *GetVirtualCircuitParams) SetVirtualCircuitID(virtualCircuitID string) {
	o.VirtualCircuitID = virtualCircuitID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVirtualCircuitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param virtualCircuitId
	if err := r.SetPathParam("virtualCircuitId", o.VirtualCircuitID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
