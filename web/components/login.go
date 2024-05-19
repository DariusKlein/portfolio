package components

import (
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
	b "github.com/willoma/bulma-gomponents"
)

func Login() g.Node {
	return FormEl(hx.Post("http://localhost:4001/login"),
		b.Section(
			Email(false, false, nil),
			Password(false, false, nil),
			b.Field(
				b.Grouped,
				b.Control(
					b.Button(b.Link, "login"),
				),
			),
		),
	)
}
