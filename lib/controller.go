package lib

type Controller struct {
    shortenLength int
    model *Model
}


var MainController Controller

func (c *Controller) Init() {
    c.shortenLength = 6
    c.model = NewModel()
}


func (c *Controller) Add(urlData UrlData) {
    urlData.ShortenUrl = c.GenerateShortenUrl()
    c.model.Add(urlData)
}

func (c *Controller) GetAll() []UrlData {
    return c.model.GetAll()
}


func (c *Controller) SearchByShortenUrl(shortenUrl string) UrlData {
    return UrlData{}
}

func (c Controller) GenerateShortenUrl() string  {
    return ""

}
