package chain

import (
	"net/http"

	"github.com/iotaledger/wasp/packages/webapi/v2/params"
	"github.com/labstack/echo/v4"
)

func (c *Controller) activateChain(e echo.Context) error {
	chainID, err := params.DecodeChainID(e)
	if err != nil {
		return err
	}

	if err := c.chainService.ActivateChain(chainID); err != nil {
		return err
	}

	return e.NoContent(http.StatusOK)
}

func (c *Controller) deactivateChain(e echo.Context) error {
	chainID, err := params.DecodeChainID(e)
	if err != nil {
		return err
	}

	if err := c.chainService.DeactivateChain(chainID); err != nil {
		return err
	}

	return e.NoContent(http.StatusOK)
}
