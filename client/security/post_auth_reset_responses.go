package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"
)

type PostAuthResetReader struct {
	formats strfmt.Registry
}

func (o *PostAuthResetReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostAuthResetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAuthResetOK creates a PostAuthResetOK with default headers values
func NewPostAuthResetOK() *PostAuthResetOK {
	return &PostAuthResetOK{}
}

/*PostAuthResetOK

成功重设
*/
type PostAuthResetOK struct {
}

func (o *PostAuthResetOK) Error() string {
	return fmt.Sprintf("[POST /auth/reset][%d] postAuthResetOK ", 200)
}

func (o *PostAuthResetOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}