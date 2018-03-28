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

// NewDeleteLocalPeeringGatewayParams creates a new DeleteLocalPeeringGatewayParams object
// with the default values initialized.
func NewDeleteLocalPeeringGatewayParams() *DeleteLocalPeeringGatewayParams {
	var ()
	return &DeleteLocalPeeringGatewayParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLocalPeeringGatewayParamsWithTimeout creates a new DeleteLocalPeeringGatewayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLocalPeeringGatewayParamsWithTimeout(timeout time.Duration) *DeleteLocalPeeringGatewayParams {
	var ()
	return &DeleteLocalPeeringGatewayParams{

		timeout: timeout,
	}
}

// NewDeleteLocalPeeringGatewayParamsWithContext creates a new DeleteLocalPeeringGatewayParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLocalPeeringGatewayParamsWithContext(ctx context.Context) *DeleteLocalPeeringGatewayParams {
	var ()
	return &DeleteLocalPeeringGatewayParams{

		Context: ctx,
	}
}

// NewDeleteLocalPeeringGatewayParamsWithHTTPClient creates a new DeleteLocalPeeringGatewayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLocalPeeringGatewayParamsWithHTTPClient(client *http.Client) *DeleteLocalPeeringGatewayParams {
	var ()
	return &DeleteLocalPeeringGatewayParams{
		HTTPClient: client,
	}
}

/*DeleteLocalPeeringGatewayParams contains all the parameters to send to the API endpoint
for the delete local peering gateway operation typically these are written to a http.Request
*/
type DeleteLocalPeeringGatewayParams struct {

	/*IfMatch
	  For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	will be updated or deleted only if the etag you provide matches the resource's current etag value.


	*/
	IfMatch *string
	/*LocalPeeringGatewayID
	  The OCID of the local peering gateway.

	*/
	LocalPeeringGatewayID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete local peering gateway params
func (o *DeleteLocalPeeringGatewayParams) WithTimeout(timeout time.Duration) *DeleteLocalPeeringGatewayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete local peering gateway params
func (o *DeleteLocalPeeringGatewayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete local peering gateway params
func (o *DeleteLocalPeeringGatewayParams) WithContext(ctx context.Context) *DeleteLocalPeeringGatewayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete local peering gateway params
func (o *DeleteLocalPeeringGatewayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete local peering gateway params
func (o *DeleteLocalPeeringGatewayParams) WithHTTPClient(client *http.Client) *DeleteLocalPeeringGatewayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete local peering gateway params
func (o *DeleteLocalPeeringGatewayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the delete local peering gateway params
func (o *DeleteLocalPeeringGatewayParams) WithIfMatch(ifMatch *string) *DeleteLocalPeeringGatewayParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the delete local peering gateway params
func (o *DeleteLocalPeeringGatewayParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithLocalPeeringGatewayID adds the localPeeringGatewayID to the delete local peering gateway params
func (o *DeleteLocalPeeringGatewayParams) WithLocalPeeringGatewayID(localPeeringGatewayID string) *DeleteLocalPeeringGatewayParams {
	o.SetLocalPeeringGatewayID(localPeeringGatewayID)
	return o
}

// SetLocalPeeringGatewayID adds the localPeeringGatewayId to the delete local peering gateway params
func (o *DeleteLocalPeeringGatewayParams) SetLocalPeeringGatewayID(localPeeringGatewayID string) {
	o.LocalPeeringGatewayID = localPeeringGatewayID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLocalPeeringGatewayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IfMatch != nil {

		// header param if-match
		if err := r.SetHeaderParam("if-match", *o.IfMatch); err != nil {
			return err
		}

	}

	// path param localPeeringGatewayId
	if err := r.SetPathParam("localPeeringGatewayId", o.LocalPeeringGatewayID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
