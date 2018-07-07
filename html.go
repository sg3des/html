package html

import (
	"bytes"
	"fmt"
	"strings"
)

type Object struct {
	TagName  string
	noEndTag bool

	ID         string
	Class      []string
	Attributes []Attribute
	Inner      string
	Childs     []*Object
}

func (o *Object) String() string {
	var w strings.Builder
	w.WriteByte('<')
	w.WriteString(o.TagName)

	if o.ID != "" {
		fmt.Fprintf(&w, ` id='%s'`, o.ID)
	}

	if len(o.Class) > 0 {
		fmt.Fprintf(&w, ` class='%s'`, strings.Join(o.Class, " "))
	}

	for _, attr := range o.Attributes {
		fmt.Fprintf(&w, ` %s='%s'`, attr.Key, attr.Val)
	}

	w.WriteByte('>')

	if o.noEndTag {
		return w.String()
	}

	w.WriteString(o.Inner)
	for _, child := range o.Childs {
		w.WriteString(child.String())
	}

	fmt.Fprintf(&w, `</%s>`, o.TagName)
	return w.String()
}

func (o *Object) SetID(id string) *Object {
	o.ID = id
	return o
}

func (o *Object) AddClass(class ...string) *Object {
	o.Class = append(o.Class, class...)
	return o
}

func (o *Object) AddAttribute(key, val string) *Object {
	o.Attributes = append(o.Attributes, Attribute{key, val})
	return o
}

func (o *Object) AddChilds(child ...*Object) *Object {
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

func NewObject(tagname string) *Object {
	return &Object{
		TagName: tagname,
	}
}

type A struct {
	*Object
}

func NewA(href string) A {
	return A{&Object{
		TagName:    "a",
		Attributes: []Attribute{Attribute{"href", href}},
	}}
}

func (a A) Target(target string) A {
	a.AddAttribute("target", target)
	return a
}

func (a A) Download(filename string) A {
	a.AddAttribute("download", filename)
	return a
}

//
// Script
//

type Script struct {
	*Object
}

func NewScript(typ string) Script {
	return Script{&Object{
		TagName:    "script",
		Attributes: []Attribute{Attribute{"type", typ}},
	}}
}

func (s Script) Src(src string) Script {
	s.AddAttribute("src", src)
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
	return Script{&Object{
		TagName:    "script",
		Attributes: []Attribute{Attribute{"type", "application/javascript"}},
	}}
}

//
// Links
//

func NewLink(rel, typ, href string) *Object {
	return &Object{
		TagName:    "link",
		noEndTag:   true,
		Attributes: []Attribute{Attribute{"rel", rel}, Attribute{"type", typ}, Attribute{"href", href}},
	}
}

func NewStyleLink(href string) *Object {
	return &Object{
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
	Head  []*Object
	Body  []*Object
}

func NewPage(title string) *Page {
	return &Page{
		Title: title,
	}
}

func (p *Page) AddToHead(o ...*Object) *Page {
	p.Head = append(p.Head, o...)
	return p
}

func (p *Page) AddToBody(o ...*Object) *Page {
	p.Body = append(p.Body, o...)
	return p
}

func (p *Page) String() string {
	var w bytes.Buffer

	w.WriteString(`<!DOCTYPE html><html><head>`)
	fmt.Fprintf(&w, "<title>%s</title>", p.Title)
	for _, o := range p.Head {
		fmt.Fprint(&w, o)
	}
	w.WriteString(`</head><body>`)
	for _, o := range p.Body {
		fmt.Fprint(&w, o)
	}

	w.WriteString(`</body></html>`)
	return w.String()
}
