package mallow

import "github.com/egoist/mygo/ui"

// MenuEntry is an item of a ContextMenu.
type MenuEntry struct {
	ID, Label string
	// Glyph is drawn before the label; nil leaves its place empty.
	Glyph func(c *ui.Context, color ui.Color)
	// Danger marks what destroys, as deleting.
	Danger bool
}

// ContextMenu gives target a menu that opens where target is clicked with
// the secondary button, or by the menu key while it has the focus: a small
// raised card of entries, the one under the pointer or the keys pressed in.
// The menu shows while *open is true; keep one bool for each target. It
// returns the ID of the entry chosen in this frame, "" for none.
func ContextMenu(c *ui.Context, target ui.Element, open *bool, entries []MenuEntry) string {
	chosen := ""
	m := ui.ContextMenuBase(c, target, open)
	m.Popup(func(panel ui.Element) {
		panel.Margin(4).Padding(6).Gap(2).MinWidth(148).Radius(RadiusSmall + 4)
		Raised(panel, Medium)
		for _, en := range entries {
			item := m.Item()
			item.Label(en.Label).Height(34).PaddingX(10).Gap(10).Radius(RadiusSmall).
				AlignItems(ui.Center).Cursor(ui.CursorPointer)
			color := Text
			if en.Danger {
				color = Danger
			}
			if item.Highlighted() {
				Recessed(item, Tiny)
				if !en.Danger {
					color = Accent
				}
			}
			item.Children(func() {
				ui.Row(c).Size(16, 16).Shrink(0).Justify(ui.Center).AlignItems(ui.Center).Children(func() {
					if en.Glyph != nil {
						en.Glyph(c, color)
					}
				})
				ui.Text(c, en.Label).FontSize(13).TextColor(color).SingleLine()
			})
			if item.Clicked() {
				chosen = en.ID
			}
		}
	})
	return chosen
}
