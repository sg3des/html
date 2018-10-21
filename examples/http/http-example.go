package main

import (
	"net/http"

	"github.com/sg3des/html"
	"github.com/sg3des/html/css"
)

func main() {
	c := NewController()

	http.HandleFunc("/", c.index)
	http.ListenAndServe(":8080", nil)
}

type Controller struct {
	users []string

	page html.Page

	header html.Object
	title  html.Object
	main   html.Object
	table  html.Table
}

func NewController() *Controller {
	page := html.NewPage("http-example")

	styles := []css.Style{
		css.NewStyle("*").Position("relative").Margin("0").FontFamily("sans-serif"),
		css.NewStyle("#title").Color("red"),
		css.NewStyle("body").Left("50%").Width("300px").Margin("50px 0 0 -150px"),
		css.NewStyle("form").Width("100%"),
	}

	page = page.AddToHead(
		html.NewStyle(styles),
	)

	floatright := css.NewStyle("").Left("200px")

	main := html.NewObject("main").AddChilds(
		html.NewObject("form").SetAttribute("method", http.MethodPost).AddChilds(
			html.NewObject("p").AddChilds(
				html.NewObject("label").AddInnerText("name:"),
				html.NewInput("text", "user").SetPlaceholder("user name"),
			),
			html.NewObject("p").AddChilds(
				html.NewInput("submit", "submit").SetValue("submit").SetStyle(floatright),
			),
		),
		html.NewObject("h2").AddInnerText("History"),
	)

	return &Controller{
		page:   page,
		header: html.NewObject("header"),
		title:  html.NewObject("h1").SetID("title"),
		main:   main,
		table:  html.NewTable(),
	}
}

func (c *Controller) index(w http.ResponseWriter, r *http.Request) {
	page := c.page

	if r.Method == http.MethodPost {
		r.ParseForm()
		user := r.PostForm.Get("user")
		c.table = c.table.AddRow(html.Text(user))

		page = page.AddToBody(
			c.header.AddChilds(
				c.title.AddInnerText("Hello, " + user),
			),
		)

		c.users = append(c.users, user)
	} else {
		page = c.page.AddToBody(
			c.header.AddChilds(
				c.title.AddInnerText("Example Title"),
			),
		)
	}

	page = page.AddToBody(
		c.main.AddChilds(c.table),
	)

	page.WriteTo(w)
}
