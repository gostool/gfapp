package testsdb

import (
	"gfapp/app/model"
	"gfapp/app/service"
	"testing"
)

var testUserServiceReq model.UserServiceSignUpReq

func init() {
	testUserServiceReq = model.UserServiceSignUpReq{
		Password: "app01",
		Passport: "app02",
		Nickname: "app02",
	}
}

func TestServiceRegister(t *testing.T) {
	err := service.User.Delete(&testUserServiceReq)
	if err != nil {
		t.Fatal(err)
	}
	err = service.User.Register(&testUserServiceReq)
	if err != nil {
		t.Fatal(err)
	}
}

func TestServiceSignIn(t *testing.T) {

}

func TestServiceDelete(t *testing.T) {
	err := service.User.Delete(&testUserServiceReq)
	if err != nil {
		t.Fatal(err)
	}
}
