package tests

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/test/gtest"
	"github.com/gogf/gf/util/gvalid"
	"testing"
)

func TestValidCheck(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// 单条数据校验
		var dataTests = []struct {
			input string
			msg string
			rule string
		}{
			{
				"123456",
				"长度为6-16",
				"length:6,16",
			},
			{
				"12346",
				"长度错误，必须小于16",
				"length:6,16",
			},
			{
				"12346789101234567",
				"长度错误，必须小于16",
				"length:6,16",
			},
			// integer 验证
			{
				"10",
				"请输入一个整数|范必须在[18:200]",
				"integer|between:6,16",
			},
			{
				"a",
				"请输入一个整数|范围必须在[18:200]",
				"integer|between:18,200",
			},
		}
		for _, data := range dataTests {
			m := gvalid.Check(data.input, data.rule, data.msg)
			if m != nil {
				t.Log(m.Error())
				t.AssertEQ(m.String(), data.msg)
			}
		}
	})
}

// Map数据校验
func TestValidStruct(t *testing.T) {
	params := map[string]interface{}{
		"passport":  "john",
		"password":  "123456",
		"password2": "1234567",
	}
	rules := map[string]string{
		"passport":  "required|length:6,16",
		"password":  "required|length:6,16|same:password2",
		"password2": "required|length:6,16",
	}
	msgs := map[string]interface{}{
		"passport": "账号不能为空|账号长度应当在:min到:max之间",
		"password": map[string]string{
			"required": "密码不能为空",
			"same":     "两次密码输入不相等",
		},
	}
	if e := gvalid.CheckMap(params, rules, msgs); e != nil {
		//g.Dump(e.Map())
		g.Dump(e.Maps())
	} else {
		t.Log("check ok!")
	}
}