package mallow

import "github.com/egoist/mygo/ui"

// NavItem is a standalone entry of a navigation list: recessed and in the
// accent while chosen, raised a little while hovered, flat otherwise. Put
// the items in a column with Role(ui.RoleTabList). Ask Clicked of it.
// NavList draws a list of them whose recess slides between the entries.
func NavItem(c *ui.Context, label string, chosen bool) ui.Element {
	item := navEntry(c, label, chosen)
	switch {
	case chosen:
		Recessed(item, Small)
	case item.Hovered():
		Raised(item, Small)
	}
	return Soft(item)
}

// navEntry is an entry of a navigation list without a surface: its text, in
// the accent while chosen and darker while hovered.
func navEntry(c *ui.Context, label string, chosen bool) ui.Element {
	item := ui.ButtonBase(c).Role(ui.RoleTab).Label(label).AccessSelected(chosen).
		Height(navItemHeight).PaddingX(16).Radius(RadiusControl).Justify(ui.Center).
		Cursor(ui.CursorPointer)
	color := Muted
	switch {
	case chosen:
		color = Accent
	case item.Hovered():
		color = Text
	}
	return item.Children(func() {
		text := ui.Text(c, label).FontSize(14).TextColor(color).SingleLine()
		if chosen {
			text.FontWeight(500)
		}
	})
}

// Item is an entry of a NavList, a SourceList or a Sidebar.
type Item struct {
	ID, Label string
	// Glyph, drawn before the label, and Note, small at the far end, as a
	// count, are for a SourceList's entries.
	Glyph func(c *ui.Context, color ui.Color)
	Note  string
	// With, if set, is called with a SourceList's entry once it is built:
	// to give it more, as a ContextMenu.
	With func(c *ui.Context, entry ui.Element)
}

const (
	navItemHeight = 44
	navItemGap    = 12
)

// NavList is a column of entries choosing one into *selected, as macOS's
// source lists: one recess marks the chosen entry and, when the choice
// changes, slides over to the new one (Slide). It reports whether the
// choice changed.
func NavList(c *ui.Context, label string, selected *string, items []Item) bool {
	idx := -1
	for i, it := range items {
		if it.ID == *selected {
			idx = i
		}
	}
	changed := false
	ui.Column(c).Role(ui.RoleTabList).Label(label).Gap(navItemGap).Children(func() {
		// The recess paints under the entries (ZIndex -1: absolute elements
		// paint over the flow otherwise) and is keyed, so that its position
		// animates from frame to frame.
		if idx >= 0 {
			mark := ui.Box(c.Key("nav-mark")).Absolute().Left(0).Right(0).Height(navItemHeight).
				Radius(RadiusControl).PassThrough().ZIndex(-1)
			Recessed(mark, Small)
			y := mark.AnimateTo("y", float32(idx*(navItemHeight+navItemGap)), Slide)
			mark.Top(y)
		}
		for i, it := range items {
			if navEntry(c.Key(it.ID), it.Label, i == idx).Clicked() && i != idx {
				*selected, changed = it.ID, true
			}
		}
	})
	return changed
}

const (
	sourceItemHeight = 36
	sourceItemGap    = 4
)

// SourceList is a NavList for many entries: lower ones set close together,
// each a glyph, its label from the left and a note at the far end, as the
// lists of a to-do app with how much each holds. Several can choose into
// one *selected: the recess shows in the list that holds the choice. It
// reports whether the choice changed.
func SourceList(c *ui.Context, label string, selected *string, items []Item) bool {
	idx := -1
	for i, it := range items {
		if it.ID == *selected {
			idx = i
		}
	}
	changed := false
	ui.Column(c).Role(ui.RoleTabList).Label(label).Gap(sourceItemGap).Children(func() {
		if idx >= 0 {
			mark := ui.Box(c.Key("source-mark")).Absolute().Left(0).Right(0).Height(sourceItemHeight).
				Radius(RadiusSmall).PassThrough().ZIndex(-1)
			Recessed(mark, Tiny)
			mark.Top(mark.AnimateTo("y", float32(idx*(sourceItemHeight+sourceItemGap)), Slide))
		}
		for i, it := range items {
			entry := sourceEntry(c.Key(it.ID), it, i == idx)
			if entry.Clicked() && i != idx {
				*selected, changed = it.ID, true
			}
			if it.With != nil {
				it.With(c.Key(it.ID+"/with"), entry)
			}
		}
	})
	return changed
}

func sourceEntry(c *ui.Context, it Item, chosen bool) ui.Element {
	item := ui.ButtonBase(c).Role(ui.RoleTab).Label(it.Label).AccessSelected(chosen).
		Height(sourceItemHeight).PaddingX(12).Gap(10).Radius(RadiusSmall).
		Justify(ui.Start).AlignItems(ui.Center).Cursor(ui.CursorPointer)
	color := Muted
	switch {
	case chosen:
		color = Accent
	case item.Hovered():
		color = Text
	}
	return item.Children(func() {
		if it.Glyph != nil {
			ui.Row(c).Size(16, 16).Shrink(0).Justify(ui.Center).AlignItems(ui.Center).Children(func() {
				it.Glyph(c, color)
			})
		}
		text := ui.Text(c, it.Label).Grow(1).MinWidth(0).FontSize(14).TextColor(color).SingleLine()
		if chosen {
			text.FontWeight(500)
		}
		if it.Note != "" {
			ui.Text(c, it.Note).FontSize(12).TextColor(Muted)
		}
	})
}

// SidebarOptions describe a Sidebar.
type SidebarOptions struct {
	// Title and Subtitle head the sidebar, as the app's name.
	Title, Subtitle string
	// Items are the entries of its navigation, choosing into Selected.
	Items    []Item
	Selected *string
	// Bottom are entries set apart at the bottom, choosing into Selected
	// too: what is not one of the app's parts, as its settings.
	Bottom []Item
	// Footer builds what goes at the bottom, as an appearance switch; nil
	// for nothing.
	Footer func()
	// Width is the sidebar's width, 184 when zero.
	Width float32
}

// Sidebar is a raised panel running down the side of a window: a title, a
// groove, a NavList whose recess slides between its entries, and at the
// bottom more entries, set apart, and a footer. It reports whether the choice changed.
func Sidebar(c *ui.Context, o SidebarOptions) bool {
	width := o.Width
	if width == 0 {
		width = 184
	}
	changed := false
	Card(c, Large).Width(width).Shrink(0).Padding(28, 16).Gap(24).Children(func() {
		if o.Title != "" || o.Subtitle != "" {
			ui.Column(c).PaddingX(12).Gap(4).Children(func() {
				if o.Title != "" {
					ui.Text(c, o.Title).FontSize(20).FontWeight(600).TextColor(Text)
				}
				if o.Subtitle != "" {
					ui.Text(c, o.Subtitle).FontSize(12).TextColor(Muted)
				}
			})
			Groove(c).MarginX(4)
		}
		ui.Column(c).Grow(1).MinHeight(0).Children(func() {
			changed = NavList(c, o.Title+"导航", o.Selected, o.Items)
		})
		if len(o.Bottom) > 0 {
			if NavList(c, o.Title+"其他", o.Selected, o.Bottom) {
				changed = true
			}
		}
		if o.Footer != nil {
			ui.Column(c).Gap(10).Children(o.Footer)
		}
	})
	return changed
}
