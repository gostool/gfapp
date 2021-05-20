package tests

import (
	"github.com/gogf/gf/test/gtest"
	"github.com/gogf/gf/util/gvalid"
	"testing"
)

func TestCheck(t *testing.T) {
	rule := "length:6,16"
	gtest.C(t, func(t *gtest.T) {
		m := gvalid.Check("123456", rule, nil)
		t.AssertNil(m)
		m2 := gvalid.Check("12356", rule, "长度错误,必须在6-16之间")
		t.AssertEQ(m2.String(), "长度错误,必须在6-16之间")
	})
}
