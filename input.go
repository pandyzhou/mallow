package mallow

import "github.com/egoist/mygo/ui"

// TextInput edits *value on a line: the field is level with the surface,
// with a groove under it that takes the accent while it is focused. It is
// not sunk into the surface, as what is chosen is: a recess says "this one",
// a line says "write here". It returns the input: give it a Placeholder and
// a Label, and ask Submitted or Changed.
func TextInput(c *ui.Context, value *string) ui.Element {
	in := ui.TextInputBase(c, value).Padding(8, 4).
		FontSize(14).TextColor(Text).PlaceholderColor(Muted).FocusRing(false)
	return Soft(underlined(in))
}

// TextArea edits *value over several lines, on a field as TextInput's.
func TextArea(c *ui.Context, value *string) ui.Element {
	in := ui.TextAreaBase(c, value).Padding(8, 4).MinHeight(72).
		FontSize(14).TextColor(Text).PlaceholderColor(Muted).FocusRing(false)
	return Soft(underlined(in))
}

// underlined draws a field's line: a groove under it, in the accent while
// the field is focused.
func underlined(in ui.Element) ui.Element {
	in.Background(Base)
	if in.Focused() {
		return in.InsetShadow(0, -2, 0, 0, Accent)
	}
	return in.InsetShadow(0, -1, 0, 0, Light).InsetShadow(0, -2, 1, 0, Dark)
}
