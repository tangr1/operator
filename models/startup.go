package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*表示创业公司资源

swagger:model Startup
*/
type Startup struct {
	Resource

	/* 创业公司许可证
	 */
	Certificate string `json:"certificate,omitempty"`

	/* 创业公司城市
	 */
	City int32 `json:"city,omitempty"`

	/* 消费ctocoins总数
	 */
	ConsumeCtoCoins int32 `json:"consumeCtoCoins,omitempty"`

	/* 封面图片
	 */
	CoverImage string `json:"coverImage,omitempty"`

	/* 封面标题
	 */
	CoverTitle string `json:"coverTitle,omitempty"`

	/* 创业公司ctocoins
	 */
	CtoCoins int32 `json:"ctoCoins,omitempty"`

	/* 创业公司介绍
	 */
	Description string `json:"description,omitempty"`

	/* 创业公司经营范围
	 */
	Domain int32 `json:"domain,omitempty"`

	/* 创始人
	 */
	Founders []*Founder `json:"founders,omitempty"`

	/* 帮助过的专家列表
	 */
	HelpedExperts []Expert `json:"helpedExperts,omitempty"`

	/* 创业公司网址
	 */
	Homepage string `json:"homepage,omitempty"`

	/* 创业公司融资阶段
	 */
	InvestmentStatus int32 `json:"investmentStatus,omitempty"`

	/* 投资人
	 */
	Investors []string `json:"investors,omitempty"`

	/* 创业公司Logo图片地址
	 */
	Logo string `json:"logo,omitempty"`

	/* 创业公司名称
	 */
	Name string `json:"name,omitempty"`

	/* 产品
	 */
	Products []*Product `json:"products,omitempty"`

	/* 创业公司省份
	 */
	Province int32 `json:"province,omitempty"`

	/* 创业公司注册名称
	 */
	RegistrationName string `json:"registrationName,omitempty"`

	/* 被解决的问题总数
	 */
	ResolvedTopicCount int32 `json:"resolvedTopicCount,omitempty"`

	/* 申请状态
	 */
	ReviewStatus int32 `json:"reviewStatus,omitempty"`

	/* 创业公司运营状态
	 */
	RunningStatus int32 `json:"runningStatus,omitempty"`

	/* 员工人数
	 */
	StaffNumber int32 `json:"staffNumber,omitempty"`

	/* 创业公司创始月份
	 */
	StartMonth int32 `json:"startMonth,omitempty"`

	/* 创业公司创始年份
	 */
	StartYear int32 `json:"startYear,omitempty"`

	/* 问题总数
	 */
	TopicCount int32 `json:"topicCount,omitempty"`

	/* 浏览次数
	 */
	ViewCount int64 `json:"viewCount,omitempty"`
}

// Validate validates this startup
func (m *Startup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Resource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFounders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHelpedExperts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvestors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProducts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Startup) validateFounders(formats strfmt.Registry) error {

	for i := 0; i < len(m.Founders); i++ {

		if m.Founders[i] != nil {

			if err := m.Founders[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Startup) validateHelpedExperts(formats strfmt.Registry) error {

	for i := 0; i < len(m.HelpedExperts); i++ {

		if err := m.HelpedExperts[i].Validate(formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *Startup) validateInvestors(formats strfmt.Registry) error {

	for i := 0; i < len(m.Investors); i++ {

		if err := validate.Required("investors"+"."+strconv.Itoa(i), "body", string(m.Investors[i])); err != nil {
			return err
		}

	}

	return nil
}

func (m *Startup) validateProducts(formats strfmt.Registry) error {

	for i := 0; i < len(m.Products); i++ {

		if m.Products[i] != nil {

			if err := m.Products[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}