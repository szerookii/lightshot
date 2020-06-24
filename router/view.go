package router

import (
    "net/http"
    "path/filepath"
    "html/template"
    "log"
    "github.com/Seyz123/lightshot/config"
    db "github.com/Seyz123/lightshot/database"
    "github.com/julienschmidt/httprouter"
)

func ViewRoute(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
    id := params.ByName("id")
    
    config, err := config.GetConfig()
    if err != nil {
        res.Write([]byte("An error occured"))
        return
    } else {
        if config.Logs {
            log.Println("New connection from " + req.RemoteAddr + " on /view")
        }
    }
    
    image, err := db.Get(id)
    if err != nil {
        res.Write([]byte("An error occured"))
        return
    }
    
    t, err := template.ParseFiles(filepath.Join("views", "view.templ"))
    
    if err != nil {
        res.Write([]byte("An error occured"))
        return
    }
        
    t.Execute(res, image)
}