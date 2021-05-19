package config

import (
	"github.com/gogf/gf/frame/g"
	"testing"
)

func TestCfg(t *testing.T) {
	address := g.Config().Get("server.Address")
	t.Logf("Address:%v\n", address)
}

func TestCfgSetPath(t *testing.T) {
	g.Cfg().SetFileName("config_tpl.toml")
	address := g.Config().Get("server.Address")
	if address.(string) != ":8299" {
		t.Logf("addr(%v) != :8299 in config_tpl.toml", address)
		t.Fail()
	}
}

func TestLogger(t *testing.T) {
	g.Cfg().SetFileName("config_tpl.toml")
	// 多个日志实例
	// 1.默认日志
	g.Log().Info("i am in logger default")
	g.Log("debug").Info("i am in logger debug")
	g.Log("test").Info("i am in logger test")
}
