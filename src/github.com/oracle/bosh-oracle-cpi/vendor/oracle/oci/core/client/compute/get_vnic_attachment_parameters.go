// Code generated by go-swagger; DO NOT EDIT.

package compute

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

// NewGetVnicAttachmentParams creates a new GetVnicAttachmentParams object
// with the default values initialized.
func NewGetVnicAttachmentParams() *GetVnicAttachmentParams {
	var ()
	return &GetVnicAttachmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVnicAttachmentParamsWithTimeout creates a new GetVnicAttachmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVnicAttachmentParamsWithTimeout(timeout time.Duration) *GetVnicAttachmentParams {
	var ()
	return &GetVnicAttachmentParams{

		timeout: timeout,
	}
}

// NewGetVnicAttachmentParamsWithContext creates a new GetVnicAttachmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVnicAttachmentParamsWithContext(ctx context.Context) *GetVnicAttachmentParams {
	var ()
	return &GetVnicAttachmentParams{

		Context: ctx,
	}
}

// NewGetVnicAttachmentParamsWithHTTPClient creates a new GetVnicAttachmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVnicAttachmentParamsWithHTTPClient(client *http.Client) *GetVnicAttachmentParams {
	var ()
	return &GetVnicAttachmentParams{
		HTTPClient: client,
	}
}

/*GetVnicAttachmentParams contains all the parameters to send to the API endpoint
for the get vnic attachment operation typically these are written to a http.Request
*/
type GetVnicAttachmentParams struct {

	/*VnicAttachmentID
	  The OCID of the VNIC attachment.

	*/
	VnicAttachmentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vnic attachment params
func (o *GetVnicAttachmentParams) WithTimeout(timeout time.Duration) *GetVnicAttachmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vnic attachment params
func (o *GetVnicAttachmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vnic attachment params
func (o *GetVnicAttachmentParams) WithContext(ctx context.Context) *GetVnicAttachmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vnic attachment params
func (o *GetVnicAttachmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vnic attachment params
func (o *GetVnicAttachmentParams) WithHTTPClient(client *http.Client) *GetVnicAttachmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vnic attachment params
func (o *GetVnicAttachmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVnicAttachmentID adds the vnicAttachmentID to the get vnic attachment params
func (o *GetVnicAttachmentParams) WithVnicAttachmentID(vnicAttachmentID string) *GetVnicAttachmentParams {
	o.SetVnicAttachmentID(vnicAttachmentID)
	return o
}

// SetVnicAttachmentID adds the vnicAttachmentId to the get vnic attachment params
func (o *GetVnicAttachmentParams) SetVnicAttachmentID(vnicAttachmentID string) {
	o.VnicAttachmentID = vnicAttachmentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVnicAttachmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param vnicAttachmentId
	if err := r.SetPathParam("vnicAttachmentId", o.VnicAttachmentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
