package entity_test

import (
	"testing"

	"github.com/tabakazu/domain-modeling-demo/domain/model/entity"
	"github.com/tabakazu/domain-modeling-demo/domain/model/valueobject"
)

func Test_NewPost(t *testing.T) {
	name := valueobject.PostName{Value: "test post"}
	post := entity.NewPost(name)

	if post == nil {
		t.Errorf("post = %v want %v", post, nil)
	}
	if post.Name.Value != name.Value {
		t.Errorf("post.Name.Value = %v want %v", post.Name.Value, name.Value)
	}
}
