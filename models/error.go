package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*Error error

swagger:model Error
*/
type Error struct {

	/* 错误代码
	 */
	Code string `json:"code,omitempty"`

	/* 错误信息
	 */
	Message string `json:"message,omitempty"`
}

// Validate validates this error
func (m *Error) Validate(formats strfmt.Registry) error {
	return nil
}
