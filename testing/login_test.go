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

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type LogInResponse struct {
	Data struct {
		Login entity.AuthToken `json:"login"`
	} `json:"data"`
}

var _ = Describe("Login", func() {
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

	It("should login a user", func() {
		loginUser := users[0]
		token := internal.GenerateJWT(loginUser)

		query := fmt.Sprintf(`{ "query": "mutation { login(username: \"%s\", password: \"%s\") { token, tokenType, expiresIn } }" }`, loginUser.Username, loginUser.Password)
		byteArray := []byte(query)

		req, err := http.NewRequest("POST", "/test-graphql", bytes.NewBuffer(byteArray))
		Expect(err).NotTo(HaveOccurred())
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))

		handler.ServeHTTP(rr, req)

		var res LogInResponse
		err = json.Unmarshal(rr.Body.Bytes(), &res)
		Expect(err).NotTo(HaveOccurred())

		Expect(rr.Code).To(Equal(http.StatusOK))
		Expect(res.Data.Login.TokenType).To(Equal("Bearer"))
		Expect(res.Data.Login.Token).To(BeAssignableToTypeOf(""))
		Expect(res.Data.Login.ExpiresIn).To(BeAssignableToTypeOf(int64(0)))
	})
})
