package gate

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func GetResponse(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	if resp.StatusCode == http.StatusNoContent {
		return "", errors.New("Empty Response Body")
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Bad Response")
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if string(response) == "null" {
		return "", errors.New("Empty Response Body")
	}

	return string(response), nil
}
