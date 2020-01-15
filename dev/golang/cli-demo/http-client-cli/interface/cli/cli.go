package cli

import (
	"github.com/tabakazu/cli-demo/http-client-cli/application/usecase"
	"github.com/tabakazu/cli-demo/http-client-cli/domain/model"
)

type Cli struct {
	usecase *usecase.SiteUsecase
}

func NewCli() *Cli {
	cli := &Cli{
		usecase: usecase.NewSiteUsecase(),
	}
	return cli
}

func (c *Cli) Get(url string) (*model.Site, error) {
	site, err := c.usecase.GetSite(url)
	if err != nil {
		return nil, err
	}
	return site, nil
}
