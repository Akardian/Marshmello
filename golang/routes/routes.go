package routes

import (
	"lookup/handler"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/api/hello", handler.HandleHello)
	http.HandleFunc("/api/data", handler.HandleData)

	http.HandleFunc("/api/user", handler.HandleUser)
	http.HandleFunc("/api/user/icon", handler.HandleUserIcon)

}
