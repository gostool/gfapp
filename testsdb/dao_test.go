package testsdb

import (
	"errors"
	"gfapp/app/dao"
	"testing"

	"github.com/gogf/gf/frame/g"
)

func TestFindOne(t *testing.T) {
	password := "654321"
	passport := "admin04"
	one, err := dao.User.FindOne("password=? and passport=?", password, passport)
	if err != nil {
		t.Fatal(err)
	}
	if one.IsEmpty() {
		t.Fatal(errors.New("user is not exist"))
	}
	id := one["id"]
	g.Dump(id)
	g.Dump(one.Json())
}
