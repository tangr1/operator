package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*表示明星资源

swagger:model Star
*/
type Star struct {

	/* 日期
	 */
	Date string `json:"date,omitempty"`

	/* 明星专家
	 */
	Experts []Expert `json:"experts,omitempty"`

	/* 明星ID
	 */
	ID int64 `json:"id,omitempty"`

	/* 明星企业
	 */
	Startups []Startup `json:"startups,omitempty"`
}

// Validate validates this star
func (m *Star) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExperts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartups(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Star) validateExperts(formats strfmt.Registry) error {

	if swag.IsZero(m.Experts) { // not required
		return nil
	}

	for i := 0; i < len(m.Experts); i++ {

		if err := m.Experts[i].Validate(formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *Star) validateStartups(formats strfmt.Registry) error {

	if swag.IsZero(m.Startups) { // not required
		return nil
	}

	for i := 0; i < len(m.Startups); i++ {

		if err := m.Startups[i].Validate(formats); err != nil {
			return err
		}

	}

	return nil
}
