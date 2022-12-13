package requests

import (
	"net/http"
	"time"

	"github.com/iotaledger/wasp/packages/webapi/v2/models"
	"github.com/iotaledger/wasp/packages/webapi/v2/params"

	"github.com/labstack/echo/v4"
)

func (c *Controller) waitForRequestToFinish(e echo.Context) error {
	chainID, err := params.DecodeChainID(e)
	if err != nil {
		return err
	}

	requestID, err := params.DecodeRequestID(e)
	if err != nil {
		return err
	}

	receipt, vmError, err := c.chainService.WaitForRequestProcessed(chainID, *requestID, 30*time.Second)
	if err != nil {
		return err
	}

	mappedReceipt := models.MapReceiptResponse(receipt, vmError)

	return e.JSON(http.StatusOK, mappedReceipt)
}
