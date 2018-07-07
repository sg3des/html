package css

import (
	"bytes"
	"fmt"
	"strings"
)

type Style struct {
	Selectors  []string
	Properties []string
}

func NewStyle(selector ...string) Style {
	return Style{
		Selectors: selector,
	}
}

func (s Style) String() string {
	var w bytes.Buffer
	fmt.Fprint(&w, strings.Join(s.Selectors, ", "))

	w.WriteByte('{')
	for i, p := range s.Properties {
		w.WriteString(p)

		w.WriteByte(';')
		if i < len(s.Properties)-1 {
			w.WriteByte(' ')
		}
	}
	w.WriteByte('}')

	return w.String()
}

func (s Style) WriteProperties(w *bytes.Buffer) {
	for i, p := range s.Properties {
		w.WriteString(p)

		w.WriteByte(';')
		if i < len(s.Properties)-1 {
			w.WriteByte(' ')
		}
	}
}

//
// Properties
//

func (s Style) AddProperty(property ...string) Style {
	for _, p := range property {
		s.Properties = append(s.Properties, strings.Trim(p, ";"))
	}

	return s
}

func (s Style) addProperty(name, value string) Style {
	s.Properties = append(s.Properties, fmt.Sprintf("%s: %s", name, value))
	return s
}

func (s Style) Position(position string) Style {
	return s.addProperty("position", position)
}

func (s Style) Margin(value string) Style {
	return s.addProperty("margin", value)
}
func (s Style) MarginLeft(value string) Style {
	return s.addProperty("margin-left", value)
}
func (s Style) MarginRight(value string) Style {
	return s.addProperty("margin-right", value)
}
func (s Style) MarginTop(value string) Style {
	return s.addProperty("margin-top", value)
}
func (s Style) MarginBottom(value string) Style {
	return s.addProperty("margin-bottom", value)
}

func (s Style) FontFamily(value string) Style {
	return s.addProperty("font-family", value)
}

func (s Style) Color(value string) Style {
	return s.addProperty("color", value)
}

func (s Style) Width(value string) Style {
	return s.addProperty("width", value)
}
func (s Style) Height(value string) Style {
	return s.addProperty("height", value)
}

func (s Style) Left(value string) Style {
	return s.addProperty("left", value)
}
func (s Style) Right(value string) Style {
	return s.addProperty("right", value)
}
func (s Style) Top(value string) Style {
	return s.addProperty("top", value)
}
func (s Style) Bottom(value string) Style {
	return s.addProperty("bottom", value)
}

func (s Style) Float(value string) Style {
	return s.addProperty("float", value)
}
