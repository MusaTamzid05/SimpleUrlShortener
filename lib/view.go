package lib

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func CreateUrlHandler(c *gin.Context) {
    var model Model

    err := c.BindJSON(&model)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H {
            "error" : err.Error(),
        })
        return
    
    }

    c.IndentedJSON(http.StatusCreated, model)

}
