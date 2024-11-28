package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"golang.org/x/term"
)

const (
	delay      = 20 * time.Millisecond
	coreString = ".,-~:;=!*#$@&§"
)

func floatMemset(arr []float64, v float64) {
	for i := range arr {
		arr[i] = v
	}
}

func byteMemset(arr []string, v string) {
	for i := range arr {
		arr[i] = v
	}
}

func getTerminalDimensions() (int, int, error) {
	return term.GetSize(0)
}

func main() {
	A := 0.0
	B := 0.0

	fmt.Print("\033[H\033[2J") // clear previous stdout

	for {
		width, height, err := getTerminalDimensions()
		if err != nil {
			panic(fmt.Sprintf("Failed to get the terminal size: %v", err))
		}

		z := make([]float64, width*height)
		b := make([]string, width*height)
		byteMemset(b, " ")
		floatMemset(z, 0)

		for j := 0.0; j < 6.28; j += 0.07 {
			for i := 0.0; i < 6.28; i += 0.02 {
				c := math.Sin(i)
				d := math.Cos(j)
				e := math.Sin(A)
				f := math.Sin(j)
				g := math.Cos(A)
				h := d + 2
				D := 1 / (c*h*e + f*g + 5)
				l := math.Cos(i)
				m := math.Cos(B)
				n := math.Sin(B)
				t := c*h*g - f*e

				x := int(float64(width)/2 + float64(width)/2*D*(l*h*m-t*n))
				y := int(float64(height)/2 + float64(height)/2*D*(l*h*n+t*m))

				o := x + width*y

				N := int(8 * ((f*e-c*d*g)*m - c*d*e - f*g - l*d*n))

				if y < height && y >= 0 && x >= 0 && x < width && D > z[o] {
					z[o] = D

					point := 0
					if N > 0 {
						point = N
					}

					b[o] = string(coreString[point])
				}
			}
		}

		print("\x1b[H")

		for k := 0; k < len(b); k++ {
			v := "\n"
			if k%width > 0 {
				v = b[k]
			}

			fmt.Print(v)

			// Generate a random float between -0.003 and -0.001
			randomFloat := rand.Float64() * (0.00002 - 0.00001)

			if rand.Intn(2) == 0 {
				A += randomFloat
			}

			if rand.Intn(2) == 0 {
				B += randomFloat
			}
		}

		time.Sleep(delay)
	}
}
