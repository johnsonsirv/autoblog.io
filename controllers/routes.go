package controllers

import "net/http"

func RegisterRoutes() {
	nc := NewNewsController()

	http.Handle("/news", nc)
	http.Handle("/news/", nc)
}
