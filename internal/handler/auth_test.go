package handler

import (
	fin "FinTransaction"
	"FinTransaction/internal/service"
	mock_service "FinTransaction/internal/service/mock"
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockAuthorization, user fin.User)

	testTable := []struct {
		name          string
		inputBody     string
		inputUser     fin.User
		mockBehaviour mockBehaviour
		expectedCode  int
		expectedBody  string
	}{
		{
			name:      "success",
			inputBody: `{"name":"test","username":"test","password":"test"}`,
			inputUser: fin.User{
				Name:     "test",
				Username: "test",
				Password: "test",
			},
			mockBehaviour: func(s *mock_service.MockAuthorization, user fin.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedCode: 200,
			expectedBody: `{"id":1}`,
		},
		{
			name:          "Empty Fields",
			inputBody:     `{"username":"test","password":"test"}`,
			mockBehaviour: func(s *mock_service.MockAuthorization, user fin.User) {},
			expectedCode:  400,
			expectedBody:  `{"massage":"invalid input body"}`,
		},
		{
			name:      "Service Failure",
			inputBody: `{"name":"test","username":"test","password":"test"}`,
			inputUser: fin.User{
				Name:     "test",
				Username: "test",
				Password: "test",
			},
			mockBehaviour: func(s *mock_service.MockAuthorization, user fin.User) {
				s.EXPECT().CreateUser(user).Return(1, errors.New("service failure"))
			},
			expectedCode: 500,
			expectedBody: `{"massage":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehaviour(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/sign-up", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedCode, w.Code)
			assert.Equal(t, testCase.expectedBody, w.Body.String())
		})
	}
}

func TestHandler_signIn(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockAuthorization, user signInInput)

	testTable := []struct {
		name          string
		inputBody     string
		inputUser     signInInput
		mockBehaviour mockBehaviour
		expectedCode  int
		expectedBody  string
	}{
		{
			name:      "success",
			inputBody: `{"username":"test","password":"test"}`,
			inputUser: signInInput{
				Username: "test",
				Password: "test",
			},
			mockBehaviour: func(s *mock_service.MockAuthorization, user signInInput) {
				s.EXPECT().GenerateToken(user.Username, user.Password).Return("token", nil)
			},
			expectedCode: 200,
			expectedBody: `{"token":"token"}`,
		},
		{
			name:          "Empty Fields",
			inputBody:     `"password":"test"}`,
			mockBehaviour: func(s *mock_service.MockAuthorization, user signInInput) {},
			expectedCode:  400,
			expectedBody:  `{"massage":"invalid input body"}`,
		},
		{
			name:      "Service Failure",
			inputBody: `{"username":"test","password":"test"}`,
			inputUser: signInInput{
				Username: "test",
				Password: "test",
			},
			mockBehaviour: func(s *mock_service.MockAuthorization, user signInInput) {
				s.EXPECT().GenerateToken(user.Username, user.Password).Return("token", errors.New("service failure"))
			},
			expectedCode: 500,
			expectedBody: `{"massage":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehaviour(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/sign-in", handler.singIn)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/sign-in", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedCode, w.Code)
			assert.Equal(t, testCase.expectedBody, w.Body.String())
		})
	}
}
