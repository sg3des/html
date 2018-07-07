package html

import "testing"

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
	p := NewObject("p")
	p.Inner = "some paragraph text"

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
}

func compare(t *testing.T, a, b string) {
	if a != b {
		t.Errorf("failed, not equal:\n\t%s\n\t%s", a, b)
	}
}
