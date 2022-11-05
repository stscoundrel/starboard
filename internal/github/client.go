package github

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const BASE_REPOSITORY_SEARCH_URL = "https://api.github.com/search/repositories"

func get(url string) ([]byte, error) {
	resp, responseErr := http.Get(url)
	if responseErr != nil {
		fmt.Println(responseErr)
		return []byte{}, responseErr
	}
	defer resp.Body.Close()
	body, parseErr := ioutil.ReadAll(resp.Body)

	if parseErr != nil {
		fmt.Println(parseErr)
		return []byte{}, parseErr
	}

	return body, nil
}

func GetStars(username string, baseApiUrl string, page int) ([]byte, error) {
	api_url := baseApiUrl + "?q=user:" + username + "+stars:>0&per_page=100&page=" + strconv.Itoa(page)
	return get(api_url)
}
