package main

import (
	"fmt"
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
	// Tall enough to show every section, so that tests can click them.
	return g, ui.NewTester(g.view, 1180, 1700)
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
	click(t, tt, "每周重复")
	if !g.repeat {
		t.Errorf("the repeat toggle is off after a click")
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

// The other look draws the same components flat, in earth tones; choosing
// neumorphism again brings its palette back.
func TestStyle(t *testing.T) {
	g, tt := newTester(t)
	click(t, tt, "自然有机")
	if g.style != string(mallow.StyleOrganic) || mallow.CurrentLook() != mallow.Organic || mallow.Base != mallow.OrganicLightPalette.Base {
		t.Fatalf("style %q, base %v after choosing 自然有机", g.style, mallow.Base)
	}
	if mallow.Canvas == mallow.Base {
		t.Errorf("the organic canvas is the color of its cards, %v", mallow.Base)
	}
	// Its components work as they do in the other look.
	click(t, tt, "月")
	click(t, tt, "深色")
	if g.period != "month" || mallow.Base != mallow.OrganicDarkPalette.Base {
		t.Errorf("in the organic look: period %q, base %v after choosing 月 and 深色", g.period, mallow.Base)
	}
	click(t, tt, "新拟态")
	if mallow.CurrentLook() != mallow.Neumorphism || mallow.Base != mallow.DarkPalette.Base || mallow.Canvas != mallow.Base {
		t.Errorf("back in neumorphism: base %v, canvas %v", mallow.Base, mallow.Canvas)
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

// The upright groove stands between the two parts of its row.
func TestGrooveV(t *testing.T) {
	_, tt := newTester(t)
	for _, s := range []string{"竖向凹槽左边", "竖向凹槽右边"} {
		if !tt.HasText(s) {
			t.Errorf("no %q beside the upright groove", s)
		}
	}
}

func TestListRowAndChip(t *testing.T) {
	g, tt := newTester(t)
	click(t, tt, "打开预约体检")
	if g.chosen != 2 {
		t.Errorf("chosen %d after opening the third row", g.chosen)
	}
	click(t, tt, "完成买菜")
	if !g.tasks[0] || g.chosen != 2 {
		t.Errorf("tasks %v, chosen %d after checking the first row", g.tasks, g.chosen)
	}
	click(t, tt, "生活")
	if g.chip != "生活" {
		t.Errorf("chip %q after choosing 生活", g.chip)
	}
}

func TestCalendar(t *testing.T) {
	g, tt := newTester(t)
	g.calMonth = time.Date(2026, 11, 23, 0, 0, 0, 0, time.Local)
	tt.Frame()
	click(t, tt, "11月30日")
	if g.calDate != "2026-11-30" {
		t.Errorf("date %q after picking the 30th", g.calDate)
	}
	click(t, tt, "下个月")
	if g.calMonth.Month() != time.December {
		t.Errorf("month %v after the next arrow", g.calMonth.Month())
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
	// MALLOW_SCREENSHOT_STYLE picks the look: neumorphism or organic.
	if s := os.Getenv("MALLOW_SCREENSHOT_STYLE"); s != "" {
		g.style = s
	}
	if os.Getenv("MALLOW_SCREENSHOT_SCALE") == "2" {
		tt.SetScale(2)
	}
	click(t, tt, "名字")
	tt.Type("Ada")
	// MALLOW_SCREENSHOT_CLICK lists what to click first, separated by
	// commas, as 打开对话框 to show the dialog.
	if s := os.Getenv("MALLOW_SCREENSHOT_CLICK"); s != "" {
		for _, label := range strings.Split(s, ",") {
			click(t, tt, label)
		}
	}
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

// The context menu opens on a right click and reports the entry chosen.
func TestContextMenu(t *testing.T) {
	g, tt := newTester(t)
	if tt.HasText("删掉它") {
		t.Fatal("the menu shows before it is asked for")
	}
	if err := tt.RightClick("在这里点右键"); err != nil {
		t.Fatal(err)
	}
	tt.Frame()
	click(t, tt, "删掉它")
	tt.Frame()
	if g.menuChoice != "删除" || g.menuOpen {
		t.Errorf("choice %q, open %v after choosing 删掉它", g.menuChoice, g.menuOpen)
	}
}

// A dialog shows over the page while it is open, with the keyboard in its
// field, and closes by Enter, by its buttons and by Escape.
func TestDialog(t *testing.T) {
	g, tt := newTester(t)
	if tt.HasText("新建相册") {
		t.Fatal("the dialog shows before it is asked for")
	}
	click(t, tt, "打开对话框")
	tt.Frame()
	if !tt.HasText("新建相册") {
		t.Fatalf("no dialog after 打开对话框: %q", tt.Texts())
	}
	// The field has the keyboard as the dialog shows.
	tt.Type("旅行")
	tt.Key(0, ui.KeyEnter)
	tt.Frame()
	if g.dialogOpen || g.dialogMade != "旅行" || tt.HasText("新建相册") {
		t.Fatalf("after typing a name and Enter: open %v, made %q", g.dialogOpen, g.dialogMade)
	}
	// A blank name makes nothing, and the dialog stays.
	click(t, tt, "打开对话框")
	tt.Frame()
	click(t, tt, "建好")
	if !g.dialogOpen || g.dialogMade != "旅行" {
		t.Fatalf("after 建好 without a name: open %v, made %q", g.dialogOpen, g.dialogMade)
	}
	// What is behind the dialog takes no click through it.
	before := g.clicks
	if err := tt.Click("普通按钮"); err == nil && g.clicks != before {
		t.Errorf("a button behind the dialog took a click")
	}
	g.dialogOpen = false
	tt.Frame()
	for name, leave := range map[string]func(){
		"取消":              func() { click(t, tt, "取消") },
		"Escape":          func() { tt.Key(0, ui.KeyEscape) },
		"a click outside": func() { tt.ClickAt(8, 400) },
	} {
		click(t, tt, "打开对话框")
		tt.Frame()
		tt.Type("不要的")
		leave()
		tt.Frame()
		if g.dialogOpen || g.dialogMade != "旅行" {
			t.Fatalf("after leaving the dialog by %s: open %v, made %q", name, g.dialogOpen, g.dialogMade)
		}
	}
}

// What shows over the window sets its lines as high as LineHeight says:
// its text inherits nothing from the window's content.
func TestLineHeightOver(t *testing.T) {
	defer func() { mallow.LineHeight = 0 }()
	g, tt := newTester(t)
	// The dialog's card, which its title names, is as high as its lines
	// make it: a title of 18 px, a field and a row of buttons of 14 px.
	height := func(lines float32) float32 {
		mallow.LineHeight = lines
		g.dialogOpen = true
		tt.Frame()
		card, ok := tt.Find("新建相册")
		if !ok {
			t.Fatalf("no dialog: %q", tt.Texts())
		}
		g.dialogOpen = false
		tt.Frame()
		return card.H
	}
	if two, three := height(2), height(3); three-two != 18+14+14 {
		t.Errorf("the dialog is %v high with lines twice their size, %v with lines three times: not 46 more", two, three)
	}
}

// A source list's entries choose, and show their notes.
func TestSourceList(t *testing.T) {
	g, tt := newTester(t)
	click(t, tt, "未整理")
	if g.source != "inbox" {
		t.Fatalf("source %q after choosing 未整理", g.source)
	}
	click(t, tt, "做完的")
	if g.source != "done" {
		t.Errorf("source %q after choosing 做完的", g.source)
	}
}

// The sidebar's entries choose the page; the recess follows the choice.
func TestSidebar(t *testing.T) {
	g, tt := newTester(t)
	click(t, tt, "表单")
	if g.nav != "表单" {
		t.Fatalf("nav %q after choosing 表单", g.nav)
	}
	click(t, tt, "设置")
	click(t, tt, "组件")
	if g.nav != "组件" {
		t.Errorf("nav %q after choosing 组件", g.nav)
	}
	// The entries at the bottom choose into the same selection.
	click(t, tt, "关于")
	if g.nav != "关于" {
		t.Errorf("nav %q after choosing 关于 at the bottom", g.nav)
	}
}

// TestSidebarFrames writes frames of the recess sliding from 组件 to 设置
// to MALLOW_FRAMES_DIR, to look at the motion.
func TestSidebarFrames(t *testing.T) { frames(t, "sidebar", "设置") }

// TestSegmentedFrames writes frames of the thumb sliding from 日 past 周 to
// 月 to MALLOW_FRAMES_DIR.
func TestSegmentedFrames(t *testing.T) {
	frames(t, "segmented", "日", "月")
}

// frames clicks all but the last of labels, lets the view settle, clicks the
// last and writes frames of what follows to MALLOW_FRAMES_DIR.
func frames(t *testing.T, name string, labels ...string) {
	dir := os.Getenv("MALLOW_FRAMES_DIR")
	if dir == "" {
		t.Skip("set MALLOW_FRAMES_DIR to write frames")
	}
	_, tt := newTester(t)
	for _, l := range labels[:len(labels)-1] {
		click(t, tt, l)
	}
	tt.Advance(time.Second)
	click(t, tt, labels[len(labels)-1])
	prev := 0
	for _, ms := range []int{0, 40, 80, 120, 160, 220, 300, 500} {
		tt.Advance(time.Duration(ms-prev) * time.Millisecond)
		prev = ms
		f, err := os.Create(fmt.Sprintf("%s/%s-%03dms.png", dir, name, ms))
		if err != nil {
			t.Fatal(err)
		}
		png.Encode(f, tt.Image())
		f.Close()
	}
}
