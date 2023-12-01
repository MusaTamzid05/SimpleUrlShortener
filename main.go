package main


import (
    "github.com/gin-gonic/gin"
    "url_shortener/lib"
)

func init() {
    lib.MainController.Init()
}

func main() {
    router := gin.Default()
    router.POST("/new", lib.CreateUrlHandler)

    router.Run()


}
