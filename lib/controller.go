package lib

import (
    "math/rand"
)

type Controller struct {
    shortenLength int
    model *Model
}


var MainController Controller

func (c *Controller) Init() {
    c.shortenLength = 6
    c.model = NewModel()
}


func (c *Controller) Add(urlData UrlData) UrlData {
    urlData.ShortenUrl = c.GenerateShortenUrl()
    c.model.Add(urlData)
    return urlData
}

func (c *Controller) GetAll() []UrlData {
    return c.model.GetAll()
}


func (c *Controller) SearchByShortenUrl(shortenUrl string) UrlData {
    urlDataList := c.GetAll()

    for _, urlData := range urlDataList {
        if urlData.ShortenUrl == shortenUrl {
            return urlData
        }
    }

    return UrlData{}

}

func (c Controller) GenerateShortenUrl() string  {
    keys := "abcdefghijklmnopqrstuvwxyz"
    shortenUrl := ""

    for i := 0 ; i < c.shortenLength;  i += 1 {
        index := rand.Intn(len(keys))
        shortenUrl += string(keys[index])

    }


    return shortenUrl

}

