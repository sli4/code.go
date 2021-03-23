package adapter

import "fmt"

type M1CPU struct {

}
func (m *M1CPU) RunApple() {
	fmt.Println("我是苹果的M1 CPU，我很强的！")
}

type X86CPU struct {}
func (x *X86CPU) Run() {
	fmt.Println("我是CPU，What can I do for you ?")
}

type CPU interface {
	Run()
}

type App struct {
}
func (app *App) UseCpu(cpu CPU) {
	cpu.Run()
}

type M1Adapter struct{
	cpu *M1CPU
}
func (m *M1Adapter) Run() {
	fmt.Println("我是M1的适配器，我将调用M1 CPU")
	m.cpu.RunApple()
}

