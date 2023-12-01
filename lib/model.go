package lib

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "os"
    "fmt"

)

type UrlData struct {
    Name string          `json:"name"`
    Url string           `json:"url"`
    ShortenUrl string    `json:"shortenUrl"`
}

type Model struct {
    db *gorm.DB

}

func NewModel() *Model{
    databasePath := "simple.db"

    _, err := os.Stat(databasePath)
    databaseMigrated := false

    if err == nil {
        databaseMigrated = true
    }

    database, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})

    if err != nil {
        panic(err.Error())
    }


    if databaseMigrated == false {
        fmt.Println("Migrating database")
        database.AutoMigrate(&UrlData{})
    } else {
        fmt.Println("Loading exitsting database")

    }


    return &Model{db:database}

}
