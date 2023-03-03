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

func TestHandlerTestSuite(t *testing.T) {
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
		bodyReader := strings.NewReader(`{"email": "abc@gmail.com", "password": "abc@12", "name": "abc" , "role": "user"}`)
		r := httptest.NewRequest(http.MethodPost, "/register", (bodyReader))
		w := httptest.NewRecorder()
		ctx := r.Context()
		respBody := domain.UserResponse{
			Message: "user register successfully",
		}
		requestBody := domain.Users{
			Email:    "abc@gmail.com",
			Password: "abc@12",
			Name:     "abc",
			Role:     "user",
		}
		s.service.On("RegisterUser", ctx, requestBody).Return(nil).Once()

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
		respBody := "Invalid mail"

		deps := Dependencies{
			bookService: s.service,
		}

		exp, _ := json.Marshal(respBody)
		got := registerUserHandler(deps)
		got.ServeHTTP(w, r)
		assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
		//assert.Equal(t, string(exp), w.Body.String())
		assert.Equal(t, strings.TrimSpace(string(exp)), strings.TrimSpace(w.Body.String()))

	})

	t.Run("when invalid register request is made", func(t *testing.T) {
		bodyReader := strings.NewReader(`{"email": "rutuja@gmail.com", "name": "Rutuja" , "role": "user"}`)
		r := httptest.NewRequest(http.MethodPost, "/register", (bodyReader))
		w := httptest.NewRecorder()
		respBody := "Invalid request body"

		deps := Dependencies{
			bookService: s.service,
		}
		exp, _ := json.Marshal(respBody)
		got := registerUserHandler(deps)
		got.ServeHTTP(w, r)
		assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
		//assert.Equal(t, string(exp), w.Body.String())
		assert.Equal(t, strings.TrimSpace(string(exp)), strings.TrimSpace(w.Body.String()))

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
		respBody := "Invalid request body"

		deps := Dependencies{
			bookService: s.service,
		}
		exp, _ := json.Marshal(respBody)
		got := loginUserHandler(deps)
		got.ServeHTTP(w, r)
		assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
		assert.Equal(t, strings.TrimSpace(string(exp)), strings.TrimSpace(w.Body.String()))

	})

	t.Run("when email is invalid", func(t *testing.T) {
		bodyReader := strings.NewReader(`{"email": "rutuja@@gmail.com", "password":"rutuja@12"}`)

		r := httptest.NewRequest(http.MethodPost, "/login", (bodyReader))
		w := httptest.NewRecorder()
		respBody := "invalid mail"
		deps := Dependencies{
			bookService: s.service,
		}
		exp, _ := json.Marshal(respBody)
		got := loginUserHandler(deps)
		got.ServeHTTP(w, r)
		assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
		assert.Equal(t, strings.TrimSpace(string(exp)), strings.TrimSpace(w.Body.String()))

	})
}

func (s *HandlerTestSuite) Test_addBooksHandler() {
	t := s.T()
	t.Run("when add request is valid", func(t *testing.T) {
		bodyReader := strings.NewReader(`{  
		"book_author":"wegeg",
		"publisher":"qyege",
		"quantity":3,
		"status":"Available"    }`)
		r := httptest.NewRequest(http.MethodPost, "/addbook", (bodyReader))
		w := httptest.NewRecorder()
		deps := Dependencies{
			bookService: s.service,
		}

		respBody := "Invalid request body"

		exp, _ := json.Marshal(respBody)
		got := addBooksHandler(deps)
		got.ServeHTTP(w, r)
		assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
		assert.Equal(t, string(exp), strings.TrimSpace(w.Body.String()))

	})

}

func (s *HandlerTestSuite) Test_getAllBooksHandler() {
	t := s.T()
	t.Run("when a valid get all books request is made ", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/books", nil)
		w := httptest.NewRecorder()
		ctx := r.Context()
		respBody := []domain.GetAllBooksResponse{
			{
				BookID:     "2",
				BookName:   "feberbr",
				BookAuthor: "dbwehe",
				Publisher:  "eddhjed",
				Quantity:   3,
				Status:     "Available",
			},
		}
		s.service.On("GetBooks", ctx).Return(respBody, nil).Once()
		deps := Dependencies{
			bookService: s.service,
		}

		exp, _ := json.Marshal(respBody)
		got := getAllBooksHandler(deps)
		got.ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, string(exp), w.Body.String())
	})

	t.Run("when getting error in GET request ", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/books", nil)
		w := httptest.NewRecorder()
		//r = r.WithContext(context.WithValue(r.Context(), "token", uint(1)))
		ctx := r.Context()
		respBody := "error in getting books"
		s.service.On("GetBooks", ctx).Return(nil, errors.New("error returning books")).Once()
		deps := Dependencies{
			bookService: s.service,
		}

		exp, _ := json.Marshal(respBody)
		got := getAllBooksHandler(deps)
		got.ServeHTTP(w, r)
		assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
		//	assert.Equal(t, string(exp), w.Body.String())
		assert.Equal(t, string(exp), strings.TrimSpace(w.Body.String()))

	})
}

