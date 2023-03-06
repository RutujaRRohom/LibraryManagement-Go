package service

import (
	"context"
	"errors"
	"library_management/domain"
	"library_management/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	service Services
	repo    *mocks.Storer
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (suite *ServiceTestSuite) SetupTest() {
	suite.repo = &mocks.Storer{}
	suite.service = NewBookService(suite.repo)
}

func (suite *ServiceTestSuite) TearDownSuite() {
	suite.repo.AssertExpectations(suite.T())
}

func (s *ServiceTestSuite) Testbookservice_Register() {
	t := s.T()

	type args struct {
		ctx  context.Context
		user domain.Users
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		prepare func(args, *mocks.Storer)
	}{
		{
			name: "check whether user data is valid ",
			args: args{
				ctx: context.TODO(),
				user: domain.Users{
					Email:    "rutujarohom@gmail.com",
					Password: "rutuja@12",
					Name:     "rutuja",
					Role:     "user",
				},
			},
			wantErr: false,
			prepare: func(a args, s *mocks.Storer) {
				s.On("CreateUser", context.TODO(), mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "when user data is invalid",
			args: args{
				ctx: context.TODO(),
				user: domain.Users{
					Email:    "rutujarrohom@gmail.com",
					Password: "rutuja1@12",
					Name:     "rutuja",
					Role:     "",
				},
			},
			wantErr: true,
			prepare: func(a args, s *mocks.Storer) {
				s.On("CreateUser", context.TODO(), mock.Anything).Return(errors.New("mocked error")).Once()

			},
		},
		{
			name: "when user already exist",
			args: args{
				ctx: context.TODO(),
				user: domain.Users{
					Email:    "rutujarohom@gmail.com",
					Password: "rutuja@12",
					Name:     "rutuja",
					Role:     "user",
				},
			},
			wantErr: true,
			prepare: func(a args, s *mocks.Storer) {
				s.On("CreateUser", context.TODO(), mock.Anything).Return(errors.New("pq: duplicate key value violates unique constraint \"email_key\"")).Once()

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(tt.args, s.repo)
			err := s.service.RegisterUser(tt.args.ctx, tt.args.user)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

		})
	}
}

func (s *ServiceTestSuite) Testbookservice_Login() {
	t := s.T()

	type args struct {
		ctx   context.Context
		login domain.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		prepare func(args, *mocks.Storer)
	}{
		{
			name: "test for valid email and password body",
			args: args{
				ctx: context.TODO(),
				login: domain.LoginRequest{
					Email:    "admin@gmail.com",
					Password: "admin@12",
				},
			},
			wantErr: false,
			prepare: func(a args, s *mocks.Storer) {
				s.On("LoginUser", context.TODO(), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("Admin", 14, nil).Once()
			},
		},
		{
			name: "test for either email or password is incorrect",
			args: args{
				ctx: context.TODO(),
				login: domain.LoginRequest{
					Email:    "admin@gmail.com",
					Password: "admin@12",
				},
			},
			wantErr: true,
			prepare: func(a args, s *mocks.Storer) {
				s.On("LoginUser", context.TODO(), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", 14, errors.New("sql: no rows in result set")).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(tt.args, s.repo)
			gotToken, err := s.service.Login(tt.args.ctx, tt.args.login)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			var token string
			assert.IsType(t, token, gotToken)
		})
	}
}

func (s *ServiceTestSuite) TestAddbook() {
	t := s.T()
	type args struct {
		ctx context.Context
		add domain.AddBook
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		prepare func(args, *mocks.Storer)
	}{
		{
			name: "testing add book request",
			args: args{
				ctx: context.TODO(),
				add: domain.AddBook{
					BookName:   "born fire",
					BookAuthor: "ehege",
					Publisher:  "shsvs",
					Quantity:   3,
					Status:     "Available",
				},
			},
			wantErr: false,
			prepare: func(a args, s *mocks.Storer) {
				s.On("AddingBook", context.TODO(), mock.Anything).Return(int(10), nil).Once()
			},
		},
		{

			name: "testing  book request",
			args: args{
				ctx: context.TODO(),
				add: domain.AddBook{
					BookName:   "born fire",
					BookAuthor: "ehege",
					Publisher:  "shsvs",
					//Quantity: ,
					Status: "Available",
				},
			},
			wantErr: true,
			prepare: func(a args, s *mocks.Storer) {
				s.On("AddingBook", context.TODO(), mock.Anything).Return(int(10), errors.New(`null value in column "quantity" violates not-null constraint`)).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(tt.args, s.repo)
			added, err := s.service.AddBooks(tt.args.ctx, tt.args.add)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.IsType(t, domain.AddBookResponse{}, added)
		})
	}
}

func (s *ServiceTestSuite) TestGetbooks() {
	t := s.T()
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		prepare func(args, *mocks.Storer)
	}{
		{
			name: "get all books",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: false,
			prepare: func(a args, s *mocks.Storer) {
				s.On("GetAllBooksFromDb", context.TODO()).Return([]domain.GetAllBooksResponse{}, nil).Once()
			},
		},
		{
			name: "if error occured while getting books",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			prepare: func(a args, s *mocks.Storer) {
				s.On("GetAllBooksFromDb", context.TODO()).Return(nil, errors.New("mocked error")).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(tt.args, s.repo)
			books, err := s.service.GetBooks(tt.args.ctx)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.IsType(t, []domain.GetAllBooksResponse{}, books)
		})
	}
}

func (s *ServiceTestSuite) TestResetpassword() {
	t := s.T()
	type args struct {
		ctx   context.Context
		email string
		pass  domain.ResetPasswordRequest
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		prepare func(args, *mocks.Storer)
	}{
		{
			name: "valid current_password with new password",
			args: args{
				ctx:   context.TODO(),
				email: "rutujarohom@gmail.com",
				pass: domain.ResetPasswordRequest{
					CurrentPassword: "rutuja@12",
					NewPassword:     "rutuja@@12",
				},
			},
			wantErr: false,
			prepare: func(a args, s *mocks.Storer) {
				s.On("UpdatePassword", context.TODO(), mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "if current password is incorrect",
			args: args{
				ctx:   context.TODO(),
				email: "rutujarohom@gmail.com",
				pass: domain.ResetPasswordRequest{
					CurrentPassword: "rutuja@12",
					NewPassword:     "rutuja@@12",
				},
			},
			wantErr: true,
			prepare: func(a args, s *mocks.Storer) {
				s.On("UpdatePassword", context.TODO(), mock.Anything, mock.Anything).Return(errors.New("invalid current password")).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(tt.args, s.repo)
			err := s.service.ResetPassword(tt.args.ctx, tt.args.email, tt.args.pass)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

		})
	}
}

func (s *ServiceTestSuite) TestUpdatename() {
	t := s.T()

	type args struct {
		ctx   context.Context
		email string
		name  domain.ResetNameRequest
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		prepare func(args, *mocks.Storer)
	}{
		{
			name: "if current name is valid",
			args: args{
				ctx:   context.TODO(),
				email: "rutujarohom@gmail.com",
				name: domain.ResetNameRequest{
					CurrentName: "Rutuja",
					NewName:     "rutu",
				},
			},
			wantErr: false,
			prepare: func(a args, s *mocks.Storer) {
				s.On("Updatename", context.TODO(), mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "if current name is invalid",
			args: args{
				ctx:   context.TODO(),
				email: "rutujarohom@gmail.com",
				name: domain.ResetNameRequest{
					CurrentName: "rutuja", //invalid name
					NewName:     "rutu",
				},
			},
			wantErr: true,
			prepare: func(a args, s *mocks.Storer) {
				s.On("Updatename", context.TODO(), mock.Anything, mock.Anything).Return(errors.New("invalid name")).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(tt.args, s.repo)
			err := s.service.UpdateName(tt.args.ctx, tt.args.email, tt.args.name)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func (s *ServiceTestSuite) TestgetUsersByEmailName() {
	t := s.T()

	type args struct {
		ctx   context.Context
		email string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		prepare func(args, *mocks.Storer)
	}{
		{
			name: "if email and name is valid",
			args: args{
				ctx:   context.TODO(),
				email: "admin@gmail.com",
			},
			wantErr: false,
			prepare: func(a args, s *mocks.Storer) {
				s.On("GetUsers", context.TODO(), mock.Anything, mock.Anything).Return([]domain.GetUsersResponse{}, nil).Once()
			},
		},
		{
			name: "if email is invalid",
			args: args{
				ctx:   context.TODO(),
				email: "admin@gmail.com",
			},
			wantErr: true,
			prepare: func(a args, s *mocks.Storer) {
				s.On("GetUsers", context.TODO(), mock.Anything, mock.Anything).Return(nil, errors.New("invalid email")).Once()
			},
		},
	}

	for _, tt := range tests {
		tt.prepare(tt.args, s.repo)
		result, err := s.service.GetUsersByEmailName(tt.args.ctx, tt.args.email)
		if tt.wantErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
		assert.IsType(t, []domain.GetUsersResponse{}, result)
	}
}

func (s *ServiceTestSuite) TestGetbooksactivity() {
	t := s.T()

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		prepare func(args, *mocks.Storer)
	}{
		{
			name: "if get all books without error from repo layer",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: false,
			prepare: func(a args, s *mocks.Storer) {
				s.On("GetBookActivity", context.TODO()).Return([]domain.GetBooksActivityResponse{}, nil).Once()
			},
		},
		{
			name: " get error from repo layer",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			prepare: func(a args, s *mocks.Storer) {
				s.On("GetBookActivity", context.TODO()).Return(nil, errors.New("mocked error")).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(tt.args, s.repo)
			result, err := s.service.GetBooksActivity(tt.args.ctx)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.IsType(t, []domain.GetBooksActivityResponse{}, result)

		})
	}
}

func (s *ServiceTestSuite) TestGetbooksusers() {
	t := s.T()

	type args struct {
		ctx   context.Context
		email string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		prepare func(args, *mocks.Storer)
	}{
		{

			name: "if valid mail",
			args: args{
				ctx:   context.TODO(),
				email: "rutujarohom@gmail.com",
			},
			wantErr: false,
			prepare: func(a args, s *mocks.Storer) {
				s.On("GetUserBooks", context.TODO(), mock.Anything).Return([]domain.GetBooksResponse{}, nil).Once()
			},
		},
		{
			name: "if valid mail",
			args: args{
				ctx:   context.TODO(),
				email: "rutujarohom@gmail.com", //invalid
			},
			wantErr: true,
			prepare: func(a args, s *mocks.Storer) {
				s.On("GetUserBooks", context.TODO(), mock.Anything).Return(nil, errors.New("invalid mail")).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(tt.args, s.repo)
			result, err := s.service.Getbooks(tt.args.ctx, tt.args.email)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.IsType(t, []domain.GetBooksResponse{}, result)
		})

	}

}

func (s *ServiceTestSuite) TestIssuebook() {
	t := s.T()
	type args struct {
		ctx    context.Context
		UserID int
		issue  domain.IssueBookRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		prepare func(args, *mocks.Storer)
	}{
		{
			name: "for valid book_id",
			args: args{
				ctx:    context.TODO(),
				UserID: 2,
				issue: domain.IssueBookRequest{

					BookID: 3,
				},
			},
			wantErr: false,
			prepare: func(a args, s *mocks.Storer) {
				s.On("IssuedBook", context.TODO(), mock.Anything, mock.Anything).Return(domain.IssuedBookResponse{}, nil).Once()
			},
		},
		{
			name: "for invalid book_id",
			args: args{
				ctx:    context.TODO(),
				UserID: 2,
				issue: domain.IssueBookRequest{

					BookID: 50,
				},
			},
			wantErr: true,
			prepare: func(a args, s *mocks.Storer) {
				s.On("IssuedBook", context.TODO(), mock.Anything, mock.Anything).Return(domain.IssuedBookResponse{}, errors.New("book not exist with this id")).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(tt.args, s.repo)

			issued, err := s.service.IssueBook(tt.args.ctx, tt.args.UserID, tt.args.issue)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.IsType(t, domain.IssuedBookResponse{}, issued)

		})
	}
}

func (s *ServiceTestSuite) TestReturnBook() {
	t := s.T()
	type args struct {
		ctx      context.Context
		UserID   int
		returned domain.ReturnBookRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		prepare func(args, *mocks.Storer)
	}{
		{
			name: "for valid book_id",
			args: args{
				ctx:    context.TODO(),
				UserID: 2,
				returned: domain.ReturnBookRequest{

					BookID: 3,
				},
			},
			wantErr: false,
			prepare: func(a args, s *mocks.Storer) {
				s.On("ReturnBooks", context.TODO(), mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "for invalid book_id",
			args: args{
				ctx:    context.TODO(),
				UserID: 2,
				returned: domain.ReturnBookRequest{

					BookID: 50,
				},
			},
			wantErr: true,
			prepare: func(a args, s *mocks.Storer) {
				s.On("ReturnBooks", context.TODO(), mock.Anything, mock.Anything).Return(errors.New("book not exist with this id")).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(tt.args, s.repo)

			err := s.service.ReturnBook(tt.args.ctx, tt.args.UserID, tt.args.returned)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

		})
	}
}
