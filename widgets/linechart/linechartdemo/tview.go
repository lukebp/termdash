package main

import (
	tcell "github.com/gdamore/tcell/v2"
	"github.com/lukebp/termdash/container"
	tcellterm "github.com/lukebp/termdash/terminal/tcell"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/rivo/tview"
)

func tviewExample() error {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	menu := newPrimitive("Menu")
	sideBar := newPrimitive("Side Bar")

	main := tview.NewTextView().
		SetTextAlign(tview.AlignCenter)

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, false).
		AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(menu, 0, 0, 0, 0, 0, 0, false).
		AddItem(main, 1, 0, 1, 3, 0, 0, false).
		AddItem(sideBar, 0, 0, 0, 0, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(menu, 1, 0, 1, 1, 0, 100, false).
		AddItem(main, 1, 1, 1, 1, 0, 100, false).
		AddItem(sideBar, 1, 2, 1, 1, 0, 100, false)

	timeSeries := []float64{1, 5, 13, 4, 15, 26, 21, 20, 21, 25, 30}

	main.Box.SetDrawFunc(func(screen tcell.Screen,
		x int, y int, width int, height int) (int, int, int, int) {

		t, err := tcellterm.NewWithScreen(screen)
		if err != nil {
			panic(err)
		}

		lc, err := linechart.New(
			linechart.AxesCellOpts(cell.FgColor(cell.ColorRed)),
			linechart.YLabelCellOpts(cell.FgColor(cell.ColorGreen)),
			linechart.YAxisAdaptive(),
			linechart.XLabelCellOpts(cell.FgColor(cell.ColorCyan)),
		)
		if err != nil {
			panic(err)
		}

		err = lc.Series("randomData", timeSeries,
			linechart.SeriesCellOpts(cell.FgColor(cell.ColorNumber(33))),
		)
		if err != nil {
			panic(err)
		}

		c, err := container.New(
			t,
			container.Border(linestyle.None),
			container.PlaceWidget(lc),
		)
		if err != nil {
			panic(err)
		}

		err = c.DrawInside(x, y, width, height)
		if err != nil {
			panic(err)
		}

		// Space for other content
		return x, y, width, height
	})

	return tview.NewApplication().SetRoot(grid, true).SetFocus(grid).Run()
}
