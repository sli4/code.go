package visitor

import "fmt"

type Info struct {
	Cluster string
	Namespace string
	Name string
}
func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info, nil)
}
type VisitorFunc func (*Info, error) error
type Visitor interface {
	Visit(visitorFunc VisitorFunc) error
}


type ClusterVisitor struct {
	visitor Visitor
}
func (v ClusterVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("ClusterVisitor before call function")
		err = fn(info, err)
		if err==nil {
			fmt.Printf("==> Cluster=%s\n", info.Cluster)
		}
		fmt.Println("ClusterVisitor after call function")
		return err
	})
}

type NamespaceVisitor struct {
	visitor Visitor
}
func (v NamespaceVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {{
		fmt.Println("SpaceVisitor befor call function")
		err = fn(info, err)
		if err==nil {
			fmt.Printf("==> Space=%s\n", info.Namespace)
		}
		fmt.Println("SpaceVisitor after call function")
		return err
	}})
}

type NameVisitor struct {
	visitor Visitor
}
func (v NameVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {{
		fmt.Println("NameVisitor befor call function")
		err = fn(info, err)
		if err==nil {
			fmt.Printf("==> Name=%s\n", info.Name)
		}
		fmt.Println("NameVisitor after call function")
		return err
	}})
}