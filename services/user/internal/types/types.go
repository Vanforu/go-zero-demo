// Code generated by goctl. DO NOT EDIT.
package types

type GetUserInfoReq struct {
	Id string `path:"id"`
}

type GetUserInfoResp struct {
	Data   UserItem `json:"data"`
	Status int64    `json:"status"`
	Msg    string   `json:"msg"`
}

type LoginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginResp struct {
	Data   UserItem `json:"data"`
	Status int64    `json:"status"`
	Token  string   `json:"token"`
	Msg    string   `json:"msg"`
}

type RegisterReq struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type RegisterResp struct {
	Status int64  `json:"status"`
	Token  string `json:"token"`
	Msg    string `json:"msg"`
}

type UserItem struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Type       int64  `json:"type"`
	CreateTime int64  `json:"createTime"`
	Status     int64  `json:"status"`
}
