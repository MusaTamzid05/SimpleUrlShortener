package main


import (
    "github.com/gin-gonic/gin"
    "url_shortener/lib"
    "math/rand"
    "time"
)

func init() {
    lib.MainController.Init()
}

func main() {
    rand.Seed(time.Now().UnixNano())
    router := gin.Default()
    router.GET("/", lib.GetAllHandler)
    router.POST("/new", lib.CreateUrlHandler)
    router.GET("/search/:shorten", lib.SearchByShortenUrl)

    router.Run()


}
