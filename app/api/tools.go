package api

import (
	"gfapp/app/model"
	"gfapp/app/service"
	"gfapp/library/file/oss"
	"gfapp/library/response"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/encoding/gbase64"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"
)

var Tools = new(toolsApi)

type toolsApi struct{}

// @Tags tools
// @Summary 生成验证码
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /api/tools/captcha [post]
func (t *toolsApi) Captcha(r *ghttp.Request) {
	result, err := service.Base.Captcha()
	//g.Log().Debugf("result:%v", result)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", result)
	}
}

// @Tags tools
// @Summary 生成二维码
// @param   entity  body model.UserApiQrcodeReq true "qrcode请求"
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"二维码获取成功"}"
// @Router /api/tools/qrcode [post]
func (t *toolsApi) Qrcode(r *ghttp.Request) {
	var (
		apiReq     *model.UserApiQrcodeReq
		serviceReq *model.UserServiceQrcodeReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, response.CODE_BAD, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, response.CODE_ERR, err.Error())
	}
	result, err := service.Base.Qrcode(serviceReq.Url, serviceReq.Size)
	result = gbase64.Encode(result)
	pre := "data:image/png;base64,"
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", pre+string(result))
	}
}

// @Tags tools
// @Summary 生成二维码
// @param   entity  body model.UserApiQrcodeReq true "qrcode请求"
// @Produce image/png
// @Success 200 {string} string "{"success":true,"data":{},"msg":"二维码获取成功"}"
// @Router /api/tools/qrcode-img [post]
func (t *toolsApi) QrcodeImg(r *ghttp.Request) {
	var (
		apiReq     *model.UserApiQrcodeReq
		serviceReq *model.UserServiceQrcodeReq
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, response.CODE_BAD, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, response.CODE_ERR, err.Error())
	}
	result, err := service.Base.Qrcode(serviceReq.Url, serviceReq.Size)
	if err != nil {
		response.JsonExit(r, response.CODE_ERR, err.Error())
	}
	r.Response.Header().Set("Content-Type", "image/png")
	r.Response.Write(result)
	r.Exit()
}

// Version
// @Tags tools
// @Summary 获取版本号
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"版本号获取成功"}"
// @Router /api/tools/version [get]
func (t *toolsApi) Version(r *ghttp.Request) {
	version := GetVersion()
	response.JsonExit(r, response.CODE_OK, "ok", g.Map{"version": version})
}

// @Tags tools
// @Summary 文件上传到本地
// @accept multipart/form-data
// @Param file formData file true "file"
// @Produce application/json
// @Success 200 {string} string response.JsonResponse
// @Router /api/tools/upload [post]
func (t *toolsApi) Upload(r *ghttp.Request) {
	file := r.GetUploadFile("file")
	file.Filename = oss.GetUUID() + "-" + file.Filename
	logger.Debugf("file:%v", file)
	filePath, err := service.File.Save(file)
	if err != nil {
		response.JsonExit(r, response.CODE_ERR, err.Error())
	}
	response.JsonExit(r, response.CODE_OK, "ok", g.Map{"link": filePath})
}

// @Tags tools
// @Summary 文件上传到oss
// @accept multipart/form-data
// @Param file formData file true "file"
// @Produce application/json
// @Success 200 {string} string response.JsonResponse
// @Router /api/tools/upload-oss [post]
func (t *toolsApi) UploadOss(r *ghttp.Request) {
	file := r.GetUploadFile("file")
	logger.Debugf("file:%v", file)
	fileUrl, _, err := service.File.SaveOss(file)
	if err != nil {
		response.JsonExit(r, response.CODE_ERR, err.Error())
	}
	response.JsonExit(r, response.CODE_OK, "ok", g.Map{"link": fileUrl})
}

//// @Tags tools
//// @Summary 生成token
//// @accept application/json
//// @Produce application/json
//// @Success 200 {string} string response.JsonResponse
//// @Router /api/tools/new-jwt [post]
//func (t *toolsApi) NewJwt(r *ghttp.Request) {
//	serviceReq := model.ClaimServiceReq{
//		Id: "1",
//	}
//	token, err := service.Base.NewJwt(&serviceReq)
//	if err != nil {
//		response.JsonExit(r, 1, err.Error())
//	} else {
//		response.JsonExit(r, 0, "ok", token)
//	}
//}

func (t *toolsApi) ParseJwt(r *ghttp.Request) {
	serviceReq := model.ClaimServiceReq{
		Id:   "1",
		Salt: "",
	}
	tokenString := r.Header.Get("Token")
	claim, err := service.Base.ParseJwt(&serviceReq, tokenString)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok", claim.Id, claim.StandardClaims.ExpiresAt)
}
