package css

import "testing"

func TestStyle(t *testing.T) {
	s := NewStyle("div").AddProperty("color: red")
	compare(t, s.String(), "div{color: red;}")
	t.Log(s)

	s = NewStyle("div p", "#id .class span").AddProperty("color: green", "position: absolute")
	compare(t, s.String(), "div p, #id .class span{color: green; position: absolute;}")
	t.Log(s)
}

func compare(t *testing.T, a, b string) {
	if a != b {
		t.Errorf("failed, not equal:\n\t%s\n\t%s", a, b)
	}
}