func (s *HandlerTestSuite) Test_issueBookHandler() {
	t := s.T()
	t.Run("when valid request is made", func(t *testing.T) {
		bodyReader := strings.NewReader(`{"book_id":0}`)
		r := httptest.NewRequest(http.MethodPost, "/issue", bodyReader)
		w := httptest.NewRecorder()
		deps := Dependencies{
			bookService: s.service,
		}

		respBody := "Invalid request body"

		exp, _ := json.Marshal(respBody)
		got := issueBookHandler(deps)
		got.ServeHTTP(w, r)
		assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
		assert.Equal(t, string(exp), strings.TrimSpace(w.Body.String()))

	})

}

func (s *HandlerTestSuite) Test_ResetPasswordHandler() {
	t := s.T()
	t.Run("when invalid request is made", func(t *testing.T) {
		bodyReader := strings.NewReader(`{"current_password":,"new_password":}`)
		r := httptest.NewRequest(http.MethodPost, "/updatepassword", bodyReader)
		w := httptest.NewRecorder()
		deps := Dependencies{
			bookService: s.service,
		}

		respBody := "Invalid request body"

		exp, _ := json.Marshal(respBody)
		got := ResetPasswordHandler(deps)
		got.ServeHTTP(w, r)
		assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
		assert.Equal(t, string(exp), strings.TrimSpace(w.Body.String()))

	})
}

func (s *HandlerTestSuite) Test_UpdateNameHandler() {
	t := s.T()
	t.Run("when invalid request is made", func(t *testing.T) {
		bodyReader := strings.NewReader(`{"current_name":,"new_name":}`)
		r := httptest.NewRequest(http.MethodPost, "/updatename", bodyReader)
		w := httptest.NewRecorder()
		deps := Dependencies{
			bookService: s.service,
		}

		respBody := "Invalid request body"

		exp, _ := json.Marshal(respBody)
		got := ResetPasswordHandler(deps)
		got.ServeHTTP(w, r)
		assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
		assert.Equal(t, string(exp), strings.TrimSpace(w.Body.String()))

	})
}

// func (s *HandlerTestSuite) Test_getUserByEmailNameHandler() {
// 	t := s.T()
// 	t.Run("when a invalid get  request is made ", func(t *testing.T) {
// 		r := httptest.NewRequest(http.MethodGet, "/users/email/name", nil)
// 		w := httptest.NewRecorder()
// 		respBody := "query parameters required"
// 		deps := Dependencies{
// 			bookService: s.service,
// 		}

// 		exp, _ := json.Marshal(respBody)
// 		got := getUserByEmailNameHandler(deps)
// 		got.ServeHTTP(w, r)
// 		assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
// 		assert.Equal(t, string(exp), strings.TrimSpace(w.Body.String()))

// 	})
// }

func (s *HandlerTestSuite) Test_getBooksActivityHandler() {
	t := s.T()
	t.Run("when a invalid get  request is made ", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/books/activity", nil)
		w := httptest.NewRecorder()
		respBody := "Authorization header is required"
		deps := Dependencies{
			bookService: s.service,
		}

		exp, _ := json.Marshal(respBody)
		got := getBooksActivityHandler(deps)
		got.ServeHTTP(w, r)
		assert.Equal(t, http.StatusBadRequest, http.StatusBadRequest)
		assert.Equal(t, string(exp), strings.TrimSpace(w.Body.String()))

	})
}

func (s *HandlerTestSuite) Test_ReturnBookHandler() {
	t := s.T()
	t.Run("when invalid request is made", func(t *testing.T) {
		bodyReader := strings.NewReader(`{"book_id":0}`)
		r := httptest.NewRequest(http.MethodPost, "/users/return", bodyReader)
		w := httptest.NewRecorder()
		deps := Dependencies{
			bookService: s.service,
		}

		respBody := "invalid request body"

		exp, _ := json.Marshal(respBody)
		got := ReturnBookHandler(deps)
		got.ServeHTTP(w, r)
		assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
		//assert.Equal(t, string(exp), w.Body.String())
		assert.Equal(t, string(exp), strings.TrimSpace(w.Body.String()))

	})

}
