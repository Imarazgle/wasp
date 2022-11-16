package chain

import (
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/labstack/echo/v4"
)

func (c *Controller) handleWebSocket(e echo.Context) error {
	chainID, err := isc.ChainIDFromString(e.Param("chainID"))
	if err != nil {
		return err
	}

	return c.webSocketHandler.ServeHTTP(chainID, e.Response(), e.Request())
}
