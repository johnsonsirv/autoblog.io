package controllers

import (
	"net/http"
	"regexp"
)

type NewsController struct {
	IDPattern *regexp.Regexp
}

func (nc NewsController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Health"))
}

// Constructor Function = OOP concept
func NewNewsController() *NewsController {
	return &NewsController{
		IDPattern: regexp.MustCompile(`^/news/(\d+)/?`),
	}
}
