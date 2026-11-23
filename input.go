package mallow

import "github.com/egoist/mygo/ui"

// TextInput edits *value on a line: the field is level with the surface,
// with a groove under it that takes the accent while it is focused. It is
// not sunk into the surface, as what is chosen is: a recess says "this one",
// a line says "write here". It returns the input: give it a Placeholder and
// a Label, and ask Submitted or Changed.
func TextInput(c *ui.Context, value *string) ui.Element {
	in := ui.TextInputBase(c, value).
		FontSize(14).TextColor(Text).PlaceholderColor(Muted).FocusRing(false)
	return Soft(underlined(in))
}

// TextArea edits *value over several lines, on a field as TextInput's.
func TextArea(c *ui.Context, value *string) ui.Element {
	in := ui.TextAreaBase(c, value).MinHeight(72).
		FontSize(14).TextColor(Text).PlaceholderColor(Muted).FocusRing(false)
	return Soft(underlined(in))
}

// underlined draws a field's line: a groove under it, in the accent while
// the field is focused.
func underlined(in ui.Element) ui.Element {
	if look == Organic {
		// A round field, a shade lighter than its card, in a hairline:
		// focused, the line turns sage, with a ring of it around.
		in.Padding(9, 16).Radius(RadiusField).Background(Control).Border(1, Line)
		if in.Focused() {
			in.Border(1, Accent).Shadow(0, 0, 0, 3, Accent.Alpha(0.2))
		}
		return in
	}
	in.Padding(8, 4).Background(Base)
	if in.Focused() {
		return in.InsetShadow(0, -2, 0, 0, Accent)
	}
	// Without blur: a blurred inner shadow shades the sides of the field
	// too, which shows as a thin upright line at each end.
	return in.InsetShadow(0, -2.5, 0, 0, Dark).InsetShadow(0, -1, 0, 0, Light)
}
