package config

import (
	"github.com/gogf/gf/frame/g"
	"testing"
)

func TestCfg(t *testing.T) {
	address := g.Config().Get("server.Address")
	t.Logf("Address:%v\n", address)
}