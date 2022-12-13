package corecontracts

import (
	"net/http"

	"github.com/iotaledger/wasp/packages/webapi/v2/params"

	"github.com/iotaledger/wasp/packages/webapi/v2/apierrors"
	"github.com/labstack/echo/v4"
)

type ErrorMessageFormatResponse struct {
	MessageFormat string
}

func (c *Controller) getErrorMessageFormat(e echo.Context) error {
	chainID, err := params.DecodeChainID(e)
	if err != nil {
		return err
	}

	contractHname, err := params.DecodeHNameFromHNameString(e, "contractHname")
	if err != nil {
		return err
	}

	errorID, err := params.DecodeUInt(e, "errorID")
	if err != nil {
		return err
	}

	messageFormat, err := c.errors.GetMessageFormat(chainID, contractHname, uint16(errorID))

	if err != nil {
		return apierrors.ContractExecutionError(err)
	}

	errorMessageFormatResponse := &ErrorMessageFormatResponse{
		MessageFormat: messageFormat,
	}

	return e.JSON(http.StatusOK, errorMessageFormatResponse)
}
