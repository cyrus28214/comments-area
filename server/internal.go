package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, Success("pong"))
}
