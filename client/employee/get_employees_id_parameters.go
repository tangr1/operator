package employee

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewGetEmployeesIDParams creates a new GetEmployeesIDParams object
// with the default values initialized.
func NewGetEmployeesIDParams() *GetEmployeesIDParams {
	return &GetEmployeesIDParams{}
}

/*GetEmployeesIDParams contains all the parameters to send to the API endpoint
for the get employees ID operation typically these are written to a http.Request
*/
type GetEmployeesIDParams struct {

	/*ID
	  创业公司员工ID

	*/
	ID int64
}

// WithID adds the id to the get employees ID params
func (o *GetEmployeesIDParams) WithID(id int64) *GetEmployeesIDParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmployeesIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

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
