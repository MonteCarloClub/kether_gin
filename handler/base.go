package handler

import (
	"github.com/gin-gonic/gin"
)

var Methods = map[string]func(r *gin.Engine, api string){
	"/get":          RegisterGetData,
	"/put":          RegisterPutData,
	"/data/account": RegisterGetAccountData,
}
