package config

import (
	mapset "github.com/deckarep/golang-set"
)

// 所有URL
const (
	ApiUrl    string = "/api"
	RegistUrl string = "/regist"
	LoginUrl  string = "/login"
)

// 允许直接通过的URL
var AllowPathSet = mapset.NewSet(ApiUrl+LoginUrl, ApiUrl+RegistUrl)
