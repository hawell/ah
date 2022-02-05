package server

import (
	"go.uber.org/zap"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
)

// Server represents API server
type Server struct {
	config     Config
	httpServer *http.Server
	router     *gin.Engine
}

// NewServer creates a new API server
func NewServer(accessLogger *zap.Logger) (*Server, error) {
	var config Config
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		return nil, err
	}

	router := newRouter(accessLogger)

	server := &http.Server{
		Addr:           config.ListenAddress,
		Handler:        router,
		ReadTimeout:    time.Duration(config.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{
		config:     config,
		httpServer: server,
		router:     router,
	}, nil
}

// ListenAndServe listens and serves a server
func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}
