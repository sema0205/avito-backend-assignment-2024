package http

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/sema0205/avito-backend-assignment-2024/internal/delivery/http/v1"
	"github.com/sema0205/avito-backend-assignment-2024/internal/service"
	"github.com/sema0205/avito-backend-assignment-2024/pkg/auth"
	"net/http"
)

type Handler struct {
	services     *service.Services
	tokenManager auth.Provider
}

func NewHandler(services *service.Services, tokenManager auth.Provider) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init() *gin.Engine {

	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.tokenManager)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
