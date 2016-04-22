package datastore

import (
  "testing"
  "reflect"

  "github.com/yhsiang/bass/models"

  "github.com/franela/goblin"
  r "github.com/dancannon/gorethink"
)

func TestUsers(t *testing.T) {
  session := openTest()
  store := From(session)

  g := goblin.Goblin(t)
  g.Describe("User", func() {
    var user models.User

    g.It("Should Add a new User", func() {
      user.Username = "hana"
      user.Email    = "hanahsu@dopemusic.com"
      user.Password = "hana1234"

      // FIXME: check return id is string
      id, err := store.CreateUser(&user)
      user.ID = id
      g.Assert(err == nil).IsTrue()
      g.Assert(reflect.TypeOf(id).String() == "string").IsTrue()
    })

    g.It("Should Update a User", func() {
      var u = map[string]interface{}{
        "Username": "hanahsu",
      }
      resp, err := store.UpdateUser(user.ID, u)
      g.Assert(err == nil).IsTrue()
      g.Assert(resp).Equal(1)
    })

    g.After(func() {
      r.Table("users").Delete().Run(session)
    })
  })
}
