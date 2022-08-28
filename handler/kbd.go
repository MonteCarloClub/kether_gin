package handler

import (
	"context"
	"net/http"

	kdb "github.com/MonteCarloClub/KBD/kitex_gen/api"
	"github.com/MonteCarloClub/kether_gin/client"
	"github.com/gin-gonic/gin"
)

func RegisterGetData(r *gin.Engine, api string) {
	registerMethod(r, api, func(ctx *gin.Context) {
		req := &kdb.GetDataRequest{
			Key: ctx.Query("key"),
		}
		resp, err := client.KBDClient.GetData(context.Background(), req)
		if err != nil {
			ctx.String(http.StatusBadRequest, "%v", err.Error())
		}
		ctx.String(http.StatusOK, resp.String())
	})
}

func RegisterPutData(r *gin.Engine, api string) {
	registerMethod(r, api, func(ctx *gin.Context) {
		req := &kdb.PutDataRequest{
			Key:   ctx.Query("key"),
			Value: ctx.Query("value"),
		}
		resp, err := client.KBDClient.PutData(context.Background(), req)
		if err != nil {
			ctx.String(http.StatusBadRequest, "%v", err.Error())
		}
		ctx.String(http.StatusOK, resp.String())
	})
}

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
