package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

var logger *glog.Logger

func init() {
	logger = g.Log("debug")
}

func GetVersion() (version string) {
	v := g.Config().GetVar("gfcli.build")
	buildConf := v.MapStrVar()
	version = buildConf["version"].String()
	return version
}
