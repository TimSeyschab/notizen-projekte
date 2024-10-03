package main

import (
	"graphPlotter/plotter"
	"graphPlotter/utils"
	"log"
)

func main() {
	params := utils.ParseCLI()

	plot, error := plotter.GetPlotter(params.FunctionType, params.Params)
	if error != nil {
		log.Fatal(error.Error())
	}

	error = plotter.PlotFunc(plot)
	if error != nil {
		log.Fatal(error.Error())
	}
}
