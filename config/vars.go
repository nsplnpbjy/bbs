package config

import (
	mapset "github.com/deckarep/golang-set"
)

// 所有URL
const (
	ApiUrl        string = "/api"
	RegistUrl     string = "/regist"
	LoginUrl      string = "/login"
	PostIdeaUrl   string = "/postidea"
	DeleteIdeaUrl string = "/deleteidea"
)

// 允许直接通过的URL
var AllowPathSet = mapset.NewSet(ApiUrl+LoginUrl, ApiUrl+RegistUrl)

// 密码加密COST
const (
	DefaultCost int = 10
)

// 请求方式
const (
	Method_GET    MethodType = "GET"
	Method_POST   MethodType = "POST"
	Method_DELETE MethodType = "DELETE"
)
