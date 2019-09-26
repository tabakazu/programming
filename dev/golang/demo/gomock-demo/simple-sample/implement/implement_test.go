package implement_test

import (
	"testing"

	"github.com/tabakazu/gomock-demo-simple-sample/mock"
)

func TestRepositoryImpl_Save(t *testing.T) {
	r := &mock.RepositoryMock{} // &implement.RepositoryImpl{} だと破壊的なロジックも含むため...
	s := r.Save()
	if s != "Success" {
		t.Errorf("Fail")
	}
}
