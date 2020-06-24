package main

import (
    "log"
    "net/http"
    "github.com/Seyz123/lightshot/database"
    "github.com/Seyz123/lightshot/config"
    "github.com/Seyz123/lightshot/router"
)


func main() {
    config, err := config.GetConfig()
    if err != nil {
        log.Fatal("Cannot get config")
        return
    }
    
    err = database.Init()
    if err != nil {
        log.Fatal("Unable to initialize database")
        return
    }
    
    router := router.Init()
    
    log.Println("Web server is running on port " + config.Port)
    log.Fatal(http.ListenAndServe(":" + config.Port, router))
}