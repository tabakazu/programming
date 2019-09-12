package infrastructure

import (
	"net/http"

	"github.com/tabakazu/cli-demo/http-client-cli/domain/model"
)

type HttpClient struct{}

func (c *HttpClient) Get(url string) (*model.Site, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	site := &model.Site{
		StatusCode: resp.StatusCode,
	}
	return site, nil
}
