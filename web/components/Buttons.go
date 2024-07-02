package components

import (
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
	b "github.com/willoma/bulma-gomponents"
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

func Save() g.Node {
	return Div(
		Class("px-3 py-2"),
		hx.Patch(BaseUrl+"/login"),
		hx.Swap("outerHTML"),
		hx.SelectOOB("true"),
		hx.Target("#main"),
		b.Button("Save"),
	)
}
