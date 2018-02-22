package main

import (
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	os.Chdir(".")

	server := gin.Default()
	server.GET(`/ping`, func(g *gin.Context) {
		s := make([]string, 1)
		s[0] = `pong`

		g.JSON(200, map[string]string{
			"result": "pong",
		})
	})

	server.Run()
}
