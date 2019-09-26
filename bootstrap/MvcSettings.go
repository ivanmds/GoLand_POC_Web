package bootstrap

import (
	"GoLand_POC_Web/controllers"
	"net/http"
)

func Load() {
	http.HandleFunc("/hello", controllers.Hello)
	http.ListenAndServe(":8080", nil)
}
