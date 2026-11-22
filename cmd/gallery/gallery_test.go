package main

import (
	"image/png"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"

	"github.com/pandyzhou/mallow"
)

// fakeWindow records what the window buttons ask of the window.
type fakeWindow struct {
	maximized bool
	calls     []string
}

func (w *fakeWindow) Minimize()         { w.calls = append(w.calls, "minimize") }
func (w *fakeWindow) Maximize()         { w.calls = append(w.calls, "maximize"); w.maximized = true }
func (w *fakeWindow) Unmaximize()       { w.calls = append(w.calls, "unmaximize"); w.maximized = false }
func (w *fakeWindow) Close()            { w.calls = append(w.calls, "close") }
func (w *fakeWindow) IsMaximized() bool { return w.maximized }

func newTester(t *testing.T) (*gallery, *ui.Tester) {
	t.Helper()
	g := newGallery()
	g.win = &fakeWindow{}
	return g, ui.NewTester(g.view, 1180, 820)
}

func click(t *testing.T, tt *ui.Tester, label string) {
	t.Helper()
	if err := tt.Click(label); err != nil {
		t.Fatalf("click %q: %v (texts %q)", label, err, tt.Texts())
	}
}

func TestButtons(t *testing.T) {
	g, tt := newTester(t)
	click(t, tt, "普通按钮")
	click(t, tt, "主要按钮")
	click(t, tt, "主要按钮")
	if g.clicks != 1 || g.primary != 2 {
		t.Errorf("clicks %d, primary %d", g.clicks, g.primary)
	}
}

func TestCheckAndSwitch(t *testing.T) {
	g, tt := newTester(t)
	click(t, tt, "读书 30 分钟")
	click(t, tt, "写周报")
	if g.checks != [3]bool{false, true, false} {
		t.Errorf("checks %v", g.checks)
	}
	click(t, tt, "自动同步")
	click(t, tt, "到点提醒")
	if g.sync || !g.notify {
		t.Errorf("sync %v, notify %v", g.sync, g.notify)
	}
}

func TestSegmented(t *testing.T) {
	g, tt := newTester(t)
	click(t, tt, "月")
	if g.period != "month" {
		t.Errorf("period %q after choosing 月", g.period)
	}
	click(t, tt, "深色")
	if g.mode != string(mallow.ModeDark) || mallow.Base != mallow.DarkPalette.Base {
		t.Errorf("mode %q, base %v after choosing 深色", g.mode, mallow.Base)
	}
}

func TestTextInput(t *testing.T) {
	g, tt := newTester(t)
	click(t, tt, "名字")
	tt.Type("Ada")
	if g.name != "Ada" {
		t.Errorf("name %q after typing", g.name)
	}
}

func TestWindowButtons(t *testing.T) {
	g, tt := newTester(t)
	for _, label := range []string{"最小化", "最大化", "还原", "关闭"} {
		click(t, tt, label)
	}
	want := "minimize maximize unmaximize close"
	if got := strings.Join(g.win.(*fakeWindow).calls, " "); got != want {
		t.Errorf("calls %q, want %q", got, want)
	}
}

// TestScreenshot renders the gallery without a window to the PNG at
// MALLOW_SCREENSHOT, in the appearance of MALLOW_SCREENSHOT_THEME (light,
// dark or system) and at the scale of MALLOW_SCREENSHOT_SCALE (1 or 2):
//
//	MALLOW_SCREENSHOT=shot.png MALLOW_SCREENSHOT_THEME=dark go test -run Screenshot ./cmd/gallery
func TestScreenshot(t *testing.T) {
	path := os.Getenv("MALLOW_SCREENSHOT")
	if path == "" {
		t.Skip("set MALLOW_SCREENSHOT to write a screenshot")
	}
	g, tt := newTester(t)
	if m := os.Getenv("MALLOW_SCREENSHOT_THEME"); m != "" {
		g.mode = m
	}
	if os.Getenv("MALLOW_SCREENSHOT_SCALE") == "2" {
		tt.SetScale(2)
	}
	click(t, tt, "名字")
	tt.Type("Ada")
	tt.Advance(time.Second)
	f, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if err := png.Encode(f, tt.Image()); err != nil {
		t.Fatal(err)
	}
}
