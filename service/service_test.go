package service

import (
	"library_management/domain"
	"library_management/mocks"
	"context"
	"errors"
	"testing"

	//"github.com/stretchr/testify/assert"
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


func (s *ServiceTestSuite) Testbookservice_Register(){
	t:=s.T()

	type args struct {
		ctx    context.Context
		book domain.Users
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		prepare func(args, *mocks.Storer)
	}{
		{
		name:"check whether user data is valid ",
		args: args{
			ctx:context.TODO(),
			user:domain.Users{
		        Email:"rutujarohom@gmail.com",
				Password :"rutuja@12",
				Name :"rutuja",
				Role :	"user",
			},
			},
			wantErr:false,
			prepare: func(args,*mocks.Storer){
				s.On("CreateUser", context.TODO(), mock.Anything).Return(nil).Once()
			},
		},
			{
				name:"when user data is false",
				args:args{
					ctx:context.TODO(),
					user:domain.Users{
						Email   :"rutujarohom@gmail.com" ,
				 		Password :"rutuja@12" ,
						Name :"rutuja" ,
						Role :	"user" ,
					},
				},
				wantErr:true,
				prepare:func(args,*mocks.Storer){
					s.On("CreateUser", context.TODO(), mock.Anything).Return(errors.New("mocked error")).Once()

				},

			},
			{
				name:"when user already exist",
				args:args{
					ctx:context.TODO(),
					user:domain.Users{
						Email   :"rutujarohom@gmail.com" ,
				 		Password :"rutuja@12" ,
						Name :"rutuja" ,
						Role :	"user" ,
					},

				},
				wantErr:true,
				prepare:func(args,*mocks.Storer){
					s.On("CreateUser", context.TODO(), mock.Anything).Return(errors.New("pq: duplicate key value violates unique constraint \"email_key\"")).Once()

			},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(tt.args, s.repo)
			err := s.service.RegisterUser(tt.args.ctx, tt.args.book)
			//t.Log("here", gotAddeduser)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			//assert.Equal(t, domain.UserResponse, "no error ")
		})
	}
}

