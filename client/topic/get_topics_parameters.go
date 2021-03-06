package topic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewGetTopicsParams creates a new GetTopicsParams object
// with the default values initialized.
func NewGetTopicsParams() *GetTopicsParams {
	return &GetTopicsParams{}
}

/*GetTopicsParams contains all the parameters to send to the API endpoint
for the get topics operation typically these are written to a http.Request
*/
type GetTopicsParams struct {

	/*Authorid
	  主题发布者ID

	*/
	Authorid int64
	/*Category
	  主题类别

	*/
	Category int64
	/*Page
	  当前页码

	*/
	Page int64
	/*Pagesize
	  每页项数

	*/
	Pagesize int64
	/*Resolved
	  是否已解决问题优先

	*/
	Resolved bool
	/*Startupid
	  创业公司ID

	*/
	Startupid int64
	/*Status
	  主题状态

	*/
	Status int64
	/*Unresolved
	  是否未解决问题优先

	*/
	Unresolved bool
	/*Wonderful
	  精彩问答数目, 若有此项则其他查询条件无效, 且只返回一页

	*/
	Wonderful int64
}

// WithAuthorid adds the authorid to the get topics params
func (o *GetTopicsParams) WithAuthorid(authorid int64) *GetTopicsParams {
	o.Authorid = authorid
	return o
}

// WithCategory adds the category to the get topics params
func (o *GetTopicsParams) WithCategory(category int64) *GetTopicsParams {
	o.Category = category
	return o
}

// WithPage adds the page to the get topics params
func (o *GetTopicsParams) WithPage(page int64) *GetTopicsParams {
	o.Page = page
	return o
}

// WithPagesize adds the pagesize to the get topics params
func (o *GetTopicsParams) WithPagesize(pagesize int64) *GetTopicsParams {
	o.Pagesize = pagesize
	return o
}

// WithResolved adds the resolved to the get topics params
func (o *GetTopicsParams) WithResolved(resolved bool) *GetTopicsParams {
	o.Resolved = resolved
	return o
}

// WithStartupid adds the startupid to the get topics params
func (o *GetTopicsParams) WithStartupid(startupid int64) *GetTopicsParams {
	o.Startupid = startupid
	return o
}

// WithStatus adds the status to the get topics params
func (o *GetTopicsParams) WithStatus(status int64) *GetTopicsParams {
	o.Status = status
	return o
}

// WithUnresolved adds the unresolved to the get topics params
func (o *GetTopicsParams) WithUnresolved(unresolved bool) *GetTopicsParams {
	o.Unresolved = unresolved
	return o
}

// WithWonderful adds the wonderful to the get topics params
func (o *GetTopicsParams) WithWonderful(wonderful int64) *GetTopicsParams {
	o.Wonderful = wonderful
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetTopicsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// query param authorid
	qrAuthorid := o.Authorid
	qAuthorid := swag.FormatInt64(qrAuthorid)
	if qAuthorid != "" && qAuthorid != "-1" {
		if err := r.SetQueryParam("authorid", qAuthorid); err != nil {
			return err
		}
	}

	// query array param authorid

	// query param category
	qrCategory := o.Category
	qCategory := swag.FormatInt64(qrCategory)
	if qCategory != "" && qCategory != "-1" {
		if err := r.SetQueryParam("category", qCategory); err != nil {
			return err
		}
	}

	// query array param category

	// query param page
	qrPage := o.Page
	qPage := swag.FormatInt64(qrPage)
	if qPage != "" && qPage != "-1" {
		if err := r.SetQueryParam("page", qPage); err != nil {
			return err
		}
	}

	// query array param page

	// query param pagesize
	qrPagesize := o.Pagesize
	qPagesize := swag.FormatInt64(qrPagesize)
	if qPagesize != "" && qPagesize != "-1" {
		if err := r.SetQueryParam("pagesize", qPagesize); err != nil {
			return err
		}
	}

	// query array param pagesize

	// query param resolved
	qrResolved := o.Resolved
	qResolved := swag.FormatBool(qrResolved)
	if qResolved != "" && qResolved != "-1" {
		if err := r.SetQueryParam("resolved", qResolved); err != nil {
			return err
		}
	}

	// query array param resolved

	// query param startupid
	qrStartupid := o.Startupid
	qStartupid := swag.FormatInt64(qrStartupid)
	if qStartupid != "" && qStartupid != "-1" {
		if err := r.SetQueryParam("startupid", qStartupid); err != nil {
			return err
		}
	}

	// query array param startupid

	// query param status
	qrStatus := o.Status
	qStatus := swag.FormatInt64(qrStatus)
	if qStatus != "" && qStatus != "-1" {
		if err := r.SetQueryParam("status", qStatus); err != nil {
			return err
		}
	}

	// query array param status

	// query param unresolved
	qrUnresolved := o.Unresolved
	qUnresolved := swag.FormatBool(qrUnresolved)
	if qUnresolved != "" && qUnresolved != "-1" {
		if err := r.SetQueryParam("unresolved", qUnresolved); err != nil {
			return err
		}
	}

	// query array param unresolved

	// query param wonderful
	qrWonderful := o.Wonderful
	qWonderful := swag.FormatInt64(qrWonderful)
	if qWonderful != "" && qWonderful != "-1" {
		if err := r.SetQueryParam("wonderful", qWonderful); err != nil {
			return err
		}
	}

	// query array param wonderful

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
