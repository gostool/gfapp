package model

// Fill with you ideas below.
// 登录请求参数，用于前后端交互参数格式约定
type UserApiQrcodeReq struct {
	Url  string `v:"required|url|length:4,1024#url不能为空|url格式错误|url长度应该在min到max之间"`
	Size int    `v:"required#size不能为空"`
}

type UserServiceQrcodeReq struct {
	Url  string `json:"url"`
	Size int    `json:"size"`
}
