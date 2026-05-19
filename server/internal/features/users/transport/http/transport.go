package users_transport_http

import (
	"net/http"

	core_http_server "github.com/VitalyCoder/todo_go/internal/core/transport/server"
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
