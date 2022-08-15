package main

import (
  "os/exec"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

  content := widget.NewButton("click me", func() {
    cmd := exec.Command("mugen.exe", "-p1", "Street Fighter/Ryu (CVS)/Ryu.def", "-p2", "Street Fighter/Dhalsim (CVS2)/Dhalsim.def")
    _, err := cmd.CombinedOutput()
    if err != nil {
      panic(err)
    }
	})

  w.SetFullScreen(true)
	w.SetContent(content)
	w.ShowAndRun()
}
