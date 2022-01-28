package event

import (
	"testing"
)

// Test som tester funksjonen PutAllB
func TestKyoHsB(t *testing.T) {
	wanted := "Sts Korn:true | Sts Rev:true | Sts Ky:false  | Sts Hs:false "
	got := BoatInfo
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q. Enter er det ingen i båten, for mange i båten eller Mennesket mangler", got, wanted)
	} else {
		t.Log("Korn og Rev er i båten")
	}
}
