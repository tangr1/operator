package startup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewGetStartupsIDParams creates a new GetStartupsIDParams object
// with the default values initialized.
func NewGetStartupsIDParams() *GetStartupsIDParams {
	return &GetStartupsIDParams{}
}

/*GetStartupsIDParams contains all the parameters to send to the API endpoint
for the get startups ID operation typically these are written to a http.Request
*/
type GetStartupsIDParams struct {

	/*ID
	  创业公司ID

	*/
	ID int64
}

// WithID adds the id to the get startups ID params
func (o *GetStartupsIDParams) WithID(id int64) *GetStartupsIDParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetStartupsIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}