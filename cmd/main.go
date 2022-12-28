package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	router.GET("/download/:file", Download)

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	Must(router.Run())
}

func Download(ctx *gin.Context) {
	filename := ctx.Param("file")
	log.Println("下载文件", filename)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+filename)
	ctx.Header("Content-Transfer-Encoding", "binary")
	file := fmt.Sprintf("file/%s", filename)
	ctx.File(file)
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
