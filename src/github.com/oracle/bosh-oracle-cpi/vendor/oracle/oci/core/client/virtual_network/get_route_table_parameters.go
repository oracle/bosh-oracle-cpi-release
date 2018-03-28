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

// NewGetRouteTableParams creates a new GetRouteTableParams object
// with the default values initialized.
func NewGetRouteTableParams() *GetRouteTableParams {
	var ()
	return &GetRouteTableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRouteTableParamsWithTimeout creates a new GetRouteTableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRouteTableParamsWithTimeout(timeout time.Duration) *GetRouteTableParams {
	var ()
	return &GetRouteTableParams{

		timeout: timeout,
	}
}

// NewGetRouteTableParamsWithContext creates a new GetRouteTableParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRouteTableParamsWithContext(ctx context.Context) *GetRouteTableParams {
	var ()
	return &GetRouteTableParams{

		Context: ctx,
	}
}

// NewGetRouteTableParamsWithHTTPClient creates a new GetRouteTableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRouteTableParamsWithHTTPClient(client *http.Client) *GetRouteTableParams {
	var ()
	return &GetRouteTableParams{
		HTTPClient: client,
	}
}

/*GetRouteTableParams contains all the parameters to send to the API endpoint
for the get route table operation typically these are written to a http.Request
*/
type GetRouteTableParams struct {

	/*RtID
	  The OCID of the route table.

	*/
	RtID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get route table params
func (o *GetRouteTableParams) WithTimeout(timeout time.Duration) *GetRouteTableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get route table params
func (o *GetRouteTableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get route table params
func (o *GetRouteTableParams) WithContext(ctx context.Context) *GetRouteTableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get route table params
func (o *GetRouteTableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get route table params
func (o *GetRouteTableParams) WithHTTPClient(client *http.Client) *GetRouteTableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get route table params
func (o *GetRouteTableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRtID adds the rtID to the get route table params
func (o *GetRouteTableParams) WithRtID(rtID string) *GetRouteTableParams {
	o.SetRtID(rtID)
	return o
}

// SetRtID adds the rtId to the get route table params
func (o *GetRouteTableParams) SetRtID(rtID string) {
	o.RtID = rtID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRouteTableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param rtId
	if err := r.SetPathParam("rtId", o.RtID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
