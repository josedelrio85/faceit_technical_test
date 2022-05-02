package faceit_cc

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddEntity(t *testing.T) {
	assert := assert.New(t)

	suite := Suite{}
	suite.NewSuite(t)

	handler := Handler{
		Database: Database{
			Db: suite.DB,
		},
	}

	kafka := KafkaMock{
		Ctx:     context.Background(),
		Brokers: []string{""},
		Topic:   "",
	}
	err := kafka.Initialize()
	assert.NoError(err)

	tests := []struct {
		Description      string
		Input            User
		Response         Response
		ExpectedResponse int
		ExpectedSucceed  bool
	}{
		{
			Description: "when input data is not valid - no data",
			Input: User{
				Kafka: &kafka,
			},
			Response:         Response{},
			ExpectedResponse: http.StatusUnprocessableEntity,
			ExpectedSucceed:  false,
		},
		{
			Description: "when input data is not valid - no uuid",
			Input: User{
				FirstName: "zzzzzzzz2",
				LastName:  "asdfasdf2",
				NickName:  "asdfasdf2",
				Password:  "asdfasdf2",
				Email:     "asdfasdf@asdfas.es",
				Country:   "ES",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Kafka:     &kafka,
			},
			Response:         Response{},
			ExpectedResponse: http.StatusUnprocessableEntity,
			ExpectedSucceed:  false,
		},
		{
			Description: "when input data is not valid - invalid email",
			Input: User{
				Id:        uuid.New().String(),
				FirstName: "zzzzzzzz2",
				LastName:  "asdfasdf2",
				NickName:  "asdfasdf2",
				Password:  "asdfasdf2",
				Email:     "sdfas",
				Country:   "ES",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Kafka:     &kafka,
			},
			Response:         Response{},
			ExpectedResponse: http.StatusUnprocessableEntity,
			ExpectedSucceed:  false,
		},
		{
			Description: "when input data is not valid - no data firstname, nickname, password",
			Input: User{
				Id:        uuid.New().String(),
				LastName:  "asdfasdf2",
				Email:     "asdfasdf@asdfas.es",
				Country:   "ES",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Kafka:     &kafka,
			},
			Response:         Response{},
			ExpectedResponse: http.StatusUnprocessableEntity,
			ExpectedSucceed:  false,
		},
		{
			Description: "when input data is valid",
			Input: User{
				Id:        uuid.New().String(),
				FirstName: "zzzzzzzz2",
				LastName:  "asdfasdf2",
				NickName:  "asdfasdf2",
				Password:  "asdfasdf2",
				Email:     "asdfasdf@asdfas.es",
				Country:   "ES",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Kafka:     &kafka,
			},
			Response:         Response{},
			ExpectedResponse: http.StatusOK,
			ExpectedSucceed:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			suite.MockInsert(test.Input)
			err := handler.Add(&test.Input, &test.Response)
			if test.ExpectedSucceed {
				assert.NoError(err)
			} else {
				assert.NotNil(err)
			}
			assert.Equal(test.Response.Status, test.ExpectedResponse)
		})
	}
}

