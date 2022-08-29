package main

import (
	"github.com/MonteCarloClub/kether_gin/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.Static("/be", "./be")
	for api, method := range handler.Methods {
		method(r, api)
	}
	r.Run(":8000")
}
