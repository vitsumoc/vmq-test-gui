// 存储并初始化所有页面, 为 app 提供管理入口
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var mainWindow fyne.Window

// ui库结构
type suis struct {
	connectList *connectList
	newConnect  *newConnect
}

// ui库实例
var uis = suis{}

// UI初始化
func Init(a fyne.App) {
	// 主窗口初始化
	mainWindow = a.NewWindow("vmqtool")
	mainWindow.SetMaster()

	// 所有页面
	uis.connectList = newConnectList()
	uis.newConnect = newNewConnect()

	// 左右布局, 左边是连接列表, 右边是点击连接后的详情
	split := container.NewHSplit(uis.connectList.content, widget.NewLabel("right"))
	split.Offset = 0.2
	mainWindow.SetContent(split)

	mainWindow.Resize(fyne.NewSize(1280, 720))
	mainWindow.ShowAndRun()
}

// 弹窗总线 弹窗名 参数
func diaBus(name string, args ...interface{}) {
	if name == "newConnect" {
		uis.newConnect.pop()
	}
}

// 向UI下发事件总线 通过 数据 改变ui的 使用该总线直接下发

// UI向上回调总线 通过 ui 影响数据的 使用 app 注册的回调函数实现
