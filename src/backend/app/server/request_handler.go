package server

import (
	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	Gin *gin.Engine
}

func NewRequestHandler() *RequestHandler {
	handler := RequestHandler{Gin: gin.Default()}

	return &handler
}
