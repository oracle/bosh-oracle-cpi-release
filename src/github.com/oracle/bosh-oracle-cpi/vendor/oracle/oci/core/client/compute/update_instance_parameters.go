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

	"oracle/oci/core/models"
)

// NewUpdateInstanceParams creates a new UpdateInstanceParams object
// with the default values initialized.
func NewUpdateInstanceParams() *UpdateInstanceParams {
	var ()
	return &UpdateInstanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateInstanceParamsWithTimeout creates a new UpdateInstanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateInstanceParamsWithTimeout(timeout time.Duration) *UpdateInstanceParams {
	var ()
	return &UpdateInstanceParams{

		timeout: timeout,
	}
}

// NewUpdateInstanceParamsWithContext creates a new UpdateInstanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateInstanceParamsWithContext(ctx context.Context) *UpdateInstanceParams {
	var ()
	return &UpdateInstanceParams{

		Context: ctx,
	}
}

// NewUpdateInstanceParamsWithHTTPClient creates a new UpdateInstanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateInstanceParamsWithHTTPClient(client *http.Client) *UpdateInstanceParams {
	var ()
	return &UpdateInstanceParams{
		HTTPClient: client,
	}
}

/*UpdateInstanceParams contains all the parameters to send to the API endpoint
for the update instance operation typically these are written to a http.Request
*/
type UpdateInstanceParams struct {

	/*UpdateInstanceDetails
	  Update instance fields

	*/
	UpdateInstanceDetails *models.UpdateInstanceDetails
	/*IfMatch
	  For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	will be updated or deleted only if the etag you provide matches the resource's current etag value.


	*/
	IfMatch *string
	/*InstanceID
	  The OCID of the instance.

	*/
	InstanceID string
	/*OpcRetryToken
	  A token that uniquely identifies a request so it can be retried in case of a timeout or
	server error without risk of executing that same action again. Retry tokens expire after 24
	hours, but can be invalidated before then due to conflicting operations (e.g., if a resource
	has been deleted and purged from the system, then a retry of the original creation request
	may be rejected).


	*/
	OpcRetryToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update instance params
func (o *UpdateInstanceParams) WithTimeout(timeout time.Duration) *UpdateInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update instance params
func (o *UpdateInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update instance params
func (o *UpdateInstanceParams) WithContext(ctx context.Context) *UpdateInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update instance params
func (o *UpdateInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update instance params
func (o *UpdateInstanceParams) WithHTTPClient(client *http.Client) *UpdateInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update instance params
func (o *UpdateInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateInstanceDetails adds the updateInstanceDetails to the update instance params
func (o *UpdateInstanceParams) WithUpdateInstanceDetails(updateInstanceDetails *models.UpdateInstanceDetails) *UpdateInstanceParams {
	o.SetUpdateInstanceDetails(updateInstanceDetails)
	return o
}

// SetUpdateInstanceDetails adds the updateInstanceDetails to the update instance params
func (o *UpdateInstanceParams) SetUpdateInstanceDetails(updateInstanceDetails *models.UpdateInstanceDetails) {
	o.UpdateInstanceDetails = updateInstanceDetails
}

// WithIfMatch adds the ifMatch to the update instance params
func (o *UpdateInstanceParams) WithIfMatch(ifMatch *string) *UpdateInstanceParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the update instance params
func (o *UpdateInstanceParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithInstanceID adds the instanceID to the update instance params
func (o *UpdateInstanceParams) WithInstanceID(instanceID string) *UpdateInstanceParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the update instance params
func (o *UpdateInstanceParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithOpcRetryToken adds the opcRetryToken to the update instance params
func (o *UpdateInstanceParams) WithOpcRetryToken(opcRetryToken *string) *UpdateInstanceParams {
	o.SetOpcRetryToken(opcRetryToken)
	return o
}

// SetOpcRetryToken adds the opcRetryToken to the update instance params
func (o *UpdateInstanceParams) SetOpcRetryToken(opcRetryToken *string) {
	o.OpcRetryToken = opcRetryToken
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UpdateInstanceDetails == nil {
		o.UpdateInstanceDetails = new(models.UpdateInstanceDetails)
	}

	if err := r.SetBodyParam(o.UpdateInstanceDetails); err != nil {
		return err
	}

	if o.IfMatch != nil {

		// header param if-match
		if err := r.SetHeaderParam("if-match", *o.IfMatch); err != nil {
			return err
		}

	}

	// path param instanceId
	if err := r.SetPathParam("instanceId", o.InstanceID); err != nil {
		return err
	}

	if o.OpcRetryToken != nil {

		// header param opc-retry-token
		if err := r.SetHeaderParam("opc-retry-token", *o.OpcRetryToken); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
