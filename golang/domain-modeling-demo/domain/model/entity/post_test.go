package entity_test

import (
	"testing"

	"github.com/tabakazu/domain-modeling-demo/domain/model/entity"
	"github.com/tabakazu/domain-modeling-demo/domain/model/valueobject"
)

var (
	name  valueobject.PostName   = valueobject.PostName{Value: "test post"}
	cusID valueobject.CustomerID = valueobject.CustomerID{Value: 1}
)

func Test_NewPost(t *testing.T) {
	post := entity.NewPost(cusID, name)

	if post == nil {
		t.Errorf("post = %v want %v", post, nil)
	}
	if post.CustomerID.Value != cusID.Value {
		t.Errorf("post.CustomerID.Value = %v want %v", post.CustomerID.Value, cusID.Value)
	}
	if post.Name.Value != name.Value {
		t.Errorf("post.Name.Value = %v want %v", post.Name.Value, name.Value)
	}
}
