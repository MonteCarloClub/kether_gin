package handler

import (
	"github.com/gin-gonic/gin"
)

var Methods = map[string]func(r *gin.Engine, api string){
	"api/get":              RegisterGetData,
	"api/put":              RegisterPutData,
	"api/data/account":     RegisterGetAccountData,
	"api/data/account/set": RegisterSetAccountData,
}
