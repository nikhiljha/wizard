package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Create Main Window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(250, 200)
	window.SetWindowTitle("Wizard")

	// TODO: Implement File Drops
	// window.SetAcceptDrops(true)

	// Use a QVBox
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	// File Selector
	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Select a file ...")
	widget.Layout().AddWidget(input)

	pickfile := widgets.NewQPushButton2("Browse", nil)
	pickfile.ConnectClicked(func(bool) {
		path := widgets.QFileDialog_GetOpenFileName(nil, "Browse...", "~", "", "", 0)
		input.SetText(path)
	})
	widget.Layout().AddWidget(pickfile)

	// Send File Button
	button := widgets.NewQPushButton2("Send", nil)
	button.ConnectClicked(func(bool) {
		code := send(input.Text())
		widgets.QMessageBox_Information(nil, "OK", "The code is... "+code, widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})
	widget.Layout().AddWidget(button)

	// Show the window and start the app.
	window.Show()
	app.Exec()
}
