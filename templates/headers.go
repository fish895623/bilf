package templates

type String struct {
	Str string
}

func (s *String) Header() {
	s.Str += "<h1>hello this is header</h1>"
}