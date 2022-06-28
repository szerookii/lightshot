package routes

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Seyz123/lightshot/config"
	"github.com/gorilla/mux"
)

func ViewRoute(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	config, err := config.GetConfig()
	if err != nil {
		res.Write([]byte("An error occured"))
		return
	} else {
		if config.Logs {
			log.Println("New connection from " + req.RemoteAddr + " on /view")
		}
	}

	t, err := template.ParseFiles(filepath.Join("views", "view.templ"))

	if err != nil {
		res.Write([]byte("An error occured"))
		return
	}

	t.Execute(res, id)
}
