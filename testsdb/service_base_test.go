package testsdb

import (
	"gfapp/app/model"
	"gfapp/app/service"
	"testing"
)

func TestServiceJwt(t *testing.T) {
	req := model.ClaimServiceReq{
		Id: "2",
	}
	token, err := service.Base.NewJwt(&req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("token:%v", token)
}
