package tests

import (
	"context"
	"testing"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/test/gtest"
	"github.com/gogf/gf/util/gvalid"
)

func TestValidCheckStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// 单条数据校验
		var dataTests = []struct {
			input string
			msg   string
			rule  string
		}{
			{
				"123456",
				"长度为6-16",
				"length:6,16",
			},
			{
				"123456",
				"长度错误，必须小于16",
				"length:6,16",
			},
			{
				"123456",
				"长度错误，必须小于16",
				"length:6,16",
			},
			{
				"23@163.com",
				"email格式错误",
				"email",
			},
		}
		for _, data := range dataTests {
			m := gvalid.CheckValue(context.TODO(), data.input, data.rule, data.msg)
			if m != nil {
				t.Log(m.Error())
				t.AssertEQ(m.String(), data.msg)
			}
		}
	})
}

func TestCheckValidInt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		// 单条数据校验
		var dataTests = []struct {
			input int
			msg   string
			rule  string
		}{
			// integer 验证
			{
				15,
				"请输入一个整数|范必须在[6:16]",
				"integer|between:6,16",
			},
			{
				18,
				"请输入一个整数|范围必须在[18:200]",
				"integer|between:18,200",
			},
		}
		for _, data := range dataTests {
			m := gvalid.CheckValue(context.TODO(), data.input, data.rule, data.msg)
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
		"passport":  "johnHello",
		"password":  "123456",
		"password2": "123456",
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
	if e := gvalid.CheckMap(context.TODO(), params, rules, msgs); e != nil {
		//g.Dump(e.Map())
		g.Dump(e.Maps())
	} else {
		t.Log("check ok!")
	}
}

func TestStruct(t *testing.T) {
	type Pass struct {
		Pass1 string `valid:"password1@required|length:6,30|same:password2#请输入您的密码|长度必须在6-30|您两次输入的密码不一致"`
		Pass2 string `valid:"password2@required|length:6,30|same:password1#请再次输入您的密码|长度必须在6-30|您两次输入的密码不一致"`
	}
	type User struct {
		Id   int    `v:"id      @integer|min:1#|请输入用户ID"`
		Name string `v:"name     @required|length:6,30#请输入用户名称|用户名称长度非法"`
		Pass
	}

	user := &User{
		Id:   1,
		Name: "john12",
		Pass: Pass{
			Pass1: "124@1534",
			Pass2: "124@1534",
		},
	}

	// 使用结构体定义的校验规则和错误提示进行校验
	if e := gvalid.CheckStruct(context.TODO(), user, nil); e != nil {
		g.Dump(e.Maps())
	}
}
