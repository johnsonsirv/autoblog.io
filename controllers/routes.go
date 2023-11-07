package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterRoutes() {
	nc := NewNewsController()

	http.Handle("/news", nc)
	http.Handle("/news/", nc)
}


func encodeResponseAsJSON(data interface{}, w io.Writer)  {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}