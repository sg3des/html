package html

import "testing"

func TestStyle(t *testing.T) {
	s := NewStyle("div").AddProperty("color: red")
	compare(t, s.String(), "div{color: red;}")
	t.Log(s)

	s = NewStyle("div p", "#id .class span").AddProperty("color: green", "position: absolute")
	compare(t, s.String(), "div p, #id .class span{color: green; position: absolute;}")
	t.Log(s)
}
