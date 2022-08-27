package handler

import (
	"github.com/gin-gonic/gin"
)

var Methods = map[string]func(r *gin.Engine, api string){
	"/login/account": RegisterGetAccountData,
}

func registerMethod(r *gin.Engine, api string, f gin.HandlerFunc) {
	r.GET(api, f)
}
