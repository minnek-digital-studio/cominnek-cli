package helper

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/inancgumus/screen"
)

func PrintName() {
	screen.Clear()
	screen.MoveTopLeft()

	myFigure := figure.NewColorFigure("COMINNEK", "", "yellow", true)
	myFigure.Print()

	print("\n")
}
