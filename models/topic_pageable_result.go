package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*资源列表的一页结果

swagger:model TopicPageableResult
*/
type TopicPageableResult struct {

	/* 当前页资源列表
	 */
	Content []Topic `json:"content,omitempty"`

	/* 是否最后一页
	 */
	First bool `json:"first,omitempty"`

	/* 是否第一页
	 */
	Last bool `json:"last,omitempty"`

	/* 当前页数
	 */
	Number int32 `json:"number,omitempty"`

	/* 当前页资源项数
	 */
	NumberOfElements int32 `json:"numberOfElements,omitempty"`

	/* 服务器当前时间
	 */
	ServerTime int64 `json:"serverTime,omitempty"`

	/* 每页资源项数
	 */
	Size int32 `json:"size,omitempty"`

	/* 总共资源项数
	 */
	TotalElements int32 `json:"totalElements,omitempty"`

	/* 总共页数
	 */
	TotalPages int32 `json:"totalPages,omitempty"`
}

// Validate validates this topic pageable result
func (m *TopicPageableResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopicPageableResult) validateContent(formats strfmt.Registry) error {

	if swag.IsZero(m.Content) { // not required
		return nil
	}

	for i := 0; i < len(m.Content); i++ {

		if err := m.Content[i].Validate(formats); err != nil {
			return err
		}

	}

	return nil
}
