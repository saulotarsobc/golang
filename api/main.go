package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, world!"})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "test"})
	})

	r.GET("/image", func(c *gin.Context) {
		c.File("images/image.png")
	})

	r.GET("/table", func(c *gin.Context) {
		dat, err := os.Open("files/table.csv")
		if err != nil {
			c.String(500, "Erro ao ler o arquivo")
			return
		}
		defer dat.Close()

		c.Render(200, render.Reader{
			Headers: map[string]string{
				"Content-Disposition": `attachment; filename="table.csv"`,
			},
			ContentType: "text/csv; charset=utf-8",
			ContentLength: func() int64 {
				info, err := dat.Stat()
				if err != nil {
					return 0
				}
				return info.Size()
			}(),
			Reader: dat,
		})
	})

	r.Run(":3000")
}
