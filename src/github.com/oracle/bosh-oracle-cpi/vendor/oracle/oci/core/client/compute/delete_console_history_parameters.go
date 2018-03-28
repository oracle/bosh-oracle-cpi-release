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

// NewDeleteConsoleHistoryParams creates a new DeleteConsoleHistoryParams object
// with the default values initialized.
func NewDeleteConsoleHistoryParams() *DeleteConsoleHistoryParams {
	var ()
	return &DeleteConsoleHistoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteConsoleHistoryParamsWithTimeout creates a new DeleteConsoleHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteConsoleHistoryParamsWithTimeout(timeout time.Duration) *DeleteConsoleHistoryParams {
	var ()
	return &DeleteConsoleHistoryParams{

		timeout: timeout,
	}
}

// NewDeleteConsoleHistoryParamsWithContext creates a new DeleteConsoleHistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteConsoleHistoryParamsWithContext(ctx context.Context) *DeleteConsoleHistoryParams {
	var ()
	return &DeleteConsoleHistoryParams{

		Context: ctx,
	}
}

// NewDeleteConsoleHistoryParamsWithHTTPClient creates a new DeleteConsoleHistoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteConsoleHistoryParamsWithHTTPClient(client *http.Client) *DeleteConsoleHistoryParams {
	var ()
	return &DeleteConsoleHistoryParams{
		HTTPClient: client,
	}
}

/*DeleteConsoleHistoryParams contains all the parameters to send to the API endpoint
for the delete console history operation typically these are written to a http.Request
*/
type DeleteConsoleHistoryParams struct {

	/*IfMatch
	  For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	will be updated or deleted only if the etag you provide matches the resource's current etag value.


	*/
	IfMatch *string
	/*InstanceConsoleHistoryID
	  The OCID of the console history.

	*/
	InstanceConsoleHistoryID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete console history params
func (o *DeleteConsoleHistoryParams) WithTimeout(timeout time.Duration) *DeleteConsoleHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete console history params
func (o *DeleteConsoleHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete console history params
func (o *DeleteConsoleHistoryParams) WithContext(ctx context.Context) *DeleteConsoleHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete console history params
func (o *DeleteConsoleHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete console history params
func (o *DeleteConsoleHistoryParams) WithHTTPClient(client *http.Client) *DeleteConsoleHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete console history params
func (o *DeleteConsoleHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the delete console history params
func (o *DeleteConsoleHistoryParams) WithIfMatch(ifMatch *string) *DeleteConsoleHistoryParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the delete console history params
func (o *DeleteConsoleHistoryParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithInstanceConsoleHistoryID adds the instanceConsoleHistoryID to the delete console history params
func (o *DeleteConsoleHistoryParams) WithInstanceConsoleHistoryID(instanceConsoleHistoryID string) *DeleteConsoleHistoryParams {
	o.SetInstanceConsoleHistoryID(instanceConsoleHistoryID)
	return o
}

// SetInstanceConsoleHistoryID adds the instanceConsoleHistoryId to the delete console history params
func (o *DeleteConsoleHistoryParams) SetInstanceConsoleHistoryID(instanceConsoleHistoryID string) {
	o.InstanceConsoleHistoryID = instanceConsoleHistoryID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteConsoleHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param instanceConsoleHistoryId
	if err := r.SetPathParam("instanceConsoleHistoryId", o.InstanceConsoleHistoryID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
