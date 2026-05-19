package users_transport_http

import (
	"net/http"

	core_http_server "95.174.104.37/gitlab/vifrolov/todo_app/internal/core/transport/server"
)

type UsersHttpHandler struct {
	usersService UsersService
}

type UsersService interface{}

func NewUsersHttpHandler(usersService UsersService) *UsersHttpHandler {
	return &UsersHttpHandler{
		usersService: usersService,
	}
}

func (h *UsersHttpHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
	}
}
