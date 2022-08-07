package chat

import (
	"github.com/ferralucho/chatter-stock/src/domain/bots"
	errors "github.com/ferralucho/chatter-stock/src/rest_errors"
	"github.com/ferralucho/chatter-stock/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendCommand(c *gin.Context) {
	var command bots.Command
	if err := c.ShouldBindJSON(&command); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := services.BotsService.CreateCommand(command)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
