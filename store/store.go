package store

import (
  // "io"
  "github.com/yhsiang/bass/models"
  "golang.org/x/net/context"
)


type Store interface {
  GetUser(string) (*models.User, error)
  GetUserList() ([]*models.User, error)
}

func GetUser(c context.Context, id string) (*models.User, error) {
  return FromContext(c).GetUser(id)
}

func GetUserList(c context.Context) ([]*models.User, error) {
  return FromContext(c).GetUserList()
}
