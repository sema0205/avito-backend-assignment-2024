package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sema0205/avito-backend-assignment-2024/internal/service"
	"net/http"
)

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/sign-in", h.userSignIn)

		authenticated := users.Group("/", h.userIdentity)
		{
			banners := authenticated.Group("/banner")
			{
				banners.GET("", h.userGetBanner)
			}
		}
	}
}

// @Summary User SignIn
// @Tags users-auth
// @Description user sign in
// @ModuleID userSignIn
// @Accept  json
// @Produce  json
// @Param input body signInInput true "sign up info"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/sign-in [post]
func (h *Handler) userSignIn(c *gin.Context) {
	var inp signInInput
	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response{"invalid input body"})

		return
	}

	res, err := h.services.User.SignIn(c.Request.Context(), service.SignInInput{
		Username: inp.Username,
		Password: inp.Password,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response{err.Error()})

		return
	}

	c.JSON(http.StatusOK, tokenResponse{
		AccessToken: res,
	})
}

type GetUserBannerQuery struct {
	FeatureId       int  `form:"feature_id" binding:"required"`
	TagId           int  `form:"tag_id" binding:"required"`
	UseLastRevision bool `form:"use_last_revision"`
}

// @Summary User Get Banner
// @Security UserAuth
// @Tags users-banners
// @Description user get banner
// @ModuleID userGetBanner
// @Accept  json
// @Produce  json
// @Param feature_id query int false "feature_id"
// @Param tag_id query int false "tag_id"
// @Param use_last_revision query string false "use_last_revision"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/banner [get]
func (h *Handler) userGetBanner(c *gin.Context) {
	var query GetUserBannerQuery
	if err := c.BindQuery(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response{err.Error()})

		return
	}

	banner, err := h.services.Banner.GetByTagIdAndFeatureId(c.Request.Context(), service.GetUserBannerInput{
		TagId:           query.TagId,
		FeatureId:       query.FeatureId,
		UseLastRevision: query.UseLastRevision,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response{err.Error()})

		return
	}

	c.JSON(http.StatusOK, banner.Content)
}
