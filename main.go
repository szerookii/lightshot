package main

import (
	"github.com/Seyz123/lightshot/database"
	"github.com/Seyz123/lightshot/router"
)

func main() {
	database.Init()
	router.Init()
}
