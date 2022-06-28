package routes

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Seyz123/lightshot/config"
	db "github.com/Seyz123/lightshot/database"
	"github.com/Seyz123/lightshot/database/models"
	"github.com/Seyz123/lightshot/utils"
	"github.com/rs/xid"
)

func UploadRoute(res http.ResponseWriter, req *http.Request) {
	addr := strings.Split(req.RemoteAddr, ":")[0]

	config, err := config.GetConfig()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("An error occured"))
		return
	} else {
		if config.Logs {
			log.Println("New connection from " + addr + " on /upload")
		}
	}

	if len(config.AuthorizedIPs) > 0 {
		if !utils.ContainsString(addr, config.AuthorizedIPs) {
			res.WriteHeader(http.StatusForbidden)
			res.Write([]byte("You are not authorized to use this service"))
			return
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

	db, err := db.GetDB()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("An error occured"))
		return
	}

	image := &models.Image{ID: xid.New().String(), Data: string(bytes), CreatedAt: time.Now()}

	err = db.Save(image).Error
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("An error occured"))
		return
	}

	var url string

	if config.Https {
		url = "https://" + config.Host + "/file/" + image.ID
	} else {
		url = "http://" + config.Host + "/file/" + image.ID
	}

	response := fmt.Sprintf("<response>\n<status>success</status>\n<share>%s</share>\n</response>", url)

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(response))
}
