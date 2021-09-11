package handlers

import (
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests
type Server struct {
}

// Start runs the HTTP server on a specific address.
func (*Server) Start(router *gin.Engine, address string) error {
	return router.Run(address)
}
