package handlers

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	b "github.com/willoma/bulma-gomponents"
	e "github.com/willoma/gomplements"
	"net/http"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {

	err := Page("Homepage", createHomeBody(w, r)).Render(w)
	if err != nil {
		return
	}

}

func createHomeBody(w http.ResponseWriter, r *http.Request) g.Node {

	return Body(
		Div(
			b.Media(
				b.MarginBottom(0),
				b.MediaLeft(
					Figure,
					b.ImageImg("assets/images/Darius_pasfoto_2023_Large.png",
						b.OnImg(Class("object-scale-down max-h-[calc(100vh_-_3.75rem)]")),
					),
				),
			),
			b.Section(
				e.H1("Darius Klein",
					Class("title backdrop-blur")),
				e.H2(
					"Backend developer die zelf zijn servers kan opzetten.",
					Class("subtitle backdrop-blur")),
			),
			Class("flex"),
		),
		Hr(Class("h-1 border-t-0 bg-neutral-300 dark:bg-white/100 flex")),
	)
}
