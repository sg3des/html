package html

import (
	"bytes"
	"testing"
	"time"

	"github.com/sg3des/html/css"
)

func TestObjectDiv(t *testing.T) {
	should := "<div></div>"

	div := NewObject("div").String()
	compare(t, should, div)

	div = NewObject("div").SetID("someid").String()
	compare(t, `<div id='someid'></div>`, div)

	div = NewObject("div").AddClass("some-class").String()
	compare(t, `<div class='some-class'></div>`, div)
}

func TestAhref(t *testing.T) {
	a := NewA("/some/page").String()
	compare(t, a, `<a href='/some/page'></a>`)

	a = NewA("/some/page").Download("filename.ext").String()
	compare(t, a, `<a href='/some/page' download='filename.ext'></a>`)
}

func TestDivWithScript(t *testing.T) {
	div := NewObject("div").AddAttribute("onclick", `console.log(this, "onclick")`).String()
	compare(t, div, `<div onclick='console.log(this, "onclick")'></div>`)
	t.Log(div)
}

func TestChilds(t *testing.T) {
	div := NewObject("div")
	p := NewObject("p").AddInnerText("some paragraph text")

	s := div.AddChilds(p).String()
	compare(t, s, "<div><p>some paragraph text</p></div>")
	t.Log(s)
}

func TestScript(t *testing.T) {
	should := `<script type='application/javascript' src='/assets/index.js'></script>`

	script := NewScript("application/javascript").Src("/assets/index.js").String()
	compare(t, should, script)
	t.Log(script)

	script = NewJavaScript().Src("/assets/index.js").String()
	compare(t, should, script)
}

func TestLink(t *testing.T) {
	should := `<link rel='stylesheet' type='text/css' href='style.css'>`

	link := NewLink("stylesheet", "text/css", "style.css").String()
	compare(t, should, link)
	t.Log(link)

	link = NewStyleLink("style.css").String()
	compare(t, should, link)
	t.Log(link)
}

func TestStyle(t *testing.T) {
	styles := []css.Style{
		css.NewStyle("*").Position("relative"),
	}
	s := NewStyle(styles).String()
	compare(t, s, "<style>*{position: relative;}</style>")
}

//
// PAGE
//

func TestPage(t *testing.T) {
	page := NewPage("page title")
	s := page.String()
	compare(t, s, `<!DOCTYPE html><html><head><title>page title</title></head><body></body></html>`)

	s = page.AddToBody(NewObject("div")).String()
	compare(t, s, `<!DOCTYPE html><html><head><title>page title</title></head><body><div></div></body></html>`)
	t.Log(s)

	var w bytes.Buffer
	n, err := page.WriteTo(&w)
	if err != nil {
		t.Error(err)
	}
	if n == 0 {
		t.Error("length should be more than 0")
	}

	t.Log(w.String())
}

func compare(t *testing.T, a, b string) {
	if a != b {
		t.Errorf("failed, not equal:\n\t%s\n\t%s", a, b)
	}
}

//
// Writer
//

func TestWriter(t *testing.T) {
	var w bytes.Buffer
	div := NewObject("div")
	n, err := div.WriteTo(&w)
	if err != nil {
		t.Error(err)
	}
	if n == 0 {
		t.Error("length should be more than 0")
	}

	t.Log(w.String())
}

// Benchmark

func BenchmarkPage(b *testing.B) {
	b.StopTimer()
	page := NewPage("title")
	page = page.AddToHead(
		NewObject("style").AddInnerText("#title{color: red;}"),
	)
	page = page.AddToBody(
		NewObject("h1").SetID("title").AddInnerText("example"),
	)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		page.AddToBody(
			NewObject("div").AddInnerText("time: " + time.Now().String()),
		)
	}
}

func BenchmarkBuffer(b *testing.B) {
	b.StopTimer()
	div := NewObject("div").SetID("id").AddClass("class-name").AddInnerText("text")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		div.buffer(nil)
	}
}
