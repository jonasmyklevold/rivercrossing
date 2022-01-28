package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jonasmyklevdold/rivercrossing/event"
	"github.com/jonasmyklevold/rivercrossing/state"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		fmt.Println("\033[H\033[2J")
		// info posisjonell commando
		if "info" == input.Text() {
			fmt.Println("Båt: " + state.StateBoat())

			// positionell command
			fmt.Println("Til Venstre: " + state.StateLandL())
			fmt.Println("Til høyre  :" + state.StateLandR())
		}

		// båt kommandoer
		if "PutBoatR" == input.Text() {
			event.PutBoatR()
			fmt.Println("Båten seilte til land på høyre side")
		}

		if "PutBoatL" == input.Text() {
			event.PutBoatL()
			fmt.Println("Båten seilte til land på venstre side")
		}

		// korn kommandoer
		if "PutKornB" == input.Text() {
			event.PutKornB()
			fmt.Println("Korn er lastet på båten.")
		}

		if "PutKornL" == input.Text() {
			event.PutKornL()
			fmt.Println("Korn er på land til venstre")
		}

		if "PutKornR" == input.Text() {
			event.PutKornR()
			fmt.Println("Korn er på land til høyre")
		}

		// kylling kommandoer
		if "PutKyB" == input.Text() {
			event.PutKyB()
			fmt.Println("Kyllingen er ombord i båten")
		}

		if "PutKyL" == input.Text() {
			event.PutKyL()
			fmt.Println("Kyllingen er på land til venstre")
		}

		if "PutKyR" == input.Text() {
			event.PutKyR()
			fmt.Println("Kyllingen er på land til høyre")
		}

		// menneske kommandoer
		if "PutHsB" == input.Text() {
			event.PutHsB()
			fmt.Println("Menneske er ombord i båten")
		}

		if "PutHsL" == input.Text() {
			event.PutHsL()
			fmt.Println("Menneske er på land til venstre")
		}

		if "PutHsR" == input.Text() {
			event.PutHsR()
			fmt.Println("Menneske er på land til høyre")
		}

		// rev kommandoer
		if "PutRevB" == input.Text() {
			event.PutRevB()
			fmt.Println("Reven er ombord i båten")
		}

		if "PutRevL" == input.Text() {
			event.PutRevL()
			fmt.Println("Reven er på land til venstre")
		}

		if "PutRevR" == input.Text() {
			event.PutRevR()
			fmt.Println("Reven er på land til høyre")
		}

		// plasserer kommandoene
		if "PutAllB" == input.Text() {
			event.PutAllB()
			fmt.Println("Alle er ombord i båten")
		}

		if "PutAllL" == input.Text() {
			event.PutAllL()
			fmt.Println("Alle er på land til venstre")
		}

		if "PutAllR" == input.Text() {
			event.PutAllR()
			fmt.Println("Alle er på land til venstre")
		}

	}

}
