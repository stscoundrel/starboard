package github

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIssues(t *testing.T) {
	fixtureResponse, _ := os.Open("../../fixtures/sample_response.json")
	fixtureBody, _ := ioutil.ReadAll(fixtureResponse)
	defer fixtureResponse.Close()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		queryParams := request.URL.Query()

		assert.Equal(t, "user:stscoundrel stars:>0", queryParams["q"][0])
		assert.Equal(t, "100", queryParams["per_page"][0])
		assert.Equal(t, "666", queryParams["page"][0])

		w.WriteHeader(http.StatusOK)
		w.Write(fixtureBody)
	}))
	defer server.Close()

	result, _ := GetStars("stscoundrel", server.URL, 666)
	assert.Equal(t, result, fixtureBody)
}
