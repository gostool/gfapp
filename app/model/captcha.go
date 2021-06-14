package model

type CaptchaReq struct {
	CaptchaId string `v:"required|length:0,64#验证码id不能为空|验证码id应当在:min到:max之间" json:"captchaId"`
	Captcha   string `v:"required|length:6,6#验证码不能为空|验证码长度应当在:min到:max之间" json:"captcha"`
}
