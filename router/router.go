package router

import (
	"log"
	"net/http"
	"time"

	"github.com/Seyz123/lightshot/config"
	"github.com/Seyz123/lightshot/router/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Init() {
	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/file/{id}", routes.FileRoute).Methods("GET")
	r.HandleFunc("/upload/{a}/{b}/", routes.UploadRoute).Methods("POST")

	log.Println("Starting server on port " + config.Port)

	srv := &http.Server{
		Handler:      handlers.ProxyHeaders(r),
		Addr:         ":" + config.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
