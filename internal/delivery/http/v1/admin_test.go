package v1

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/sema0205/avito-backend-assignment-2024/internal/domain"
	"github.com/sema0205/avito-backend-assignment-2024/internal/mocks/authmocks"
	"github.com/sema0205/avito-backend-assignment-2024/internal/mocks/servicemocks"
	"github.com/sema0205/avito-backend-assignment-2024/internal/service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_adminCreateBanner(t *testing.T) {
	type mockBehavior func(r *servicemocks.MockBanner, ctx context.Context, banner domain.Banner)

	tests := []struct {
		name             string
		body             string
		banner           domain.Banner
		prepare          mockBehavior
		expectedCode     int
		expectedResponse string
	}{
		{
			name: "OK",
			body: `{"tag_ids":[1,2],"feature_id":10,"content":"New Banner Content","is_active":true}`,
			banner: domain.Banner{
				Content:   "New Banner Content",
				IsActive:  true,
				TagsIds:   []int{1, 2},
				FeatureId: 10,
			},
			prepare: func(r *servicemocks.MockBanner, ctx context.Context, banner domain.Banner) {
				r.EXPECT().Create(ctx, banner).Return(1, nil)
			},
			expectedCode:     http.StatusCreated,
			expectedResponse: `{"banner_id":1}`,
		},
		{
			name: "Invalid Input - Missing Fields",
			body: `{"content":"New Banner Content","is_active":true}`,
			prepare: func(r *servicemocks.MockBanner, ctx context.Context, banner domain.Banner) {

			},
			expectedCode:     http.StatusBadRequest,
			expectedResponse: `{"message":"invalid input body"}`,
		},
		{
			name: "Service Failure",
			body: `{"tag_ids":[1,2],"feature_id":10,"content":"New Banner Content","is_active":true}`,
			banner: domain.Banner{
				Content:   "New Banner Content",
				IsActive:  true,
				TagsIds:   []int{1, 2},
				FeatureId: 10,
			},
			prepare: func(r *servicemocks.MockBanner, ctx context.Context, banner domain.Banner) {
				r.EXPECT().Create(ctx, banner).
					Return(0, errors.New("internal server error"))
			},
			expectedCode:     http.StatusInternalServerError,
			expectedResponse: `{"message":"internal server error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockBannerService := servicemocks.NewMockBanner(ctrl)
			authManager := authmocks.NewMockProvider(ctrl)
			tt.prepare(mockBannerService, context.Background(), tt.banner)

			handler := NewHandler(&service.Services{Banner: mockBannerService}, authManager)
			r := gin.New()
			r.POST("/admins/banners", func(c *gin.Context) {}, handler.adminCreateBanner)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/admins/banners", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Equal(t, tt.expectedResponse, w.Body.String())
		})
	}
}

func TestHandler_adminGetFilteredBanners(t *testing.T) {
	type args struct {
		FeatureId *int
		TagId     *int
		Limit     int
		Offset    int
	}
	type mockBehavior func(r *servicemocks.MockBanner, input service.GetFilteredBannerInput)

	tests := []struct {
		name             string
		query            args
		prepare          mockBehavior
		expectedCode     int
		expectedResponse string
	}{
		{
			name: "OK - Full Query",
			query: args{
				FeatureId: ptrToInt(10),
				TagId:     ptrToInt(20),
				Limit:     5,
				Offset:    0,
			},
			prepare: func(r *servicemocks.MockBanner, input service.GetFilteredBannerInput) {
				r.EXPECT().GetFilteredBanners(gomock.Any(), input).Return([]domain.Banner{{Content: "Banner Content", IsActive: true}}, nil)
			},
			expectedCode:     http.StatusOK,
			expectedResponse: "[{\"banner_id\":0,\"content\":\"Banner Content\",\"is_active\":true,\"updated_at\":\"0001-01-01T00:00:00Z\",\"created_at\":\"0001-01-01T00:00:00Z\",\"tag_ids\":null,\"feature_id\":0}]",
		},
		{
			name: "Service Error",
			query: args{
				FeatureId: ptrToInt(10),
				TagId:     ptrToInt(20),
				Limit:     5,
				Offset:    0,
			},
			prepare: func(r *servicemocks.MockBanner, input service.GetFilteredBannerInput) {
				r.EXPECT().GetFilteredBanners(gomock.Any(), input).Return(nil, errors.New("database error"))
			},
			expectedCode:     http.StatusInternalServerError,
			expectedResponse: "{\"message\":\"database error\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockBannerService := servicemocks.NewMockBanner(ctrl)
			authManager := authmocks.NewMockProvider(ctrl)
			tt.prepare(mockBannerService, service.GetFilteredBannerInput{
				FeatureId: tt.query.FeatureId,
				TagID:     tt.query.TagId,
				Limit:     tt.query.Limit,
				Offset:    tt.query.Offset,
			})

			handler := NewHandler(&service.Services{Banner: mockBannerService}, authManager)
			router := gin.Default()
			router.GET("/admins/banners", handler.adminGetFilteredBanners)

			w := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/admins/banners?feature_id=%d&tag_id=%d&limit=%d&offset=%d",
				*tt.query.FeatureId, *tt.query.TagId, tt.query.Limit, tt.query.Offset), nil)
			req.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.Equal(t, tt.expectedResponse, w.Body.String())
		})
	}
}

func ptrToInt(i int) *int {
	return &i
}
