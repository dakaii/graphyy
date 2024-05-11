package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"graphyy/controller"
	"graphyy/database"
	"graphyy/entity"
	"graphyy/repository"
	"graphyy/view"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MeUpTestSuite struct {
	suite.Suite
	rr      *httptest.ResponseRecorder
	handler http.Handler
}

type MeResponse struct {
	Data struct {
		Signup entity.AuthToken `json:"signup"`
	} `json:"data"`
}

func (suite *MeUpTestSuite) SetupTest() {
	db := database.GetDatabase()
	repos := repository.InitRepositories(db)
	controllers := controller.InitControllers(repos)
	schema := view.Schema(controllers)
	suite.handler = view.GraphqlHandlfunc(schema)

	suite.rr = httptest.NewRecorder()
}

func (suite *MeUpTestSuite) TearDownTest() {
	TruncateAllTables()
}
func (suite *MeUpTestSuite) TestMeEndpoint() {

	username := "testuser"
	password := "password"

	// First, create a user
	signupQuery := fmt.Sprintf(`{ "query": "mutation { signup(username: \"%s\", password: \"%s\") { token, tokenType, expiresIn } }" }`, username, password)
	byteArray := []byte(signupQuery)

	req, err := http.NewRequest("POST", "/test-graphql", bytes.NewBuffer(byteArray))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	suite.handler.ServeHTTP(suite.rr, req)

	var signupRes SignUpResponse
	err = json.Unmarshal(suite.rr.Body.Bytes(), &signupRes)
	suite.NoError(err)

	// Then, use the token to query the me endpoint
	meQuery := `{ "query": "{ me { username } }" }`
	byteArray = []byte(meQuery)

	req, err = http.NewRequest("POST", "/test-graphql", bytes.NewBuffer(byteArray))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", signupRes.Data.Signup.Token)) // set the Authorization header
	suite.rr = httptest.NewRecorder()
	suite.handler.ServeHTTP(suite.rr, req)

	var meRes struct {
		Data struct {
			Me struct {
				Username string
			}
		}
	}
	err = json.Unmarshal(suite.rr.Body.Bytes(), &meRes)
	suite.NoError(err)

	// // Check the response
	assert.Equal(suite.T(), http.StatusOK, suite.rr.Code)
	assert.Equal(suite.T(), username, meRes.Data.Me.Username)
}

func TestMeTestSuite(t *testing.T) {
	suite.Run(t, new(MeUpTestSuite))
}
