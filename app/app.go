package app

import (
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/vitsumoc/vmq-tool/ui"
)

func Run() {
	// 主窗口
	a := fyneApp.NewWithID("vitsumoc.github.vmqtool")
	w := a.NewWindow("vmqtool")
	w.SetMaster()

	label2 := widget.NewLabel("right")

	// 左右布局, 左边是连接列表, 右边是点击连接后的详情
	split := container.NewHSplit(ui.NewConnectList(), label2)
	split.Offset = 0.2
	w.SetContent(split)

	w.Resize(fyne.NewSize(1280, 720))
	w.ShowAndRun()
}
