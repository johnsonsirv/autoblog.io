package main

import (
	"net/http"

	"github.com/johnsonsirv/auto-blog/controllers"
)

func main() {
	controllers.RegisterRoutes()

	http.ListenAndServe(":3000", nil)
}
