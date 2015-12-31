package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*表示更新创业公司员工请求

swagger:model UpdateEmployeeRequest
*/
type UpdateEmployeeRequest struct {

	/* 是否为管理员
	 */
	Admin *bool `json:"admin,omitempty"`

	/* 用户头像

	Max Length: 200
	*/
	Avatar string `json:"avatar,omitempty"`

	/* 用户介绍
	 */
	Description string `json:"description,omitempty"`

	/* 用户使用的移动设备token

	Max Length: 100
	*/
	DeviceToken string `json:"deviceToken,omitempty"`

	/* 用户邮件地址

	Max Length: 45
	Pattern: ^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$
	*/
	Email string `json:"email,omitempty"`

	/* 用户姓名

	Max Length: 45
	*/
	Name string `json:"name,omitempty"`

	/* 评论有评论时邮件通知
	 */
	NotifyCommentNewCommentByEmail *bool `json:"notifyCommentNewCommentByEmail,omitempty"`

	/* 评论有评论时推送通知
	 */
	NotifyCommentNewCommentByPush *bool `json:"notifyCommentNewCommentByPush,omitempty"`

	/* 主题有回复时邮件通知
	 */
	NotifyNewReplyByEmail *bool `json:"notifyNewReplyByEmail,omitempty"`

	/* 主题有回复时推送通知
	 */
	NotifyNewReplyByPush *bool `json:"notifyNewReplyByPush,omitempty"`

	/* 只当我的主题有响应时通知
	 */
	NotifyOnlyMyTopic *bool `json:"notifyOnlyMyTopic,omitempty"`

	/* 主题有评论时邮件通知
	 */
	NotifyTopicNewCommentByEmail *bool `json:"notifyTopicNewCommentByEmail,omitempty"`

	/* 主题有评论时推送通知
	 */
	NotifyTopicNewCommentByPush *bool `json:"notifyTopicNewCommentByPush,omitempty"`

	/* 用户当前密码, 如修改密码需要提供

	Max Length: 20
	*/
	OldPassword string `json:"oldPassword,omitempty"`

	/* 用户密码

	Max Length: 20
	*/
	Password string `json:"password,omitempty"`

	/* 用户电话号码

	Max Length: 45
	*/
	Phone string `json:"phone,omitempty"`

	/* 职位

	Max Length: 45
	*/
	Position string `json:"position,omitempty"`
}

// Validate validates this update employee request
func (m *UpdateEmployeeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvatar(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDeviceToken(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOldPassword(formats); err != nil {
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

	if err := m.validatePosition(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateEmployeeRequest) validateAvatar(formats strfmt.Registry) error {

	if swag.IsZero(m.Avatar) { // not required
		return nil
	}

	if err := validate.MaxLength("avatar", "body", string(m.Avatar), 200); err != nil {
		return err
	}

	return nil
}

func (m *UpdateEmployeeRequest) validateDeviceToken(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceToken) { // not required
		return nil
	}

	if err := validate.MaxLength("deviceToken", "body", string(m.DeviceToken), 100); err != nil {
		return err
	}

	return nil
}

func (m *UpdateEmployeeRequest) validateEmail(formats strfmt.Registry) error {

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

func (m *UpdateEmployeeRequest) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 45); err != nil {
		return err
	}

	return nil
}

func (m *UpdateEmployeeRequest) validateOldPassword(formats strfmt.Registry) error {

	if swag.IsZero(m.OldPassword) { // not required
		return nil
	}

	if err := validate.MaxLength("oldPassword", "body", string(m.OldPassword), 20); err != nil {
		return err
	}

	return nil
}

func (m *UpdateEmployeeRequest) validatePassword(formats strfmt.Registry) error {

	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.MaxLength("password", "body", string(m.Password), 20); err != nil {
		return err
	}

	return nil
}

func (m *UpdateEmployeeRequest) validatePhone(formats strfmt.Registry) error {

	if swag.IsZero(m.Phone) { // not required
		return nil
	}

	if err := validate.MaxLength("phone", "body", string(m.Phone), 45); err != nil {
		return err
	}

	return nil
}

func (m *UpdateEmployeeRequest) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := validate.MaxLength("position", "body", string(m.Position), 45); err != nil {
		return err
	}

	return nil
}
