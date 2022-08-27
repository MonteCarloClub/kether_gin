package client

import (
	"log"

	"github.com/MonteCarloClub/KBD/kitex_gen/api/kanbandatabase"
	"github.com/MonteCarloClub/kether_gin/constant"
	"github.com/cloudwego/kitex/client"
)

var KBDClient kanbandatabase.Client

func init() {
	var err error
	KBDClient, err = kanbandatabase.NewClient("kdb", client.WithHostPorts(constant.KDBServer))
	if err != nil {
		log.Fatal(err)
	}
}
