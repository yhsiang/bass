package main

import (
	"net/http"
	// "time"

	"github.com/yhsiang/bass/router"
	"github.com/yhsiang/bass/router/middleware"

	// "github.com/Sirupsen/logrus"
	// "github.com/gin-gonic/contrib/ginrus"
	"github.com/ianschenck/envflag"
	// _ "github.com/joho/godotenv/autoload"
)

var (
	addr = envflag.String("SERVER_ADDR", ":8000", "")
	cert = envflag.String("SERVER_CERT", "", "")
	key  = envflag.String("SERVER_KEY", "", "")

	debug = envflag.Bool("DEBUG", false, "")
)

func main() {
	handler := router.Load(
		middleware.Store(),
	)

	http.ListenAndServe(*addr, handler)
}
