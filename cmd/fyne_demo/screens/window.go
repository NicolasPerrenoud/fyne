package screens

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func confirmCallback(response bool) {
	fmt.Println("Responded with", response)
}

func fileOpened(f fyne.File) {
	if f == nil {
		log.Println("Cancelled")
		return
	}

	if f.URI()[len(f.URI())-4:] == ".png" {
		showImage(f)
	}
}

func fileSaved(f fyne.File) {
	if f == nil {
		log.Println("Cancelled")
		return
	}

	log.Println("Save to...", f.URI())
}

func loadImage(f fyne.File) *canvas.Image {
	in, err := f.Open()
	if err != nil {
		fyne.LogError("Could not open image stream", err)
		return nil
	}

	data, err := ioutil.ReadAll(in)
	if err != nil {
		fyne.LogError("Failed to load image data", err)
		return nil
	}
	res := fyne.NewStaticResource(f.Name(), data)

	return canvas.NewImageFromResource(res)
}

func showImage(f fyne.File) {
	img := loadImage(f)
	if img == nil {
		return
	}
	img.FillMode = canvas.ImageFillOriginal

	w := fyne.CurrentApp().NewWindow(f.Name())
	w.SetContent(widget.NewScrollContainer(img))
	w.Resize(fyne.NewSize(320, 240))
	w.Show()
}

// DialogScreen loads a panel that lists the dialog windows that can be tested.
func DialogScreen(win fyne.Window) fyne.CanvasObject {
	dialogs := widget.NewGroup("Dialogs",
		widget.NewButton("Info", func() {
			dialog.ShowInformation("Information", "You should know this thing...", win)
		}),
		widget.NewButton("Error", func() {
			err := errors.New("A dummy error message")
			dialog.ShowError(err, win)
		}),
		widget.NewButton("Confirm", func() {
			cnf := dialog.NewConfirm("Confirmation", "Are you enjoying this demo?", confirmCallback, win)
			cnf.SetDismissText("Nah")
			cnf.SetConfirmText("Oh Yes!")
			cnf.Show()
		}),
		widget.NewButton("Progress", func() {
			prog := dialog.NewProgress("MyProgress", "Nearly there...", win)

			go func() {
				num := 0.0
				for num < 1.0 {
					time.Sleep(50 * time.Millisecond)
					prog.SetValue(num)
					num += 0.01
				}

				prog.SetValue(1)
				prog.Hide()
			}()

			prog.Show()
		}),
		widget.NewButton("ProgressInfinite", func() {
			prog := dialog.NewProgressInfinite("MyProgress", "Closes after 5 seconds...", win)

			go func() {
				time.Sleep(time.Second * 5)
				prog.Hide()
			}()

			prog.Show()
		}),
		widget.NewButton("File Open", func() {
			dialog.ShowFileOpen(fileOpened, win)
		}),
		widget.NewButton("File Save", func() {
			dialog.ShowFileSave(fileSaved, win)
		}),
		widget.NewButton("Custom", func() {
			entry := widget.NewEntry()
			entry.SetPlaceHolder("Type something here")
			entry.OnChanged = func(text string) {
				fmt.Println("Entered", text)
			}
			sel := widget.NewSelect([]string{"Option A", "Option B", "Option C"}, func(o string) {
				fmt.Println("Selected", o)
			})
			content := widget.NewVBox(entry, sel)
			dialog.ShowCustom("Custom dialog", "Done", content, win)
		}),
	)

	windowGroup := widget.NewGroup("Windows",
		widget.NewButton("New window", func() {
			w := fyne.CurrentApp().NewWindow("Hello")
			w.SetContent(widget.NewLabel("Hello World!"))
			w.Show()
		}),
		widget.NewButton("Fixed size window", func() {
			w := fyne.CurrentApp().NewWindow("Fixed")
			w.SetContent(fyne.NewContainerWithLayout(layout.NewCenterLayout(), widget.NewLabel("Hello World!")))

			w.Resize(fyne.NewSize(240, 180))
			w.SetFixedSize(true)
			w.Show()
		}),
		widget.NewButton("Centered window", func() {
			w := fyne.CurrentApp().NewWindow("Central")
			w.SetContent(fyne.NewContainerWithLayout(layout.NewCenterLayout(), widget.NewLabel("Hello World!")))

			w.CenterOnScreen()
			w.Show()
		}))

	drv := fyne.CurrentApp().Driver()
	if drv, ok := drv.(desktop.Driver); ok {
		windowGroup.Append(
			widget.NewButton("Splash Window (only use on start)", func() {
				w := drv.CreateSplashWindow()
				w.SetContent(widget.NewLabelWithStyle("Hello World!\n\nMake a splash!",
					fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
				w.Show()

				go func() {
					time.Sleep(time.Second * 3)
					w.Close()
				}()
			}))
	}
	windows := widget.NewVBox(dialogs, windowGroup)

	return fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(2), windows, LayoutPanel())
}
