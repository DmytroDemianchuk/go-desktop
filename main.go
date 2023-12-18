package main

import (
	"fmt"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/xuri/excelize/v2"
)

func createTable(data [][]string) fyne.CanvasObject {
	rows := len(data)
	cols := 0
	if rows > 0 {
		cols = len(data[0])
	}

	table := widget.NewTable(
		func() (int, int) {
			return rows, cols
		},
		func() fyne.CanvasObject {
			label := widget.NewLabel("")
			return label
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Row < len(data) && i.Col < len(data[i.Row]) {
				label, ok := o.(*widget.Label)
				if ok {
					label.SetText(data[i.Row][i.Col])
				}
			}
		},
	)

	for i := 0; i < cols; i++ {
		table.SetColumnWidth(i, 150)
	}
	for i := 0; i < rows; i++ {
		table.SetRowHeight(i, 30)
	}
	table.Refresh()

	scrollableTable := container.NewScroll(table)

	return scrollableTable
}

func main() {
	myApp := app.New()

	myWindow := myApp.NewWindow("Resizable & Movable Table")
	myWindow.Resize(fyne.NewSize(800, 600))

	updateContent := func(data [][]string) {
		tableContent := createTable(data)
		tableWindow := myApp.NewWindow("Table Content")
		tableWindow.Resize(fyne.NewSize(800, 600))
		tableWindow.SetContent(tableContent)
		tableWindow.Show()
	}

	openFile := func(reader io.ReadCloser) {
		f, err := excelize.OpenReader(reader)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		rows, err := f.GetRows(f.GetSheetList()[0])
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		updateContent(rows)
	}

	openFileButton := widget.NewButton("Open File 1", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil || reader == nil {
				fmt.Println("Error:", err)
				return
			}
			defer reader.Close()

			openFile(reader)
		}, myWindow)
	})

	openFileButton2 := widget.NewButton("Open File 2", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil || reader == nil {
				fmt.Println("Error:", err)
				return
			}
			defer reader.Close()

			openFile(reader)
		}, myWindow)
	})

	initialContent := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		openFileButton,
		openFileButton2,
		widget.NewLabel("Select an XLSX file to open..."),
	)
	myWindow.SetContent(initialContent)

	myWindow.ShowAndRun()
}
