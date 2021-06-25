package service

import (
	"gfapp/app/model"
	"gfapp/library/jwt"
	"time"
)

var Token = jwtService{}

type jwtService struct{}

func (t *jwtService) GenAccessToken(r *model.TokenServiceGenTokenReq) (token string, err error) {
	secret := jwt.SecSecret(r.Uid, jwtSalt)
	logger.Debugf("uid:%v, secret:%v, exp:%v", r.Uid, secret, r.Exp)
	token, err = jwt.CreateToken(r.Uid, secret, r.Exp)
	if err != nil {
		logger.Errorf("req:%v, err:%v, token:%v", r, err, token)
		return "", err
	}
	return token, nil
}

func (t *jwtService) GenToken(uid string, exp int64) (token string, err error) {
	if exp <= 0 {
		exp = time.Now().Add(time.Duration(jwtExp) * time.Second).Unix()
	}
	return t.GenAccessToken(&model.TokenServiceGenTokenReq{
		Uid: uid,
		Exp: exp,
	})
}

func (t *jwtService) CheckToken(token string) (uid string, err error) {
	uid, err = jwt.GetUid(token)
	if err != nil {
		return uid, err
	}
	secret := jwt.SecSecret(uid, jwtSalt)
	uid, err = jwt.AuthToken(token, secret)
	if err != nil {
		logger.Errorf("req:%v, err:%v, uid:%v", token, err, uid)
		return uid, err
	}
	return uid, nil
}
