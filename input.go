package mallow

import "github.com/egoist/mygo/ui"

// TextInput edits *value in a recessed field, which keeps its recess when
// focused and takes a tint of the accent instead of a ring. It returns the
// input: give it a Placeholder and a Label, and ask Submitted or Changed.
func TextInput(c *ui.Context, value *string) ui.Element {
	in := ui.TextInputBase(c, value).Padding(8, 12).Radius(RadiusSmall).
		FontSize(14).TextColor(Text).PlaceholderColor(Muted).FocusRing(false)
	Recessed(in, 3)
	if in.Focused() {
		in.InsetShadow(0, 0, 0, 1.5, Accent.Alpha(0.45))
	}
	return Soft(in)
}
