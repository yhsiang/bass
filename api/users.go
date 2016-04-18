package api

import (
  "net/http"

  "github.com/gin-gonic/gin"

  "github.com/yhsiang/bass/models"
  "github.com/yhsiang/bass/store"
)

func GetUsers(c *gin.Context) {
  users, err := store.GetUserList(c)
  if err != nil {
    c.String(500, "Error getting user list. %s", err)
  } else {
    c.JSON(http.StatusOK, users)
  }
}

func PostUser(c *gin.Context) {
  in := &models.User{}
  err := c.Bind(in)
  if err != nil {
    c.String(http.StatusBadRequest, err.Error())
    return
  }

  user := &models.User{}
  user.Username = in.Username
  user.Email = in.Email
  user.Avatar = in.Avatar
  user.Active = true

  id, err := store.CreateUser(c, user)
  if err != nil {
    c.String(http.StatusInternalServerError, err.Error())
    return
  }

  c.JSON(http.StatusOK, id)
}
