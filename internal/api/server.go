package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db "task-manage/internal/db/sqlc"
	"task-manage/internal/token"
	"task-manage/internal/utils"
)

type Server struct {
	config     utils.Config
	queries    *db.Queries
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config utils.Config, queries *db.Queries) (*Server, error) {
	tokenMaker, err := token.NewJwtMaker(config.TokenSymmetricKey)
	server := &Server{queries: queries, config: config, tokenMaker: tokenMaker}
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	setupRouter(server)
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
