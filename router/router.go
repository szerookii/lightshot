package router

import "github.com/julienschmidt/httprouter"

func Init() *httprouter.Router {
    router := httprouter.New()
    
    router.GET("/view/:id", ViewRoute)
    router.GET("/file/:id", FileRoute)
    router.POST("/upload/:a/:b/", UploadRoute)
    
    return router
}