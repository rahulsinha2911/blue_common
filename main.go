package main

import (
	"blue_common/base"
	"blue_common/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	base.Init()
	router := gin.Default()
	routes.InitRoutes(router)
	err := router.Run(fmt.Sprintf(":%v", os.Getenv("GO_PORT")))
	if err != nil {
		panic(err)
	}
}
