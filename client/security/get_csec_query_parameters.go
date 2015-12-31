package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewGetCsecQueryParams creates a new GetCsecQueryParams object
// with the default values initialized.
func NewGetCsecQueryParams() *GetCsecQueryParams {
	return &GetCsecQueryParams{}
}

/*GetCsecQueryParams contains all the parameters to send to the API endpoint
for the get csec query operation typically these are written to a http.Request
*/
type GetCsecQueryParams struct {

	/*Businessid
	  business ID

	*/
	Businessid int64
	/*Captype
	  captcha type

	*/
	Captype int64
	/*Sceneid
	  scene ID

	*/
	Sceneid int64
	/*Userid
	  user ID

	*/
	Userid int64
	/*Userip
	  user IP

	*/
	Userip string
}

// WithBusinessid adds the businessid to the get csec query params
func (o *GetCsecQueryParams) WithBusinessid(businessid int64) *GetCsecQueryParams {
	o.Businessid = businessid
	return o
}

// WithCaptype adds the captype to the get csec query params
func (o *GetCsecQueryParams) WithCaptype(captype int64) *GetCsecQueryParams {
	o.Captype = captype
	return o
}

// WithSceneid adds the sceneid to the get csec query params
func (o *GetCsecQueryParams) WithSceneid(sceneid int64) *GetCsecQueryParams {
	o.Sceneid = sceneid
	return o
}

// WithUserid adds the userid to the get csec query params
func (o *GetCsecQueryParams) WithUserid(userid int64) *GetCsecQueryParams {
	o.Userid = userid
	return o
}

// WithUserip adds the userip to the get csec query params
func (o *GetCsecQueryParams) WithUserip(userip string) *GetCsecQueryParams {
	o.Userip = userip
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetCsecQueryParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param businessid
	if err := r.SetPathParam("businessid", swag.FormatInt64(o.Businessid)); err != nil {
		return err
	}

	// path param captype
	if err := r.SetPathParam("captype", swag.FormatInt64(o.Captype)); err != nil {
		return err
	}

	// path param sceneid
	if err := r.SetPathParam("sceneid", swag.FormatInt64(o.Sceneid)); err != nil {
		return err
	}

	// path param userid
	if err := r.SetPathParam("userid", swag.FormatInt64(o.Userid)); err != nil {
		return err
	}

	// path param userip
	if err := r.SetPathParam("userip", o.Userip); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
