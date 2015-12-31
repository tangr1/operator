package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*表示主题资源

swagger:model Topic
*/
type Topic struct {
	Post

	/* 主题类别
	 */
	Category int32 `json:"category,omitempty"`

	/* 主题悬赏分
	 */
	CtoCoins int32 `json:"ctoCoins,omitempty"`

	/* 回复总数
	 */
	ReplyCount int32 `json:"replyCount,omitempty"`

	/* 发布者采纳回复时的评论
	 */
	ResolvedComment string `json:"resolvedComment,omitempty"`

	/* 采纳的回复ID
	 */
	ResolvedReplyID int64 `json:"resolvedReplyId,omitempty"`

	/* 创业公司
	 */
	Startup Startup `json:"startup,omitempty"`

	/* 主题发布者所属创业公司
	 */
	StartupID int64 `json:"startupId,omitempty"`

	/* 摘要
	 */
	Summary string `json:"summary,omitempty"`

	/* 主题标签
	 */
	Tags []string `json:"tags,omitempty"`
}

// Validate validates this topic
func (m *Topic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Post.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Topic) validateTags(formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.Required("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i])); err != nil {
			return err
		}

	}

	return nil
}