package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new security API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for security API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*获取captcha

获取captcha

*/
func (a *Client) GetCsecQuery(params *GetCsecQueryParams, authInfo client.AuthInfoWriter) (*GetCsecQueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCsecQueryParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:       "GetCsecQuery",
		Params:   params,
		Reader:   &GetCsecQueryReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCsecQueryOK), nil
}

/*忘记密码

忘记密码

*/
func (a *Client) PostAuthForgot(params *PostAuthForgotParams, authInfo client.AuthInfoWriter) (*PostAuthForgotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAuthForgotParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:       "PostAuthForgot",
		Params:   params,
		Reader:   &PostAuthForgotReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAuthForgotOK), nil
}

/*用户登录

通过邮件地址和密码获取token

*/
func (a *Client) PostAuthLogin(params *PostAuthLoginParams, authInfo client.AuthInfoWriter) (*PostAuthLoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAuthLoginParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:       "PostAuthLogin",
		Params:   params,
		Reader:   &PostAuthLoginReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAuthLoginOK), nil
}

/*用户登出

用户登出，使token无效

*/
func (a *Client) PostAuthLogout(params *PostAuthLogoutParams, authInfo client.AuthInfoWriter) (*PostAuthLogoutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAuthLogoutParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:       "PostAuthLogout",
		Params:   params,
		Reader:   &PostAuthLogoutReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAuthLogoutOK), nil
}

/*刷新token

在token的有效时间内刷新token

*/
func (a *Client) PostAuthRefresh(params *PostAuthRefreshParams, authInfo client.AuthInfoWriter) (*PostAuthRefreshOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAuthRefreshParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:       "PostAuthRefresh",
		Params:   params,
		Reader:   &PostAuthRefreshReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAuthRefreshOK), nil
}

/*重设密码

重设密码

*/
func (a *Client) PostAuthReset(params *PostAuthResetParams, authInfo client.AuthInfoWriter) (*PostAuthResetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAuthResetParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:       "PostAuthReset",
		Params:   params,
		Reader:   &PostAuthResetReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAuthResetOK), nil
}

/*验证aptcha

验证captcha

*/
func (a *Client) PostCsecCheck(params *PostCsecCheckParams, authInfo client.AuthInfoWriter) (*PostCsecCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCsecCheckParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:       "PostCsecCheck",
		Params:   params,
		Reader:   &PostCsecCheckReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCsecCheckOK), nil
}

/*创业公司员工登录

通过邮件地址和密码获取token

*/
func (a *Client) PostEmployeeLogin(params *PostEmployeeLoginParams, authInfo client.AuthInfoWriter) (*PostEmployeeLoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEmployeeLoginParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:       "PostEmployeeLogin",
		Params:   params,
		Reader:   &PostEmployeeLoginReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostEmployeeLoginOK), nil
}

/*专家登录

通过邮件地址和密码获取token

*/
func (a *Client) PostExpertLogin(params *PostExpertLoginParams, authInfo client.AuthInfoWriter) (*PostExpertLoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostExpertLoginParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:       "PostExpertLogin",
		Params:   params,
		Reader:   &PostExpertLoginReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostExpertLoginOK), nil
}

/*运营登录

通过邮件地址和密码获取token

*/
func (a *Client) PostOperatorLogin(params *PostOperatorLoginParams, authInfo client.AuthInfoWriter) (*PostOperatorLoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOperatorLoginParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:       "PostOperatorLogin",
		Params:   params,
		Reader:   &PostOperatorLoginReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostOperatorLoginOK), nil
}

/*获取OSS签名

获取OSS签名

*/
func (a *Client) PostOssSignature(params *PostOssSignatureParams, authInfo client.AuthInfoWriter) (*PostOssSignatureOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOssSignatureParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:       "PostOssSignature",
		Params:   params,
		Reader:   &PostOssSignatureReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostOssSignatureOK), nil
}

/*获取STS临时token

获取STS临时token

*/
func (a *Client) PostSts(params *PostStsParams, authInfo client.AuthInfoWriter) (*PostStsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostStsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:       "PostSts",
		Params:   params,
		Reader:   &PostStsReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostStsOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}

// NewAPIError creates a new API error
func NewAPIError(opName string, response interface{}, code int) APIError {
	return APIError{
		OperationName: opName,
		Response:      response,
		Code:          code,
	}
}

// APIError wraps an error model and captures the status code
type APIError struct {
	OperationName string
	Response      interface{}
	Code          int
}

func (a APIError) Error() string {
	return fmt.Sprintf("%s (status %d): %+v ", a.OperationName, a.Code, a.Response)
}
