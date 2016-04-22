package datastore

import (
  "testing"

  "github.com/yhsiang/bass/models"

  "github.com/franela/goblin"
  r "github.com/dancannon/gorethink"
)

func TestUsers(t *testing.T) {
    session := openTest()
    store := From(session)

    g := goblin.Goblin(t)
    g.Describe("User", func() {
      g.BeforeEach(func() {
        // FIXME: prepare data here
      })

      g.It("Should Add a new User", func() {
        user := models.User{
          Username: "hana",
          Email: "hanahsu@dopemusic.com",
          Password: "hana1234",
        }
        // FIXME: check return id is string
        _, err := store.CreateUser(&user)
        g.Assert(err == nil).IsTrue()
        // g.Assert().IsTrue()
      })
      g.AfterEach(func() {
        r.Table("users").Delete().Run(session)
      })
    })


}
