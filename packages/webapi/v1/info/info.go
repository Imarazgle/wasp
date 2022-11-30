package info

import (
	"github.com/iotaledger/wasp/packages/webapi/v1/model"
	"github.com/iotaledger/wasp/packages/webapi/v1/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"

	"github.com/iotaledger/wasp/packages/peering"
	"github.com/iotaledger/wasp/packages/wasp"
)

type infoService struct {
	network       peering.NetworkProvider
	publisherPort int
}

func AddEndpoints(server echoswagger.ApiRouter, network peering.NetworkProvider, publisherPort int) {
	s := &infoService{
		network:       network,
		publisherPort: publisherPort,
	}

	server.GET(routes.Info(), s.handleInfo).
		SetSummary("Get information about the node").
		AddResponse(http.StatusOK, "Node properties", model.InfoResponse{}, nil)
}

func (s *infoService) handleInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, model.InfoResponse{
		Version:       wasp.Version,
		VersionHash:   wasp.VersionHash,
		NetworkID:     s.network.Self().NetID(),
		PublisherPort: s.publisherPort,
	})
}