package main

import (
	"fmt"
	// "log"
	"os"

	"github.com/urfave/cli"
	// "github.com/asaskevich/govalidator"
	r "github.com/dancannon/gorethink"
	"golang.org/x/crypto/bcrypt"

	"github.com/yhsiang/bass/model"
	// "github.com/yhsiang/bass/store"
)

func Init() *r.Session {
	// FIXME: handle connection failed
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "test",
	})
	if err != nil {
		fmt.Println(err)
	}
	return session
}

var userCmd = cli.Command{
	Name:  "user",
	Usage: "manage users",
	Subcommands: []cli.Command{
		// userListCmd,
		// userInfoCmd,
		userAddCmd,
		// userRemoveCmd,
	},
}

var userAddCmd = cli.Command{
	Name:   "add",
	Usage:  "adds a user",
	Action: userAdd,
}

func userAdd(c *cli.Context) error {
	username := c.Args().First()
	password := c.Args()[1]
	user := &model.User{
		Username: username,
	}
	session := Init()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	resp, err := r.Table("users").Insert(user).RunWrite(session)
	if err != nil {
		return err
	}
	fmt.Println(resp.GeneratedKeys[0])
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "bass"
	app.Usage = "command line utility"
	app.Commands = []cli.Command{
		userCmd,
	}
	app.Run(os.Args)
}
