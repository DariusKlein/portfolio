package handlers

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

func Page(title string, body []g.Node) g.Node {

	return c.HTML5(c.HTML5Props{
		Title:    title,
		Language: "en",
		Head: []g.Node{
			Script(Src("https://cdn.tailwindcss.com?plugins=typography")),
			Link(Rel("icon"), Type("image/x-icon"), Href("assets/images/favicon.ico")),
			Link(Rel("stylesheet"), Href("https://cdn.jsdelivr.net/npm/bulma@1.0.0/css/bulma.min.css")),
		},
		Body: body,
	})
}
