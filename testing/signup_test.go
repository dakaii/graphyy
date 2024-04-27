package testing

import (
	"bytes"
	"encoding/json"
	"graphyy/controller"
	"graphyy/database"
	"graphyy/entity"
	"graphyy/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SignUpResponse struct {
	Data struct {
		Signup entity.AuthToken `json:"signup"`
	} `json:"data"`
}

func TestCreateUser(t *testing.T) {
	db := database.GetDatabase()
	repos := repository.InitRepositories(db)
	controllers := controller.InitControllers(repos)
	schema := controller.Schema(controllers)
	jsonStr := []byte(`{
        "query": "mutation { signup(username: \"secondtestuser\", password: \"testpass\") { token, tokenType, expiresIn } }"
    }`)

	req, err := http.NewRequest("POST", "/test-graphql", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := controller.GraphqlHandlfunc(schema)
	handler.ServeHTTP(rr, req)

	var res SignUpResponse
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Bearer", res.Data.Signup.TokenType)
	assert.NotEmpty(t, res.Data.Signup.Token)
	assert.NotEmpty(t, res.Data.Signup.ExpiresIn)
}
