package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var btn = widgets.NewQPushButton2("Hello World", nil)
	btn.Resize2(180, 44)
	btn.ConnectClicked(func(flag bool) {
		widgets.QMessageBox_Information(nil, "OK", "Hello, world!", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})

	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Hello World Example")
	window.Layout().AddWidget(btn)
	window.Show()

	widgets.QApplication_Exec()
}
