package components

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type PageLink struct {
	Path string
	Name string
}

func Navbar(currentPath string, links []PageLink) g.Node {
	return Nav(Class("bg-gray-700 mb-4"),
		container(
			Div(Class("flex items-center space-x-4 h-16"),
				navbarLink("/", "Home", currentPath == "/"),

				// We can Map custom slices to Nodes
				g.Group(g.Map(links, func(l PageLink) g.Node {
					return navbarLink(l.Path, l.Name, currentPath == l.Path)
				})),
			),
		),
	)
}

// NavbarLink is a link in the Navbar.
func navbarLink(path, text string, active bool) g.Node {
	return A(Href(path), g.Text(text),
		// Apply CSS classes conditionally
		c.Classes{
			"px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:text-white focus:bg-gray-700": true,
			"text-white bg-gray-900":                           active,
			"text-gray-300 hover:text-white hover:bg-gray-700": !active,
		},
	)
}

func container(children ...g.Node) g.Node {
	return Div(Class("mx-auto px-2 sm:px-6 lg:px-8"), g.Group(children))
}
