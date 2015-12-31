package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/tangr1/hicto/models"
)

type PostCsecCheckReader struct {
	formats strfmt.Registry
}

func (o *PostCsecCheckReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCsecCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostCsecCheckBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostCsecCheckInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCsecCheckOK creates a PostCsecCheckOK with default headers values
func NewPostCsecCheckOK() *PostCsecCheckOK {
	return &PostCsecCheckOK{}
}

/*PostCsecCheckOK

成功验证
*/
type PostCsecCheckOK struct {
}

func (o *PostCsecCheckOK) Error() string {
	return fmt.Sprintf("[POST /csec/check][%d] postCsecCheckOK ", 200)
}

func (o *PostCsecCheckOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostCsecCheckBadRequest creates a PostCsecCheckBadRequest with default headers values
func NewPostCsecCheckBadRequest() *PostCsecCheckBadRequest {
	return &PostCsecCheckBadRequest{}
}

/*PostCsecCheckBadRequest

参数无效
*/
type PostCsecCheckBadRequest struct {
	Payload *models.Error
}

func (o *PostCsecCheckBadRequest) Error() string {
	return fmt.Sprintf("[POST /csec/check][%d] postCsecCheckBadRequest  %+v", 400, o.Payload)
}

func (o *PostCsecCheckBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPostCsecCheckInternalServerError creates a PostCsecCheckInternalServerError with default headers values
func NewPostCsecCheckInternalServerError() *PostCsecCheckInternalServerError {
	return &PostCsecCheckInternalServerError{}
}

/*PostCsecCheckInternalServerError

请求CSEC失败
*/
type PostCsecCheckInternalServerError struct {
	Payload *models.Error
}

func (o *PostCsecCheckInternalServerError) Error() string {
	return fmt.Sprintf("[POST /csec/check][%d] postCsecCheckInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCsecCheckInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}