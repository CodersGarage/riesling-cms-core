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
	v1.HandleFunc("/user/list/{hash}", api.GetUser).Methods("GET")
	v1.HandleFunc("/user/update/{hash}", api.SelfAuth(api.UpdateUser)).Methods("PUT")

	v1.HandleFunc("/auth", api.CreateSession).Methods("POST")   // Create Session
	v1.HandleFunc("/auth", api.CheckSession).Methods("GET")     // Verify Session
	v1.HandleFunc("/auth", api.DeleteSession).Methods("DELETE") // Delete Session
	v1.HandleFunc("/auth", api.ReCreateSession).Methods("PUT")  // Refresh Session

	go http.ListenAndServe(viper.GetString("app.uri"), RootRoute)
	WaitGroup.Add(1)
}
