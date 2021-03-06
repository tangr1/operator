package operation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewDeleteStarsIDParams creates a new DeleteStarsIDParams object
// with the default values initialized.
func NewDeleteStarsIDParams() *DeleteStarsIDParams {
	return &DeleteStarsIDParams{}
}

/*DeleteStarsIDParams contains all the parameters to send to the API endpoint
for the delete stars ID operation typically these are written to a http.Request
*/
type DeleteStarsIDParams struct {

	/*ID
	  明星ID

	*/
	ID int64
}

// WithID adds the id to the delete stars ID params
func (o *DeleteStarsIDParams) WithID(id int64) *DeleteStarsIDParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteStarsIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

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
