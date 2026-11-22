// Package mallow is a neumorphism (soft UI) component library for MyGo's
// native UI. Every element is the color of the surface it sits on; a light
// from the top left gives raised elements a highlight above and a shadow
// below, and recessed elements the same pair inside. No borders, large radii,
// one accent.
//
// Every component is drawn by mallow on MyGo's headless bases (ButtonBase,
// CheckboxBase, TextInputBase, …), which bring the pointer, keyboard, focus,
// input methods and accessibility; none of MyGo's styled widgets is used.
//
// Call Use (or UseMode) at the start of each frame, then c.SetTheme(Theme()),
// and build with the components and the colors of the palette in use.
//
// Look: https://www.stylekit.top/zh/styles/neumorphism
package mallow

import "github.com/egoist/mygo/ui"

// Palette is the colors of one variant of the look.
type Palette struct {
	IsDark bool
	// Base is the color of the surface and of every element on it.
	Base ui.Color
	// Dark is the shadow on the side away from the light; Light the highlight.
	Dark, Light ui.Color
	Text, Muted ui.Color
	// Accent marks the primary action and what is chosen; OnAccent is text
	// and marks on it. Danger is for destructive actions, Success for what
	// went well.
	Accent, OnAccent ui.Color
	Danger, Success  ui.Color
}

// LightPalette is the look's own palette.
var LightPalette = Palette{
	Base:     ui.Hex("#e0e5ec"),
	Dark:     ui.Hex("#b8bcc2"),
	Light:    ui.Hex("#ffffff"),
	Text:     ui.Hex("#3d4451"),
	Muted:    ui.Hex("#6b7280"),
	Accent:   ui.Hex("#6d5dfc"),
	OnAccent: ui.Hex("#ffffff"),
	Danger:   ui.Hex("#ff6b6b"),
	Success:  ui.Hex("#4ecdc4"),
}

// DarkPalette carries the same light into a dim room: a charcoal surface, a
// near-black shadow and a highlight only a little lighter than the surface,
// as white would glare; text and accent are lifted to keep their contrast.
var DarkPalette = Palette{
	IsDark:   true,
	Base:     ui.Hex("#2a2d33"),
	Dark:     ui.Hex("#1b1d21"),
	Light:    ui.Hex("#393d44"),
	Text:     ui.Hex("#dde1e7"),
	Muted:    ui.Hex("#8d939d"),
	Accent:   ui.Hex("#8b7dff"),
	OnAccent: ui.Hex("#ffffff"),
	Danger:   ui.Hex("#ff7b7b"),
	Success:  ui.Hex("#5ad8cf"),
}

// The colors of the palette in use, which Use sets: read them while building
// a view.
var (
	Base, Dark, Light, Text, Muted ui.Color
	Accent, OnAccent               ui.Color
	Danger, Success                ui.Color
	isDark                         bool
)

func init() { Use(LightPalette) }

// Use makes p the palette in use.
func Use(p Palette) {
	Base, Dark, Light, Text, Muted = p.Base, p.Dark, p.Light, p.Text, p.Muted
	Accent, OnAccent, Danger, Success = p.Accent, p.OnAccent, p.Danger, p.Success
	isDark = p.IsDark
}

// IsDark reports whether the palette in use is a dark one.
func IsDark() bool { return isDark }

// Mode is the appearance the user chose.
type Mode string

const (
	ModeSystem Mode = "system" // follow the system's light or dark appearance
	ModeLight  Mode = "light"
	ModeDark   Mode = "dark"
)

// PaletteFor returns the palette of mode, given whether the system is dark.
func PaletteFor(mode Mode, systemDark bool) Palette {
	if mode == ModeDark || (mode != ModeLight && systemDark) {
		return DarkPalette
	}
	return LightPalette
}

// UseMode applies mode for the frame and sets the window's theme to match.
// Call it first in the view: before SetTheme, the frame's theme is MyGo's
// default, which follows the system's appearance as it changes.
func UseMode(c *ui.Context, mode Mode) {
	Use(PaletteFor(mode, c.Theme().Dark))
	c.SetTheme(Theme())
}

// Font is the family mallow sets first. Text in a fallback font gets no
// strikethrough in MyGo, so a font with CJK glyphs comes first.
var Font = "Microsoft YaHei UI"

// Theme returns the window theme for the palette in use. Mallow's components
// take their colors from the palette, not from it; it colors what the bases
// and the window draw on their own: the window background, text defaults,
// the caret, selections, scroll bars and focus rings.
func Theme() *ui.Theme {
	t := ui.LightTheme()
	if isDark {
		t = ui.DarkTheme()
	}
	t = t.WithAccent(Accent)
	t.Background = Base
	t.Surface = Base
	t.SurfaceHover = Base
	t.SurfacePressed = Base
	t.Border = ui.Transparent
	t.Text = Text
	t.TextMuted = Muted
	t.Danger = Danger
	t.Success = Success
	t.InactiveAccent = Accent
	t.Scrollbar = Dark
	t.Radius = RadiusSmall
	t.FontSize = 14
	t.Font = Font
	t.Inverse, t.InverseText = Text, Base
	return t
}
