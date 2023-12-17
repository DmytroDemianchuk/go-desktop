package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/xuri/excelize/v2"
)

func main() {
	myApp := app.New()

	// Create a window
	myWindow := myApp.NewWindow("XLSX Reader")

	// Create a label to display the content
	contentLabel := widget.NewLabel("Select an XLSX file to display its content")

	// Create a button to open the file dialog
	openButton := widget.NewButton("Open File", func() {
		fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				fmt.Println(err)
				return
			}
			if reader == nil {
				return
			}
			defer reader.Close()

			// Read the content of the selected XLSX file
			f, err := excelize.OpenReader(reader)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Get all the rows in the first sheet
			rows, err := f.GetRows(f.GetSheetList()[0])
			if err != nil {
				fmt.Println(err)
				return
			}

			// Display the content in the label
			var content string
			for _, row := range rows {
				content += fmt.Sprintf("%v\n", row)
			}
			contentLabel.SetText(content)
		}, myWindow)

		// Show the file dialog
		fileDialog.Show()
	})

	// Create a container to hold the widgets
	content := container.NewVBox(
		openButton,
		contentLabel,
	)

	// Set the content of the window
	myWindow.SetContent(content)

	// Show and run the application
	myWindow.ShowAndRun()
}
