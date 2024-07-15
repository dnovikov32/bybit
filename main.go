package main

import (
	"bybit/internal/bybit"
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file")
	}

	application := app.New()
	window := application.NewWindow("Hello window")
	window.Resize(fyne.NewSize(800, 1000))

	/* clock := widget.NewLabel("") */

	/* 	go func() {
		for range time.Tick(time.Second) {
			fmt.Println("Tick")
			updateTime(clock)
		}
	}() */

	client := bybit.NewTestClient()
	params := &bybit.OrderBookRequest{Category: "spot", Symbol: "ETHUSDT", Limit: 10}

	orderBookList, err := bybit.GetOrderBook(client, params)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("orderBookList", orderBookList)

	table := GetOrderBookTable(orderBookList)

	window.SetContent(table)
	window.ShowAndRun()
}

func GetOrderBookTable(items []bybit.OrderBookItem) *widget.Table {
	return widget.NewTable(
		func() (int, int) {
			return len(items), 2
		},
		func() fyne.CanvasObject {
			bg := new(canvas.Rectangle)
			bg.SetMinSize(fyne.NewSize(70, 20))

			text := canvas.NewText("", color.White)
			text.TextSize = 12

			return container.NewStack(bg, text)

		},
		func(id widget.TableCellID, o fyne.CanvasObject) {
			container := o.(*fyne.Container)
			bg := container.Objects[0].(*canvas.Rectangle)
			text := container.Objects[1].(*canvas.Text)
			bidColor := color.NRGBA{R: 101, G: 27, B: 27, A: 255}
			askColor := color.NRGBA{R: 47, G: 76, B: 101, A: 255}
			bidBestColor := color.NRGBA{R: 176, G: 31, B: 33, A: 255}
			askBestColor := color.NRGBA{R: 71, G: 127, B: 180, A: 255}

			item := items[id.Row]

			if item.IsBid {
				if item.IsBestPrice {
					bg.FillColor = bidBestColor
				} else {
					bg.FillColor = bidColor
				}
			} else {
				if item.IsBestPrice {
					bg.FillColor = askBestColor
				} else {
					bg.FillColor = askColor
				}
			}

			switch id.Col {
			case 0:
				text.Text = item.Size
			case 1:
				text.Text = item.Price
			}
		},
	)
}
