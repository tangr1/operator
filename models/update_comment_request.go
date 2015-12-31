package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*表示更新评论请求

swagger:model UpdateCommentRequest
*/
type UpdateCommentRequest struct {

	/* 发布内容

	Max Length: 65535
	*/
	Content string `json:"content,omitempty"`
}

// Validate validates this update comment request
func (m *UpdateCommentRequest) Validate(formats strfmt.Registry) error {
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

func (m *UpdateCommentRequest) validateContent(formats strfmt.Registry) error {

	if swag.IsZero(m.Content) { // not required
		return nil
	}

	if err := validate.MaxLength("content", "body", string(m.Content), 65535); err != nil {
		return err
	}

	return nil
}