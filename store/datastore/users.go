package datastore

import (
  "fmt"
  "golang.org/x/crypto/bcrypt"

  "github.com/yhsiang/bass/models"
  r "github.com/dancannon/gorethink"
)

func (db *datastore) GetUser(id string) (*models.User, error) {
  cursor, err := r.Table(userTable).Get(id).Run(db.session)

  var user = new(models.User)
  cursor.One(&user)
  cursor.Close()
  return user, err
}

func (db *datastore) GetUserList() ([]*models.User, error) {
  rows, err := r.Table(userTable).Run(db.session)

  if err != nil {
    fmt.Println(err)
  }

  var users = []*models.User{}
  err2 := rows.All(&users)
  return users, err2
}

func (db *datastore) CreateUser(user *models.User) (string, error) {
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
  if err != nil {
      return "", err
  }

  user.Password = string(hashedPassword)
  resp, err := r.Table(userTable).Insert(user).RunWrite(db.session)
  return resp.GeneratedKeys[0], err
}

func (db *datastore) UpdateUser(id string, user map[string]interface{}) (int, error) {
  resp, err := r.Table(userTable).Get(id).Update(user).RunWrite(db.session)
  return resp.Replaced, err
}

const userTable = "users"
