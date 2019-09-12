package usecase

import (
	"github.com/tabakazu/cli-demo/http-client-cli/domain/model"
	"github.com/tabakazu/cli-demo/http-client-cli/infrastructure"
)

type SiteUsecase struct {
	httpClient *infrastructure.HttpClient
}

func NewSiteUsecase() *SiteUsecase {
	usecase := &SiteUsecase{
		httpClient: &infrastructure.HttpClient{},
	}
	return usecase
}

func (u *SiteUsecase) GetSite(url string) (*model.Site, error) {
	site, err := u.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	return site, nil
}
