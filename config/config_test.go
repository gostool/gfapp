package config

import (
	"github.com/gogf/gf/frame/g"
	"testing"
)

func TestCfg(t *testing.T) {
	address := g.Config().Get("server.Address")
	t.Logf("Address:%v\n", address)
	g.Log("abc")
}

func TestCfgSetPath(t *testing.T) {
	g.Cfg().SetFileName("config_tpl.toml")
	address := g.Config().Get("server.Address")
	if address.(string) != ":8299" {
		t.Logf("addr(%v) != :8299 in config_tpl.toml", address)
		t.Fail()
	}
}