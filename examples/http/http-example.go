package main

import (
	"net/http"
	"time"

	"github.com/sg3des/html"
)

func main() {
	c := NewController()

	http.HandleFunc("/", c.index)
	http.ListenAndServe(":8080", nil)
}

type Controller struct {
	page html.Page
}

func NewController() *Controller {
	page := html.NewPage("http-example")

	page = page.AddToHead(
		html.NewObject("style").SetInner("#title{color: red;}"),
	)

	page = page.AddToBody(
		html.NewObject("h1").SetID("title").SetInner("example"),
	)

	return &Controller{
		page: page,
	}
}

func (c *Controller) index(w http.ResponseWriter, r *http.Request) {
	page := c.page.AddToBody(
		html.NewObject("div").SetInner("time: "+time.Now().String()),
		html.NewObject("div").SetInner("addr: "+r.RemoteAddr),
	)

	page.WriteTo(w)
}
