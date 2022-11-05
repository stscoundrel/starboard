package stars

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseStarsFromJson(t *testing.T) {
	fixtureResponse, _ := os.Open("../fixtures/sample_response.json")
	fixtureBody, _ := ioutil.ReadAll(fixtureResponse)
	defer fixtureResponse.Close()

	result := starsFromJson(fixtureBody)

	assert.Equal(t, len(result), 27)

	assert.Equal(t, "stscoundrel/struct", result[0].Repository)
	assert.Equal(t, 6, result[0].Count)
	assert.Equal(t, "https://github.com/stscoundrel/struct", result[0].Link)
	assert.Equal(t, "https://api.github.com/repos/stscoundrel/struct/stargazers", result[0].StarGazersLink)

	assert.Equal(t, "stscoundrel/cleasby-vigfusson-dictionary", result[4].Repository)
	assert.Equal(t, 2, result[4].Count)
	assert.Equal(t, "https://github.com/stscoundrel/cleasby-vigfusson-dictionary", result[4].Link)
	assert.Equal(t, "https://api.github.com/repos/stscoundrel/cleasby-vigfusson-dictionary/stargazers", result[4].StarGazersLink)
}
