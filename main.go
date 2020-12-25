package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

const (
	height      = 16
	stumpHeight = 2
	stumpWidth  = 3
	blank       = 7
)

func main() {
	red := color.New(color.FgRed)
	green := color.New(color.FgGreen)
	yellow := color.New(color.FgYellow)
	blue := color.New(color.FgBlue)
	magenta := color.New(color.FgMagenta)
	cyan := color.New(color.FgCyan)
	white := color.New(color.FgWhite)

	left, right := height, height

	fmt.Println()

	for i := 0; i < height; i++ {
		// snow
		for j := 0; j < (height-i)+blank; j++ {
			left += 3
			if j%5 == 0 && i%2 == 0 && left%2 == 0 {
				white.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		// tree
		for j := height - 2*i; j <= height; j++ {
			switch {
			case i == 0:
				yellow.Print("C")
			case j%11 == 0 && i%2 == 0:
				yellow.Print("*")
			case j%9 == 0 && i%5 == 0:
				blue.Print("-")
			case j%5 == 0:
				red.Print("~")
			case j%3 == 0:
				cyan.Print("*")
			default:
				green.Print("*")
			}
		}

		// snow
		for j := 0; j < (height-i)+blank+2; j++ {
			right += 3
			if j%5 == 0 && i%2 == 0 && right%2 == 0 {
				white.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()

		// stump
		if i == height-1 {
			for h := 0; h < stumpHeight; h++ {
				for j := 0; j < (((height+i-stumpWidth)/2)+1)+blank; j++ {
					fmt.Print(" ")
				}
				for w := 0; w < stumpWidth; w++ {
					yellow.Print("|")
				}
				fmt.Println()
			}
		}
	}

	// message
	width := blank + height/2
	fmt.Print("\n" + strings.Repeat(" ", width))

	for _, char := range "Merry Christmas!" {
		magenta.Printf("%c", char)
		time.Sleep(time.Second / 5)
	}
	fmt.Println()
}