func TestUpdateEntity(t *testing.T) {
	assert := assert.New(t)

	suite := Suite{}
	suite.NewSuite(t)

	handler := Handler{
		Database: Database{
			Db: suite.DB,
		},
	}

	kafka := KafkaMock{
		Ctx:     context.Background(),
		Brokers: []string{""},
		Topic:   "",
	}
	err := kafka.Initialize()
	assert.NoError(err)

	tests := []struct {
		Description      string
		Input            User
		Response         Response
		ExpectedResponse int
		ExpectedSucceed  bool
	}{
		{
			Description: "when input data is not valid - no data",
			Input: User{
				Kafka: &kafka,
			},
			Response:         Response{},
			ExpectedResponse: http.StatusUnprocessableEntity,
			ExpectedSucceed:  false,
		},
		{
			Description: "when input data is not valid - no uuid",
			Input: User{
				FirstName: "zzzzzzzz2",
				LastName:  "asdfasdf2",
				NickName:  "asdfasdf2",
				Password:  "asdfasdf2",
				Email:     "asdfasdf@asdfas.es",
				Country:   "ES",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Kafka:     &kafka,
			},
			Response:         Response{},
			ExpectedResponse: http.StatusUnprocessableEntity,
			ExpectedSucceed:  false,
		},
		{
			Description: "when input data is not valid - invalid email",
			Input: User{
				Id:        uuid.New().String(),
				FirstName: "zzzzzzzz2",
				LastName:  "asdfasdf2",
				NickName:  "asdfasdf2",
				Password:  "asdfasdf2",
				Email:     "sdfas",
				Country:   "ES",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Kafka:     &kafka,
			},
			Response:         Response{},
			ExpectedResponse: http.StatusUnprocessableEntity,
			ExpectedSucceed:  false,
		},
		{
			Description: "when input data is not valid - no data firstname, nickname, password",
			Input: User{
				Id:        uuid.New().String(),
				LastName:  "asdfasdf2",
				Email:     "asdfasdf@asdfas.es",
				Country:   "ES",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Kafka:     &kafka,
			},
			Response:         Response{},
			ExpectedResponse: http.StatusUnprocessableEntity,
			ExpectedSucceed:  false,
		},
		{
			Description: "when input data is valid",
			Input: User{
				Id:        uuid.New().String(),
				FirstName: "zzzzzzzz2",
				LastName:  "asdfasdf2",
				NickName:  "asdfasdf2",
				Password:  "asdfasdf2",
				Email:     "asdfasdf@asdfas.es",
				Country:   "ES",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Kafka:     &kafka,
			},
			Response:         Response{},
			ExpectedResponse: http.StatusOK,
			ExpectedSucceed:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			suite.MockUpdate(test.Input)
			err := handler.Update(&test.Input, &test.Response)
			if test.ExpectedSucceed {
				assert.NoError(err)
			} else {
				assert.NotNil(err)
			}
			assert.Equal(test.Response.Status, test.ExpectedResponse)
		})
	}
}

func TestDeleteEntity(t *testing.T) {
	assert := assert.New(t)

	suite := Suite{}
	suite.NewSuite(t)

	handler := Handler{
		Database: Database{
			Db: suite.DB,
		},
	}

	kafka := KafkaMock{
		Ctx:     context.Background(),
		Brokers: []string{""},
		Topic:   "",
	}
	err := kafka.Initialize()
	assert.NoError(err)

	tests := []struct {
		Description      string
		Input            User
		Response         Response
		ExpectedResponse int
		ExpectedSucceed  bool
	}{
		{
			Description: "when input data is not valid - no data",
			Input: User{
				Kafka: &kafka,
			},
			Response:         Response{},
			ExpectedResponse: http.StatusUnprocessableEntity,
			ExpectedSucceed:  false,
		},
		{
			Description: "when input data is not valid - no uuid",
			Input: User{
				Kafka: &kafka,
			},
			Response:         Response{},
			ExpectedResponse: http.StatusUnprocessableEntity,
			ExpectedSucceed:  false,
		},
		{
			Description: "when input data is valid",
			Input: User{
				Id:    uuid.New().String(),
				Kafka: &kafka,
			},
			Response:         Response{},
			ExpectedResponse: http.StatusOK,
			ExpectedSucceed:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			suite.MockDelete(test.Input)
			err := handler.Delete(&test.Input, &test.Response)
			if test.ExpectedSucceed {
				assert.NoError(err)
			} else {
				assert.NotNil(err)
			}
			assert.Equal(test.Response.Status, test.ExpectedResponse)
		})
	}
}

func TestListEntity(t *testing.T) {
	assert := assert.New(t)

	suite := Suite{}
	suite.NewSuite(t)

	handler := Handler{
		Database: Database{
			Db: suite.DB,
		},
	}

	kafka := KafkaMock{
		Ctx:     context.Background(),
		Brokers: []string{""},
		Topic:   "",
	}
	err := kafka.Initialize()
	assert.NoError(err)

	tests := []struct {
		Description           string
		Input                 User
		Response              Response
		ExpectedResponse      int
		ExpectedNumberResults int
		ExpectedSucceed       bool
	}{
		{
			Description: "when pagination property is empty, return default number of results",
			Input: User{
				Kafka: &kafka,
			},
			Response:              Response{},
			ExpectedNumberResults: 3,
			ExpectedSucceed:       true,
		},
		{
			Description: "when input data is valid",
			Input: User{
				Pagination: Pagination{
					SearchBy:    "country",
					SearchValue: "UK",
					ResultsPage: 3,
					Offset:      0,
				},
				Kafka: &kafka,
			},
			ExpectedNumberResults: 3,
			ExpectedSucceed:       true,
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			suite.MockList(test.Input)
			err := handler.List(&test.Input, &test.Response)
			if test.ExpectedSucceed {
				assert.NoError(err)
			} else {
				assert.NotNil(err)
			}
			resultlength := len(test.Response.Data.UserList)
			assert.Equal(resultlength, test.ExpectedNumberResults)
		})
	}
}
