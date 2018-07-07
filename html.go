package html

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func name(w http.ResponseWriter, r *http.Request) {
	r.URL.Query().Get("user")
}

type Object struct {
	TagName  string
	noEndTag bool

	ID         string
	Class      []string
	Attributes []Attribute
	Inner      string
	Childs     []Object
}

func (o Object) String() string {
	return o.buffer(nil).String()
}

func (o Object) buffer(p []byte) *bytes.Buffer {
	w := bytes.NewBuffer(p)

	w.WriteByte('<')
	w.WriteString(o.TagName)

	if o.ID != "" {
		fmt.Fprintf(w, ` id='%s'`, o.ID)
	}

	if len(o.Class) > 0 {
		fmt.Fprintf(w, ` class='%s'`, strings.Join(o.Class, " "))
	}

	for _, attr := range o.Attributes {
		fmt.Fprintf(w, ` %s='%s'`, attr.Key, attr.Val)
	}

	w.WriteByte('>')

	if o.noEndTag {
		return w
	}

	w.WriteString(o.Inner)

	for _, child := range o.Childs {
		w.WriteString(child.String())
	}

	fmt.Fprintf(w, `</%s>`, o.TagName)
	return w
}

func (o Object) WriteTo(w io.Writer) (n int64, err error) {
	return o.buffer(nil).WriteTo(w)
}

func (o Object) SetID(id string) Object {
	o.ID = id
	return o
}

func (o Object) AddClass(class ...string) Object {
	o.Class = append(o.Class, class...)
	return o
}

func (o Object) AddAttribute(key, val string) Object {
	o.Attributes = append(o.Attributes, Attribute{key, val})
	return o
}

func (o Object) SetInner(s string) Object {
	o.Inner = s
	return o
}

func (o Object) AddChilds(child ...Object) Object {
	o.Childs = append(o.Childs, child...)
	return o
}

type Attribute struct {
	Key string
	Val string
}

//
// Objects
//

func NewObject(tagname string) Object {
	return Object{
		TagName: tagname,
	}
}

type A struct {
	Object
}

func NewA(href string) A {
	return A{Object{
		TagName:    "a",
		Attributes: []Attribute{Attribute{"href", href}},
	}}
}

func (a A) Target(target string) A {
	a.Object = a.AddAttribute("target", target)
	return a
}

func (a A) Download(filename string) A {
	a.Object = a.AddAttribute("download", filename)
	return a
}

//
// Script
//

type Script struct {
	Object
}

func NewScript(typ string) Script {
	return Script{Object{
		TagName:    "script",
		Attributes: []Attribute{Attribute{"type", typ}},
	}}
}

func (s Script) Src(src string) Script {
	s.Object = s.AddAttribute("src", src)
	return s
}

func (s Script) Code(code string) Script {
	s.Inner = code
	return s
}

func (s Script) String() string {
	return s.Object.String()
}

func NewJavaScript() Script {
	return Script{Object{
		TagName:    "script",
		Attributes: []Attribute{Attribute{"type", "application/javascript"}},
	}}
}

//
// Links
//

func NewLink(rel, typ, href string) Object {
	return Object{
		TagName:    "link",
		noEndTag:   true,
		Attributes: []Attribute{Attribute{"rel", rel}, Attribute{"type", typ}, Attribute{"href", href}},
	}
}

func NewStyleLink(href string) Object {
	return Object{
		TagName:    "link",
		noEndTag:   true,
		Attributes: []Attribute{Attribute{"rel", "stylesheet"}, Attribute{"type", "text/css"}, Attribute{"href", href}},
	}
}

//
// HEADERS
//

type Page struct {
	Title string
	Head  []Object
	Body  Object
}

func NewPage(title string) Page {
	return Page{
		Title: title,
		Body:  NewObject("body"),
	}
}

func (p Page) AddToHead(o ...Object) Page {
	p.Head = append(p.Head, o...)
	return p
}

func (p Page) AddToBody(o ...Object) Page {
	p.Body = p.Body.AddChilds(o...)
	return p
}

func (p Page) String() string {
	return p.buffer(nil).String()
}

func (page Page) buffer(p []byte) *bytes.Buffer {
	w := bytes.NewBuffer(p)

	w.WriteString(`<!DOCTYPE html><html><head>`)
	fmt.Fprintf(w, "<title>%s</title>", page.Title)
	for _, o := range page.Head {
		fmt.Fprint(w, o)
	}
	w.WriteString(`</head>`)

	page.Body.WriteTo(w)

	w.WriteString(`</html>`)
	return w
}

func (p Page) WriteTo(w io.Writer) (n int64, err error) {
	return p.buffer(nil).WriteTo(w)
}
