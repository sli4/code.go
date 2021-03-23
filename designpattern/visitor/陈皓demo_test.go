package visitor

import (
	"fmt"
	"testing"
)

func TestCNN(t *testing.T) {
	info := Info{
		Name: "陈皓",
		Namespace: "Test",
		Cluster:"酒仙桥",
	}
	var v Visitor = &info
	v = NameVisitor{v}
	v = NamespaceVisitor{v}
	v = ClusterVisitor{v}

	fn := func(info *Info, err error) error {
		info.Name = "李帅"
		info.Namespace = "PROD"
		info.Cluster = "蓝汛"
		fmt.Println("我是谁，我在哪？")
		return nil
	}
	v.Visit(fn)
}
