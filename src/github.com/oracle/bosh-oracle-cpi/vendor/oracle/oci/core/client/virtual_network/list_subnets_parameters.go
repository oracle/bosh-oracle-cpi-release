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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListSubnetsParams creates a new ListSubnetsParams object
// with the default values initialized.
func NewListSubnetsParams() *ListSubnetsParams {
	var ()
	return &ListSubnetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSubnetsParamsWithTimeout creates a new ListSubnetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSubnetsParamsWithTimeout(timeout time.Duration) *ListSubnetsParams {
	var ()
	return &ListSubnetsParams{

		timeout: timeout,
	}
}

// NewListSubnetsParamsWithContext creates a new ListSubnetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSubnetsParamsWithContext(ctx context.Context) *ListSubnetsParams {
	var ()
	return &ListSubnetsParams{

		Context: ctx,
	}
}

// NewListSubnetsParamsWithHTTPClient creates a new ListSubnetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSubnetsParamsWithHTTPClient(client *http.Client) *ListSubnetsParams {
	var ()
	return &ListSubnetsParams{
		HTTPClient: client,
	}
}

/*ListSubnetsParams contains all the parameters to send to the API endpoint
for the list subnets operation typically these are written to a http.Request
*/
type ListSubnetsParams struct {

	/*CompartmentID
	  The OCID of the compartment.

	*/
	CompartmentID string
	/*DisplayName
	  A filter to return only resources that match the given display name exactly.


	*/
	DisplayName *string
	/*LifecycleState
	  A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.


	*/
	LifecycleState *string
	/*Limit
	  The maximum number of items to return in a paginated "List" call.

	Example: `500`


	*/
	Limit *int64
	/*Page
	  The value of the `opc-next-page` response header from the previous "List" call.


	*/
	Page *string
	/*SortBy
	  The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	sort order is case sensitive.

	**Note:** In general, some "List" operations (for example, `ListInstances`) let you
	optionally filter by Availability Domain if the scope of the resource type is within a
	single Availability Domain. If you call one of these "List" operations without specifying
	an Availability Domain, the resources are grouped by Availability Domain, then sorted.


	*/
	SortBy *string
	/*SortOrder
	  The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	is case sensitive.


	*/
	SortOrder *string
	/*VcnID
	  The OCID of the VCN.

	*/
	VcnID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list subnets params
func (o *ListSubnetsParams) WithTimeout(timeout time.Duration) *ListSubnetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list subnets params
func (o *ListSubnetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list subnets params
func (o *ListSubnetsParams) WithContext(ctx context.Context) *ListSubnetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list subnets params
func (o *ListSubnetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list subnets params
func (o *ListSubnetsParams) WithHTTPClient(client *http.Client) *ListSubnetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list subnets params
func (o *ListSubnetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompartmentID adds the compartmentID to the list subnets params
func (o *ListSubnetsParams) WithCompartmentID(compartmentID string) *ListSubnetsParams {
	o.SetCompartmentID(compartmentID)
	return o
}

// SetCompartmentID adds the compartmentId to the list subnets params
func (o *ListSubnetsParams) SetCompartmentID(compartmentID string) {
	o.CompartmentID = compartmentID
}

// WithDisplayName adds the displayName to the list subnets params
func (o *ListSubnetsParams) WithDisplayName(displayName *string) *ListSubnetsParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the list subnets params
func (o *ListSubnetsParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WithLifecycleState adds the lifecycleState to the list subnets params
func (o *ListSubnetsParams) WithLifecycleState(lifecycleState *string) *ListSubnetsParams {
	o.SetLifecycleState(lifecycleState)
	return o
}

// SetLifecycleState adds the lifecycleState to the list subnets params
func (o *ListSubnetsParams) SetLifecycleState(lifecycleState *string) {
	o.LifecycleState = lifecycleState
}

// WithLimit adds the limit to the list subnets params
func (o *ListSubnetsParams) WithLimit(limit *int64) *ListSubnetsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list subnets params
func (o *ListSubnetsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the list subnets params
func (o *ListSubnetsParams) WithPage(page *string) *ListSubnetsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list subnets params
func (o *ListSubnetsParams) SetPage(page *string) {
	o.Page = page
}

// WithSortBy adds the sortBy to the list subnets params
func (o *ListSubnetsParams) WithSortBy(sortBy *string) *ListSubnetsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the list subnets params
func (o *ListSubnetsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the list subnets params
func (o *ListSubnetsParams) WithSortOrder(sortOrder *string) *ListSubnetsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the list subnets params
func (o *ListSubnetsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithVcnID adds the vcnID to the list subnets params
func (o *ListSubnetsParams) WithVcnID(vcnID string) *ListSubnetsParams {
	o.SetVcnID(vcnID)
	return o
}

// SetVcnID adds the vcnId to the list subnets params
func (o *ListSubnetsParams) SetVcnID(vcnID string) {
	o.VcnID = vcnID
}

// WriteToRequest writes these params to a swagger request
func (o *ListSubnetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param compartmentId
	qrCompartmentID := o.CompartmentID
	qCompartmentID := qrCompartmentID
	if qCompartmentID != "" {
		if err := r.SetQueryParam("compartmentId", qCompartmentID); err != nil {
			return err
		}
	}

	if o.DisplayName != nil {

		// query param displayName
		var qrDisplayName string
		if o.DisplayName != nil {
			qrDisplayName = *o.DisplayName
		}
		qDisplayName := qrDisplayName
		if qDisplayName != "" {
			if err := r.SetQueryParam("displayName", qDisplayName); err != nil {
				return err
			}
		}

	}

	if o.LifecycleState != nil {

		// query param lifecycleState
		var qrLifecycleState string
		if o.LifecycleState != nil {
			qrLifecycleState = *o.LifecycleState
		}
		qLifecycleState := qrLifecycleState
		if qLifecycleState != "" {
			if err := r.SetQueryParam("lifecycleState", qLifecycleState); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage string
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := qrPage
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.SortBy != nil {

		// query param sortBy
		var qrSortBy string
		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {
			if err := r.SetQueryParam("sortBy", qSortBy); err != nil {
				return err
			}
		}

	}

	if o.SortOrder != nil {

		// query param sortOrder
		var qrSortOrder string
		if o.SortOrder != nil {
			qrSortOrder = *o.SortOrder
		}
		qSortOrder := qrSortOrder
		if qSortOrder != "" {
			if err := r.SetQueryParam("sortOrder", qSortOrder); err != nil {
				return err
			}
		}

	}

	// query param vcnId
	qrVcnID := o.VcnID
	qVcnID := qrVcnID
	if qVcnID != "" {
		if err := r.SetQueryParam("vcnId", qVcnID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
