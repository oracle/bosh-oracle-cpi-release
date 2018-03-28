// Code generated by go-swagger; DO NOT EDIT.

package blockstorage

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

// NewGetBootVolumeParams creates a new GetBootVolumeParams object
// with the default values initialized.
func NewGetBootVolumeParams() *GetBootVolumeParams {
	var ()
	return &GetBootVolumeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBootVolumeParamsWithTimeout creates a new GetBootVolumeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBootVolumeParamsWithTimeout(timeout time.Duration) *GetBootVolumeParams {
	var ()
	return &GetBootVolumeParams{

		timeout: timeout,
	}
}

// NewGetBootVolumeParamsWithContext creates a new GetBootVolumeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBootVolumeParamsWithContext(ctx context.Context) *GetBootVolumeParams {
	var ()
	return &GetBootVolumeParams{

		Context: ctx,
	}
}

// NewGetBootVolumeParamsWithHTTPClient creates a new GetBootVolumeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBootVolumeParamsWithHTTPClient(client *http.Client) *GetBootVolumeParams {
	var ()
	return &GetBootVolumeParams{
		HTTPClient: client,
	}
}

/*GetBootVolumeParams contains all the parameters to send to the API endpoint
for the get boot volume operation typically these are written to a http.Request
*/
type GetBootVolumeParams struct {

	/*BootVolumeID
	  The OCID of the boot volume.

	*/
	BootVolumeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get boot volume params
func (o *GetBootVolumeParams) WithTimeout(timeout time.Duration) *GetBootVolumeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get boot volume params
func (o *GetBootVolumeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get boot volume params
func (o *GetBootVolumeParams) WithContext(ctx context.Context) *GetBootVolumeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get boot volume params
func (o *GetBootVolumeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get boot volume params
func (o *GetBootVolumeParams) WithHTTPClient(client *http.Client) *GetBootVolumeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get boot volume params
func (o *GetBootVolumeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBootVolumeID adds the bootVolumeID to the get boot volume params
func (o *GetBootVolumeParams) WithBootVolumeID(bootVolumeID string) *GetBootVolumeParams {
	o.SetBootVolumeID(bootVolumeID)
	return o
}

// SetBootVolumeID adds the bootVolumeId to the get boot volume params
func (o *GetBootVolumeParams) SetBootVolumeID(bootVolumeID string) {
	o.BootVolumeID = bootVolumeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBootVolumeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bootVolumeId
	if err := r.SetPathParam("bootVolumeId", o.BootVolumeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
