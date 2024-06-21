package main

import (
    "github.com/gin-gonic/gin"
    "myapp/routers"
    "myapp/config"
)

func main() {
    config.Init()

    r := gin.Default()
    routers.SetupRoutes(r)

    r.Static("/static", "./static")
    r.GET("/", func(c *gin.Context) {
        c.File("./static/index.html")
    })

    r.Run(":8080")
}
