package service

import (
	"encoding/json"
	"errors"
	"library_management/domain"
	"library_management/mocks"
	"net/http"
	"net/http/httptest"

	//"strings"

	//"net/http/httptest"

	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	//"gotest.tools/assert"
)

type HandlerTestSuite struct {
	suite.Suite
	service *mocks.Services
}

//deps := Dependencies{bookService: suite.service}

func TestlibarayHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (suite *HandlerTestSuite) SetupTest() {
	suite.service = &mocks.Services{}
}

func (suite *HandlerTestSuite) TearDownTest() {
	suite.service.AssertExpectations(suite.T())
}
func (s *HandlerTestSuite) Test_registerUserHandler() {
	t := s.T()

	t.Run("when valid register request is made", func(t *testing.T) {
		bodyReader := strings.NewReader(`{"email": "rutuja@gmail.com", "password": "rutuja", "name": "Rutuja" , "role": "user"}`)
		r := httptest.NewRequest(http.MethodPost, "/register", (bodyReader))
		w := httptest.NewRecorder()
		ctx := r.Context()
		respBody := domain.UserResponse{
			Message: "user register successfully",
		}
		requestBody := domain.Users{
			Email:    "rutuja@gmail.com",
			Password: "rutuja@12",
			Name:     "rutuja",
			Role:     "user",
		}
		s.service.On("RegisterUser", ctx, requestBody).Return(respBody, nil).Once()

		deps := Dependencies{
			bookService: s.service,
		}
		exp, _ := json.Marshal(respBody)
		got := registerUserHandler(deps)
		got.ServeHTTP(w, r)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
		assert.Equal(t, string(exp), w.Body.String())
	})

	t.Run("when invalid email", func(t *testing.T) {
		bodyReader := strings.NewReader(`{"email": "rutuja@@gmail.com", "password": "rutuja", "name": "Rutuja" , "role": "user"}`)
		r := httptest.NewRequest(http.MethodPost, "/register", (bodyReader))
		w := httptest.NewRecorder()
		ctx := r.Context()
		respBody := domain.UserResponse{
			Message: "invalid mail",
		}
		requestBody := domain.Users{
			Email:    "rutuja@@gmail.com",
			Password: "rutuja@12",
			Name:     "rutuja",
			Role:     "user",
		}
		s.service.On("RegisterUser", ctx, requestBody).Return(respBody, nil).Once()

		deps := Dependencies{
			bookService: s.service,
		}
		exp, _ := json.Marshal(respBody)
		got := registerUserHandler(deps)
		got.ServeHTTP(w, r)
		assert.Equal(t, w.Result().StatusCode, http.StatusCreated)
		assert.Equal(t, string(exp), w.Body.String())
	})

	t.Run("when valid register request is made", func(t *testing.T) {
		bodyReader := strings.NewReader(`{"email": "rutuja@gmail.com", "name": "Rutuja" , "role": "user"}`)
		r := httptest.NewRequest(http.MethodPost, "/register", (bodyReader))
		w := httptest.NewRecorder()
		ctx := r.Context()
		respBody := domain.UserResponse{
			Message: "invalid request body",
		}
		requestBody := domain.Users{
			Email: "rutuja@gmail.com",
			//Password: "rutuja@12",
			Name: "rutuja",
			Role: "user",
		}
		s.service.On("RegisterUser", ctx, requestBody).Return(respBody, nil).Once()

		deps := Dependencies{
			bookService: s.service,
		}
		exp, _ := json.Marshal(respBody)
		got := registerUserHandler(deps)
		got.ServeHTTP(w, r)
		assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
		assert.Equal(t, string(exp), w.Body.String())
	})

}

func (s *HandlerTestSuite) Test_loginUserHandler() {
	t := s.T()
	t.Run("when login request is valid", func(t *testing.T) {
		bodyReader := strings.NewReader(`{"email": "rutuja@gmail.com", "password":"rutuja@12"}`)
		r := httptest.NewRequest(http.MethodPost, "/login", (bodyReader))
		w := httptest.NewRecorder()
		ctx := r.Context()
		respBody := domain.LoginResponse{
			Message: "login successful",
			Token:   "token",
		}
		requestBody := domain.LoginRequest{
			Email:    "rutuja@gmail.com",
			Password: "rutuja@12",
		}
		s.service.On("Login", ctx, requestBody).Return("token", nil).Once()

		deps := Dependencies{
			bookService: s.service,
		}
		exp, _ := json.Marshal(respBody)
		got := loginUserHandler(deps)
		got.ServeHTTP(w, r)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
		assert.Equal(t, string(exp), w.Body.String())

	})

	t.Run("when login request is invalid", func(t *testing.T) {
		bodyReader := strings.NewReader(`{"email": "rutuja@gmail.com"}`)
		r := httptest.NewRequest(http.MethodPost, "/login", (bodyReader))
		w := httptest.NewRecorder()
		ctx := r.Context()
		respBody := domain.LoginResponse{
			Message: "login successful",
			Token:   "token",
		}
		requestBody := domain.LoginRequest{
			Email: "rutuja@gmail.com",
			//Password: "rutuja@12",
		}
		s.service.On("Login", ctx, requestBody).Return("token", nil).Once()

		deps := Dependencies{
			bookService: s.service,
		}
		exp, _ := json.Marshal(respBody)
		got := loginUserHandler(deps)
		got.ServeHTTP(w, r)
		assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
		assert.Equal(t, string(exp), w.Body.String())

	})

	t.Run("when login request and token invalid", func(t *testing.T) {
		bodyReader := strings.NewReader(`{"email": "rutuja@gmail.com","}`)
		r := httptest.NewRequest(http.MethodPost, "/login", (bodyReader))
		w := httptest.NewRecorder()
		ctx := r.Context()
		respBody := domain.LoginResponse{
			Message: "invalid token",
		}
		requestBody := domain.LoginRequest{
			Email:    "rutuja@gmail.com",
			Password: "rutuja@12",
		}
		s.service.On("Login", ctx, requestBody).Return(nil, errors.New("invalid token")).Once()

		deps := Dependencies{
			bookService: s.service,
		}
		exp, _ := json.Marshal(respBody)
		got := loginUserHandler(deps)
		got.ServeHTTP(w, r)
		assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
		assert.Equal(t, string(exp), w.Body.String())

	})

}
