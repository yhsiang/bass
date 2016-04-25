package router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/yhsiang/bass/api"
)

func Load(middleware ...gin.HandlerFunc) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())

	e.Use(middleware...)
	users := e.Group("/api/users")
	{
		users.GET("", api.GetUsers)
		users.POST("", api.PostUser)
	}
	files := e.Group("/api/files")
	{
		files.PUT("", api.UploadFile)
	}

	return normalize(e)
}

// normalize is a helper function to work around the following
// issue with gin. https://github.com/gin-gonic/gin/issues/388
func normalize(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		parts := strings.Split(r.URL.Path, "/")[1:]
		switch parts[0] {
		case "settings", "bots", "repos", "api", "login", "logout", "", "authorize", "hook", "static", "gitlab":
			// no-op
		default:

			if len(parts) > 2 && parts[2] != "settings" {
				parts = append(parts[:2], append([]string{"builds"}, parts[2:]...)...)
			}

			// prefix the URL with /repo so that it
			// can be effectively routed.
			parts = append([]string{"", "repos"}, parts...)

			// reconstruct the path
			r.URL.Path = strings.Join(parts, "/")
		}

		h.ServeHTTP(w, r)
	})
}
