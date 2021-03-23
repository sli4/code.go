package observer

import (
	"fmt"
	"testing"
)

func TestObserver(t *testing.T) {
	deployer := NewDeployer("istio deployer 12345")
	tester := &Tester{id: "TestAB"}
	securiter := &Securitier{id: "Sec360"}
	developers := &Developers{id: "Dev996"}
	deployer.register(tester)
	deployer.register(securiter)
	deployer.register(developers)

	fmt.Println("======Deploy Task Success======")
	deployer.notifyAll()
}
