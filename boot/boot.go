package boot

import (
	"gfapp/library/utils"
	_ "gfapp/packed"
	"os"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})

	//static files
	v := g.Config().GetVar("upload")
	uploadConf := v.MapStrVar()
	enabled := uploadConf["enabled"].Bool()
	if enabled {
		uploadDir := uploadConf["dir"].String()
		uploadUrl := uploadConf["url"].String()
		if !utils.IsDir(uploadDir) {
			_ = os.MkdirAll(uploadDir, os.ModePerm)
		}
		s.SetIndexFolder(enabled)
		s.SetServerRoot(uploadDir)
		s.AddStaticPath(uploadUrl, uploadDir)
	}
}
