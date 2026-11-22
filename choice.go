package mallow

import "github.com/egoist/mygo/ui"

// CheckDisc is the mark of a check: a recessed disc, filled with the accent
// and checked when on.
func CheckDisc(c *ui.Context, on bool) ui.Element {
	disc := ui.Box(c).Size(18, 18).Shrink(0).Radius(RadiusPill).AlignItems(ui.Center).Justify(ui.Center)
	if on {
		disc.Background(Accent).Shadow(2, 2, 4, 0, Dark).Shadow(-2, -2, 4, 0, Light)
	} else {
		Recessed(disc, Tiny)
	}
	return Soft(disc).Children(func() {
		if on {
			GlyphCheck(c, OnAccent)
		}
	})
}

// Check toggles *on: a check disc and its label, the whole row clickable.
// While done, the label is muted and struck through, as for a to-do. Ask
// Changed of it for work that follows a click.
func Check(c *ui.Context, on *bool, label string) ui.Element {
	row := ui.CheckboxBase(c, on).Label(label).Padding(6, 4).Gap(8).AlignItems(ui.Start).
		Radius(RadiusSmall).Cursor(ui.CursorPointer)
	return row.Children(func() {
		CheckDisc(c, *on).Margin(1, 0, 0, 0)
		text := ui.Text(c, label).Grow(1).MinWidth(0).FontSize(14).TextColor(Text)
		if *on {
			text.TextColor(Muted).Strikethrough()
		}
	})
}

// Switch toggles *on: a recessed track whose raised knob slides over,
// filled with the accent when on, with an optional label before it.
func Switch(c *ui.Context, on *bool, label string) ui.Element {
	row := ui.SwitchBase(c, on).Label(label).Gap(12).AlignItems(ui.Center).Cursor(ui.CursorPointer)
	return row.Children(func() {
		if label != "" {
			ui.Text(c, label).Grow(1).FontSize(14).TextColor(Text)
		}
		track := ui.Row(c).Size(46, 26).Shrink(0).Padding(4).Radius(RadiusPill).AlignItems(ui.Center)
		if *on {
			track.Justify(ui.End)
		}
		Recessed(track, 3).Children(func() {
			knob := ui.Box(c).Size(18, 18).Radius(RadiusPill)
			Raised(knob, Tiny)
			if *on {
				knob.Background(Accent)
			}
			Soft(knob)
		})
	})
}

// Option is one choice of a Segmented control.
type Option struct {
	Value, Label string
}

// Segmented chooses one of options into *value: a recessed track whose chosen
// segment rises out of it, a radio group the arrows move through. It reports
// whether the choice changed.
func Segmented(c *ui.Context, label string, value *string, options []Option) bool {
	idx := 0
	for i, o := range options {
		if o.Value == *value {
			idx = i
		}
	}
	before := idx
	seg := ui.SegmentedBase(c.Key(label), &idx, len(options))
	seg.Track.Label(label).Padding(4).Gap(4).Radius(RadiusPill)
	Recessed(seg.Track, 3).Children(func() {
		for i, o := range options {
			s := seg.Segment(i).Label(o.Label).Grow(1).Padding(6, 10).Radius(RadiusPill).
				AlignItems(ui.Center).Justify(ui.Center).Cursor(ui.CursorPointer)
			color := Muted
			if i == idx {
				Raised(s, Tiny)
				color = Accent
			}
			Soft(s).Children(func() {
				ui.Text(c, o.Label).FontSize(12).FontWeight(500).TextColor(color).SingleLine()
			})
		}
	})
	if idx != before && idx < len(options) {
		*value = options[idx].Value
		return true
	}
	return false
}
