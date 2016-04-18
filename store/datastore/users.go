package datastore

import (
  "fmt"

  "github.com/yhsiang/bass/models"
  // "github.com/russross/meddler"
  r "github.com/dancannon/gorethink"
)

func (db *datastore) GetUser(id string) (*models.User, error) {
  cursor, err := r.Table("users").Get(id).Run(db.session)

  var user = new(models.User)
  cursor.One(&user)
  cursor.Close()
  // var usr = new(model.User)
  // var err = meddler.Load(db, userTable, usr, id)
  return user, err
}

func (db *datastore) GetUserList() ([]*models.User, error) {
  rows, err := r.Table("users").Run(db.session)

  if err != nil {
    fmt.Println(err)
    // return
  }

  var users = []*models.User{}
  err2 := rows.All(&users)
  // if err2 != nil {
  //   fmt.Println(err2)
  //   return
  // }
  return users, err2
}
