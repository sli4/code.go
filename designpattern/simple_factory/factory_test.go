package simple_factory

import "testing"

func TestSimpleFactory(t *testing.T) {
	factory := &Factory{}
	factory.GetApi(2).Print()
	factory.GetApi(3).Print()
}
