package config

import (
	mapset "github.com/deckarep/golang-set"
)

// 所有URL
const (
	ApiUrl                    string = "/api"
	RegistUrl                 string = "/regist"
	LoginUrl                  string = "/login"
	PostIdeaUrl               string = "/postidea"
	DeleteIdeaUrl             string = "/deleteidea"
	ShowNewestIdeasUrl        string = "/shownewestideas"
	ShowAllIdeasByUserInfoUrl string = "/showallideasbyuserinfo"
)

// 允许直接通过的URL
var AllowPathSet = mapset.NewSet(ApiUrl+LoginUrl, ApiUrl+RegistUrl)

// 用户密码加密COST
const (
	DefaultCost int = 10
)

// 请求方式
const (
	Method_GET    MethodType = "GET"
	Method_POST   MethodType = "POST"
	Method_DELETE MethodType = "DELETE"
)

// 每次返回帖子数量
const (
	IdeasNum int64 = 30
)
