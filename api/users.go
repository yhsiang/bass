package api

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"github.com/yhsiang/bass/model"
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
	user := &model.User{}
	err := c.Bind(user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	_, err = govalidator.ValidateStruct(user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	id, err := store.CreateUser(c, user)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
}
