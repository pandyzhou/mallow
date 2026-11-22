# mallow

**在线展示：<https://pandyzhou.github.io/mallow/>**（组件编译成 WebAssembly 在浏览器里运行，可以直接点、直接切换浅色和深色）

[MyGo](https://github.com/egoist/mygo) 原生 UI 的**新拟态（Neumorphism / Soft UI）组件库**，纯 Go 编写，不依赖 WebView。

所有元素都与所在的表面同色，光从左上方照来：凸起的元素上方是高光、下方是阴影，凹陷的元素则把这对阴影收进内侧。没有边框，大圆角，只用一种强调色。自带浅色、深色两套配色，可以跟随系统切换。

每个组件的外观都由 mallow 自己绘制。交互基于 MyGo 的无样式底座（`ButtonBase`、`CheckboxBase`、`TextInputBase` 等），由底座负责鼠标、键盘、焦点、输入法和无障碍，不使用 MyGo 自带的成品控件。

## 组件

| 分类 | 组件 |
|---|---|
| 主题 | `Use` / `UseMode` / `Theme`，`LightPalette` / `DarkPalette`，当前配色 `Base`、`Text`、`Muted`、`Accent`… |
| 表面 | `Raised` 凸起、`Recessed` 凹陷、`Card` 卡片、`Tray` 托盘、`Groove` 凹槽分隔线、`Soft` 过渡、`Label` 小标题 |
| 按钮 | `Button`、`PrimaryButton`、`IconButton`、`GhostIconButton`，以及 `Glyph*` 图形 |
| 选择 | `Check` 勾选、`CheckDisc`、`Switch` 开关、`Segmented` 分段选择 |
| 输入 | `TextInput` |
| 导航 | `NavItem` |
| 浮层 | `Tooltip` |
| 窗口 | `TitleBar` 自绘标题栏、`WindowButtons` 窗口按钮 |

## 使用

```sh
go get github.com/pandyzhou/mallow
```

```go
func (a *app) view(c *ui.Context) {
	// 每帧开始时应用外观（跟随系统 / 浅色 / 深色），并设置窗口主题。
	mallow.UseMode(c, mallow.ModeSystem)

	mallow.Card(c, mallow.Large).Padding(24).Gap(16).Children(func() {
		ui.Text(c, "设置").FontSize(20).FontWeight(600).TextColor(mallow.Text)
		mallow.Switch(c, &a.sync, "自动同步")
		mallow.TextInput(c, &a.name).Placeholder("你的名字").Label("名字")
		if mallow.PrimaryButton(c, "保存").Clicked() {
			a.save()
		}
	})
}
```

用自绘标题栏时，窗口要隐藏系统标题栏和按钮：

```go
win := mygo.NewWindow(mygo.WindowOptions{
	TitleBarStyle:     mygo.TitleBarHidden,
	HideWindowButtons: true,
	Content:           ui.View(a.view),
})
// 视图里：mallow.TitleBar(c, win, nil)
```

## 组件展示

```sh
go run ./cmd/gallery                            # 桌面窗口
go tool mygo dev -web cmd/gallery               # 浏览器里运行，改代码自动刷新
go tool mygo build -platform web cmd/gallery    # 编译成静态网站，输出到 cmd/gallery/build/web
```

推送到 main 后，GitHub Actions 会把展示页编译成网页版，部署到 GitHub Pages。

不开窗口，把展示页渲染成图片：

```sh
MALLOW_SCREENSHOT=shot.png MALLOW_SCREENSHOT_THEME=dark go test -run Screenshot ./cmd/gallery
```

## 设计参考

[stylekit · Neumorphism](https://www.stylekit.top/zh/styles/neumorphism)。浅色配色取自该规范，深色配色是按同样的光影规则推出来的。
