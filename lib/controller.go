package lib

type Controller struct {
    model *Model
}


var MainController Controller

func (c *Controller) Init() {
    c.model = NewModel()
}


func (c *Controller) Add(model Model) {

}


func (c *Controller) SearchByShortenUrl(shortenUrl string) UrlData {
    return UrlData{}

}
