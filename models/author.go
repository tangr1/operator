package models

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*表示发布者资源

swagger:model Author
*/
type Author struct {

	/* 用户头像
	 */
	Avatar string `json:"avatar,omitempty"`

	/* 发布者ID
	 */
	ID int64 `json:"id,omitempty"`

	/* 用户姓名
	 */
	Name string `json:"name,omitempty"`

	/* 是否为创业公司员工
	 */
	Startup *bool `json:"startup,omitempty"`
}

// Validate validates this author
func (m *Author) Validate(formats strfmt.Registry) error {
	return nil
}