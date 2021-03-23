package simple_factory

import "fmt"

type Api interface {
	Print()
}

type implA struct {}
func (a *implA) Print() {
	fmt.Println("I am Impl A")
}

type implB struct{}
func (b *implB) Print() {
	fmt.Println("I am Impl B")
}

type Factory struct{}
func (f *Factory) GetApi(t int) Api {
	if t%2==0 {
		return &implA{}
	} else {
		return &implB{}
	}
}