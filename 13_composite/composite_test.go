package composite

import (
	"fmt"
	"testing"
)

func ExampleComposite() {
	root := NewComponent(CompositeNode, "root")
	c1 := NewComponent(CompositeNode, "c1")
	c2 := NewComponent(CompositeNode, "c2")
	c3 := NewComponent(CompositeNode, "c3")

	l1 := NewComponent(LeafNode, "l1")
	l2 := NewComponent(LeafNode, "l2")
	l3 := NewComponent(LeafNode, "l3")

	root.AddChild(c1)
	root.AddChild(c2)
	c1.AddChild(c3)
	c1.AddChild(l1)
	c2.AddChild(l2)
	c2.AddChild(l3)

	root.Print("")
	// Output:
	// +root
	//  +c1
	//   +c3
	//   -l1
	//  +c2
	//   -l2
	//   -l3
}

func TestNewEmployee(t *testing.T) {
	CEO := NewEmployee("John", "CEO", 30000)
	headSales := NewEmployee("Robert", "Head Sales", 20000)
	headMarketing := NewEmployee("Michel", "Head Marketing", 20000)

	clerk1 := NewEmployee("Laura", "Marketing", 10000)
	clerk2 := NewEmployee("Bob", "Marketing", 10000)

	salesExecutive1 := NewEmployee("Richard", "Sales", 10000)
	salesExecutive2 := NewEmployee("Rob", "Sales", 10000)

	CEO.Add(headSales)
	CEO.Add(headMarketing)

	headSales.Add(salesExecutive1)
	headSales.Add(salesExecutive2)

	headMarketing.Add(clerk1)
	headMarketing.Add(clerk2)

	fmt.Println(CEO.ToString())
	for _, e := range CEO.GetSubordinates() {
		fmt.Println(e.ToString())
		for _, ec := range e.GetSubordinates() {
			fmt.Println(ec.ToString())
		}
	}
}
