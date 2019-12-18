package controllers

import (
	"net/http"
	"regexp"
)

type UserController struct {
	pattern *regexp.Regexp
}

func (userController UserController) serveHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from user controller"))
}

func newUserController() *UserController {
	return &UserController{pattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
