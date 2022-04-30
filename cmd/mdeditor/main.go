package main

import (
	"bufio"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const WIDTH = 920
const HEIGHT = 700
const PATH = "data"

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
	input := widget.NewEntryWithData(str)
	input.MultiLine = true
	input.TextStyle.Monospace = true
	input.Resize(fyne.NewSize(WIDTH/2, HEIGHT))
	input.OnChanged = func(s string) {
		richText.ParseMarkdown(s)
	}

	// tool bar
	toolBar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			runPopup(window, input.Text)
		}),
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {}),
	)

	// container
	editor := container.NewWithoutLayout(
		input,
		richText,
	)

	editor.Resize(fyne.NewSize(WIDTH, HEIGHT))
	editor.Add(layout.NewSpacer())

	content := container.NewBorder(toolBar, nil, nil, nil, editor)

	window.SetContent(content)
	window.Resize(fyne.NewSize(WIDTH, HEIGHT))

	window.ShowAndRun()

}

func saveToFile(text string, filename string) {
	// check if the directory for data files exists
	if _, err := os.Stat(PATH); err != nil {
		// if not, create the directory
		os.Mkdir(PATH, 0700)
	}

	// create file
	file, err := os.Create(fmt.Sprintf("%s/%s.md", PATH, filename))
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer file.Close()
	//create buffers
	writer := bufio.NewWriter(file)
	result, err := writer.WriteString(text)
	if err != nil {
		fmt.Println("Failed to save text to a file")
		return
	}

	fmt.Println("Successfully saved to a file: ", text)
	fmt.Println(result)
	// free up the buffers
	writer.Flush()
}

func runPopup(w fyne.Window, text string) (modal *widget.PopUp) {

	fileNameEntry := widget.NewEntry()

	modal = widget.NewModalPopUp(
		container.NewVBox(
			widget.NewLabel("Save As ..."),
			fileNameEntry,
			container.NewHBox(
				widget.NewButton("Save", func() {
					saveToFile(text, fileNameEntry.Text)
					modal.Hide()
				}),
				widget.NewButton("Cancel", func() { modal.Hide() }),
			),
		),
		w.Canvas(),
	)

	modal.Resize(fyne.NewSize(300, 300))

	modal.Show()
	return modal
}
