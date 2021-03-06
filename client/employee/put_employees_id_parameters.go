package employee

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/tangr1/hicto/models"
)

// NewPutEmployeesIDParams creates a new PutEmployeesIDParams object
// with the default values initialized.
func NewPutEmployeesIDParams() *PutEmployeesIDParams {
	return &PutEmployeesIDParams{}
}

/*PutEmployeesIDParams contains all the parameters to send to the API endpoint
for the put employees ID operation typically these are written to a http.Request
*/
type PutEmployeesIDParams struct {

	/*Body*/
	Body *models.UpdateEmployeeRequest
	/*ID
	  创业公司员工ID

	*/
	ID int64
}

// WithBody adds the body to the put employees ID params
func (o *PutEmployeesIDParams) WithBody(body *models.UpdateEmployeeRequest) *PutEmployeesIDParams {
	o.Body = body
	return o
}

// WithID adds the id to the put employees ID params
func (o *PutEmployeesIDParams) WithID(id int64) *PutEmployeesIDParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PutEmployeesIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.UpdateEmployeeRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
