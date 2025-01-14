package apierrors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GenericError struct {
	Error string
}

// HTTPErrorHandler must be hooked to an echo server to render instances
// of HTTPError as JSON
func HTTPErrorHandler(err error, c echo.Context) {
	if echoError, ok := err.(*echo.HTTPError); ok {
		mappedError := HTTPErrorFromEchoError(echoError)
		err = c.JSON(mappedError.HTTPCode, mappedError.GetErrorResult())
	}

	if apiError, ok := err.(*HTTPError); ok {
		if !c.Response().Committed {
			if c.Request().Method == http.MethodHead { // Issue #608
				err = c.NoContent(apiError.HTTPCode)
			} else {
				err = c.JSON(apiError.HTTPCode, apiError.GetErrorResult())
			}
		}
	}

	c.Echo().DefaultHTTPErrorHandler(err, c)
}
