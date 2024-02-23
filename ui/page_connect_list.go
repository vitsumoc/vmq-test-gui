// 连接管理列表
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type connectList struct {
	// 顶层布局
	content *fyne.Container
	// 交互部分
	btnAdd *widget.Button // 添加连接

	// 数据部分
	connList *widget.List    // 列表容器
	listData []*connectLabel // 数据内容
}

func newConnectList() *connectList {
	cl := &connectList{}

	cl.listData = []*connectLabel{
		NewConnectLabel("192.168.1.1:1883"),
		NewConnectLabel("192.168.1.1:1883"),
		NewConnectLabel("192.168.1.1:1883"),
	}
	cl.connList = widget.NewList(
		func() int {
			return len(cl.listData)
		},
		func() fyne.CanvasObject {
			return NewConnectLabel("192.168.1.1:1883")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*connectLabel).update()
		},
	)
	cl.btnAdd = widget.NewButton("创建连接", func() { diaBus("newConnect") })
	cl.content = container.NewBorder(cl.btnAdd, nil, nil, nil, cl.connList)

	return cl
}
