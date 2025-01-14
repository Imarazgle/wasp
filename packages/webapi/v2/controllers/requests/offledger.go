package requests

import (
	"net/http"

	"github.com/labstack/echo/v4"

	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/webapi/v2/apierrors"
	"github.com/iotaledger/wasp/packages/webapi/v2/models"
)

func (c *Controller) handleOffLedgerRequest(e echo.Context) error {
	request := new(models.OffLedgerRequest)
	if err := e.Bind(request); err != nil {
		return apierrors.InvalidOffLedgerRequestError(err)
	}

	chainID, err := isc.ChainIDFromString(request.ChainID)
	if err != nil {
		return apierrors.InvalidPropertyError("ChainID", err)
	}

	requestDecoded, err := iotago.DecodeHex(request.Request)
	if err != nil {
		return apierrors.InvalidPropertyError("Request", err)
	}

	err = c.offLedgerService.EnqueueOffLedgerRequest(chainID, requestDecoded)
	if err != nil {
		return apierrors.ContractExecutionError(err)
	}

	return e.NoContent(http.StatusAccepted)
}
