package components

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
	"portfolio/web/types"
)

func ProjectList(projects []types.Project) g.Node {
	return Nav(Class("bg-gray-700 mb-4"),
		container(
			Div(Class("flex items-center space-x-4 h-16"),
				// We can Map custom slices to Nodes
				g.Group(g.Map(projects, func(p types.Project) g.Node {
					return Project(p)
				})),
			),
		),
	)
}

func Project(project types.Project) g.Node {
	return A(g.Text(project.Name),
		// Apply CSS classes conditionally
		c.Classes{
			"px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:text-white focus:bg-gray-700": true,
		},
	)
}
