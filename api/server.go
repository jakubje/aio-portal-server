package api

import (
	db "server/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Server server HTTP requests for our service
type Server struct {
	store db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/update", server.updateUser)
	router.GET("/users/:id", server.getUser)
	router.GET("/users", server.listUsers)
	router.DELETE("/users/:id", server.deleteUser)


	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}