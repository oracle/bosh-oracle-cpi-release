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

	models "oracle/oci/core/models"
)

// NewUpdatePublicIPParams creates a new UpdatePublicIPParams object
// with the default values initialized.
func NewUpdatePublicIPParams() *UpdatePublicIPParams {
	var ()
	return &UpdatePublicIPParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePublicIPParamsWithTimeout creates a new UpdatePublicIPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePublicIPParamsWithTimeout(timeout time.Duration) *UpdatePublicIPParams {
	var ()
	return &UpdatePublicIPParams{

		timeout: timeout,
	}
}

// NewUpdatePublicIPParamsWithContext creates a new UpdatePublicIPParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePublicIPParamsWithContext(ctx context.Context) *UpdatePublicIPParams {
	var ()
	return &UpdatePublicIPParams{

		Context: ctx,
	}
}

// NewUpdatePublicIPParamsWithHTTPClient creates a new UpdatePublicIPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePublicIPParamsWithHTTPClient(client *http.Client) *UpdatePublicIPParams {
	var ()
	return &UpdatePublicIPParams{
		HTTPClient: client,
	}
}

/*UpdatePublicIPParams contains all the parameters to send to the API endpoint
for the update public Ip operation typically these are written to a http.Request
*/
type UpdatePublicIPParams struct {

	/*UpdatePublicIPDetails
	  Public IP details.

	*/
	UpdatePublicIPDetails *models.UpdatePublicIPDetails
	/*IfMatch
	  For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	will be updated or deleted only if the etag you provide matches the resource's current etag value.


	*/
	IfMatch *string
	/*PublicIPID
	  The OCID of the public IP.

	*/
	PublicIPID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update public Ip params
func (o *UpdatePublicIPParams) WithTimeout(timeout time.Duration) *UpdatePublicIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update public Ip params
func (o *UpdatePublicIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update public Ip params
func (o *UpdatePublicIPParams) WithContext(ctx context.Context) *UpdatePublicIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update public Ip params
func (o *UpdatePublicIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update public Ip params
func (o *UpdatePublicIPParams) WithHTTPClient(client *http.Client) *UpdatePublicIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update public Ip params
func (o *UpdatePublicIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdatePublicIPDetails adds the updatePublicIPDetails to the update public Ip params
func (o *UpdatePublicIPParams) WithUpdatePublicIPDetails(updatePublicIPDetails *models.UpdatePublicIPDetails) *UpdatePublicIPParams {
	o.SetUpdatePublicIPDetails(updatePublicIPDetails)
	return o
}

// SetUpdatePublicIPDetails adds the updatePublicIpDetails to the update public Ip params
func (o *UpdatePublicIPParams) SetUpdatePublicIPDetails(updatePublicIPDetails *models.UpdatePublicIPDetails) {
	o.UpdatePublicIPDetails = updatePublicIPDetails
}

// WithIfMatch adds the ifMatch to the update public Ip params
func (o *UpdatePublicIPParams) WithIfMatch(ifMatch *string) *UpdatePublicIPParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the update public Ip params
func (o *UpdatePublicIPParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithPublicIPID adds the publicIPID to the update public Ip params
func (o *UpdatePublicIPParams) WithPublicIPID(publicIPID string) *UpdatePublicIPParams {
	o.SetPublicIPID(publicIPID)
	return o
}

// SetPublicIPID adds the publicIpId to the update public Ip params
func (o *UpdatePublicIPParams) SetPublicIPID(publicIPID string) {
	o.PublicIPID = publicIPID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePublicIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UpdatePublicIPDetails != nil {
		if err := r.SetBodyParam(o.UpdatePublicIPDetails); err != nil {
			return err
		}
	}

	if o.IfMatch != nil {

		// header param if-match
		if err := r.SetHeaderParam("if-match", *o.IfMatch); err != nil {
			return err
		}

	}

	// path param publicIpId
	if err := r.SetPathParam("publicIpId", o.PublicIPID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
