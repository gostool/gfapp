package testsdb

import (
	"gfapp/app/model"
	"gfapp/app/service"
	"testing"
)

var testUserServiceReq model.UserRegisterServiceSignUpReq

func init() {
	testUserServiceReq = model.UserRegisterServiceSignUpReq{
		Password:     "123456",
		Passport:     "admin123",
		RegisterType: 1,
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
