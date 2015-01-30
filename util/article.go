package util

type Heading string
type Paragraph string

type Article struct {
	Title string
	Text  []interface{}
}

func (a *Article) Append(v interface{}) {
	a.Text = append(a.Text, v)
}

func (a *Article) Prepend(v interface{}) {
	a.Text = append([]interface{}{v}, a.Text...)
}
