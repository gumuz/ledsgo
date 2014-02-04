package ledsgo

import (
	"fmt"
)

type Matrix struct {
	strip    LEDStrip
	width    int
	height   int
	lookupxy map[string]int
}

func (m *Matrix) Set(x, y int, color Color) {
	pos := m.lookupxy[fmt.Sprintf("%d,%d", x, y)]
	m.strip.Set(pos, color)
	m.strip.Update()
}

func CalculateLookupTable(width, height int) map[string]int {
	lookupxy := make(map[string]int)
	pos := width * height

	for y := 0; y < height; y++ {
		if y%2 == 0 {
			for x := 0; x < width; x++ {
				lookupxy[fmt.Sprintf("%d,%d", x, y)] = pos - 1
				pos--
			}
		} else {
			for x := width - 1; x >= 0; x-- {
				lookupxy[fmt.Sprintf("%d,%d", x, y)] = pos - 1
				pos--
			}
		}
	}
	return lookupxy
}

func NewMatrix(width, height int) *Matrix {
	m := &Matrix{
		strip:    NewLPD8806Strip(width * height),
		width:    width,
		height:   height,
		lookupxy: CalculateLookupTable(width, height),
	}
	m.strip.Reset()
	return m
}
