package main

import "fmt"

func main() {
	label := Label{Widget{100, 200}, "name"}
	button1 := Button{
		Label{
			Widget{
				X: 10,
				Y: 70,
			},
			"OK",
		},
	}

	button2 := Button{Label{Widget{50, 70}, "Cancel"}}
	listBox := ListBox{Widget{10, 40}, []string{"AL", "AK", "AZ", "AR"}, 0}

	for _, painter := range []Painter{label, listBox, button1, button2} {
		painter.Paint()
	}

	for _, widget := range []interface{}{label, listBox, button1, button2} {
		widget.(Painter).Paint()

		if clicker, ok := widget.(Clicker); ok {
			clicker.Click()
		}
	}

	fmt.Println()
}
