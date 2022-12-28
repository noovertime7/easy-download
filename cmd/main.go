package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var port int

func init() {
	flag.IntVar(&port, "port", 40000, "port")
}

func main() {
	flag.Parse()
	router := gin.Default()

	router.GET("/download/:file", Download)

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	Must(router.Run(fmt.Sprintf(":%d", port)))
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
