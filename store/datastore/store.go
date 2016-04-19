package datastore

import (
  "fmt"

  "github.com/yhsiang/bass/store"
  r "github.com/dancannon/gorethink"
)


type datastore struct {
  session *r.Session
}

func New() store.Store {
	return From(
    Init(),
  )
}

func From(session *r.Session) store.Store {
  return &datastore{session: session}
}

func Init() *r.Session {
  var err error
  var session *r.Session
  // FIXME: handle connection failed
  session, err = r.Connect(r.ConnectOpts{
    Address:  "localhost:28015",
    Database: "test",
  })
  if err != nil {
    fmt.Println(err)
  }
  return session
}
