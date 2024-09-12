package components

import (
	g "github.com/maragudk/gomponents"
	b "github.com/willoma/bulma-gomponents"
	e "github.com/willoma/gomplements"
)

func Navbar() g.Node {
	return b.Navbar(
		e.AriaNavigation,
		e.AriaLabel("main navigation"),
		b.NavbarBrand(
			b.NavbarAHref(
				"/",
				e.ImgSrc(
					"/assets/images/favicon.ico",
				),
			),
		),
		b.NavbarStart(
			b.NavbarAHref("/", "Home"),
			b.NavbarAHref("/about", "Wie ben ik"),
			b.NavbarAHref("/projects", "projecten"),
		),
		b.NavbarEnd(
			b.NavbarItem(
				b.Buttons(
					b.ButtonAHref(
						"/login",
						b.Primary,
						"Log in",
					),
				),
			),
		),
	)
}
