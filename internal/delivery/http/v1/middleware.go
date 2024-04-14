package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sema0205/avito-backend-assignment-2024/pkg/auth"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"

	adminCtx = "adminId"
	userCtx  = "userId"
)

func (h *Handler) adminIdentity(c *gin.Context) {
	claims, err := h.parseAuthHeader(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response{err.Error()})

		return
	}

	if claims.Role != auth.Admin {
		c.AbortWithStatusJSON(http.StatusForbidden, response{"insufficient permissions"})
		return
	}

	c.Set("role", claims.Role)
	c.Set(adminCtx, claims.Id)
}

func (h *Handler) userIdentity(c *gin.Context) {
	claims, err := h.parseAuthHeader(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response{err.Error()})

		return
	}

	if claims.Role != auth.User && claims.Role != auth.Admin {
		c.AbortWithStatusJSON(http.StatusForbidden, response{"insufficient permissions"})
		return
	}

	c.Set("role", claims.Role)
	c.Set(userCtx, claims.Id)
}

func (h *Handler) parseAuthHeader(c *gin.Context) (auth.TokenClaims, error) {
	header := c.GetHeader(authHeader)
	if header == "" {
		return auth.TokenClaims{}, errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return auth.TokenClaims{}, errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return auth.TokenClaims{}, errors.New("token is empty")
	}

	return h.tokenManager.Parse(headerParts[1])
}
