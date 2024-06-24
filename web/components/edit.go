package components

import (
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
)

func Edit() g.Node {
	return Div(
		Class("px-3 py-2"),
		hx.Post("/projects/edit"),
		hx.Swap("outerHTML"),
		hx.SelectOOB("true"),
		hx.Target("#main"),
		Div(
			hx.Trigger("load"),
			hx.Get(BaseUrl+"/htmx/canEdit"),
			hx.Target("this"),
			hx.Swap("outerHTML")),
	)
}
