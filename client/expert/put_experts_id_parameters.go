package expert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/tangr1/hicto/models"
)

// NewPutExpertsIDParams creates a new PutExpertsIDParams object
// with the default values initialized.
func NewPutExpertsIDParams() *PutExpertsIDParams {
	return &PutExpertsIDParams{}
}

/*PutExpertsIDParams contains all the parameters to send to the API endpoint
for the put experts ID operation typically these are written to a http.Request
*/
type PutExpertsIDParams struct {

	/*Body*/
	Body *models.UpdateExpertRequest
	/*ID
	  专家ID

	*/
	ID int64
}

// WithBody adds the body to the put experts ID params
func (o *PutExpertsIDParams) WithBody(body *models.UpdateExpertRequest) *PutExpertsIDParams {
	o.Body = body
	return o
}

// WithID adds the id to the put experts ID params
func (o *PutExpertsIDParams) WithID(id int64) *PutExpertsIDParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PutExpertsIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.UpdateExpertRequest)
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