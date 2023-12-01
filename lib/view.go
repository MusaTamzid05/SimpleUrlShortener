package lib

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "fmt"
)

func CreateUrlHandler(c *gin.Context) {
    var urlData UrlData

    err := c.BindJSON(&urlData)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H {
            "error" : err.Error(),
        })
        return
    
    }

    MainController.Add(urlData)


    c.IndentedJSON(http.StatusCreated, urlData)

}

func GetAllHandler(c *gin.Context) {
    urlInfos  := MainController.GetAll()
    fmt.Println("total data ", len(urlInfos))
    c.IndentedJSON(http.StatusOK, urlInfos)

}
