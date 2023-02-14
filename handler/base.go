package handler

import (
	"github.com/gin-gonic/gin"
)

var Methods = map[string]func(r *gin.Engine, api string){
	"api/get/block":               RegisterGetBlockByHash,
	"api/get/account/transaction": RegisterGetAccountTransaction,
	"api/data/account":            RegisterGetAccountData,
	"api/data/account/set":        RegisterTransaction,
	"api/get":                     RegisterGet,
	"api/put":                     RegisterPut,
}
