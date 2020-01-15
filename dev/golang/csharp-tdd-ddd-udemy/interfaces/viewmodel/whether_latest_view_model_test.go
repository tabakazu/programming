package viewmodel_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tabakazu/csharp-tdd-ddd-udemy/interfaces/viewmodel"
	"github.com/tabakazu/csharp-tdd-ddd-udemy/mocks"
)

func TestWhetherLatestViewModel(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	whetherMock := mocks.NewMockIWhetherRepository(mockCtrl)
	viewModel := viewmodel.NewWhetherLatestViewModel(whetherMock)

	if viewModel.AreaIdText != "" {
		t.Errorf("viewModel.AreaIdText = %v want %v", viewModel.AreaIdText, "")
	}

	if viewModel.DataDateText != "" {
		t.Errorf("viewModel.DataDateText = %v want %v", viewModel.DataDateText, "")
	}

	if viewModel.ConditionText != "" {
		t.Errorf("viewModel.ConditionText = %v want %v", viewModel.ConditionText, "")
	}

	if viewModel.ConditionText != "" {
		t.Errorf("viewModel.ConditionText = %v want %v", viewModel.ConditionText, "")
	}
}
