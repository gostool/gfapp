package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

var logger *glog.Logger

func init() {
	logger = g.Log("debug")
}
