package user

import (
	reflect "reflect"

	"github.com/golang/mock/gomock"
)

// MockJWT - это структура, которая будет имитировать поведение JWT
type MockJWT struct {
	ctrl     *gomock.Controller
	recorder *MockJWTMockRecorder
}

// MockJWTMockRecorder - это структура для записи вызовов методов
type MockJWTMockRecorder struct {
	mock *MockJWT
}

// NewMockJWT создает новый мок для JWT
func NewMockJWT(ctrl *gomock.Controller) *MockJWT {
	mock := &MockJWT{ctrl: ctrl}
	mock.recorder = &MockJWTMockRecorder{mock}
	return mock
}

// EXPECT возвращает объект для настройки ожиданий
func (m *MockJWT) EXPECT() *MockJWTMockRecorder {
	return m.recorder
}

// GenerateToken имитирует метод GenerateToken
func (m *MockJWT) GenerateToken(userID uint, email, login string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", userID, email, login)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken records a call to GenerateToken
func (mr *MockJWTMockRecorder) GenerateToken(userID, email, login interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockJWT)(nil).GenerateToken), userID, email, login)
}

// ParseToken имитирует метод ParseToken
func (m *MockJWT) ParseToken(token string) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken records a call to ParseToken
func (mr *MockJWTMockRecorder) ParseToken(token interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockJWT)(nil).ParseToken), token)
}
