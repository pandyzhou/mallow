package mallow

import "github.com/egoist/mygo/ui"

// press styles b by its state: raised by d at rest, half as deep while
// hovered, pressed in while held.
func press(b ui.Element, d Depth) {
	switch {
	case b.Pressed():
		Recessed(b, d/2)
	case b.Hovered():
		Raised(b, d/2)
	default:
		Raised(b, d)
	}
	Soft(b)
}

// Button is a raised button with a label. It returns the element: ask
// Clicked of it, and add Tooltip as for any element.
func Button(c *ui.Context, label string) ui.Element {
	b := ui.ButtonBase(c).Label(label).Padding(8, 16).Radius(RadiusControl).
		AlignItems(ui.Center).Justify(ui.Center).Cursor(ui.CursorPointer)
	press(b, Small)
	b.Children(func() {
		ui.Text(c, label).FontSize(14).FontWeight(500).TextColor(Text).SingleLine()
	})
	return b
}

// PrimaryButton is a button filled with the accent, for the one action a
// view leads to.
func PrimaryButton(c *ui.Context, label string) ui.Element {
	b := ui.ButtonBase(c).Label(label).Padding(8, 18).Radius(RadiusControl).
		AlignItems(ui.Center).Justify(ui.Center).Cursor(ui.CursorPointer)
	press(b, Small)
	b.Background(Accent)
	b.Children(func() {
		ui.Text(c, label).FontSize(14).FontWeight(500).TextColor(OnAccent).SingleLine()
	})
	return b
}

// Glyph draws a small mark in a color, as the icons of IconButton.
type Glyph func(c *ui.Context, color ui.Color)

// IconButtonStyle varies an IconButton.
type IconButtonStyle struct {
	// Size is the diameter, 26 when zero.
	Size float32
	// Danger turns the glyph the danger color while hovered, as a close or
	// delete button.
	Danger bool
}

// IconButton is a round raised button holding a glyph, labelled for
// assistive technology and shown as a tooltip.
func IconButton(c *ui.Context, label string, glyph Glyph, style IconButtonStyle) ui.Element {
	size := style.Size
	if size == 0 {
		size = 26
	}
	b := ui.ButtonBase(c.Key(label)).Label(label).Size(size, size).Radius(RadiusPill).
		AlignItems(ui.Center).Justify(ui.Center).Cursor(ui.CursorPointer)
	color := Muted
	if b.Hovered() && !b.Pressed() {
		color = Text
		if style.Danger {
			color = Danger
		}
	}
	d := Depth(3)
	if size >= 36 {
		d = Small
	}
	press(b, d)
	Tooltip(c, b, label)
	b.Children(func() { glyph(c, color) })
	return b
}

// GhostIconButton is a glyph that is a button but no surface of its own,
// for small actions inside a row, as deleting an item.
func GhostIconButton(c *ui.Context, label string, glyph Glyph, style IconButtonStyle) ui.Element {
	size := style.Size
	if size == 0 {
		size = 24
	}
	b := ui.ButtonBase(c.Key(label)).Label(label).Size(size, size).Radius(RadiusPill).
		AlignItems(ui.Center).Justify(ui.Center).Cursor(ui.CursorPointer).Background(Base)
	color := Muted
	if b.Hovered() {
		color = Text
		if style.Danger {
			color = Danger
		}
	}
	Tooltip(c, b, label)
	b.Children(func() { glyph(c, color) })
	return b
}
