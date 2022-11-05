package stars

import (
	"github.com/stscoundrel/starboard/internal/github"
)

func GetStars(username string) ([]Star, error) {
	// TODO: Paginate responses, if more than 100 starred repos.
	rawStars, error := github.GetStars(username, github.BASE_REPOSITORY_SEARCH_URL, 1)

	if error != nil {
		return []Star{}, error
	}

	return starsFromJson(rawStars), nil
}
