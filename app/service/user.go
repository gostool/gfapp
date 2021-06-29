package service

import (
	"context"
	"errors"
	"gfapp/app/dao"
	"gfapp/app/model"
	"gfapp/library/response"

	"github.com/gogf/gf/frame/g"
)

// 中间件管理服务
var User = userService{}

type userService struct{}

func (s *userService) InsertAndGetID(data *g.Map) (id int64, err error) {
	id, err = dao.User.InsertAndGetId(data)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// 用户注册
func (s *userService) RegisterAccount(r *model.UserRegisterServiceReq) (id int64, err error) {
	return s.InsertAndGetID(r.ToMap())
}

func (s *userService) RegisterEmail(r *model.UserRegisterServiceReq) (id int64, err error) {
	return s.InsertAndGetID(r.ToMap())
}

func (s *userService) RegisterPhone(r *model.UserRegisterServiceReq) (id int64, err error) {
	return s.InsertAndGetID(r.ToMap())
}

func (s *userService) Find(pk int64) (user *model.User, err error) {
	fields := g.ArrayStr{
		"id",
		"type",
		"passport",
		"name",
	}
	one, err := dao.User.Fields(fields).FindOne("id=? and is_deleted=?", pk, 0)
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, errors.New("user is not exist")
	}
	user = (*model.User)(nil)
	one.Struct(&user)
	return user, nil
}

// 用户注册 db.Table("user").Delete("uid", 10)
func (s *userService) Delete(r *model.UserRegisterServiceReq) error {
	data := g.Map{}
	if len(r.Passport) != 0 {
		data = g.Map{"passport": r.Passport}
	} else if len(r.Email) != 0 {
		data = g.Map{"passport": r.Email}
	} else {
		data = g.Map{"passport": r.Phone}
	}
	if _, err := dao.User.Delete(data); err != nil {
		return err
	}
	return nil
}

func (s *userService) UpdateName(uid int64, name string) error {
	_, err := dao.User.Where("id=?", uid).Update(g.Map{"name": name})
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) UpdatePwd(uid int64, password string) error {
	_, err := dao.User.Where("id=?", uid).Update(g.Map{"password": password})
	if err != nil {
		return err
	}
	return nil
}

// 判断用户是否已经登录
func (s *userService) IsSignedIn(ctx context.Context) bool {
	if v := Context.Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// 判断用户是否已经登录
func (s *userService) IsSignedInWeb(token string, id string, salt string) bool {
	serviceReq := model.ClaimServiceReq{
		Id:   id,
		Salt: salt,
	}
	_, err := Base.ParseJwt(&serviceReq, token)
	if err != nil {
		return false
	}
	return true
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func (s *userService) Login(r *model.UserServiceLoginReq) (data *response.User, err error) {
	fields := g.ArrayStr{
		"id",
		"passport",
	}
	one, err := dao.User.Fields(fields).FindOne("passport=? and is_deleted=? and password=?", r.Passport, 0, r.Password)
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, errors.New("账号或密码错误")
	}
	data = (*response.User)(nil)
	one.Struct(&data)
	return data, nil
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
//func (s *userService) LoginWeb(ctx context.Context, r *model.UserServiceSignInWebReq) (user *model.User, err error) {
//	user, err = dao.User.FindOne("passport=? and password=?", r.Username, r.Password)
//	if err != nil {
//		return nil, err
//	}
//	if user == nil {
//		return nil, errors.New("账号或密码错误")
//	}
//	Context.SetUserWeb(ctx, user)
//	return user, nil
//}

// 用户注销
func (s *userService) LogOut(ctx context.Context) error {
	return Session.RemoveUser(ctx)
}

// 获得用户信息详情
func (s *userService) GetProfile(ctx context.Context) *model.User {
	return Session.GetUser(ctx)
}

// 获得用户信息详情
func (s *userService) GetProfileWeb(ctx context.Context) *model.User {
	return Context.Get(ctx).UserWeb
}
