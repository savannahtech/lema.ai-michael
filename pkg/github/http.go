package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (gh *GHClient) listCommits(owner, repo string, since string, perPage int, page int, expectedResponse interface{}) error {
	var endPointURL string
	endPointURL = fmt.Sprintf("repos/%s/%s/commits?since=%s&per_page=%d&page=%d", owner, repo, since, perPage, page)

	return gh.get(endPointURL, expectedResponse)
}

func (gh *GHClient) getRepo(owner, repo string, expectedResponse interface{}) error {
	endPointURl := fmt.Sprintf("repos/%s/%s", owner, repo)
	return gh.get(endPointURl, expectedResponse)
}
func (gh *GHClient) getRepos(expectedResponse interface{}) error {
	endPointURl := "repositories"
	return gh.get(endPointURl, expectedResponse)
}

func (gh *GHClient) post(endPointURL string, reqBody, expectedResponse interface{}) error {
	return gh.sendRequest(http.MethodPost, endPointURL, reqBody, expectedResponse)
}
func (gh *GHClient) get(endPointURL string, expectedResponse interface{}) error {
	return gh.sendRequest(http.MethodGet, endPointURL, nil, expectedResponse)
}
func (gh *GHClient) generateRequest(method, endpoint string, body interface{}) (*http.Request, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyByte, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		bodyReader = bytes.NewReader(bodyByte)
	}
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", gh.BaseURL, endpoint), bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", gh.token))

	return req, nil
}
func (gh *GHClient) sendRequest(method, endpoint string, body, output interface{}) error {
	req, err := gh.generateRequest(method, endpoint, body)
	if err != nil {
		return err
	}

	res, err := gh.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request %+v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}(res.Body)

	if err := json.NewDecoder(res.Body).Decode(output); err != nil {
		return fmt.Errorf("error marshalling client response: %s", err)
	}
	if res.StatusCode >= http.StatusBadRequest {

		return fmt.Errorf("client response with status code: %v message: %v", res.StatusCode, output)
	}

	return nil
}
