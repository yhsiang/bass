package datastore

import (
  "fmt"

  "github.com/yhsiang/bass/models"
  r "github.com/dancannon/gorethink"
)

func (db *datastore) GetUser(id string) (*models.User, error) {
  cursor, err := r.Table("users").Get(id).Run(db.session)

  var user = new(models.User)
  cursor.One(&user)
  cursor.Close()
  return user, err
}

func (db *datastore) GetUserList() ([]*models.User, error) {
  rows, err := r.Table("users").Run(db.session)

  if err != nil {
    fmt.Println(err)
  }

  var users = []*models.User{}
  err2 := rows.All(&users)
  return users, err2
}

func (db *datastore) CreateUser(user *models.User) (string, error) {
  res, err := r.Table("users").Insert(user).RunWrite(db.session)
  return res.GeneratedKeys[0], err
}
