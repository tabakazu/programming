package viewmodel

import "github.com/tabakazu/csharp-tdd-ddd-udemy/domain/repository"

type WhetherLatestViewModel struct {
	whether          repository.IWhetherRepository
	AreaIdText       string
	DataDateText     string
	ConditionText    string
	TemperateureText string
}

func NewWhetherLatestViewModel(whether repository.IWhetherRepository) *WhetherLatestViewModel {
	viewModel := &WhetherLatestViewModel{}
	viewModel.whether = whether
	return viewModel
}
