package html

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

//
// Properties
//

type Property struct {
	Property string
	Value    string
}

func (s Style) AddProperty(property ...string) Style {
	for _, p := range property {
		s.Properties = append(s.Properties, strings.Trim(p, ";"))
	}

	return s
}
