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
)

type LogInResponse struct {
	Data struct {
		Login entity.AuthToken `json:"login"`
	} `json:"data"`
}

func TestLogin(t *testing.T) {
	db := database.GetDatabase()
	repos := repository.InitRepositories(db)
	controllers := controller.InitControllers(repos)
	schema := controller.Schema(controllers)

	users := factory.CreateUsers(db, 5)
	loginUser := users[0]
	token := internal.GenerateJWT(*loginUser)

	query := fmt.Sprintf(`{ "query": "mutation { login(username: \"%s\", password: \"%s\") { token, tokenType, expiresIn } }" }`, loginUser.Username, loginUser.Password)
	byteArray := []byte(query)

	req, err := http.NewRequest("POST", "/test-graphql", bytes.NewBuffer(byteArray))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))
	rr := httptest.NewRecorder()

	handler := controller.GraphqlHandlfunc(schema)
	handler.ServeHTTP(rr, req)

	var res LogInResponse
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(rr)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Bearer", res.Data.Login.TokenType)
	assert.NotEmpty(t, res.Data.Login.Token)
	assert.NotEmpty(t, res.Data.Login.ExpiresIn)
}
