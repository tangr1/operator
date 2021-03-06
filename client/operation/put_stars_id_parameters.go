package operation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/tangr1/hicto/models"
)

// NewPutStarsIDParams creates a new PutStarsIDParams object
// with the default values initialized.
func NewPutStarsIDParams() *PutStarsIDParams {
	return &PutStarsIDParams{}
}

/*PutStarsIDParams contains all the parameters to send to the API endpoint
for the put stars ID operation typically these are written to a http.Request
*/
type PutStarsIDParams struct {

	/*Body*/
	Body *models.UpdateStarRequest
	/*ID
	  明星ID

	*/
	ID int64
}

// WithBody adds the body to the put stars ID params
func (o *PutStarsIDParams) WithBody(body *models.UpdateStarRequest) *PutStarsIDParams {
	o.Body = body
	return o
}

// WithID adds the id to the put stars ID params
func (o *PutStarsIDParams) WithID(id int64) *PutStarsIDParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PutStarsIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.UpdateStarRequest)
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
