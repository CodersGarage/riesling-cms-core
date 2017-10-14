package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/spf13/viper"
	"sync"
	"riesling-cms-core/app/api"
)

var RootRoute = mux.NewRouter()
var WaitGroup = sync.WaitGroup{}

func InitRoutes() {
	v1 := RootRoute.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/user/create", api.CreateUser).Methods("POST")

	go http.ListenAndServe(viper.GetString("app.uri"), RootRoute)
	WaitGroup.Add(1)
}
