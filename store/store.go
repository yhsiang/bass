package store

import (
	// "io"
	"github.com/yhsiang/bass/model"
	"golang.org/x/net/context"
)

type Store interface {
	GetUser(string) (*model.User, error)
	GetUserList() ([]*model.User, error)
	CreateUser(*model.User) (string, error)
	UpdateUser(string, map[string]interface{}) (int, error)
}

func GetUser(c context.Context, id string) (*model.User, error) {
	return FromContext(c).GetUser(id)
}

func GetUserList(c context.Context) ([]*model.User, error) {
	return FromContext(c).GetUserList()
}

func CreateUser(c context.Context, user *model.User) (string, error) {
	return FromContext(c).CreateUser(user)
}

func UpdateUser(c context.Context, id string, user map[string]interface{}) (int, error) {
	return FromContext(c).UpdateUser(id, user)
}
