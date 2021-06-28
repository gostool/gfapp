package testsdb

import (
	"gfapp/app/model"
	"gfapp/app/service"
	"testing"
)

var testUserServiceReq model.UserRegisterServiceReq
var testUserServiceMailReq model.UserRegisterServiceReq
var testUserServicePhoneReq model.UserRegisterServiceReq

func init() {
	testUserServiceReq = model.UserRegisterServiceReq{
		Password:     "123456",
		Passport:     "admin04",
		RegisterType: 1,
	}
	testUserServiceMailReq = model.UserRegisterServiceReq{
		Email:        "1353176359@qq.com",
		Password:     "123456",
		RegisterType: 2,
	}
	testUserServicePhoneReq = model.UserRegisterServiceReq{
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
	id, err := service.User.RegisterAccount(&testUserServiceReq)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("ID:%v\n", id)
	}
}

func TestServiceMailRegister(t *testing.T) {
	err := service.User.Delete(&testUserServiceMailReq)
	if err != nil {
		t.Fatal(err)
	}
	id, err := service.User.RegisterEmail(&testUserServiceMailReq)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("ID:%v\n", id)
	}
}

func TestServicePhoneRegister(t *testing.T) {
	err := service.User.Delete(&testUserServicePhoneReq)
	if err != nil {
		t.Fatal(err)
	}
	id, err := service.User.RegisterPhone(&testUserServicePhoneReq)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("ID:%v\n", id)
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
