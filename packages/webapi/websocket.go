package webapi

import (
	_ "embed"

	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/publisher/publisherws"
	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"
)

type webSocketAPI struct {
	pws *publisherws.PublisherWebSocket
}

func addWebSocketEndpoint(e echoswagger.ApiGroup) *webSocketAPI {
	api := &webSocketAPI{
		pws: publisherws.New(),
	}

	e.GET("/chain/:chainid/ws", api.handleWebSocket)

	return api
}

func (w *webSocketAPI) handleWebSocket(c echo.Context) error {
	chainID, err := iscp.ChainIDFromBase58(c.Param("chainid"))
	if err != nil {
		return err
	}
	w.pws.GetHandler(chainID).ServeHTTP(c.Response(), c.Request())
	return nil
}

func (w *webSocketAPI) startWebSocketForwarder() {
	w.pws.Start("state", "vmmsg")
}
