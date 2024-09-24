package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/letsmakecakes/github-activity/internal/model"
)

type Request struct {
	URL string
}

type Response struct {
	Event []model.Event
}

func NewRequest(url string) *Request {
	return &Request{URL: url}
}

func NewResponse() *Response {
	return &Response{}
}

func (r Request) SendClientRequest() (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, r.URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	return resp, nil
}

func (respObj *Response) GetEvents(resp *http.Response) error {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &respObj.Event); err != nil {
		return err
	}

	return nil
}
