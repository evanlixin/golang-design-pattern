package composite

import (
	"fmt"
	"reflect"
)

// 接口
type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	Print(string)
}

const (
	LeafNode = iota
	CompositeNode
)

func NewComponent(kind int, name string) Component {
	var c Component
	switch kind {
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()
	}

	c.SetName(name)
	return c
}

// 基类
type component struct {
	parent Component
	name   string
}

func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(parent Component) {
	c.parent = parent
}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

func (c *component) AddChild(Component) {}

func (c *component) Print(string) {}

// 匿名组合 继承
type Leaf struct {
	component
}

func NewLeaf() *Leaf {
	return &Leaf{}
}

// 方法重写
func (c *Leaf) Print(pre string) {
	fmt.Printf("%s-%s\n", pre, c.Name())
}

// 匿名组合,(继承)  整体部分关系
type Composite struct {
	component
	childs []Component
}

func NewComposite() *Composite {
	return &Composite{
		childs: make([]Component, 0),
	}
}

// 方法重写
func (c *Composite) AddChild(child Component) {
	child.SetParent(c)
	c.childs = append(c.childs, child)
}

func (c *Composite) Print(pre string) {
	fmt.Printf("%s+%s\n", pre, c.Name())
	pre += " "
	for _, comp := range c.childs {
		comp.Print(pre)
	}
}

// 员工
type Employee struct {
	name string
	dept string
	salary int

	subordinates []*Employee
}

func NewEmployee(name, dept string, sal int) *Employee {
	return &Employee{
		name:         name,
		dept:         dept,
		salary:       sal,
		subordinates: []*Employee{},
	}
}

func (em *Employee) Add(e *Employee) {
	em.subordinates = append(em.subordinates, e)
}

func (em *Employee) Remove(e *Employee) {
	em.subordinates = em.remove(e)
}

func (em *Employee) remove(e *Employee) []*Employee {
	newitems := []*Employee{}

	for _, sm := range em.subordinates {
		if !reflect.DeepEqual(e, sm) {
			newitems = append(newitems, sm)
		}
	}

	return newitems
}

func (em *Employee) GetSubordinates() []*Employee{
	return em.subordinates
}

func (em *Employee) ToString() string {
	return fmt.Sprintf("Employee :[ Name : %s] [dept: %s] [salary: %v]", em.name, em.dept, em.salary)
}

