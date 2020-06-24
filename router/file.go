package router

import (
    "net/http"
    "log"
    "github.com/Seyz123/lightshot/config"
    db "github.com/Seyz123/lightshot/database"
    "github.com/julienschmidt/httprouter"
)

func FileRoute(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
    id := params.ByName("id")
    
    config, err := config.GetConfig()
    if err != nil {
        res.Write([]byte("An error occured"))
        return
    } else {
        if config.Logs {
            log.Println("New connection from " + req.RemoteAddr + " on /file")
        }
    }
    
    image, err := db.Get(id)
    if err != nil {
        res.Write([]byte("An error occured"))
        
        return
    }
        
    res.Header().Set("Content-Type", "image/jpeg")
        
    res.Write(image.Data)
}