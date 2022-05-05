// Code generated by MockGen. DO NOT EDIT.
// Source: article_service.go

// Package mock_service is a generated GoMock package.
package service

import (
	reflect "reflect"

	ent "github.com/efectn/fiber-boilerplate/internal/ent"
	request "github.com/efectn/fiber-boilerplate/module/article/request"
	gomock "github.com/golang/mock/gomock"
)

// MockIArticleService is a mock of IArticleService interface.
type MockIArticleService struct {
	ctrl     *gomock.Controller
	recorder *MockIArticleServiceMockRecorder
}

// MockIArticleServiceMockRecorder is the mock recorder for MockIArticleService.
type MockIArticleServiceMockRecorder struct {
	mock *MockIArticleService
}

// NewMockIArticleService creates a new mock instance.
func NewMockIArticleService(ctrl *gomock.Controller) *MockIArticleService {
	mock := &MockIArticleService{ctrl: ctrl}
	mock.recorder = &MockIArticleServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIArticleService) EXPECT() *MockIArticleServiceMockRecorder {
	return m.recorder
}

// CreateArticle mocks base method.
func (m *MockIArticleService) CreateArticle(request request.ArticleRequest) (*ent.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArticle", request)
	ret0, _ := ret[0].(*ent.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateArticle indicates an expected call of CreateArticle.
func (mr *MockIArticleServiceMockRecorder) CreateArticle(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArticle", reflect.TypeOf((*MockIArticleService)(nil).CreateArticle), request)
}

// DeleteArticle mocks base method.
func (m *MockIArticleService) DeleteArticle(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteArticle", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteArticle indicates an expected call of DeleteArticle.
func (mr *MockIArticleServiceMockRecorder) DeleteArticle(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteArticle", reflect.TypeOf((*MockIArticleService)(nil).DeleteArticle), id)
}

// GetArticleByID mocks base method.
func (m *MockIArticleService) GetArticleByID(id int) (*ent.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleByID", id)
	ret0, _ := ret[0].(*ent.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticleByID indicates an expected call of GetArticleByID.
func (mr *MockIArticleServiceMockRecorder) GetArticleByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleByID", reflect.TypeOf((*MockIArticleService)(nil).GetArticleByID), id)
}

// GetArticles mocks base method.
func (m *MockIArticleService) GetArticles() ([]*ent.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticles")
	ret0, _ := ret[0].([]*ent.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticles indicates an expected call of GetArticles.
func (mr *MockIArticleServiceMockRecorder) GetArticles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticles", reflect.TypeOf((*MockIArticleService)(nil).GetArticles))
}

// UpdateArticle mocks base method.
func (m *MockIArticleService) UpdateArticle(id int, request request.ArticleRequest) (*ent.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateArticle", id, request)
	ret0, _ := ret[0].(*ent.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateArticle indicates an expected call of UpdateArticle.
func (mr *MockIArticleServiceMockRecorder) UpdateArticle(id, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateArticle", reflect.TypeOf((*MockIArticleService)(nil).UpdateArticle), id, request)
}
