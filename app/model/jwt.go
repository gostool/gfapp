package model

//jwt generate token
type TokenServiceGenTokenReq struct {
	Uid string `json:"uid,omitempty"`
	Exp int64  `json:"exp,omitempty"`
}

type TokenServiceCheckTokenReq struct {
	Token string `json:"token"`
}
