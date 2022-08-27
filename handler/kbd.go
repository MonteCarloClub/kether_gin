package handler

import (
	"context"
	"net/http"

	kdb "github.com/MonteCarloClub/KBD/kitex_gen/api"
	"github.com/MonteCarloClub/kether_gin/client"
	"github.com/gin-gonic/gin"
)

func RegisterGetAccountData(r *gin.Engine, api string) {
	registerMethod(r, api, func(ctx *gin.Context) {
		req := &kdb.GetAccountDataRequest{
			Address: ctx.Query("address"),
		}
		resp, err := client.KBDClient.GetAccountData(context.Background(), req)
		if err != nil {
			ctx.String(http.StatusBadRequest, "%v", err.Error())
		}
		ctx.String(http.StatusOK, resp.String())
	})
}
