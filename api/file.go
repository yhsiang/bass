package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	filename := header.Filename
	fmt.Println(header.Filename)
	out, err := os.Create("./tmp/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusOK, "Done")
}
