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
					b.ImageImg("assets/images/darius_circle_noEdge_crop_upscayl_1x_realesrgan-x4plus-anime.png",
						b.OnImg(Class("object-scale-down grayscale "+
							"max-h-[calc(25vh_-_3.75rem)] "+
							"sm:max-h-[calc(40vh_-_3.75rem)] "+
							"md:max-h-[calc(60vh_-_3.75rem)] "+
							"lg:max-h-[calc(100vh_-_3.75rem)] ")),
					),
				),
			),
			b.Section(
				b.Margin(100),
				e.H1("Darius Klein",
					Class("title backdrop-blur "+
						"text-2xl "+
						"sm:text-4xl "+
						"md:text-6xl "+
						"lg:text-8xl "),
				),
				e.H2(
					"Backend developer die zelf zijn servers kan opzetten.",
					Class("subtitle backdrop-blur"),
				),

				b.ButtonAHref("/projects", "Mijn projecten.", b.Primary, b.Centered, b.Large),
			),
			Class("flex"),
		),
		Hr(Class("h-1 border-t-0 bg-neutral-300 dark:bg-white/100 flex")),
	)
}
