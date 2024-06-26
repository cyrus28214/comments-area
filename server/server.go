package server

import (
	"github.com/cyrus28214/comments-area/logging"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Server struct {
	Config ServerConfig
	Log    *logrus.Logger
	DB     *gorm.DB
	Router *gin.Engine
}

func NewServer(Config ServerConfig, Log *logrus.Logger, DB *gorm.DB, Router *gin.Engine) *Server {
	s := &Server{Config: Config, Log: Log, DB: DB, Router: Router}
	s.Router.Use(logging.LogMiddleware(s.Log))
	s.Router.GET("/ping", s.Ping)
	s.Router.GET("/comment/get", s.GetComments)
	s.Router.POST("/comment/add", s.CreateComment)
	s.Router.POST("/comment/delete", s.DeleteComment)
	return s
}

func (s *Server) Start() {
	url := s.Config.Host + ":" + s.Config.Port
	s.Log.Info("Server starting on ", url)
	err := s.Router.Run(url)
	if err != nil {
		s.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to start server")
	}
}
