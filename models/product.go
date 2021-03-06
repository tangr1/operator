package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*表示产品资源

swagger:model Product
*/
type Product struct {

	/* 产品Android APP链接

	Max Length: 200
	*/
	AndroidLink string `json:"androidLink,omitempty"`

	/* 产品描述

	Max Length: 255
	*/
	Description string `json:"description,omitempty"`

	/* 产品主页

	Max Length: 200
	*/
	Homepage string `json:"homepage,omitempty"`

	/* 产品iOS APP链接

	Max Length: 200
	*/
	IosLink string `json:"iosLink,omitempty"`

	/* 产品logo

	Max Length: 200
	*/
	Logo string `json:"logo,omitempty"`

	/* 产品名称

	Max Length: 45
	*/
	Name string `json:"name,omitempty"`
}

// Validate validates this product
func (m *Product) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAndroidLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHomepage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIosLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLogo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Product) validateAndroidLink(formats strfmt.Registry) error {

	if swag.IsZero(m.AndroidLink) { // not required
		return nil
	}

	if err := validate.MaxLength("androidLink", "body", string(m.AndroidLink), 200); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 255); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateHomepage(formats strfmt.Registry) error {

	if swag.IsZero(m.Homepage) { // not required
		return nil
	}

	if err := validate.MaxLength("homepage", "body", string(m.Homepage), 200); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateIosLink(formats strfmt.Registry) error {

	if swag.IsZero(m.IosLink) { // not required
		return nil
	}

	if err := validate.MaxLength("iosLink", "body", string(m.IosLink), 200); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateLogo(formats strfmt.Registry) error {

	if swag.IsZero(m.Logo) { // not required
		return nil
	}

	if err := validate.MaxLength("logo", "body", string(m.Logo), 200); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 45); err != nil {
		return err
	}

	return nil
}
