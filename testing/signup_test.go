package testing

import (
	"bytes"
	"graphyy/controller"
	"graphyy/database"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCreateUser(t *testing.T) {
	collectionName, exists := os.LookupEnv("MONGODB_COLLECTION_NAME")
	if !exists {
		collectionName = "testingCollection"
	}

	ctx, db := database.GetDatabase(collectionName)
	userRepo := database.NewUserRepo(db, ctx, db.Collection(collectionName))
	h := controller.NewBaseHandler(userRepo)

	jsonStr := []byte(`{"username":"secondtestuser","password":"testpass"}`)

	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(h.Signup)
	// handler.ServeHTTP(rr, req)
	// if status := rr.Code; status != http.StatusOK {
	// 	t.Errorf("handler returned wrong status code: got %v want %v",
	// 		status, http.StatusOK)
	// }
	// res := model.AuthToken{}
	// json.Unmarshal([]byte(rr.Body.String()), &res)

	// expected := `Bearer`
	// if res.TokenType != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }
}
