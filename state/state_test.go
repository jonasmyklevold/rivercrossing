package state

import (
	"testing"
)

//Tester om alle er her
func TestBoat(t *testing.T) {

	wanted := "Sts Ky B:true | Sts Rev B:true | Sts Korn B:true | Sts HS B:true"

	state := StateBoat()
	if state == wanted {
		t.Log("Alle er på Båt")
	} else {
		t.Errorf("Ingen er på båten %q, ", state)
	}
}
func TestLandV(t *testing.T) {
	wanted := "Sts Ky V:true | Sts Rev V:true | Sts Korn V:true | Sts HS V:true | Sts Boat V: true"

	var state = StateLandL()
	if state == wanted {
		t.Log("Alle er på land til venstre")
	} else {
		t.Errorf("Ingen er på land til venstre %q, ", state)
	}
}

func TestLandE(t *testing.T) {

	wanted := "Sts Ky E:true | Sts Rev E:true | Sts Korn E:true | Sts HS E:true | Sts Boat E: true"

	state := StateLandR()
	if state == wanted {
		t.Log("Alle er på land til høyre")
	} else {
		t.Errorf("Ingen er på land til høyre %q, ", state)
	}
}
