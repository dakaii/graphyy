package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"graphyy/controller"
	"graphyy/database"
	"graphyy/domain"
	"graphyy/internal/auth"
	"graphyy/repository"
	"graphyy/testing/factory"
	"graphyy/view"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MeUpTestSuite struct {
	suite.Suite
	users   []domain.User
	rr      *httptest.ResponseRecorder
	handler http.Handler
}

type MeResponse struct {
	Data struct {
		Me struct {
			Username string `json:"username"`
		} `json:"me"`
	} `json:"data"`
}

func (suite *MeUpTestSuite) SetupTest() {
	db := database.GetDatabase()
	repos := repository.InitRepositories(db)
	controllers := controller.InitControllers(repos)
	schema := view.Schema(controllers)
	suite.handler = view.GraphqlHandlfunc(schema)

	suite.users = factory.CreateUsers(db, 5)
	suite.rr = httptest.NewRecorder()
}

func (suite *MeUpTestSuite) TearDownTest() {
	TruncateAllTables()
}
func (suite *MeUpTestSuite) TestMeEndpoint() {
	loginUser := suite.users[0]
	token := auth.GenerateJWT(loginUser)

	meQuery := `{ "query": "{ me { username } }" }`
	byteArray := []byte(meQuery)

	req, err := http.NewRequest("POST", "/test-graphql", bytes.NewBuffer(byteArray))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))
	suite.rr = httptest.NewRecorder()
	suite.handler.ServeHTTP(suite.rr, req)

	var res MeResponse
	err = json.Unmarshal(suite.rr.Body.Bytes(), &res)
	suite.NoError(err)

	assert.Equal(suite.T(), http.StatusOK, suite.rr.Code)
	assert.Equal(suite.T(), loginUser.Username, res.Data.Me.Username)
}

func TestMeTestSuite(t *testing.T) {
	suite.Run(t, new(MeUpTestSuite))
}
