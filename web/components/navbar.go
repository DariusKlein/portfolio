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
			b.NavbarDropdown(
				"placeholder",
				b.Hoverable,
				b.NavbarAHref("#1", "item 1"),
				b.NavbarAHref("#2", "item 2"),
				b.NavbarAHref("#3", "item 3"),
				b.NavbarDivider(),
				b.NavbarAHref("#4", "divided item"),
			),
		),
		b.NavbarEnd(
			b.NavbarItem(
				b.Buttons(
					b.ButtonA(
						b.Primary,
						"Log in",
					),
				),
			),
		),
	)
}
