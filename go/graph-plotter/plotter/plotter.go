package plotter

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"golang.org/x/term"
)

type Plotter interface {
	Plot(x int) float64
}

const (
	scaleX = 1
	scaleY = 1
)

func PlotFunc(p Plotter) error {
	fd := int(os.Stdout.Fd())
	width, height, error := term.GetSize(fd)
	if error != nil {
		return fmt.Errorf("Cant get terminal for function plot: %w", error)
	}
	centerX, centerY := width/2, height/2

	canvas := createCanvas(height, width)

	drawAxis(width, centerX, canvas, centerY, height)

	plotFunction(centerX, p, centerY, width, height, canvas)

	printCanvas(canvas)

	return nil
}

func plotFunction(centerX int, p Plotter, centerY int, width int, height int, canvas [][]rune) {
	for x := -centerX / scaleX; x < centerX/scaleX; x++ {
		y := p.Plot(x)
		screenY := centerY - int(math.Round(y*scaleY))
		screenX := centerX + x*scaleX

		if screenX >= 0 && screenX < width && screenY >= 0 && screenY < height {
			canvas[screenY][screenX] = '*'
		}
	}
}

func createCanvas(height int, width int) [][]rune {
	canvas := make([][]rune, height)
	for i := range canvas {
		canvas[i] = make([]rune, width)
		for j := range canvas[i] {
			canvas[i][j] = ' '
		}
	}
	return canvas
}

func drawAxis(width int, centerX int, canvas [][]rune, centerY int, height int) {
	for x := 0; x < width; x++ {
		if (x-centerX)%5 == 0 {
			canvas[centerY][x] = '+'
		} else {
			canvas[centerY][x] = '-'
		}
	}
	for y := 0; y < height; y++ {
		if (y-centerY)%5 == 0 {
			canvas[y][centerX] = '+'
		} else {
			canvas[y][centerX] = '|'
		}
	}
	canvas[centerY][centerX] = '+'
}

func printCanvas(canvas [][]rune) {
	writer := bufio.NewWriter(os.Stdout)

	for i, row := range canvas {
		for _, cell := range row {
			fmt.Fprintf(writer, "%c", cell)
		}
		if i < len(canvas)-1 {
			writer.WriteByte('\n')
		}
	}

	writer.Flush()

	var buf [1]byte
	os.Stdin.Read(buf[:])
}

func GetPlotter(funcType string, params []float64) (Plotter, error) {
	switch funcType {
	case "linear":
		if len(params) < 2 {
			return nil, fmt.Errorf("linear function requires 2 parameters: slope and intercept")
		}
		return Linear{Slope: params[0], Intercept: params[1]}, nil
	case "quadratic":
		if len(params) < 3 {
			return nil, fmt.Errorf("quadratic function requires 3 parameters: a, b, and c")
		}
		return Quadratic{A: params[0], B: params[1], C: params[2]}, nil
	case "cubic":
		if len(params) < 4 {
			return nil, fmt.Errorf("cubic function requires 4 parameters: a, b, c, and d")
		}
		return Cubic{A: params[0], B: params[1], C: params[2], D: params[3]}, nil

	default:
		return nil, fmt.Errorf("unknown function type: %s", funcType)
	}
}
