package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)
func main() {
	var w *walk.MainWindow;
	var b *walk.PushButton;
	MainWindow{
		AssignTo:&w,
		Title:   "TimeManager",
		Layout:Flow{},
		MinSize: Size{600, 400},

		Children: []Widget{
			PushButton{
				AssignTo:&b,
				MaxSize:Size{50,20},
				Text:"弹出",
				OnClicked: func() {
					walk.MsgBox(w,"提示","一个提示消息", walk.MsgBoxIconInformation)
				},
			},
		},
	}.Run()
}