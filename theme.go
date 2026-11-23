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
	// Look is how surfaces are drawn: by light and shadow, or flat.
	Look Look
	// Canvas is the window behind the cards; Control fills buttons and
	// fields, Hover those under the pointer, Select what is chosen or
	// pressed, and Line draws borders and dividers. Only a flat look uses
	// them: neumorphism draws everything in Base.
	Canvas, Control, Hover, Select, Line ui.Color
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

// Look is a way of drawing surfaces.
type Look uint8

const (
	// Neumorphism is soft UI: everything is the color of the surface, and
	// stands out of it or sinks into it by a highlight and a shadow.
	Neumorphism Look = iota
	// Organic is flat and warm: earth tones, cards a shade off the canvas
	// under a faint shadow, round controls with a hairline around them,
	// and a tint of sage for what is chosen.
	// Look: https://www.stylekit.top/zh/styles/natural-organic
	Organic
)

var (
	bark  = ui.Hex("#5c4033")
	cream = ui.Hex("#faf6f1")
	linen = ui.Hex("#e9e0d4")
	sage  = ui.Hex("#8b9d77")
)

// OrganicLightPalette is the natural, organic look: bark for text, never
// black; a cream canvas and linen cards; sage for what is chosen. Sage
// itself is too light to read as text on linen, so the accent is a deeper
// sage; the danger color is a brick red, as the look names none.
var OrganicLightPalette = Palette{
	Look:     Organic,
	Canvas:   cream,
	Base:     linen,
	Control:  cream,
	Hover:    cream.Mix(sage, 0.22),
	Select:   linen.Mix(sage, 0.38),
	Line:     linen.Mix(bark, 0.22),
	Dark:     linen.Mix(bark, 0.30),
	Light:    cream,
	Text:     bark,
	Muted:    bark.Mix(linen, 0.34),
	Accent:   ui.Hex("#586c44"),
	OnAccent: cream,
	Danger:   ui.Hex("#b4533b"),
	Success:  sage,
}

// OrganicDarkPalette carries the organic look into a dim room, which the
// look itself does not describe: the same earth, as dark as bark at night.
var OrganicDarkPalette = Palette{
	IsDark:   true,
	Look:     Organic,
	Canvas:   ui.Hex("#241e1a"),
	Base:     ui.Hex("#322a24"),
	Control:  ui.Hex("#3d342d"),
	Hover:    ui.Hex("#3d342d").Mix(sage, 0.20),
	Select:   ui.Hex("#322a24").Mix(sage, 0.34),
	Line:     ui.Hex("#322a24").Mix(cream, 0.16),
	Dark:     ui.Hex("#322a24").Mix(cream, 0.22),
	Light:    ui.Hex("#3d342d"),
	Text:     ui.Hex("#efe6da"),
	Muted:    ui.Hex("#b3a698"),
	Accent:   ui.Hex("#b2c49d"),
	OnAccent: ui.Hex("#241e1a"),
	Danger:   ui.Hex("#e0876c"),
	Success:  sage,
}

// The colors of the palette in use, which Use sets: read them while building
// a view.
var (
	Base, Dark, Light, Text, Muted ui.Color
	Accent, OnAccent               ui.Color
	Danger, Success                ui.Color
	// Those of a flat look; in neumorphism all are Base.
	Canvas, Control, Hover, Select, Line ui.Color
	isDark                               bool
	look                                 Look
)

// Radii, which the look in use sets: Use changes them.
var (
	RadiusCard    float32 = 24
	RadiusControl float32 = 16
	RadiusSmall   float32 = 12
	RadiusField   float32 = 12 // text fields
)

// RadiusPill rounds an element's ends fully.
const RadiusPill = 999

func init() { Use(LightPalette) }

// Use makes p the palette in use.
func Use(p Palette) {
	Base, Dark, Light, Text, Muted = p.Base, p.Dark, p.Light, p.Text, p.Muted
	Accent, OnAccent, Danger, Success = p.Accent, p.OnAccent, p.Danger, p.Success
	isDark, look = p.IsDark, p.Look
	Canvas, Control, Hover, Select, Line = p.Canvas, p.Control, p.Hover, p.Select, p.Line
	RadiusCard, RadiusControl, RadiusSmall, RadiusField = 24, 16, 12, 12
	switch look {
	case Neumorphism:
		Canvas, Control, Hover, Select, Line = Base, Base, Base, Base, Dark
	case Organic:
		// No sharp corners: cards 32, fields 20, buttons round.
		RadiusCard, RadiusControl, RadiusSmall, RadiusField = 32, RadiusPill, 16, 20
	}
}

// CurrentLook is the look of the palette in use.
func CurrentLook() Look { return look }

// IsDark reports whether the palette in use is a dark one.
func IsDark() bool { return isDark }

// Mode is the appearance the user chose.
type Mode string

const (
	ModeSystem Mode = "system" // follow the system's light or dark appearance
	ModeLight  Mode = "light"
	ModeDark   Mode = "dark"
)

// Style names a look, for a setting: StyleNeumorphism or StyleOrganic.
type Style string

const (
	StyleNeumorphism Style = "neumorphism"
	StyleOrganic     Style = "organic"
)

// PaletteFor returns the neumorphism palette of mode, given whether the
// system is dark.
func PaletteFor(mode Mode, systemDark bool) Palette {
	return PaletteOf(StyleNeumorphism, mode, systemDark)
}

// PaletteOf returns the palette of a style in mode, given whether the
// system is dark. A style it does not know is neumorphism.
func PaletteOf(style Style, mode Mode, systemDark bool) Palette {
	dark := mode == ModeDark || (mode != ModeLight && systemDark)
	switch {
	case style == StyleOrganic && dark:
		return OrganicDarkPalette
	case style == StyleOrganic:
		return OrganicLightPalette
	case dark:
		return DarkPalette
	}
	return LightPalette
}

// UseStyle applies a style in mode for the frame, as UseMode does for
// neumorphism.
func UseStyle(c *ui.Context, style Style, mode Mode) {
	Use(PaletteOf(style, mode, c.Theme().Dark))
	c.SetTheme(Theme())
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
	t.Background = Canvas
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
