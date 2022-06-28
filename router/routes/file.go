package routes

import (
	"log"
	"net/http"

	"github.com/Seyz123/lightshot/config"
	"github.com/gorilla/mux"
)

func FileRoute(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	config, err := config.GetConfig()
	if err != nil {
		res.Write([]byte("An error occured"))
		return
	} else {
		if config.Logs {
			log.Println("New connection from " + req.RemoteAddr + " on /file")
		}
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(id))
}
