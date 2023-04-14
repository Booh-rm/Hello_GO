package main

import (
	"fmt"
	"strconv"

	"github.com/fyne-io/fyne"
	"github.com/fyne-io/fyne/app"
	"github.com/fyne-io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculadora")

	// Crear widgets
	num1Entry := widget.NewEntry()
	num2Entry := widget.NewEntry()
	operatorEntry := widget.NewEntry()
	resultLabel := widget.NewLabel("Resultado: ")

	calculateButton := widget.NewButton("Calcular", func() {
		num1Str := num1Entry.Text
		num2Str := num2Entry.Text
		operator := operatorEntry.Text

		num1, err1 := strconv.ParseFloat(num1Str, 64)
		num2, err2 := strconv.ParseFloat(num2Str, 64)

		if err1 != nil || err2 != nil {
			resultLabel.SetText("Error: Entrada inválida")
			return
		}

		switch operator {
		case "+":
			resultLabel.SetText(fmt.Sprintf("Resultado: %.2f", num1+num2))
		case "-":
			resultLabel.SetText(fmt.Sprintf("Resultado: %.2f", num1-num2))
		case "*":
			resultLabel.SetText(fmt.Sprintf("Resultado: %.2f", num1*num2))
		case "/":
			if num2 == 0 {
				resultLabel.SetText("Error: No se puede dividir por cero")
			} else {
				resultLabel.SetText(fmt.Sprintf("Resultado: %.2f", num1/num2))
			}
		default:
			resultLabel.SetText("Error: Operador inválido")
		}
	})

	// Crear layout
	content := fyne.NewContainerWithLayout(
		fyne.NewGridLayout(2),
		widget.NewLabel("Número 1: "), num1Entry,
		widget.NewLabel("Número 2: "), num2Entry,
		widget.NewLabel("Operador: "), operatorEntry,
		widget.NewLabel(""), calculateButton,
		widget.NewLabel(""), resultLabel,
	)

	w.SetContent(content)
	w.ShowAndRun()
}
