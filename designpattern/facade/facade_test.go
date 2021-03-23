package facade

import (
	"fmt"
	"testing"
)

func TestFacade(t *testing.T) {
	f := &Facade{}
	f.Deploy(1)
	fmt.Println("**********************")
	f.Deploy(2)
}
