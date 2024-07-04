package components

import (
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
	b "github.com/willoma/bulma-gomponents"
)

var BaseUrl = "https://api.portfolio.dariusklein.nl"

// todo var BaseUrl = "http://localhost:4001"

func Login() g.Node {
	return FormEl(hx.Post(BaseUrl+"/login"), //https://api.portfolio.dariusklein.nl/login
		Class("max-w-xl m-auto py-32"),
		b.Box(
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
