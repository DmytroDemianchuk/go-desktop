# Simple Go Desktop App

## Description

Simple application using Fyne to create a GUI interface for opening and displaying the content of an XLSX file in a table format

- Application Initialization: The code initializes a Fyne application and a window titled "Resizable & Movable Table".

- Table Creation: The createTable function generates a Fyne table based on the provided XLSX data, adjusting the column widths and row heights.

- Content Update: The updateContent function creates a new window to display the table content obtained from the XLSX file.

- Open File Functionality: Two buttons (Open File 1 and Open File 2) are provided to select XLSX files. These buttons utilize a file dialog (dialog.ShowFileOpen) to allow users to pick an XLSX file.

- File Reading and Display: When a file is chosen, the code reads the content of the selected XLSX file and updates the content in a new window using the updateContent function.

- Initial Content: The initial content of the main window includes the two "Open File" buttons and a label prompting the user to select an XLSX file.

Overall, this code demonstrates a simple Fyne-based application to open XLSX files and display their content in a resizable and movable table format.

## Prerequisites
- go 1.20
- fyne.io/fyne
- github.com/xuri/excelize/v2

## Start

Use `go mod tidy` and `go run main.go`

to build&run project

```
go mod tidy
```

```
go run main.go
```

## Screenshots
#### The main window in which we open files
![image](../main/assets/desktop_image.png)
#### Windows with open files
![image](../main/assets/excelFiles_image.png)