package components

import (
	"github.com/delaneyj/gomponents-iconify/iconify/mdi"
	g "github.com/maragudk/gomponents"
	b "github.com/willoma/bulma-gomponents"
	e "github.com/willoma/gomplements"
)

func Email(valid bool, invalid bool, color e.Element) e.Element {
	return b.Field(
		b.Label("Email"),
		b.Control(
			b.IconsLeft,
			b.InputEmail(
				e.Name("email"),
				color,
				e.Placeholder("Email input"),
			),
			b.Icon(mdi.Email(), b.Left),
		),
		g.If(invalid, b.Help(b.Danger, "This email is invalid")),
		g.If(valid, b.Help(b.Success, "This email is valid")),
	)
}

func Password(valid bool, invalid bool, color e.Element) e.Element {
	return b.Field(
		b.Label("Email"),
		b.Control(
			b.IconsLeft,
			b.InputPassword(
				e.Name("password"),
				color,
				e.Placeholder("password input"),
			),
			b.Icon(mdi.FormTextboxPassword(), b.Left),
		),
		g.If(invalid, b.Help(b.Danger, "This password is invalid")),
		g.If(valid, b.Help(b.Success, "This password is valid")),
	)
}
