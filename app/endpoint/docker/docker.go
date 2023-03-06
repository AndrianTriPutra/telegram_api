package docker

import (
	"api-telegram/app/middleware"
	"api-telegram/pkg/telegram"
	"api-telegram/pkg/utils/logger"
	"api-telegram/pkg/utils/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

type telegramHandler struct {
	ucase telegram.RepositoryI
}

func NewHandler(e *echo.Echo, endpoint string, ucase telegram.RepositoryI) {
	handler := telegramHandler{
		ucase: ucase,
	}
	e.POST(endpoint, middleware.ErrorMiddleware(handler.Docker))
}

func (h telegramHandler) Docker(c echo.Context) error {
	ctx := c.Request().Context()

	message := c.FormValue("message")
	if len(message) < 1 {
		logger.Level("error", "[Docker] ", "The message field is required")
		return util.CustomError{
			ErrorType: util.ErrBadRequest,
			Message:   "The given data was invalid.",
			Cause:     "The message field is required",
		}
	}

	err := h.ucase.Send(ctx, message)
	if err != nil {
		logger.Level("error", "[Docker] ", "[Telegram-Send] msg:"+message+"->"+err.Error())
		return util.CustomError{
			ErrorType: util.ErrInternalServer,
			Message:   "Failed send message",
			Cause:     err.Error(),
		}
	}

	response := util.WrapSuccessResponse("success", "command succes")

	return c.JSON(http.StatusOK, response)
}
