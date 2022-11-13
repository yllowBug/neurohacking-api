// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/kirill0909/neurohacking-api/models"
)

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

// CheckUserIdExists mocks base method.
func (m *MockUser) CheckUserIdExists(id int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserIdExists", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserIdExists indicates an expected call of CheckUserIdExists.
func (mr *MockUserMockRecorder) CheckUserIdExists(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserIdExists", reflect.TypeOf((*MockUser)(nil).CheckUserIdExists), id)
}

// Create mocks base method.
func (m *MockUser) Create(user models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUser)(nil).Create), user)
}

// Delete mocks base method.
func (m *MockUser) Delete(userId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserMockRecorder) Delete(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUser)(nil).Delete), userId)
}

// GenerateToken mocks base method.
func (m *MockUser) GenerateToken(email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockUserMockRecorder) GenerateToken(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockUser)(nil).GenerateToken), email, password)
}

// ParseToken mocks base method.
func (m *MockUser) ParseToken(token string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockUserMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockUser)(nil).ParseToken), token)
}

// Update mocks base method.
func (m *MockUser) Update(input models.UserUpdateInput, userId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", input, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserMockRecorder) Update(input, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUser)(nil).Update), input, userId)
}

// MockCategory is a mock of Category interface.
type MockCategory struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryMockRecorder
}

// MockCategoryMockRecorder is the mock recorder for MockCategory.
type MockCategoryMockRecorder struct {
	mock *MockCategory
}

// NewMockCategory creates a new mock instance.
func NewMockCategory(ctrl *gomock.Controller) *MockCategory {
	mock := &MockCategory{ctrl: ctrl}
	mock.recorder = &MockCategoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategory) EXPECT() *MockCategoryMockRecorder {
	return m.recorder
}

// CheckCategoryIdExists mocks base method.
func (m *MockCategory) CheckCategoryIdExists(userId, categoryId int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCategoryIdExists", userId, categoryId)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckCategoryIdExists indicates an expected call of CheckCategoryIdExists.
func (mr *MockCategoryMockRecorder) CheckCategoryIdExists(userId, categoryId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCategoryIdExists", reflect.TypeOf((*MockCategory)(nil).CheckCategoryIdExists), userId, categoryId)
}

// Create mocks base method.
func (m *MockCategory) Create(category models.Category, userId int) (models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", category, userId)
	ret0, _ := ret[0].(models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCategoryMockRecorder) Create(category, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCategory)(nil).Create), category, userId)
}

// Delete mocks base method.
func (m *MockCategory) Delete(userId, categoryId int) (models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userId, categoryId)
	ret0, _ := ret[0].(models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockCategoryMockRecorder) Delete(userId, categoryId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCategory)(nil).Delete), userId, categoryId)
}

// GetAll mocks base method.
func (m *MockCategory) GetAll(userId int) ([]models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", userId)
	ret0, _ := ret[0].([]models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockCategoryMockRecorder) GetAll(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockCategory)(nil).GetAll), userId)
}

// GetById mocks base method.
func (m *MockCategory) GetById(userId, categoryId int) (models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", userId, categoryId)
	ret0, _ := ret[0].(models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockCategoryMockRecorder) GetById(userId, categoryId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockCategory)(nil).GetById), userId, categoryId)
}

// Update mocks base method.
func (m *MockCategory) Update(input models.CategoryUpdateInput, userId, categoryId int) (models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", input, userId, categoryId)
	ret0, _ := ret[0].(models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCategoryMockRecorder) Update(input, userId, categoryId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCategory)(nil).Update), input, userId, categoryId)
}

// MockWord is a mock of Word interface.
type MockWord struct {
	ctrl     *gomock.Controller
	recorder *MockWordMockRecorder
}

// MockWordMockRecorder is the mock recorder for MockWord.
type MockWordMockRecorder struct {
	mock *MockWord
}

// NewMockWord creates a new mock instance.
func NewMockWord(ctrl *gomock.Controller) *MockWord {
	mock := &MockWord{ctrl: ctrl}
	mock.recorder = &MockWordMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWord) EXPECT() *MockWordMockRecorder {
	return m.recorder
}

// CheckCategoryOwner mocks base method.
func (m *MockWord) CheckCategoryOwner(userId, categoryId int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCategoryOwner", userId, categoryId)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckCategoryOwner indicates an expected call of CheckCategoryOwner.
func (mr *MockWordMockRecorder) CheckCategoryOwner(userId, categoryId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCategoryOwner", reflect.TypeOf((*MockWord)(nil).CheckCategoryOwner), userId, categoryId)
}

// Create mocks base method.
func (m *MockWord) Create(word models.Word, userId, categoryId int) (models.Word, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", word, userId, categoryId)
	ret0, _ := ret[0].(models.Word)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockWordMockRecorder) Create(word, userId, categoryId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWord)(nil).Create), word, userId, categoryId)
}

// Delete mocks base method.
func (m *MockWord) Delete(userId, categoryId, wordId int) (models.Word, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userId, categoryId, wordId)
	ret0, _ := ret[0].(models.Word)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockWordMockRecorder) Delete(userId, categoryId, wordId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWord)(nil).Delete), userId, categoryId, wordId)
}

// GetAll mocks base method.
func (m *MockWord) GetAll(userId, categoryId int) ([]models.Word, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", userId, categoryId)
	ret0, _ := ret[0].([]models.Word)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockWordMockRecorder) GetAll(userId, categoryId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockWord)(nil).GetAll), userId, categoryId)
}

// GetById mocks base method.
func (m *MockWord) GetById(userId, categoryId, wordId int) (models.Word, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", userId, categoryId, wordId)
	ret0, _ := ret[0].(models.Word)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockWordMockRecorder) GetById(userId, categoryId, wordId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockWord)(nil).GetById), userId, categoryId, wordId)
}

// Update mocks base method.
func (m *MockWord) Update(input models.WordUpdateInput, userId, categoryId, wordId int) (models.Word, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", input, userId, categoryId, wordId)
	ret0, _ := ret[0].(models.Word)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockWordMockRecorder) Update(input, userId, categoryId, wordId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWord)(nil).Update), input, userId, categoryId, wordId)
}
