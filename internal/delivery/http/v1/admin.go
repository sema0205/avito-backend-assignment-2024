package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sema0205/avito-backend-assignment-2024/internal/domain"
	"github.com/sema0205/avito-backend-assignment-2024/internal/service"
	"net/http"
)

func (h *Handler) initAdminRoutes(api *gin.RouterGroup) {
	admins := api.Group("/admins")
	{
		admins.POST("/sign-in", h.adminSignIn)

		authenticated := admins.Group("/", h.adminIdentity)
		{
			banners := authenticated.Group("/banners")
			{
				banners.GET("", h.adminGetFilteredBanners)
				banners.POST("", h.adminCreateBanner)
				banners.PATCH("/:id", h.adminUpdateBanner)
				banners.DELETE("/:id", h.adminDeleteBanner)
			}
		}
	}
}

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary Admin SignIn
// @Tags admins-auth
// @Description admin sign in
// @ModuleID adminSignIn
// @Accept  json
// @Produce  json
// @Param input body signInInput true "sign up info"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admins/sign-in [post]
func (h *Handler) adminSignIn(c *gin.Context) {
	var inp signInInput
	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response{"invalid input body"})

		return
	}

	res, err := h.services.Admin.SignIn(c.Request.Context(), service.SignInInput{
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

type createBannerInput struct {
	TagIds    []int  `json:"tag_ids" binding:"required"`
	FeatureID int    `json:"feature_id" binding:"required"`
	Content   string `json:"content" binding:"required"`
	IsActive  bool   `json:"is_active"`
}

// @Summary Admin Create New Banner
// @Security AdminAuth
// @Tags admins-banners
// @Description admin create new banner
// @ModuleID adminCreateBanner
// @Accept  json
// @Produce  json
// @Param input body createBannerInput true "banner info"
// @Success 201 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admins/banners [post]
func (h *Handler) adminCreateBanner(c *gin.Context) {
	var inp createBannerInput
	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response{"invalid input body"})

		return
	}

	bannerId, err := h.services.Banner.Create(c.Request.Context(), domain.Banner{
		Content:   inp.Content,
		IsActive:  inp.IsActive,
		TagsIds:   inp.TagIds,
		FeatureId: inp.FeatureID,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response{err.Error()})

		return
	}

	c.JSON(http.StatusCreated, bannerIdResponse{bannerId})
}

type updateBannerInput struct {
	TagIds    []*int  `json:"tag_ids"`
	FeatureID *int    `json:"feature_id"`
	Content   *string `json:"content"`
	IsActive  *bool   `json:"is_active"`
}

// @Summary Admin Update Banner
// @Security AdminAuth
// @Tags admins-banners
// @Description admin update banner
// @ModuleID adminUpdateBanner
// @Accept  json
// @Produce  json
// @Param id path string true "banner id"
// @Param input body updateBannerInput true "banner update info"
// @Success 200 {string} string "ok"
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admins/banners/{id} [put]
func (h *Handler) adminUpdateBanner(c *gin.Context) {
	bannerId, err := parseIdFromPath(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response{err.Error()})

		return
	}

	var inp updateBannerInput
	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response{"invalid input body"})

		return
	}

	err = h.services.Banner.Update(c.Request.Context(), service.UpdateBannerInput{
		Id:        bannerId,
		Content:   inp.Content,
		TagIds:    inp.TagIds,
		FeatureId: inp.FeatureID,
		IsActive:  inp.IsActive,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response{err.Error()})

		return
	}

	c.Status(http.StatusOK)
}

// @Summary Admin Delete Banner
// @Security AdminAuth
// @Tags admins-banners
// @Description admin delete banner
// @ModuleID adminDeleteBanner
// @Accept  json
// @Produce  json
// @Param id path string true "banner id"
// @Success 200 {string} string "ok"
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admins/banners/{id} [delete]
func (h *Handler) adminDeleteBanner(c *gin.Context) {
	bannerId, err := parseIdFromPath(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response{err.Error()})

		return
	}

	err = h.services.Banner.Delete(c.Request.Context(), bannerId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response{err.Error()})

		return
	}

	c.Status(http.StatusOK)
}

type GetBannersQuery struct {
	FeatureId *int `form:"feature_id"`
	TagId     *int `form:"tag_id"`
	Limit     int  `form:"limit"`
	Offset    int  `form:"offset"`
}

// @Summary Admin Get Filtered Banners
// @Security AdminAuth
// @Tags admins-orders
// @Description admin get filtered banners
// @ModuleID adminGetFilteredBanners
// @Accept  json
// @Produce  json
// @Param feature_id query int false "feature_id"
// @Param tag_id query int false "tag_id"
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admins/banners [get]
func (h *Handler) adminGetFilteredBanners(c *gin.Context) {
	var query GetBannersQuery
	if err := c.BindQuery(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response{err.Error()})

		return
	}

	filteredBanners, err := h.services.Banner.GetFilteredBanners(c.Request.Context(), service.GetFilteredBannerInput{
		FeatureId: query.FeatureId,
		TagID:     query.TagId,
		Limit:     query.Limit,
		Offset:    query.Offset,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response{err.Error()})

		return
	}

	c.JSON(http.StatusOK, filteredBanners)
}
