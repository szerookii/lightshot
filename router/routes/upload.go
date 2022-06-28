package routes

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Seyz123/lightshot/config"
	db "github.com/Seyz123/lightshot/database"
	"github.com/Seyz123/lightshot/database/models"
	"github.com/Seyz123/lightshot/utils"
)

func UploadRoute(res http.ResponseWriter, req *http.Request) {
	config, err := config.GetConfig()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("An error occured"))
		return
	} else {
		if config.Logs {
			log.Println("New connection from " + req.RemoteAddr + " on /upload")
		}
	}

	file, _, err := req.FormFile("image")
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Bad request"))
		return
	}

	buf := new(bytes.Buffer)
	io.Copy(buf, file)
	bytes := buf.Bytes()

	fmt.Println(bytes)

	db := db.GetDB()
	image := &models.Image{ID: utils.GenerateId(12), Data: string(bytes) CreatedAt: time.Now()}

	err = db.Save(image).Error
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("An error occured"))
		return
	}

	var url string

	if config.Https {
		url = "https://" + config.Host + "/view/" + image.ID
	} else {
		url = "http://" + config.Host + "/view/" + image.ID
	}

	response := fmt.Sprintf("<response>\n<status>success</status>\n<share>%s</share>\n</response>", url)

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(response))
}
