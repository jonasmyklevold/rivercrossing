package event

import (
	"github.com/rivercrossing/state"
)

var BoatInfo = state.StateBoat()
var LeftInfo = state.StateLandL()
var RightInfo = state.StateLandR()

//Kylling i båt
func PutKyB() {
	state.KyB = true
	state.KyR = false
	state.KyL = false
}

//Rev i båt
func PutRevB() {
	state.RevB = true
	state.RevR = false
	state.RevL = false
}

//Korn i båt
func PutKornB() {
	state.KornB = true
	state.KornR = false
	state.KornL = false
}

//Menneske (Hs) i båt
func PutHsB() {
	state.HsB = true
	state.HsR = false
	state.HsL = false
}

//Kylling i land mot venstre
func PutKyL() {
	state.KyB = false
	state.KyR = false
	state.KyL = true
}

//Rev i land mot venstre
func PutRevL() {
	state.RevB = false
	state.RevR = false
	state.RevL = true
}

//Korn i land mot venstre
func PutKornL() {
	state.KornB = false
	state.KornR = false
	state.KornL = true
}

//Menneske (Hs) i land mot venstre
func PutHsL() {
	state.HsB = false
	state.HsR = false
	state.HsL = true
}

//Tar menneske i land mot venstre
func PutBoatL() {
	state.BoatR = false
	state.BoatL = true
}

//Tar Kylling i land mot høyre
func PutKyR() {
	state.KyB = false
	state.KyR = true
	state.KyL = false
}

//Tar Rev i land mot høyre
func PutRevR() {
	state.RevB = false
	state.RevR = true
	state.RevL = false
}

//Tar Korn i land mot høyre
func PutKornR() {
	state.KornB = false
	state.KornR = true
	state.KornL = false
}

//Menneske (Hs) i land mot Høyre
func PutHsR() {
	state.HsB = false
	state.HsR = true
	state.HsL = false
}

//Menneske (Hs) i land mot venstre
func PutBoatR() {
	state.BoatR = true
	state.BoatL = false
}

//Alle i båten
func PutAllB() {
	state.KyB = true
	state.KyR = false
	state.KyL = false
	state.RevB = true
	state.RevR = false
	state.RevL = false
	state.KornB = true
	state.KornR = false
	state.KornL = false
	state.HsB = true
	state.HsR = false
	state.HsL = false
}

//Tar alle i land til venstre + båten
func PutAllL() {
	state.KyB = false
	state.KyR = false
	state.KyL = true
	state.RevB = false
	state.RevR = false
	state.RevL = true
	state.KornB = false
	state.KornR = false
	state.KornL = true
	state.HsB = false
	state.HsR = false
	state.HsL = true
	state.BoatR = false
	state.BoatL = true
}

//Tar alle i Land Øst + båten
func PutAllR() {
	state.KyB = false
	state.KyR = true
	state.KyL = false
	state.RevB = false
	state.RevR = true
	state.RevL = false
	state.KornB = false
	state.KornR = true
	state.KornL = false
	state.HsB = false
	state.HsR = true
	state.HsL = false
	state.BoatR = true
	state.BoatL = false
}
