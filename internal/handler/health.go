package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) checkHealth(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
