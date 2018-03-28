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

// NewGetPublicIPByPrivateIPIDParams creates a new GetPublicIPByPrivateIPIDParams object
// with the default values initialized.
func NewGetPublicIPByPrivateIPIDParams() *GetPublicIPByPrivateIPIDParams {
	var ()
	return &GetPublicIPByPrivateIPIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicIPByPrivateIPIDParamsWithTimeout creates a new GetPublicIPByPrivateIPIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicIPByPrivateIPIDParamsWithTimeout(timeout time.Duration) *GetPublicIPByPrivateIPIDParams {
	var ()
	return &GetPublicIPByPrivateIPIDParams{

		timeout: timeout,
	}
}

// NewGetPublicIPByPrivateIPIDParamsWithContext creates a new GetPublicIPByPrivateIPIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicIPByPrivateIPIDParamsWithContext(ctx context.Context) *GetPublicIPByPrivateIPIDParams {
	var ()
	return &GetPublicIPByPrivateIPIDParams{

		Context: ctx,
	}
}

// NewGetPublicIPByPrivateIPIDParamsWithHTTPClient creates a new GetPublicIPByPrivateIPIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicIPByPrivateIPIDParamsWithHTTPClient(client *http.Client) *GetPublicIPByPrivateIPIDParams {
	var ()
	return &GetPublicIPByPrivateIPIDParams{
		HTTPClient: client,
	}
}

/*GetPublicIPByPrivateIPIDParams contains all the parameters to send to the API endpoint
for the get public Ip by private Ip Id operation typically these are written to a http.Request
*/
type GetPublicIPByPrivateIPIDParams struct {

	/*GetPublicIPByPrivateIPIDDetails
	  Private IP details for fetching the public IP.

	*/
	GetPublicIPByPrivateIPIDDetails *models.GetPublicIPByPrivateIPIDDetails

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public Ip by private Ip Id params
func (o *GetPublicIPByPrivateIPIDParams) WithTimeout(timeout time.Duration) *GetPublicIPByPrivateIPIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public Ip by private Ip Id params
func (o *GetPublicIPByPrivateIPIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public Ip by private Ip Id params
func (o *GetPublicIPByPrivateIPIDParams) WithContext(ctx context.Context) *GetPublicIPByPrivateIPIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public Ip by private Ip Id params
func (o *GetPublicIPByPrivateIPIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public Ip by private Ip Id params
func (o *GetPublicIPByPrivateIPIDParams) WithHTTPClient(client *http.Client) *GetPublicIPByPrivateIPIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public Ip by private Ip Id params
func (o *GetPublicIPByPrivateIPIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGetPublicIPByPrivateIPIDDetails adds the getPublicIPByPrivateIPIDDetails to the get public Ip by private Ip Id params
func (o *GetPublicIPByPrivateIPIDParams) WithGetPublicIPByPrivateIPIDDetails(getPublicIPByPrivateIPIDDetails *models.GetPublicIPByPrivateIPIDDetails) *GetPublicIPByPrivateIPIDParams {
	o.SetGetPublicIPByPrivateIPIDDetails(getPublicIPByPrivateIPIDDetails)
	return o
}

// SetGetPublicIPByPrivateIPIDDetails adds the getPublicIpByPrivateIpIdDetails to the get public Ip by private Ip Id params
func (o *GetPublicIPByPrivateIPIDParams) SetGetPublicIPByPrivateIPIDDetails(getPublicIPByPrivateIPIDDetails *models.GetPublicIPByPrivateIPIDDetails) {
	o.GetPublicIPByPrivateIPIDDetails = getPublicIPByPrivateIPIDDetails
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicIPByPrivateIPIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.GetPublicIPByPrivateIPIDDetails != nil {
		if err := r.SetBodyParam(o.GetPublicIPByPrivateIPIDDetails); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
