package bridge

import "testing"

func TestBridge(t *testing.T) {
	mac := new(Mac)
	win := new(Windows)
	hp := new(HPPrinter)
	cannon := new(CannonPrinter)

	mac.setPrinter(hp)
	mac.print()

	mac.setPrinter(cannon)
	mac.print()

	win.setPrinter(hp)
	win.print()

	win.setPrinter(cannon)
	win.print()
}
