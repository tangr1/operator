package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*资源列表的一页结果

swagger:model RecommendationPageableResult
*/
type RecommendationPageableResult struct {

	/* 当前页资源列表
	 */
	Content []*Recommendation `json:"content,omitempty"`

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

// Validate validates this recommendation pageable result
func (m *RecommendationPageableResult) Validate(formats strfmt.Registry) error {
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

func (m *RecommendationPageableResult) validateContent(formats strfmt.Registry) error {

	if swag.IsZero(m.Content) { // not required
		return nil
	}

	for i := 0; i < len(m.Content); i++ {

		if m.Content[i] != nil {

			if err := m.Content[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
