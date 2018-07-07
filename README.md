# HTML code generator

This package just generate html code, it can be used as insertions in templates, or even how template engine.

### Install

```sh
go get github.com/sg3des/html
```

### Usage

```go
div := html.NewObject("div").SetID("yourid").AddClass("classname")

fmt.Println(div) // <div id="yourid" class="classname"></div>

a := html.NewA("http://somedomain.com/")
fmt.Println(a) // <a href="http://somedomain.com/"></a>

div = div.AddChilds(a)
fmt.Println(div) // <div id="yourid" class="classname"><a href="http://somedomain.com/"></a></div>


script := html.NewJavaScript().Src("/assets/script.js").String()
//<script type="application/javascript" src="/assets/script.js"></script>

linkstyle := html.NewStyleLink("/assets/style.css").String()
//<link rel="stylesheet" type="text/css" href="/assets/style.css">
```

Generate full page:

```go
page := html.NewPage("page title")
page.AddToHead(...)
page.AddToBody(...)
s := page.String()
//<!DOCTYPE html><html><head><title>page title</title>...</head><body>...</body></html>
```

Inner text of any html objects can be passet by pointer, it allow to change it on the fly.

```go
text := "text"
div := NewObject("div").SetInnerPointer(&text)
div.String() // <div>text</div>

//then change value of text variable
text = "another text"
dev.String() // <div>another text</div>
```


Work how template engine for http hander:

```go
func handler(w http.ResponseWriter, r *http.Request) {
	page := html.NewPage("title")
	div := html.NewObject("div").SetInner("Hello, World!")
	page := page.AddToBody(div)
	
	page.WriteTo(w)
}
```



It\`s can be before prepared:

```go
type Controller struct {
	Page *html.Page
}

func (c *Controller) InitPage() {
	style := html.NewStyleLink("/assets/style.css")
	script := html.NewJavaScript().Src("/assets/script.js")

	main := html.NewObject("main")
	main.AddToBody(html.NewObject("h1"))

	footer := html.NewObject("footer")

	c.Page = html.NewPage("title")
	c.Page.AddToBody(main, footer)
	s.Page.AddToHead(script, style)
	// etc
	// ...
}

func (c *Controller) render(w io.Writer, username string) {

}

func (c *Controller) handler(w http.ResponseWriter, r *http.Request) {
	c.User = r.URL.Query().Get("user")
	page.WriteTo(w)
} 
```

This approach will allow substitue variable to 