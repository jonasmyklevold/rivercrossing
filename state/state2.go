package state

import (
	"strconv"
)

//Om de er på båten (Boat)
var KyB = false
var RevB = false
var KornB = false
var HsB = false

//Om de er på land til venstre
var KyL = true
var RevL = true
var KornL = true
var HsL = true
var BoatL = true

//om der er på land til høyre
var KyR = false
var RevR = false
var KornR = false
var HsR = false
var BoatR = false

//State som bestemmer hvem er i båten
func StateBoat() (FinalState string) {

	FinalState = "Sts Korn B:" + strconv.FormatBool(KornB) + " | " + "Sts Hs B:" + strconv.FormatBool(HsB) + " | " + "Sts Korn B:" + strconv.FormatBool(KornB) + " | " + "Sts Rev B:" + strconv.FormatBool(RevB)

	return
}

// State som bestemmer hvem som er på land til venstre
func StateLandL() (FinalState string) {
	FinalState = "Sts Korn L:" + strconv.FormatBool(KornL) + " | " + "Sts Hs L:" + strconv.FormatBool(HsL) + " | " + "Sts Korn L:" + strconv.FormatBool(KornL) + " | " + "Sts Rev L:" + strconv.FormatBool(RevL)

	return
}

// State som bestemmer hvem som er på land til høyre
func StateLandR() (FinalState string) {
	FinalState = "Sts Korn R:" + strconv.FormatBool(KornR) + " | " + "Sts Hs R:" + strconv.FormatBool(HsR) + " | " + "Sts Korn R:" + strconv.FormatBool(KornR) + " | " + "Sts Rev R:" + strconv.FormatBool(RevR)

	return
}

// State som putter alle sammen
func FinalState() (FinalState string) {
	FinalState = StateBoat() + StateLandL() + StateLandR()
	return
}

func ViewState() string {
	return FinalState()
}
