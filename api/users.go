package api

import (
  "github.com/gin-gonic/gin"

  "github.com/yhsiang/bass/store"
)

func GetUsers(c *gin.Context) {
  users, err := store.GetUserList(c)
  if err != nil {
    c.String(500, "Error getting user list. %s", err)
  } else {
    c.JSON(200, users)
  }
}
