package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"graphyy/controller"
	"graphyy/database"
	"graphyy/entity"
	"graphyy/internal"
	"graphyy/repository"
	"graphyy/testing/factory"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LoginTestSuite struct {
	suite.Suite
	users   []entity.User
	rr      *httptest.ResponseRecorder
	handler http.Handler
}

type LogInResponse struct {
	Data struct {
		Login entity.AuthToken `json:"login"`
	} `json:"data"`
}

func (suite *LoginTestSuite) SetupTest() {
	db := database.GetDatabase()
	repos := repository.InitRepositories(db)
	controllers := controller.InitControllers(repos)
	schema := controller.Schema(controllers)
	suite.handler = controller.GraphqlHandlfunc(schema)

	suite.users = factory.CreateUsers(db, 5)
	suite.rr = httptest.NewRecorder()
}

func (suite *LoginTestSuite) TearDownTest() {
	TruncateAllTables()
}

func (suite *LoginTestSuite) TestLoginUser() {
	loginUser := suite.users[0]
	token := internal.GenerateJWT(loginUser)

	query := fmt.Sprintf(`{ "query": "mutation { login(username: \"%s\", password: \"%s\") { token, tokenType, expiresIn } }" }`, loginUser.Username, loginUser.Password)
	byteArray := []byte(query)

	req, err := http.NewRequest("POST", "/test-graphql", bytes.NewBuffer(byteArray))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))

	suite.handler.ServeHTTP(suite.rr, req)

	var res LogInResponse
	err = json.Unmarshal(suite.rr.Body.Bytes(), &res)
	suite.NoError(err)

	assert.Equal(suite.T(), http.StatusOK, suite.rr.Code)
	assert.Equal(suite.T(), "Bearer", res.Data.Login.TokenType)
	assert.IsType(suite.T(), "", res.Data.Login.Token)
	assert.IsType(suite.T(), int64(0), res.Data.Login.ExpiresIn)
}

func TestLoginTestSuite(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}
