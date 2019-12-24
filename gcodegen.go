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

	writeLine("(Cut test v1 for PS1 CNC plasma machine)")
	writeLine("G20")
	writeLine("F1")
	writeLine("G53 G90 G40") // g90 absolute positioning
	writeLine("M666")

	initialY := 0.0
	incrementY := 0.2468

	// Speed in DECIMAL!
	writeLine("(18 GAUGE)")
	initialF := 300.0 // THIS IS WHAT YOU WANT TO CHANGE TO AN INITIAL VALUE!!!
	incrementF := 5.0

	numPasses := 10
	for i := 0; i < numPasses; i++ {
		writeLine(fmt.Sprintf("(Pass %d - Y is %.4f and IPM is %.2f)", (i + 1), initialY, initialF))

		writeLine(fmt.Sprintf("G00 X0.0000 Y%.4f", initialY))

		writeLine("G31 Z -100 F19.685")
		writeLine("G92 Z0.0 (Set Z 0)")                  // Set Z 0
		writeLine("G00 Z0.1000 (Lift up a little)")      // lift up a little from the table...also something to look into
		writeLine("G92 Z0.0 (Set Z to 0 again)")         // Set Z 0 again
		writeLine("G00 Z0.1500 (Move to pierce height)") // Move to pierce height
		writeLine("M667 (THC on)")                       // THC on
		writeLine("M03 (TORCH ON)")                      // This turns the torch on
		writeLine("G04 P0.0 (Pierce Delay)")             // Pierce delay
		writeLine("G00 Z0.0600 (Move to cut height)")    // Move to cut height

		// Now start moving
		writeLine(fmt.Sprintf("G01 X3.7300 F%.2f (Do the move at F%.2f)", initialF, initialF))
		writeLine("M04 (TORCH OFF)") // Torch off
		writeLine("M666 (THC OFF)")  // thc off

		// move up
		writeLine("G00 Z1.5000 (Move up and out of the way)")

		initialY += incrementY
		initialF += incrementF
	}

	// Put us back at 0,0
	writeLine("X0.0000 Y0.0000") // this is part of the G00 above

	// and finish
	writeLine("M04 M30")

	for _, element := range lines {
		fmt.Printf("%s\n", element)
	}
}
