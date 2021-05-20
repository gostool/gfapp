package config

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/test/gtest"
	"testing"
)
// 测试针对: config_tpl.toml 进行测试.
// 项目中真实使用.config.toml.
// config_tpl.toml 必须实时同步.
// 更新config_tpl.toml 需要更新本测试文件. git action会进行测试
// 本地测试: go test -v ./config

func TestCfgSetPath(t *testing.T) {
	g.Cfg().SetFileName("config_tpl.toml")
	v := g.Config().GetVar("server.Address")
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(v.String(), ":8299")
	})
}

func TestLogger(t *testing.T) {
	g.Cfg().SetFileName("config_tpl.toml")
	// 多个日志实例
	// 1.默认日志
	g.Log().Info("i am in logger default")
	g.Log("debug").Info("i am in logger debug")
	g.Log("test").Info("i am in logger test")
}

// add redis conf test
func TestCfgRedis(t *testing.T) {
	g.Cfg().SetFileName("config_tpl.toml")
	v := g.Config().GetVar("redis")
	redis := v.MapStrStr()
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(redis["default"], "127.0.0.1:6379,0")
		t.AssertEQ(redis["cache"], "127.0.0.1:6379,1,123456?idleTimeout=600")
	})
}