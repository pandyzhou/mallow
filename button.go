package mallow

import "github.com/egoist/mygo/ui"

// press styles b by its state: raised by d at rest, half as deep while
// hovered, pressed in while held.
func press(b ui.Element, d Depth) {
	if look == Organic {
		// Nothing moves: the fill deepens under the pointer, and takes
		// the tint of what is chosen while held.
		Raised(b, d)
		switch {
		case b.Pressed():
			b.Background(Select)
		case b.Hovered():
			b.Background(Hover)
		}
		Soft(b)
		return
	}
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
	if look == Organic {
		b.Border(1, ui.Transparent)
		if b.Hovered() {
			b.Background(Accent.Mix(Base, 0.12))
		}
	}
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
	// Active shows the glyph in the accent, for a button that is on, as a
	// toggle whose setting is set.
	Active bool
	// Primary fills an IconButton with the accent, for the one action a
	// view leads to, as sending what was written.
	Primary bool
}

// glyphColor is the color of an icon button's glyph.
func glyphColor(b ui.Element, style IconButtonStyle, pressed bool) ui.Color {
	switch {
	case style.Primary:
		return OnAccent
	case b.Hovered() && !pressed && style.Danger:
		return Danger
	case style.Active:
		return Accent
	case b.Hovered() && !pressed:
		return Text
	}
	return Muted
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
	color := glyphColor(b, style, b.Pressed())
	d := Depth(3)
	if size >= 36 {
		d = Small
	}
	press(b, d)
	if style.Primary {
		b.Background(Accent)
		if look == Organic {
			b.Border(1, ui.Transparent)
			if b.Hovered() {
				b.Background(Accent.Mix(Base, 0.12))
			}
		}
	}
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
		AlignItems(ui.Center).Justify(ui.Center).Cursor(ui.CursorPointer)
	if look == Neumorphism {
		b.Background(Base)
	}
	color := glyphColor(b, style, false)
	Tooltip(c, b, label)
	b.Children(func() { glyph(c, color) })
	return b
}
