package main

import (
	"fmt"
)

var currentRowNum = 10
var lines []string

func writeLine(line string) {
	lines = append(lines, fmt.Sprintf("N%04d %s", currentRowNum, line))
	currentRowNum = currentRowNum + 10
}

func main() {

	writeLine("(Cut test v1 for PS1 CNC plasma machine - HEIGHTS)")
	writeLine("G20")
	writeLine("F1")
	writeLine("G53 G90 G40") // g90 absolute positioning
	writeLine("M666")

	initialY := 0.0
	incrementY := 0.2468

	// Speed in DECIMAL!
	fixedF := 300.0
	zHeight := 0.10

	for i := 0; i < 10; i++ {
		writeLine(fmt.Sprintf("(Pass %d - Y is %.4f and IPM is %.2f and Z is %.2f)", (i + 1), initialY, fixedF, zHeight))

		writeLine(fmt.Sprintf("G00 X0.0000 Y%.4f", initialY))

		writeLine("G31 Z -100 F19.685")
		writeLine("G92 Z0.0")
		writeLine(fmt.Sprintf("G00 Z%.4f (Z OFFSET FROM WORK)", zHeight)) // lift up a little from the table...also something to look into
		writeLine("M03 (TORCH ON)")                                       // This turns the torch on
		writeLine("G04 P0.3")
		writeLine("M667")

		writeLine(fmt.Sprintf("G01 X3.7300 F%.2f", fixedF))
		writeLine("M04 (TORCH OFF)") // Torch off

		// move up
		writeLine("G00 Z1.5000")

		initialY += incrementY
		zHeight += .01

	}

	// Put us back at 0,0
	writeLine("X0.0000 Y0.0000") // this is part of the G00 above

	// and finish
	writeLine("M04 M30")

	for _, element := range lines {
		fmt.Printf("%s\n", element)
	}
}
