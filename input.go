package mallow

import "github.com/egoist/mygo/ui"

// TextInput edits *value in a field sunk only a little into the surface: a
// field is wide, and a recess as deep as a small control's reads as a slot.
// Focused, it keeps its recess and takes a line of the accent. It returns the
// input: give it a Placeholder and a Label, and ask Submitted or Changed.
func TextInput(c *ui.Context, value *string) ui.Element {
	in := ui.TextInputBase(c, value).Padding(8, 12).Radius(RadiusSmall).
		FontSize(14).TextColor(Text).PlaceholderColor(Muted).FocusRing(false)
	in.Background(Base).InsetShadow(1, 1, 3, 0, Dark).InsetShadow(-1, -1, 3, 0, Light)
	if in.Focused() {
		in.InsetShadow(0, 0, 0, 1.5, Accent.Alpha(0.45))
	}
	return Soft(in)
}

// TextArea edits *value over several lines, in a field as TextInput's.
func TextArea(c *ui.Context, value *string) ui.Element {
	in := ui.TextAreaBase(c, value).Padding(8, 12).Radius(RadiusSmall).MinHeight(72).
		FontSize(14).TextColor(Text).PlaceholderColor(Muted).FocusRing(false)
	in.Background(Base).InsetShadow(1, 1, 3, 0, Dark).InsetShadow(-1, -1, 3, 0, Light)
	if in.Focused() {
		in.InsetShadow(0, 0, 0, 1.5, Accent.Alpha(0.45))
	}
	return Soft(in)
}
