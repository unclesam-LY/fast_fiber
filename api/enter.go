package api

import "FastFiber/api/user_api"

type Api struct {
	UserApi user_api.UserApi
}

var App = new(Api)
