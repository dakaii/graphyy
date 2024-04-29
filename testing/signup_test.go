package testing

import (
	"bytes"
	"encoding/json"
	"graphyy/controller"
	"graphyy/database"
	"graphyy/entity"
	"graphyy/internal"
	"graphyy/repository"
	"graphyy/testing/factory"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type SignUpResponse struct {
	Data struct {
		Signup entity.AuthToken `json:"signup"`
	} `json:"data"`
}

var _ = Describe("CreateUser", func() {
	var (
		users   []entity.User
		rr      *httptest.ResponseRecorder
		handler http.Handler
	)

	BeforeEach(func() {
		db := database.GetDatabase()
		repos := repository.InitRepositories(db)
		controllers := controller.InitControllers(repos)
		schema := controller.Schema(controllers)
		handler = controller.GraphqlHandlfunc(schema)

		users = factory.CreateUsers(db, 5)
		rr = httptest.NewRecorder()
	})

	AfterEach(func() {
		truncateAllTables()
	})

	It("should create a user", func() {
		loginUser := users[0]
		token := internal.GenerateJWT(loginUser)

		query := `{
            "query": "mutation { signup(username: \"secondtestuser\", password: \"testpass\") { token, tokenType, expiresIn } }"
        }`
		byteArray := []byte(query)

		req, err := http.NewRequest("POST", "/test-graphql", bytes.NewBuffer(byteArray))
		Expect(err).NotTo(HaveOccurred())
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token.Token)

		handler.ServeHTTP(rr, req)

		var res SignUpResponse
		err = json.Unmarshal(rr.Body.Bytes(), &res)
		Expect(err).NotTo(HaveOccurred())

		Expect(rr.Code).To(Equal(http.StatusOK))
		Expect(res.Data.Signup.TokenType).To(Equal("Bearer"))
		Expect(res.Data.Signup.Token).To(BeAssignableToTypeOf(""))
		Expect(res.Data.Signup.ExpiresIn).To(BeAssignableToTypeOf(int64(0)))
	})
})
