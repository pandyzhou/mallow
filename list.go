package mallow

import "github.com/egoist/mygo/ui"

// CheckMark toggles *on as Check does, but is the disc alone, for a row
// whose text does something else, as a task's that opens it. The label is
// for assistive technology.
func CheckMark(c *ui.Context, on *bool, label string) ui.Element {
	b := ui.CheckboxBase(c, on).Label(label).Padding(4).Radius(RadiusPill).Cursor(ui.CursorPointer)
	return b.Children(func() { CheckDisc(c, *on) })
}

// ListRow is a row of a list that can be the one chosen: it sinks a little
// into the surface while it is. Give it children, as a CheckMark and a
// Plain button that chooses it.
func ListRow(c *ui.Context, chosen bool) ui.Element {
	row := ui.Row(c).Padding(2, 8, 2, 4).Gap(4).AlignItems(ui.Center).Radius(RadiusSmall)
	switch {
	case look == Organic && chosen:
		row.Background(Select)
	case look == Organic:
	case chosen:
		row.Background(Base).InsetShadow(2, 2, 4, 0, Dark).InsetShadow(-2, -2, 4, 0, Light)
	default:
		row.Background(Base)
	}
	return Soft(row)
}

// Plain is a button without a surface, a row of the children it is given:
// what a ListRow holds to be chosen by. The label is for assistive
// technology.
func Plain(c *ui.Context, label string) ui.Element {
	return ui.ButtonBase(c).Label(label).Justify(ui.Start).AlignItems(ui.Center).
		Radius(RadiusSmall).Cursor(ui.CursorPointer)
}

// Chip is a small pill that is on or off, for choosing among a few things
// that wrap over lines, as a task's list: raised while off, pressed in and
// in the accent while on. Ask Clicked of it.
func Chip(c *ui.Context, label string, on bool) ui.Element {
	b := ui.ButtonBase(c.Key(label)).Label(label).Padding(5, 12).Radius(RadiusPill).
		AlignItems(ui.Center).Justify(ui.Center).Cursor(ui.CursorPointer)
	color := Muted
	switch {
	case on:
		Recessed(b, Tiny)
		color = Accent
	case b.Hovered():
		Raised(b, Tiny)
		color = Text
	default:
		Raised(b, 3)
	}
	Soft(b).Children(func() {
		ui.Text(c, label).FontSize(12).FontWeight(500).TextColor(color).SingleLine()
	})
	return b
}
