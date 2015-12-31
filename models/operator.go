package models

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*表示审批者资源

swagger:model Operator
*/
type Operator struct {
	User
}

// Validate validates this operator
func (m *Operator) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.User.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}