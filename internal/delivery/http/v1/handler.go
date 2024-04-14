package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sema0205/avito-backend-assignment-2024/internal/service"
	"github.com/sema0205/avito-backend-assignment-2024/pkg/auth"
	"strconv"
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

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initUserRoutes(v1)
		h.initAdminRoutes(v1)
	}
}

func parseIdFromPath(c *gin.Context) (int, error) {
	idParam := c.Param("id")
	if idParam == "" {
		return 0, errors.New("empty id param")
	}

	bannerId, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, errors.New("invalid id param")
	}

	return bannerId, nil
}
