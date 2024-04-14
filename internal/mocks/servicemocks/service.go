// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/service.go

// Package servicemocks is a generated GoMock package.
package servicemocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/sema0205/avito-backend-assignment-2024/internal/domain"
	service "github.com/sema0205/avito-backend-assignment-2024/internal/service"
)

// MockBanner is a mock of Banner interface.
type MockBanner struct {
	ctrl     *gomock.Controller
	recorder *MockBannerMockRecorder
}

// MockBannerMockRecorder is the mock recorder for MockBanner.
type MockBannerMockRecorder struct {
	mock *MockBanner
}

// NewMockBanner creates a new mock instance.
func NewMockBanner(ctrl *gomock.Controller) *MockBanner {
	mock := &MockBanner{ctrl: ctrl}
	mock.recorder = &MockBannerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBanner) EXPECT() *MockBannerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBanner) Create(ctx context.Context, banner domain.Banner) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, banner)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBannerMockRecorder) Create(ctx, banner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBanner)(nil).Create), ctx, banner)
}

// Delete mocks base method.
func (m *MockBanner) Delete(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBannerMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBanner)(nil).Delete), ctx, id)
}

// GetByTagIdAndFeatureId mocks base method.
func (m *MockBanner) GetByTagIdAndFeatureId(ctx context.Context, input service.GetUserBannerInput) (domain.Banner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTagIdAndFeatureId", ctx, input)
	ret0, _ := ret[0].(domain.Banner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTagIdAndFeatureId indicates an expected call of GetByTagIdAndFeatureId.
func (mr *MockBannerMockRecorder) GetByTagIdAndFeatureId(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTagIdAndFeatureId", reflect.TypeOf((*MockBanner)(nil).GetByTagIdAndFeatureId), ctx, input)
}

// GetFilteredBanners mocks base method.
func (m *MockBanner) GetFilteredBanners(ctx context.Context, input service.GetFilteredBannerInput) ([]domain.Banner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilteredBanners", ctx, input)
	ret0, _ := ret[0].([]domain.Banner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilteredBanners indicates an expected call of GetFilteredBanners.
func (mr *MockBannerMockRecorder) GetFilteredBanners(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilteredBanners", reflect.TypeOf((*MockBanner)(nil).GetFilteredBanners), ctx, input)
}

// Update mocks base method.
func (m *MockBanner) Update(ctx context.Context, input service.UpdateBannerInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockBannerMockRecorder) Update(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBanner)(nil).Update), ctx, input)
}

// MockAdmin is a mock of Admin interface.
type MockAdmin struct {
	ctrl     *gomock.Controller
	recorder *MockAdminMockRecorder
}

// MockAdminMockRecorder is the mock recorder for MockAdmin.
type MockAdminMockRecorder struct {
	mock *MockAdmin
}

// NewMockAdmin creates a new mock instance.
func NewMockAdmin(ctrl *gomock.Controller) *MockAdmin {
	mock := &MockAdmin{ctrl: ctrl}
	mock.recorder = &MockAdminMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdmin) EXPECT() *MockAdminMockRecorder {
	return m.recorder
}

// SignIn mocks base method.
func (m *MockAdmin) SignIn(ctx context.Context, input service.SignInInput) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", ctx, input)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockAdminMockRecorder) SignIn(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockAdmin)(nil).SignIn), ctx, input)
}

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// SignIn mocks base method.
func (m *MockUser) SignIn(ctx context.Context, input service.SignInInput) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", ctx, input)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockUserMockRecorder) SignIn(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockUser)(nil).SignIn), ctx, input)
}
