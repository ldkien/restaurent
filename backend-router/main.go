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
	rtr.HandleFunc(app.ApiLogin, handler.Login).Methods("POST")
	rtr.HandleFunc(app.ApiRegister, handler.Register).Methods("POST")
	rtr.HandleFunc(app.ApiOrder, handler.CreateOrder).Methods("POST")
	rtr.HandleFunc(app.ApiOrder, handler.UpdateOrder).Methods("PUT")

	http.Handle("/", rtr)
	log.Logger.Info("Start service router ....")
	log.Logger.Info(http.ListenAndServe(":8000", nil))
}
