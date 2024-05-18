package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (s *Server) HandleBindingError(c *gin.Context, err error) {
	s.Log.WithFields(logrus.Fields{
		"error": err.Error(),
	}).Warn("Invalid request: binding failed")
	c.JSON(http.StatusOK, InvalidRequest("Invalid request")) //传递到前端的错误信息没有log记录的详细，是为了避免敏感信息通过err被泄露
}

func (s *Server) HandleDatabaseError(c *gin.Context, err error) {
	s.Log.WithFields(logrus.Fields{
		"error": err.Error(),
	}).Error("Database error occurred")
	c.JSON(http.StatusOK, InternalServerError("Internal server error"))
}
