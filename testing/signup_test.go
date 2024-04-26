package testing

import (
	"bytes"
	"fmt"
	"graphyy/controller"
	"graphyy/database"
	"graphyy/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db := database.GetDatabase()
	repos := repository.InitRepositories(db)
	controllers := controller.InitControllers(repos)
	schema := controller.Schema(controllers)
	jsonStr := []byte(`{
        "query": "mutation { signup(username: \"secondtestuser\", password: \"testpass\") { token, tokenType } }"
    }`)

	req, err := http.NewRequest("POST", "/test-graphql", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := controller.GraphqlHandlfunc(schema)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	fmt.Println(rr.Body.String())
	fmt.Println("sadf")
	// assert.Equal(t, `Bearer`, rr.Body.)

	// res := model.AuthToken{}
	// json.Unmarshal([]byte(rr.Body.String()), &res)

	// expected := `Bearer`
	// if res.TokenType != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }
}
