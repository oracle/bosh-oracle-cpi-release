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

// NewDeleteCrossConnectParams creates a new DeleteCrossConnectParams object
// with the default values initialized.
func NewDeleteCrossConnectParams() *DeleteCrossConnectParams {
	var ()
	return &DeleteCrossConnectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCrossConnectParamsWithTimeout creates a new DeleteCrossConnectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCrossConnectParamsWithTimeout(timeout time.Duration) *DeleteCrossConnectParams {
	var ()
	return &DeleteCrossConnectParams{

		timeout: timeout,
	}
}

// NewDeleteCrossConnectParamsWithContext creates a new DeleteCrossConnectParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCrossConnectParamsWithContext(ctx context.Context) *DeleteCrossConnectParams {
	var ()
	return &DeleteCrossConnectParams{

		Context: ctx,
	}
}

// NewDeleteCrossConnectParamsWithHTTPClient creates a new DeleteCrossConnectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCrossConnectParamsWithHTTPClient(client *http.Client) *DeleteCrossConnectParams {
	var ()
	return &DeleteCrossConnectParams{
		HTTPClient: client,
	}
}

/*DeleteCrossConnectParams contains all the parameters to send to the API endpoint
for the delete cross connect operation typically these are written to a http.Request
*/
type DeleteCrossConnectParams struct {

	/*CrossConnectID
	  The OCID of the cross-connect.

	*/
	CrossConnectID string
	/*IfMatch
	  For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	will be updated or deleted only if the etag you provide matches the resource's current etag value.


	*/
	IfMatch *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cross connect params
func (o *DeleteCrossConnectParams) WithTimeout(timeout time.Duration) *DeleteCrossConnectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cross connect params
func (o *DeleteCrossConnectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cross connect params
func (o *DeleteCrossConnectParams) WithContext(ctx context.Context) *DeleteCrossConnectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cross connect params
func (o *DeleteCrossConnectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cross connect params
func (o *DeleteCrossConnectParams) WithHTTPClient(client *http.Client) *DeleteCrossConnectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cross connect params
func (o *DeleteCrossConnectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrossConnectID adds the crossConnectID to the delete cross connect params
func (o *DeleteCrossConnectParams) WithCrossConnectID(crossConnectID string) *DeleteCrossConnectParams {
	o.SetCrossConnectID(crossConnectID)
	return o
}

// SetCrossConnectID adds the crossConnectId to the delete cross connect params
func (o *DeleteCrossConnectParams) SetCrossConnectID(crossConnectID string) {
	o.CrossConnectID = crossConnectID
}

// WithIfMatch adds the ifMatch to the delete cross connect params
func (o *DeleteCrossConnectParams) WithIfMatch(ifMatch *string) *DeleteCrossConnectParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the delete cross connect params
func (o *DeleteCrossConnectParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCrossConnectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crossConnectId
	if err := r.SetPathParam("crossConnectId", o.CrossConnectID); err != nil {
		return err
	}

	if o.IfMatch != nil {

		// header param if-match
		if err := r.SetHeaderParam("if-match", *o.IfMatch); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
