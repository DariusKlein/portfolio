package components

import (
	"github.com/delaneyj/gomponents-iconify/iconify/mdi"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	b "github.com/willoma/bulma-gomponents"
	e "github.com/willoma/gomplements"
	"portfolio/database/ent"
)

func ProjectList(projects []*ent.Project) g.Node {
	return Div(Class("py-2 px-2"), g.Group(g.Map(projects, func(p *ent.Project) g.Node {
		return Project(p)
	})),
	)
}

func Project(project *ent.Project) g.Node {
	return b.Card(
		b.Media(
			b.MediaLeft(
				b.ImageImg(
					project.ImageURL,
					e.Alt("project image"),
					b.ImgSq64,
				),
			),
			b.Title(4, project.Name),
			b.Subtitle(
				6,
				A(Class("flex"), Href(project.URL), mdi.Github(), g.Text("Checkout repo")),
				A(Class("flex"), Href(project.DocURL), mdi.Document(), g.Text("Docs"))),
		),
		b.Content(
			g.Raw(project.Description),
		),
	)
}
