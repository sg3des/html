# HTML code generator

This package just generate html and css code, it can be used as insertions in templates, or even how template engine.

### Install

```sh
go get github.com/sg3des/html
```

### Usage

#### HTML

Examples:

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


Generate page:

```go
page := html.NewPage("page title")
page.AddToHead(...)
page.AddToBody(...)
s := page.String()
//<!DOCTYPE html><html><head><title>page title</title>...</head><body>...</body></html>
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

func (c *Controller) handler(w http.ResponseWriter, r *http.Request) {
	page := c.page.AddToBody(
		html.NewObject("div").SetInner("time: "+time.Now().String()),
		html.NewObject("div").SetInner("addr: "+r.RemoteAddr),
	)

	page.WriteTo(w)
} 
```


#### CSS

Genereate CSS style:

```go
style := NewStyle("#someid", ".some-class").AddProperty("color: red","position: relative")
style.String() // #someid, .some-class{color: red; position: relative;}
```