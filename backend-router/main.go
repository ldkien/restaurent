package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"restaurant/backend-base/app"
	log "restaurant/backend-base/logger"
	"restaurant/backend-router/handler"
)

func main() {
	rtr := mux.NewRouter()

	rtr.Use(handler.AuthMiddleware)
	rtr.HandleFunc(app.API_LOGIN, handler.Login).Methods("POST")
	rtr.HandleFunc("/welcome", handler.Welcome)

	http.Handle("/", rtr)
	log.Logger.Info("Start service router ....")
	log.Logger.Info(http.ListenAndServe(":8000", nil))
}
