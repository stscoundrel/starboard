package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stscoundrel/starboard/stars"
)

func TestStarsFromGithub(t *testing.T) {
	result, _ := stars.GetStars("stscoundrel")

	// Should always have at least test star of Starboard repo.
	assert.True(t, len(result) > 0)

	found_expected_star := false

	for _, star := range result {
		if star.Repository == "stscoundrel/starboard" {
			found_expected_star = true
		}
	}

	assert.True(t, found_expected_star)
}
