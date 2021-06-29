package model

// Fill with you ideas below.
// 登录请求参数，用于前后端交互参数格式约定
type UserApiQrcodeReq struct {
	Url  string `v:"required#url不能为空"`
	Size int    `v:"required#size不能为空"`
}

type UserServiceQrcodeReq struct {
	Url  string `json:"url"`
	Size int    `json:"size"`
}
