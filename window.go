package mallow

import "github.com/egoist/mygo/ui"

// Window is what the window buttons ask of a window: *mygo.Window has these.
type Window interface {
	Minimize()
	Maximize()
	Unmaximize()
	Close()
	IsMaximized() bool
}

// TitleBarHeight is the room TitleBar takes at the top of the window.
const TitleBarHeight = 44

// TitleBar draws a window's title bar, for a window opened with
// TitleBarStyle: mygo.TitleBarHidden and HideWindowButtons: true: a strip
// that drags the window, which a double click maximizes, holding what
// leading builds (nil for nothing) and, at its right, the window buttons.
// A nil win draws no buttons, as in tests.
func TitleBar(c *ui.Context, win Window, leading func()) {
	ui.Row(c).Height(TitleBarHeight).Shrink(0).Padding(0, 20, 0, 24).AlignItems(ui.Center).Gap(14).
		DragWindow().Children(func() {
		ui.Row(c).Grow(1).AlignItems(ui.Center).Children(func() {
			if leading != nil {
				leading()
			}
		})
		if win != nil {
			WindowButtons(c, win)
		}
	})
}

// WindowButtons are minimize, maximize (restore while maximized) and close,
// as soft round buttons.
func WindowButtons(c *ui.Context, win Window) {
	if IconButton(c, "最小化", GlyphMinimize, IconButtonStyle{}).Clicked() {
		win.Minimize()
	}
	if win.IsMaximized() {
		if IconButton(c, "还原", GlyphRestore, IconButtonStyle{}).Clicked() {
			win.Unmaximize()
		}
	} else if IconButton(c, "最大化", GlyphMaximize, IconButtonStyle{}).Clicked() {
		win.Maximize()
	}
	if IconButton(c, "关闭", GlyphClose, IconButtonStyle{Danger: true}).Clicked() {
		win.Close()
	}
}
