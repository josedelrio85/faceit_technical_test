package faceit_cc

import (
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

type Suite struct {
	DB     *gorm.DB
	Mock   sqlmock.Sqlmock
	Assert *assert.Assertions
	User   *User
}

func (s *Suite) NewSuite(t *testing.T) {
	s.Assert = assert.New(t)
	db, mock, err := sqlmock.New()
	s.Assert.NoError(err)

	s.DB, err = gorm.Open("sqlmock", db)
	s.Assert.NoError(err)

	s.Mock = mock
}

func (s *Suite) MockInsert(user User) sqlmock.Sqlmock {
	s.Mock.ExpectBegin()
	s.Mock.ExpectExec(`INSERT INTO (.+)`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.Mock.ExpectCommit()
	return s.Mock
}

func (s *Suite) MockGetById(user User) sqlmock.Sqlmock {
	s.Mock.ExpectBegin()
	s.Mock.ExpectQuery(`SELECT (.+)`).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(user.Id))
	s.Mock.ExpectCommit()
	return s.Mock
}

func (s *Suite) MockUpdate(user User) sqlmock.Sqlmock {
	s.Mock.ExpectBegin()
	s.Mock.ExpectExec(`UPDATE (.+)`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.Mock.ExpectCommit()
	return s.Mock
}

func (s *Suite) MockDelete(user User) sqlmock.Sqlmock {

	s.Mock.ExpectBegin()
	s.Mock.ExpectExec(`DELETE (.+)`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.Mock.ExpectCommit()
	return s.Mock
}

func (s *Suite) MockList(user User) sqlmock.Sqlmock {
	s.Mock.ExpectBegin()
	s.Mock.ExpectQuery(`SELECT (.+)`).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(user.Id).AddRow(user.Id).AddRow(user.Id))
	s.Mock.ExpectCommit()
	return s.Mock
}

func TestAdd(t *testing.T) {

	tests := []struct {
		Description string
		Input       User
	}{
		{
			Description: "",
			Input: User{
				Id:        "asdfasdf",
				FirstName: "asdfasdf",
				LastName:  "asdfasdf",
				NickName:  "asdfasdf",
				Password:  "asdfasdf",
				Email:     "asdfasdf",
				Country:   "asdfasdf",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}

	suite := Suite{}
	suite.NewSuite(t)

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {

			suite.MockInsert(test.Input)

			err := test.Input.Add(suite.DB)
			suite.Assert.NoError(err)
		})
	}
}

func TestGetById(t *testing.T) {

	tests := []struct {
		Description string
		Input       User
	}{
		{
			Description: "",
			Input: User{
				Id:        "asdfasdf",
				FirstName: "asdfasdf",
				LastName:  "asdfasdf",
				NickName:  "asdfasdf",
				Password:  "asdfasdf",
				Email:     "asdfasdf",
				Country:   "asdfasdf",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}

	suite := Suite{}
	suite.NewSuite(t)

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			suite.MockGetById(test.Input)
			err := test.Input.GetById(suite.DB)
			suite.Assert.NoError(err)
		})
	}
}

func TestUpdate(t *testing.T) {

	tests := []struct {
		Description string
		Input       User
	}{
		{
			Description: "",
			Input: User{
				Id:        "asdfasdf",
				FirstName: "asdfasdf",
				LastName:  "asdfasdf",
				NickName:  "asdfasdf",
				Password:  "asdfasdf",
				Email:     "asdfasdf",
				Country:   "asdfasdf",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}

	suite := Suite{}
	suite.NewSuite(t)

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {

			suite.MockUpdate(test.Input)

			err := test.Input.Update(suite.DB)
			suite.Assert.NoError(err)
		})
	}
}

func TestDelete(t *testing.T) {

	tests := []struct {
		Description string
		Input       User
	}{
		{
			Description: "",
			Input: User{
				Id: "asdfasdf",
			},
		},
	}

	suite := Suite{}
	suite.NewSuite(t)

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			suite.MockDelete(test.Input)

			err := test.Input.Delete(suite.DB)
			suite.Assert.NoError(err)
		})
	}
}

func TestList(t *testing.T) {

	tests := []struct {
		Description string
		Input       User
	}{
		{
			Description: "",
			Input: User{
				Id:        "asdfasdf",
				FirstName: "asdfasdf",
			},
		},
	}

	suite := Suite{}
	suite.NewSuite(t)

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			suite.MockList(test.Input)
			pagination := Pagination{
				SearchBy:    "",
				SearchValue: "",
				Offset:      0,
				ResultsPage: 10,
			}
			_, err := test.Input.List(suite.DB, pagination)
			suite.Assert.NoError(err)
		})
	}
}
