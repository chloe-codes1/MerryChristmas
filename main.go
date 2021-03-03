package main

import (
	"fmt"
	"strings"
	"sync"
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

	// These variables essentially act as global variables since
	// all the functions in the main for loop can access any of them
	// It's usually ok for async stuff inside a function like we're doing
	// here. To pass them into functions, you would have to pass a pointer
	// to the WaitGroup variables
	var (
		// WaitGroup variable for the first snow function
		firstSnowWg sync.WaitGroup

		// WaitGroup variable for the tree function
		treeWg sync.WaitGroup

		// WaitGroup variable for the second snow function
		secondSnowWg sync.WaitGroup

		// WaitGroup variable for the stump function
		stumpWg sync.WaitGroup
	)

	for i := 0; i < height; i++ {
		// Add a queue of 1 to all the WaitGroup variables
		// We're adding 1 but it can be other integers as well
		firstSnowWg.Add(1)
		treeWg.Add(1)
		secondSnowWg.Add(1)
		stumpWg.Add(1)

		// Anonymous go function to wrap around the snow for-loop
		// This makes the snow code be asynchronous
		go func() {
			// snow
			for j := 0; j < (height-i)+blank; j++ {
				left += 3
				if j%5 == 0 && i%2 == 0 && left%2 == 0 {
					white.Print("*")
				} else {
					fmt.Print(" ")
				}
			}

			// We're done with the first snow function. Calling the
			// Done() function negates the WaitGroup variable by 1.
			firstSnowWg.Done()
		}() // The '()' syntax calls the anonymous function

		// Anonymous go function for the tree loop
		go func() {
			// Wait until the firstSnowWg is at 0.
			firstSnowWg.Wait()

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

			// Mark the firstSnowWg as finished
			treeWg.Done()
		}()

		// Anonymous go function for the tree loop
		go func() {
			// Wait until the treeWg is at 0. Since the tree goroutine will
			// wait until the firstSnowWg is at 0, this function will only start
			// after the tree goroutine is finished.
			treeWg.Wait()

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

			// secondSnowWg is finished now
			secondSnowWg.Done()
		}()

		// Anonymous go function for the stump loop. Since this function
		// depends on the loop integer from the main loop, we have to
		// ask for it.
		go func(i int) {
			// Wait until the second snow goroutine function is finished
			secondSnowWg.Wait()

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

			// Mark stumpWg as Done
			stumpWg.Done()

		}(i) // pass the i variable from the main loop to this anonymous function

		// We must tell the main loop to wait for the stump goroutine to finish
		// if we don't wait for the stumpWg here, the main loop will just iterate
		stumpWg.Wait()
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
