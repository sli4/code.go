package bridge

import "fmt"

type printer interface {
	print()
}

type Mac struct {
	printer printer
}
func (m *Mac) setPrinter(p printer) {
	m.printer = p
}
func (m *Mac) print() {
	fmt.Print("Mac 打印，")
	m.printer.print()
}

type Windows struct {
	printer printer
}
func (w *Windows) setPrinter(p printer) {
	w.printer =p
}
func (w *Windows) print() {
	fmt.Print("Windows 打印，")
	w.printer.print()
}

type HPPrinter struct {}
func (h *HPPrinter) print() {
	fmt.Println("惠普打印机！")
}

type CannonPrinter struct{}
func (c *CannonPrinter) print() {
	fmt.Println("佳能打印机！")
}