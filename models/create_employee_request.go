package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*表示创建创业员工资源

swagger:model CreateEmployeeRequest
*/
type CreateEmployeeRequest struct {

	/* 是否为管理员

	Required: true
	*/
	Admin bool `json:"admin,omitempty"`

	/* 用户头像

	Max Length: 200
	*/
	Avatar string `json:"avatar,omitempty"`

	/* 用户介绍

	Max Length: 255
	*/
	Description string `json:"description,omitempty"`

	/* 用户使用的移动设备token

	Max Length: 100
	*/
	DeviceToken string `json:"deviceToken,omitempty"`

	/* 用户邮件地址

	Required: true
	Max Length: 45
	Pattern: ^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$
	*/
	Email string `json:"email,omitempty"`

	/* token失效时间, 单位为分钟
	 */
	Expiration int32 `json:"expiration,omitempty"`

	/* 使用的邀请码

	Max Length: 45
	*/
	InviteCode string `json:"inviteCode,omitempty"`

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

	/* 用户密码

	Required: true
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

	/* 创业公司ID
	 */
	StartupID int64 `json:"startupId,omitempty"`
}

// Validate validates this create employee request
func (m *CreateEmployeeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdmin(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAvatar(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
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

	if err := m.validateInviteCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *CreateEmployeeRequest) validateAdmin(formats strfmt.Registry) error {

	if err := validate.Required("admin", "body", bool(m.Admin)); err != nil {
		return err
	}

	return nil
}

func (m *CreateEmployeeRequest) validateAvatar(formats strfmt.Registry) error {

	if swag.IsZero(m.Avatar) { // not required
		return nil
	}

	if err := validate.MaxLength("avatar", "body", string(m.Avatar), 200); err != nil {
		return err
	}

	return nil
}

func (m *CreateEmployeeRequest) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 255); err != nil {
		return err
	}

	return nil
}

func (m *CreateEmployeeRequest) validateDeviceToken(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceToken) { // not required
		return nil
	}

	if err := validate.MaxLength("deviceToken", "body", string(m.DeviceToken), 100); err != nil {
		return err
	}

	return nil
}

func (m *CreateEmployeeRequest) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", string(m.Email)); err != nil {
		return err
	}

	if err := validate.MaxLength("email", "body", string(m.Email), 45); err != nil {
		return err
	}

	if err := validate.Pattern("email", "body", string(m.Email), `^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$`); err != nil {
		return err
	}

	return nil
}

func (m *CreateEmployeeRequest) validateInviteCode(formats strfmt.Registry) error {

	if swag.IsZero(m.InviteCode) { // not required
		return nil
	}

	if err := validate.MaxLength("inviteCode", "body", string(m.InviteCode), 45); err != nil {
		return err
	}

	return nil
}

func (m *CreateEmployeeRequest) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 45); err != nil {
		return err
	}

	return nil
}

func (m *CreateEmployeeRequest) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", string(m.Password)); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", string(m.Password), 20); err != nil {
		return err
	}

	return nil
}

func (m *CreateEmployeeRequest) validatePhone(formats strfmt.Registry) error {

	if swag.IsZero(m.Phone) { // not required
		return nil
	}

	if err := validate.MaxLength("phone", "body", string(m.Phone), 45); err != nil {
		return err
	}

	return nil
}

func (m *CreateEmployeeRequest) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := validate.MaxLength("position", "body", string(m.Position), 45); err != nil {
		return err
	}

	return nil
}