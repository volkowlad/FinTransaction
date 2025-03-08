package handler

import (
	"FinTransaction/internal/service"
	mock_service "FinTransaction/internal/service/mock"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_userIdentity(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockAuthorization, token string)

	testTable := []struct {
		name          string
		headerName    string
		headerValue   string
		token         string
		mockBehaviour mockBehaviour
		expectedCode  int
		expectedBody  string
	}{
		{
			name:        "OK",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehaviour: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedCode: 200,
			expectedBody: "1",
		},
		{
			name:          "No Header",
			headerName:    "",
			mockBehaviour: func(s *mock_service.MockAuthorization, token string) {},
			expectedCode:  401,
			expectedBody:  `{"massage":"empty auth header"}`,
		},
		{
			name:        "Invalid Header",
			headerName:  "Authorization",
			headerValue: "Bearr token",
			token:       "token",
			mockBehaviour: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedCode: 401,
			expectedBody: `{"massage":"invalid auth header"}`,
		},
		{
			name:        "Invalid token",
			headerName:  "Authorization",
			headerValue: "Bearer ",
			mockBehaviour: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedCode: 401,
			expectedBody: `{"massage":"token is empty"}`,
		},
		{
			name:        "Service failure",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehaviour: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, errors.New("failed to parse token"))
			},
			expectedCode: 401,
			expectedBody: `{"massage":"failed to parse token"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehaviour(auth, testCase.token)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			r.GET("/protected", handler.userIdentity, func(ctx *gin.Context) {
				id, _ := ctx.Get(userCtx)
				ctx.String(200, fmt.Sprintf("%d", id.(int)))
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/protected", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedCode, w.Code)
			assert.Equal(t, testCase.expectedBody, w.Body.String())

		})
	}
}

func TestHandler_GetUserId(t *testing.T) {
	var getContext = func(id int) *gin.Context {
		ctx := &gin.Context{}
		ctx.Set(userCtx, id)
		return ctx
	}

	testTable := []struct {
		name string
		ctx  *gin.Context
		id   int
		Fail bool
	}{
		{
			name: "OK",
			ctx:  getContext(1),
			id:   1,
		},
		{
			name: "Empty Context",
			ctx:  &gin.Context{},
			Fail: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			id, err := getUserID(testCase.ctx)
			if testCase.Fail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

			}

			assert.Equal(t, testCase.id, id)
		})
	}
}
