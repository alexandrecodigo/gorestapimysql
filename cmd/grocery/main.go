package main

import (
	"log"

	"github.com/alexandrecodigo/gorestapimysql/internal/app/grocery/webapi"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Print("Starting the grocery app")
	r := gin.Default()
	webapi.Route(r)
	r.Run()
}
