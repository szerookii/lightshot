package router

import (
    "fmt"
    "bytes"
    "io"
    "net/http"
    "log"
    "github.com/Seyz123/lightshot/config"
    db "github.com/Seyz123/lightshot/database"
    "github.com/julienschmidt/httprouter"
)

func UploadRoute(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    config, err := config.GetConfig()
    if err != nil {
        res.Write([]byte("An error occured"))
        return
    } else {
        if config.Logs {
            log.Println("New connection from " + req.RemoteAddr + " on /upload")
        }
    }
    
    file, _, err := req.FormFile("image")
    if err != nil {
        res.Write([]byte("An error occured"))
        return
    }
    
    buf := new(bytes.Buffer)
    io.Copy(buf, file)
    bytes := buf.Bytes()
    
    id, err := db.Save(bytes)
    if err != nil {
        res.Write([]byte("An error occured"))
        return
    }
    
    var url string
    
    if config.Https {
        url = "https://" + config.Host + "/view/" + id
    } else {
        url = "http://" + config.Host + "/view/" + id
    }
    
    response := fmt.Sprintf("<response>\n<status>success</status>\n<share>%s</share>\n</response>", url)
    
    res.Write([]byte(response))
}