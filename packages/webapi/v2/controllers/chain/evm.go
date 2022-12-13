package chain

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/iotaledger/wasp/packages/webapi/v2/apierrors"
	"github.com/iotaledger/wasp/packages/webapi/v2/params"
)

func (c *Controller) handleJSONRPC(e echo.Context) error {
	chainID, err := params.DecodeChainID(e)
	if err != nil {
		return err
	}

	return c.evmService.HandleJSONRPC(chainID, e.Request(), e.Response())
}

func (c *Controller) getRequestID(e echo.Context) error {
	chainID, err := params.DecodeChainID(e)
	if err != nil {
		return err
	}

	txHash := e.Param("txHash")
	requestID, err := c.evmService.GetRequestID(chainID, txHash)
	if err != nil {
		return apierrors.InvalidPropertyError("txHash", err)
	}

	return e.JSON(http.StatusOK, requestID)
}
