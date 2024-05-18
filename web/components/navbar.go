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
				"https://bulma.io",
				e.ImgSrc(
					"https://bulma.io/assets/images/bulma-logo.png",
					e.Width("112"), e.Height("28"),
				),
			),
		),
		b.NavbarStart(
			b.NavbarAHref("#", "Home"),
			b.NavbarAHref("#", "Documentation"),
			b.NavbarDropdown(
				"More",
				b.Hoverable,
				b.NavbarAHref("#", "About"),
				b.NavbarAHref("#", "Jobs"),
				b.NavbarAHref("#", "Contact"),
				b.NavbarDivider(),
				b.NavbarAHref("#", "Report an issue"),
			),
		),
		b.NavbarEnd(
			b.NavbarItem(
				b.Buttons(
					b.ButtonA(
						b.Primary,
						e.Strong("Sign up"),
					),
					b.ButtonA(
						b.Light,
						"Log in",
					),
				),
			),
		),
	)
}
