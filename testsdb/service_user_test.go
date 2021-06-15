package testsdb

import (
	"gfapp/app/model"
	"gfapp/app/service"
	"testing"
)

var testUserServiceReq model.UserRegisterServiceSignUpReq
var testUserServiceMailReq model.UserRegisterServiceSignUpReq
var testUserServicePhoneReq model.UserRegisterServiceSignUpReq

func init() {
	testUserServiceReq = model.UserRegisterServiceSignUpReq{
		Password:     "123456",
		Passport:     "admin04",
		RegisterType: 1,
	}
	testUserServiceMailReq = model.UserRegisterServiceSignUpReq{
		Email:        "1353176359@qq.com",
		Password:     "123456",
		RegisterType: 2,
	}
	testUserServicePhoneReq = model.UserRegisterServiceSignUpReq{
		Phone:        "17792301520",
		Password:     "123456",
		RegisterType: 3,
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

func TestServiceMailRegister(t *testing.T) {
	err := service.User.Delete(&testUserServiceMailReq)
	if err != nil {
		t.Fatal(err)
	}
	err = service.User.Register(&testUserServiceMailReq)
	if err != nil {
		t.Fatal(err)
	}
}

func TestServicePhoneRegister(t *testing.T) {
	err := service.User.Delete(&testUserServicePhoneReq)
	if err != nil {
		t.Fatal(err)
	}
	err = service.User.Register(&testUserServicePhoneReq)
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
	err = service.User.Delete(&testUserServiceMailReq)
	if err != nil {
		t.Fatal(err)
	}
	err = service.User.Delete(&testUserServicePhoneReq)
	if err != nil {
		t.Fatal(err)
	}
}
