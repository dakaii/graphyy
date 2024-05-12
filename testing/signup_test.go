package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"graphyy/controller"
	"graphyy/database"
	"graphyy/domain"
	"graphyy/repository"
	"graphyy/view"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SignUpTestSuite struct {
	suite.Suite
	rr      *httptest.ResponseRecorder
	handler http.Handler
}

type SignUpResponse struct {
	Data struct {
		Signup domain.AuthToken `json:"signup"`
	} `json:"data"`
}

func (suite *SignUpTestSuite) SetupTest() {
	db := database.GetDatabase()
	repos := repository.InitRepositories(db)
	controllers := controller.InitControllers(repos)
	schema := view.Schema(controllers)
	suite.handler = view.GraphqlHandlfunc(schema)

	suite.rr = httptest.NewRecorder()
}

func (suite *SignUpTestSuite) TearDownTest() {
	TruncateAllTables()
}

func (suite *SignUpTestSuite) TestCreateUser() {
	username := "testuser"
	password := "password"

	query := fmt.Sprintf(`{ "query": "mutation { signup(username: \"%s\", password: \"%s\") { token, tokenType, expiresIn } }" }`, username, password)
	byteArray := []byte(query)

	req, err := http.NewRequest("POST", "/test-graphql", bytes.NewBuffer(byteArray))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	suite.handler.ServeHTTP(suite.rr, req)

	var res SignUpResponse
	err = json.Unmarshal(suite.rr.Body.Bytes(), &res)
	suite.NoError(err)

	assert.Equal(suite.T(), http.StatusOK, suite.rr.Code)
	assert.Equal(suite.T(), "Bearer", res.Data.Signup.TokenType)
	assert.IsType(suite.T(), "", res.Data.Signup.Token)
	assert.IsType(suite.T(), int64(0), res.Data.Signup.ExpiresIn)
}

func TestSignUpTestSuite(t *testing.T) {
	suite.Run(t, new(SignUpTestSuite))
}
