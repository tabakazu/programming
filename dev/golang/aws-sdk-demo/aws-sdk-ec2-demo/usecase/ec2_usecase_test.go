package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tabakazu/aws-sdk-ec2-demo/domain"
	"github.com/tabakazu/aws-sdk-ec2-demo/mocks"
	"github.com/tabakazu/aws-sdk-ec2-demo/usecase"
)

func TestGetById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	resourceMock := mocks.NewMockIEC2Resource(mockCtrl)
	u := &usecase.EC2Usecase{Resource: resourceMock}

	targetId := "123"
	target := &domain.EC2{ID: targetId}
	resourceMock.EXPECT().FindById(targetId).Return(target, nil)

	ec2, _ := u.GetById(targetId)
	if ec2.ID != targetId {
		t.Errorf("GetById(\"%v\") = %v want %v", targetId, ec2.ID, targetId)
	}
}
