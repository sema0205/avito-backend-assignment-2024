package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/sema0205/avito-backend-assignment-2024/internal/domain"
	"github.com/sema0205/avito-backend-assignment-2024/internal/mocks/servicemocks"
	"github.com/sema0205/avito-backend-assignment-2024/internal/service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_userGetBanner(t *testing.T) {
	type args struct {
		FeatureId       int
		TagId           int
		UseLastRevision bool
	}
	type mockBehavior func(r *servicemocks.MockBanner, input service.GetUserBannerInput)

	tests := []struct {
		name             string
		query            args
		prepare          mockBehavior
		expectedCode     int
		expectedResponse string
	}{
		{
			name: "OK",
			query: args{
				FeatureId:       10,
				TagId:           20,
				UseLastRevision: true,
			},
			prepare: func(r *servicemocks.MockBanner, input service.GetUserBannerInput) {
				r.EXPECT().GetByTagIdAndFeatureId(gomock.Any(), input).Return(domain.Banner{Content: "Valid Content"}, nil)
			},
			expectedCode:     http.StatusOK,
			expectedResponse: `"Valid Content"`,
		},
		{
			name: "Invalid Input - Missing TagId",
			query: args{
				FeatureId:       10,
				UseLastRevision: false,
			},
			prepare: func(r *servicemocks.MockBanner, input service.GetUserBannerInput) {

			},
			expectedCode:     http.StatusBadRequest,
			expectedResponse: `{"message":"Key: 'GetUserBannerQuery.TagId' Error:Field validation for 'TagId' failed on the 'required' tag"}`,
		},
		{
			name: "Service Error",
			query: args{
				FeatureId:       10,
				TagId:           20,
				UseLastRevision: true,
			},
			prepare: func(r *servicemocks.MockBanner, input service.GetUserBannerInput) {
				r.EXPECT().GetByTagIdAndFeatureId(gomock.Any(), input).Return(domain.Banner{}, errors.New("internal server error"))
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
			tt.prepare(mockBannerService, service.GetUserBannerInput{
				FeatureId:       tt.query.FeatureId,
				TagId:           tt.query.TagId,
				UseLastRevision: tt.query.UseLastRevision,
			})

			handler := NewHandler(&service.Services{Banner: mockBannerService}, nil)
			router := gin.Default()
			router.GET("/users/banner", handler.userGetBanner)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/banner?feature_id=%d&tag_id=%d&use_last_revision=%t",
				tt.query.FeatureId, tt.query.TagId, tt.query.UseLastRevision), nil)
			req.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			assert.JSONEq(t, tt.expectedResponse, w.Body.String())
		})
	}
}
