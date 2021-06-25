package model

import "github.com/gogf/gf/frame/g"

// 注册请求参数，用于前后端交互参数格式约定
type UserRegisterApiSignUpReq struct {
	Passport  string `v:"required|length:6,16#账号不能为空|账号长度应当在:min到:max之间"`
	Password  string `v:"required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间"`
	Password2 string `v:"required|length:6,16|same:Password#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等"`
	CaptchaReq
	RegisterType int `v:"required|integer|between:1,3" json:"type"`
}

// 注册请求参数，用于前后端交互参数格式约定
type UserRegisterApiMailSignUpReq struct {
	Email     string `v:"required|email"`
	Password  string `v:"required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间"`
	Password2 string `v:"required|length:6,16|same:Password#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等"`
	CaptchaReq
	RegisterType int `v:"required|integer|between:1,3" json:"type"`
}

// 注册请求参数，用于前后端交互参数格式约定
type UserRegisterApiPhoneSignUpReq struct {
	Phone        string `v:"required|phone"`
	Password     string `v:"required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间"`
	Password2    string `v:"required|length:6,16|same:Password#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等"`
	RegisterType int    `v:"required|integer|between:1,3" json:"type"`
	SmsCode      string `v:"required"|length:6`
}

// 注册输入参数
type UserRegisterServiceReq struct {
	Passport     string
	Password     string
	Email        string
	Phone        string
	RegisterType int `json:"type"`
}

func (r *UserRegisterServiceReq) ToMap() (data *g.Map) {
	if r.RegisterType == 1 { //account
		data = &g.Map{
			"passport":   r.Passport,
			"password":   r.Password,
			"type":       r.RegisterType,
			"is_deleted": 0,
		}
	} else if r.RegisterType == 2 { //email
		data = &g.Map{
			"passport":   r.Email,
			"password":   r.Password,
			"type:":      r.RegisterType,
			"email":      r.Email,
			"is_deleted": 0,
		}
	} else { //phone
		data = &g.Map{
			"passport":   r.Phone,
			"password":   r.Password,
			"type":       r.RegisterType,
			"phone":      r.Phone,
			"is_deleted": 0,
		}
	}
	return
}
