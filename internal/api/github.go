package api

import (
	"encoding/json"
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

func (req Request) SendClientRequest() (*http.Response, error) {
	resp, err := http.Get(req.URL)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (respObj *Response) GetEvents(resp *http.Response) error {
	err := json.NewDecoder(resp.Body).Decode(&respObj.Event)
	if err != nil {
		return err
	}

	return nil
}
