package operation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewGetStarsIDParams creates a new GetStarsIDParams object
// with the default values initialized.
func NewGetStarsIDParams() *GetStarsIDParams {
	return &GetStarsIDParams{}
}

/*GetStarsIDParams contains all the parameters to send to the API endpoint
for the get stars ID operation typically these are written to a http.Request
*/
type GetStarsIDParams struct {

	/*ID
	  明星ID

	*/
	ID int64
}

// WithID adds the id to the get stars ID params
func (o *GetStarsIDParams) WithID(id int64) *GetStarsIDParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetStarsIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

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