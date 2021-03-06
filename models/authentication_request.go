package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*表示认证请求, 用户可以通过邮件地址或电话号码登录, 二者必选其一

swagger:model AuthenticationRequest
*/
type AuthenticationRequest struct {

	/* 用户使用移动设备登录时的设备token

	Max Length: 100
	*/
	DeviceToken string `json:"deviceToken,omitempty"`

	/* 用户注册的邮件地址

	Max Length: 45
	Pattern: ^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$
	*/
	Email string `json:"email,omitempty"`

	/* token失效时间, 单位为分钟
	 */
	Expiration int32 `json:"expiration,omitempty"`

	/* 密码

	Required: true
	Max Length: 20
	*/
	Password string `json:"password,omitempty"`

	/* 用户注册的电话号码

	Max Length: 45
	*/
	Phone string `json:"phone,omitempty"`
}

// Validate validates this authentication request
func (m *AuthenticationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceToken(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePhone(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthenticationRequest) validateDeviceToken(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceToken) { // not required
		return nil
	}

	if err := validate.MaxLength("deviceToken", "body", string(m.DeviceToken), 100); err != nil {
		return err
	}

	return nil
}

func (m *AuthenticationRequest) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.MaxLength("email", "body", string(m.Email), 45); err != nil {
		return err
	}

	if err := validate.Pattern("email", "body", string(m.Email), `^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$`); err != nil {
		return err
	}

	return nil
}

func (m *AuthenticationRequest) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", string(m.Password)); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", string(m.Password), 20); err != nil {
		return err
	}

	return nil
}

func (m *AuthenticationRequest) validatePhone(formats strfmt.Registry) error {

	if swag.IsZero(m.Phone) { // not required
		return nil
	}

	if err := validate.MaxLength("phone", "body", string(m.Phone), 45); err != nil {
		return err
	}

	return nil
}
