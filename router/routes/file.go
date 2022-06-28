package routes

import (
	"log"
	"net/http"

	"github.com/Seyz123/lightshot/config"
	"github.com/Seyz123/lightshot/database"
	"github.com/Seyz123/lightshot/database/models"
	"github.com/gorilla/mux"
)

func FileRoute(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	config, err := config.GetConfig()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("An error occured"))
		return
	} else {
		if config.Logs {
			log.Println("New connection from " + req.RemoteAddr + " on /file")
		}
	}

	db, err := database.GetDB()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("An error occured"))
		return
	}

	var image models.Image

	err = db.Where("id = ?", id).First(&image).Error
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Image not found"))
		return
	}

	res.Header().Set("Content-Type", "image/png")
	res.WriteHeader(http.StatusOK)

	res.Write([]byte(image.Data))
}
