package testsdb

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/test/gtest"
	"github.com/gogf/gf/util/gconv"
	"testing"
)

func TestRedisSetGetDo(t *testing.T) {
	g.Redis().Do("SET", "k", "v")
	gtest.C(t, func(t *gtest.T) {
		v, _ := g.Redis().Do("GET", "k")
		t.AssertEQ(gconv.String(v), "v")
	})
}

func TestRedisSetGetDoVar(t *testing.T) {
	g.Redis().Do("SET", "k", "v")
	gtest.C(t, func(t *gtest.T) {
		v, _ := g.Redis().DoVar("GET", "k")
		t.AssertEQ(v.String(), "v")
	})
}