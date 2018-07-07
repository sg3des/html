# HTML code generator

This package just generate html code, it can be used as insertions in templates, or even how template engine.

### Install

```sh
go get github.com/sg3des/html
```

### Usage

```go
div := NewObject("div").SetID("yourid").AddClass("classname")

fmt.Println(div) // <div id="yourid" class="classname"></div>

a := NewA("http://somedomain.com/")
fmt.Println(a) // <a href="http://somedomain.com/"></a>

div = div.AddChilds(a)
fmt.Println(div) // <div id="yourid" class="classname"><a href="http://somedomain.com/"></a></div>


script := NewJavaScript().Src("/assets/script.js").String()
//<script type="application/javascript" src="/assets/script.js"></script>

linkstyle := NewStyleLink("/assets/style.css").String()
//<link rel="stylesheet" type="text/css" href="/assets/style.css">
```

Generate full page

```go
page := NewPage("page title")
page.AddToHead(...)
page.AddToBody(...)
s := page.String()
//<!DOCTYPE html><html><head><title>page title</title>...</head><body>...</body></html>
```