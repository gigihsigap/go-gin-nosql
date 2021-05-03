package api

import (
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests.
type Server struct {
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer() *Server {
	server := &Server{}

	router := gin.New()

	router.Use(gin.Logger())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/accounts", server.createAccount)
	// router.GET("/accounts/:id", server.getAccount)
	// router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

// Start runs the HTTP server on the input address to start listening on API requests
// it takes address as input and error as output
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// Gives an error response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
