package mallow

import "github.com/egoist/mygo/ui"

// NavItem is an entry of a navigation list, as a sidebar's: recessed and in
// the accent while chosen, raised a little while hovered, flat otherwise.
// Put the items in a column with Role(ui.RoleTabList). Ask Clicked of it.
func NavItem(c *ui.Context, label string, chosen bool) ui.Element {
	item := ui.ButtonBase(c).Role(ui.RoleTab).Label(label).AccessSelected(chosen).
		Padding(12, 16).Radius(RadiusControl).Cursor(ui.CursorPointer)
	switch {
	case chosen:
		Recessed(item, Small)
	case item.Hovered():
		Raised(item, Small)
	}
	return Soft(item).Children(func() {
		text := ui.Text(c, label).FontSize(14).TextColor(Muted)
		if chosen {
			text.TextColor(Accent).FontWeight(500)
		}
	})
}
