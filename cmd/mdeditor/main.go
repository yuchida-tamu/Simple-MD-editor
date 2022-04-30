package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const WIDTH = 920
const HEIGHT = 1080

func main() {

	a := app.New()
	window := a.NewWindow("Simple Markdown Editor")

	str := binding.NewString()
	str.Set("")

	// preview field
	richText := widget.NewRichTextFromMarkdown("")
	richText.Resize(fyne.NewSize(WIDTH/2, HEIGHT))
	richText.Move(fyne.NewPos(WIDTH/2, 0))

	// editor area (input field)
	entry := widget.NewEntryWithData(str)
	entry.MultiLine = true
	entry.TextStyle.Monospace = true
	entry.Resize(fyne.NewSize(WIDTH/2, HEIGHT))
	entry.OnChanged = func(s string) {
		richText.ParseMarkdown(s)
	}

	// container
	content := container.NewWithoutLayout(
		entry,
		richText,
	)
	content.Resize(fyne.NewSize(WIDTH, HEIGHT))
	content.Add(layout.NewSpacer())

	window.SetContent(content)
	window.Resize(fyne.NewSize(WIDTH, HEIGHT))

	window.ShowAndRun()

}
