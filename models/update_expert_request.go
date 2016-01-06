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

/*表示更新专家请求

swagger:model UpdateExpertRequest
*/
type UpdateExpertRequest struct {

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

	/* 拥有的悬赏分
	 */
	CtoCoins int32 `json:"ctoCoins,omitempty"`

	/* 用户介绍

	Max Length: 255
	*/
	Description string `json:"description,omitempty"`

	/* 用户使用的移动设备token

	Max Length: 200
	*/
	DeviceToken string `json:"deviceToken,omitempty"`

	/* 用户邮件地址

	Max Length: 45
	Pattern: ^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$
	*/
	Email string `json:"email,omitempty"`

	/* 专家特长
	 */
	Expertise []int32 `json:"expertise,omitempty"`

	/* 是否为内部专家
	 */
	Internal bool `json:"internal,omitempty"`

	/* 专家管理技能
	 */
	ManagementSkill int32 `json:"managementSkill,omitempty"`

	/* 修改原因

	Max Length: 255
	*/
	ModifyReason string `json:"modifyReason,omitempty"`

	/* 用户姓名

	Max Length: 45
	*/
	Name string `json:"name,omitempty"`

	/* 评论被评论时邮件通知
	 */
	NotifyCommentNewCommentByEmail bool `json:"notifyCommentNewCommentByEmail,omitempty"`

	/* 评论被评论时推送通知
	 */
	NotifyCommentNewCommentByPush bool `json:"notifyCommentNewCommentByPush,omitempty"`

	/* 新主题邮件通知
	 */
	NotifyNewTopicByEmail bool `json:"notifyNewTopicByEmail,omitempty"`

	/* 新主题推送通知
	 */
	NotifyNewTopicByPush bool `json:"notifyNewTopicByPush,omitempty"`

	/* 是否只在空闲时间通知
	 */
	NotifyOnlyFreeTime bool `json:"notifyOnlyFreeTime,omitempty"`

	/* 有回复被采纳时邮件通知
	 */
	NotifyReplyAcceptedByEmail bool `json:"notifyReplyAcceptedByEmail,omitempty"`

	/* 有回复被采纳时推送通知
	 */
	NotifyReplyAcceptedByPush bool `json:"notifyReplyAcceptedByPush,omitempty"`

	/* 答案或评论被评论时邮件通知
	 */
	NotifyReplyNewCommentByEmail bool `json:"notifyReplyNewCommentByEmail,omitempty"`

	/* 答案或评论被评论时推送通知
	 */
	NotifyReplyNewCommentByPush bool `json:"notifyReplyNewCommentByPush,omitempty"`

	/* 用户当前密码, 如修改密码需要提供

	Max Length: 45
	*/
	OldPassword string `json:"oldPassword,omitempty"`

	/* 用户新密码

	Max Length: 20
	*/
	Password string `json:"password,omitempty"`

	/* 用户电话号码
	 */
	Phone string `json:"phone,omitempty"`

	/* 专家职位

	Max Length: 45
	*/
	Position string `json:"position,omitempty"`

	/* 专家威望
	 */
	Reputation int32 `json:"reputation,omitempty"`

	/* 审核状态
	 */
	ReviewStatus int32 `json:"reviewStatus,omitempty"`
}

// Validate validates this update expert request
func (m *UpdateExpertRequest) Validate(formats strfmt.Registry) error {
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

	if err := m.validateModifyReason(formats); err != nil {
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

	if err := m.validatePosition(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateExpertRequest) validateAvatar(formats strfmt.Registry) error {

	if swag.IsZero(m.Avatar) { // not required
		return nil
	}

	if err := validate.MaxLength("avatar", "body", string(m.Avatar), 200); err != nil {
		return err
	}

	return nil
}

func (m *UpdateExpertRequest) validateCompany(formats strfmt.Registry) error {

	if swag.IsZero(m.Company) { // not required
		return nil
	}

	if err := validate.MaxLength("company", "body", string(m.Company), 45); err != nil {
		return err
	}

	return nil
}

func (m *UpdateExpertRequest) validateCoverImage(formats strfmt.Registry) error {

	if swag.IsZero(m.CoverImage) { // not required
		return nil
	}

	if err := validate.MaxLength("coverImage", "body", string(m.CoverImage), 200); err != nil {
		return err
	}

	return nil
}

func (m *UpdateExpertRequest) validateCoverTitle(formats strfmt.Registry) error {

	if swag.IsZero(m.CoverTitle) { // not required
		return nil
	}

	if err := validate.MaxLength("coverTitle", "body", string(m.CoverTitle), 100); err != nil {
		return err
	}

	return nil
}

func (m *UpdateExpertRequest) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 255); err != nil {
		return err
	}

	return nil
}

func (m *UpdateExpertRequest) validateDeviceToken(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceToken) { // not required
		return nil
	}

	if err := validate.MaxLength("deviceToken", "body", string(m.DeviceToken), 200); err != nil {
		return err
	}

	return nil
}

func (m *UpdateExpertRequest) validateEmail(formats strfmt.Registry) error {

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

func (m *UpdateExpertRequest) validateExpertise(formats strfmt.Registry) error {

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

func (m *UpdateExpertRequest) validateModifyReason(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifyReason) { // not required
		return nil
	}

	if err := validate.MaxLength("modifyReason", "body", string(m.ModifyReason), 255); err != nil {
		return err
	}

	return nil
}

func (m *UpdateExpertRequest) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 45); err != nil {
		return err
	}

	return nil
}

func (m *UpdateExpertRequest) validateOldPassword(formats strfmt.Registry) error {

	if swag.IsZero(m.OldPassword) { // not required
		return nil
	}

	if err := validate.MaxLength("oldPassword", "body", string(m.OldPassword), 45); err != nil {
		return err
	}

	return nil
}

func (m *UpdateExpertRequest) validatePassword(formats strfmt.Registry) error {

	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.MaxLength("password", "body", string(m.Password), 20); err != nil {
		return err
	}

	return nil
}

func (m *UpdateExpertRequest) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := validate.MaxLength("position", "body", string(m.Position), 45); err != nil {
		return err
	}

	return nil
}
