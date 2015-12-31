package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*表示创建专家请求

swagger:model CreateExpertRequest
*/
type CreateExpertRequest struct {

	/* 用户头像

	Max Length: 200
	*/
	Avatar string `json:"avatar,omitempty"`

	/* 城市id
	 */
	City int32 `json:"city,omitempty"`

	/* 专家任职公司

	Max Length: 45
	*/
	Company string `json:"company,omitempty"`

	/* 封面图片

	Max Length: 200
	*/
	CoverImage string `json:"coverImage,omitempty"`

	/* 封面标题

	Max Length: 100
	*/
	CoverTitle string `json:"coverTitle,omitempty"`

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

	/* 专家特长
	 */
	Expertise []int32 `json:"expertise,omitempty"`

	/* token失效时间, 单位为分钟
	 */
	Expiration int32 `json:"expiration,omitempty"`

	/* 使用的邀请码

	Max Length: 45
	*/
	InviteCode string `json:"inviteCode,omitempty"`

	/* 专家管理技能
	 */
	ManagementSkill int32 `json:"managementSkill,omitempty"`

	/* 用户姓名

	Max Length: 45
	*/
	Name string `json:"name,omitempty"`

	/* 评论被评论时邮件通知
	 */
	NotifyCommentNewCommentByEmail *bool `json:"notifyCommentNewCommentByEmail,omitempty"`

	/* 评论被评论时推送通知
	 */
	NotifyCommentNewCommentByPush *bool `json:"notifyCommentNewCommentByPush,omitempty"`

	/* 新主题邮件通知
	 */
	NotifyNewTopicByEmail *bool `json:"notifyNewTopicByEmail,omitempty"`

	/* 新主题推送通知
	 */
	NotifyNewTopicByPush *bool `json:"notifyNewTopicByPush,omitempty"`

	/* 是否只在空闲时间通知
	 */
	NotifyOnlyFreeTime *bool `json:"notifyOnlyFreeTime,omitempty"`

	/* 有回复被采纳时邮件通知
	 */
	NotifyReplyAcceptedByEmail *bool `json:"notifyReplyAcceptedByEmail,omitempty"`

	/* 有回复被采纳时推送通知
	 */
	NotifyReplyAcceptedByPush *bool `json:"notifyReplyAcceptedByPush,omitempty"`

	/* 答案或评论被评论时邮件通知
	 */
	NotifyReplyNewCommentByEmail *bool `json:"notifyReplyNewCommentByEmail,omitempty"`

	/* 答案或评论被评论时推送通知
	 */
	NotifyReplyNewCommentByPush *bool `json:"notifyReplyNewCommentByPush,omitempty"`

	/* 用户新密码

	Required: true
	Max Length: 20
	*/
	Password string `json:"password,omitempty"`

	/* 用户电话号码

	Max Length: 45
	*/
	Phone string `json:"phone,omitempty"`

	/* 专家职位

	Max Length: 45
	*/
	Position string `json:"position,omitempty"`
}

// Validate validates this create expert request
func (m *CreateExpertRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvatar(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCompany(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCoverImage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCoverTitle(formats); err != nil {
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

	if err := m.validateExpertise(formats); err != nil {
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

func (m *CreateExpertRequest) validateAvatar(formats strfmt.Registry) error {

	if swag.IsZero(m.Avatar) { // not required
		return nil
	}

	if err := validate.MaxLength("avatar", "body", string(m.Avatar), 200); err != nil {
		return err
	}

	return nil
}

func (m *CreateExpertRequest) validateCompany(formats strfmt.Registry) error {

	if swag.IsZero(m.Company) { // not required
		return nil
	}

	if err := validate.MaxLength("company", "body", string(m.Company), 45); err != nil {
		return err
	}

	return nil
}

func (m *CreateExpertRequest) validateCoverImage(formats strfmt.Registry) error {

	if swag.IsZero(m.CoverImage) { // not required
		return nil
	}

	if err := validate.MaxLength("coverImage", "body", string(m.CoverImage), 200); err != nil {
		return err
	}

	return nil
}

func (m *CreateExpertRequest) validateCoverTitle(formats strfmt.Registry) error {

	if swag.IsZero(m.CoverTitle) { // not required
		return nil
	}

	if err := validate.MaxLength("coverTitle", "body", string(m.CoverTitle), 100); err != nil {
		return err
	}

	return nil
}

func (m *CreateExpertRequest) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 255); err != nil {
		return err
	}

	return nil
}

func (m *CreateExpertRequest) validateDeviceToken(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceToken) { // not required
		return nil
	}

	if err := validate.MaxLength("deviceToken", "body", string(m.DeviceToken), 100); err != nil {
		return err
	}

	return nil
}

func (m *CreateExpertRequest) validateEmail(formats strfmt.Registry) error {

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

func (m *CreateExpertRequest) validateExpertise(formats strfmt.Registry) error {

	if swag.IsZero(m.Expertise) { // not required
		return nil
	}

	for i := 0; i < len(m.Expertise); i++ {

		if err := validate.Required("expertise"+"."+strconv.Itoa(i), "body", int32(m.Expertise[i])); err != nil {
			return err
		}

	}

	return nil
}

func (m *CreateExpertRequest) validateInviteCode(formats strfmt.Registry) error {

	if swag.IsZero(m.InviteCode) { // not required
		return nil
	}

	if err := validate.MaxLength("inviteCode", "body", string(m.InviteCode), 45); err != nil {
		return err
	}

	return nil
}

func (m *CreateExpertRequest) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 45); err != nil {
		return err
	}

	return nil
}

func (m *CreateExpertRequest) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", string(m.Password)); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", string(m.Password), 20); err != nil {
		return err
	}

	return nil
}

func (m *CreateExpertRequest) validatePhone(formats strfmt.Registry) error {

	if swag.IsZero(m.Phone) { // not required
		return nil
	}

	if err := validate.MaxLength("phone", "body", string(m.Phone), 45); err != nil {
		return err
	}

	return nil
}

func (m *CreateExpertRequest) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := validate.MaxLength("position", "body", string(m.Position), 45); err != nil {
		return err
	}

	return nil
}