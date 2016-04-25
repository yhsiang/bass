package middleware

import (
	// "golang.org/x/net/context"
	"github.com/yhsiang/bass/store"
	"github.com/yhsiang/bass/store/datastore"

	// "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	// "github.com/ianschenck/envflag"
)

// Store is a middleware function that initializes the Datastore and attaches to
// the context of every http.Request.
func Store() gin.HandlerFunc {
	db := datastore.New()

	// logrus.Infof("using database driver %s", *database)
	// logrus.Infof("using database config %s", *datasource)

	return func(c *gin.Context) {
		store.ToContext(c, db)
		c.Next()
	}
}
